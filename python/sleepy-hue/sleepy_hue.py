#!/usr/bin/env python3
import argparse
import datetime
import json
import os
import re
import subprocess
import sys
import time
from collections import defaultdict

import requests

requests.packages.urllib3.disable_warnings(
    requests.packages.urllib3.exceptions.InsecureRequestWarning
)

CONFIG_PATH = os.path.expanduser("~/.sleepy_hue.json")
SESSION_LOG_PATH = os.path.expanduser("~/.sleepy_hue_sessions.json")
DISCOVERY_URL = "https://discovery.meethue.com/"
DURATION_SECONDS = 25 * 60
SCENE_NAME = "work"
END_SCENE_NAME = "Natural light"

SPOTIFY_DEFAULT_URI = "spotify:playlist:37i9dQZF1DWZeKCadgRdKQ"  # Deep Focus

DND_ON_SHORTCUT = "DND On"
DND_OFF_SHORTCUT = "DND Off"

ENGLISH_MONTHS = [
    "Jan", "Feb", "Mar", "Apr", "May", "Jun",
    "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
]


def load_config():
    if os.path.exists(CONFIG_PATH):
        with open(CONFIG_PATH) as f:
            return json.load(f)
    return {}


def save_config(config):
    with open(CONFIG_PATH, "w") as f:
        json.dump(config, f, indent=2)


def ensure_spotify_playlist(config):
    if "spotify_playlist" not in config:
        print(
            f"\nSpotify integration: enter a playlist URI (Enter for Deep Focus, 'skip' to disable):"
        )
        uri = input(f"  [{SPOTIFY_DEFAULT_URI}]: ").strip()
        if uri.lower() == "skip":
            config["spotify_playlist"] = ""
        else:
            config["spotify_playlist"] = uri or SPOTIFY_DEFAULT_URI
        save_config(config)
    return config


def spotify_play(uri):
    if not uri:
        return
    script = f'tell application "Spotify" to play track "{uri}"'
    try:
        subprocess.run(["osascript", "-e", script], check=False, capture_output=True, timeout=5)
    except Exception:
        pass


def spotify_pause():
    script = 'tell application "Spotify" to pause'
    try:
        subprocess.run(["osascript", "-e", script], check=False, capture_output=True, timeout=5)
    except Exception:
        pass


def shortcut_exists(name):
    try:
        result = subprocess.run(
            ["shortcuts", "list"], capture_output=True, text=True, timeout=10
        )
        return name in result.stdout.splitlines()
    except Exception:
        return False


def ensure_dnd(config):
    if "dnd_enabled" not in config:
        if shortcut_exists(DND_ON_SHORTCUT) and shortcut_exists(DND_OFF_SHORTCUT):
            print("\nDo Not Disturb: found shortcuts, enabling during sessions.")
            config["dnd_enabled"] = True
        else:
            print(
                f"\nDo Not Disturb integration: create two macOS Shortcuts named "
                f"'{DND_ON_SHORTCUT}' and '{DND_OFF_SHORTCUT}' (each with a 'Set Focus' action)."
            )
            choice = input("Enable Do Not Disturb during sessions? [y/N]: ").strip().lower()
            config["dnd_enabled"] = choice == "y"
        save_config(config)
    return config


def set_dnd(config, on):
    if not config.get("dnd_enabled"):
        return
    name = DND_ON_SHORTCUT if on else DND_OFF_SHORTCUT
    try:
        subprocess.run(
            ["shortcuts", "run", name], check=False, capture_output=True, timeout=10
        )
    except Exception:
        pass


def ensure_logseq_vault(config):
    if "logseq_vault" not in config:
        print(
            "\nLogseq integration: enter the path to your vault (press Enter to skip):"
        )
        path = input("Logseq vault path: ").strip()
        if path:
            config["logseq_vault"] = path
            save_config(config)
    return config


def discover_bridge_arp():
    result = subprocess.run(["arp", "-a"], capture_output=True, text=True)
    for line in result.stdout.splitlines():
        if re.search(r"0?0:17:88:", line, re.IGNORECASE):
            match = re.search(r"\((\d+\.\d+\.\d+\.\d+)\)", line)
            if match:
                return match.group(1)
    return None


def discover_bridge():
    print("Searching for Hue Bridge on the network...")

    ip = discover_bridge_arp()
    if ip:
        print(f"Found bridge via network: {ip}")
        return ip

    try:
        resp = requests.get(DISCOVERY_URL, timeout=10)
        bridges = resp.json()
        if bridges:
            ip = bridges[0]["internalipaddress"]
            print(f"Found bridge via cloud: {ip}")
            return ip
    except Exception:
        pass

    print("Automatic discovery failed.")
    ip = input("Enter bridge IP address manually: ").strip()
    if not ip:
        sys.exit("No IP address provided.")
    return ip


def create_api_key(bridge_ip):
    print("\nPress the button on your Hue Bridge, then press Enter here...")
    input()

    url = f"https://{bridge_ip}/api"
    resp = requests.post(
        url, json={"devicetype": "sleepy_hue#script"}, timeout=10, verify=False
    )
    result = resp.json()

    if isinstance(result, list) and "success" in result[0]:
        username = result[0]["success"]["username"]
        print("API key created.")
        return username
    elif isinstance(result, list) and "error" in result[0]:
        error_desc = result[0]["error"].get("description", "unknown error")
        sys.exit(f"Bridge error: {error_desc}")
    else:
        sys.exit(f"Unexpected response from bridge: {result}")


def get_scenes(bridge_ip, username):
    url = f"https://{bridge_ip}/api/{username}/scenes"
    resp = requests.get(url, timeout=10, verify=False)
    return resp.json()


def get_smart_scenes(bridge_ip, username):
    url = f"https://{bridge_ip}/clip/v2/resource/smart_scene"
    headers = {"hue-application-key": username}
    resp = requests.get(url, headers=headers, timeout=10, verify=False)
    return resp.json().get("data", [])


def get_groups(bridge_ip, username):
    url = f"https://{bridge_ip}/api/{username}/groups"
    resp = requests.get(url, timeout=10, verify=False)
    return resp.json()


def ensure_rooms(config, bridge_ip, username):
    if "rooms" not in config:
        groups = get_groups(bridge_ip, username)
        room_list = sorted(
            [(gid, g["name"]) for gid, g in groups.items() if g.get("type") in ("Room", "Zone")],
            key=lambda x: x[1],
        )
        print("\nAvailable rooms:")
        for _, name in room_list:
            print(f"  - {name}")
        print("Enter room names to use (comma-separated), or Enter to use all:")
        choice = input(": ").strip()
        config["rooms"] = [r.strip() for r in choice.split(",")] if choice else []
        save_config(config)
    return config


def resolve_scene(scenes, smart_scenes, name):
    for scene_id, scene in scenes.items():
        if scene.get("name", "").lower() == name.lower():
            return ("v1", scene_id, scene.get("group", "0"))
    for s in smart_scenes:
        if s.get("metadata", {}).get("name", "").lower() == name.lower():
            return ("v2", s["id"], None)
    return None


def resolve_scenes(scenes, smart_scenes, name, groups=None, room_names=None):
    if not room_names:
        result = resolve_scene(scenes, smart_scenes, name)
        return [result] if result else []
    room_group_ids = {gid for gid, g in (groups or {}).items() if g.get("name") in room_names}
    results = []
    for scene_id, scene in scenes.items():
        if scene.get("name", "").lower() == name.lower() and scene.get("group") in room_group_ids:
            results.append(("v1", scene_id, scene["group"]))
    if not results:
        for s in smart_scenes:
            if s.get("metadata", {}).get("name", "").lower() == name.lower():
                results.append(("v2", s["id"], None))
    return results


def activate_scene(bridge_ip, username, scene_id, group_id):
    url = f"https://{bridge_ip}/api/{username}/groups/{group_id}/action"
    requests.put(url, json={"scene": scene_id}, timeout=10, verify=False)


def activate_smart_scene(bridge_ip, username, smart_scene_id):
    url = f"https://{bridge_ip}/clip/v2/resource/smart_scene/{smart_scene_id}"
    headers = {"hue-application-key": username}
    requests.put(
        url,
        headers=headers,
        json={"recall": {"action": "activate"}},
        timeout=10,
        verify=False,
    )


def activate_resolved_scene(bridge_ip, username, resolved):
    if resolved[0] == "v1":
        activate_scene(bridge_ip, username, resolved[1], resolved[2])
    else:
        activate_smart_scene(bridge_ip, username, resolved[1])


def blink_group(bridge_ip, username, group_id, times=3):
    url = f"https://{bridge_ip}/api/{username}/groups/{group_id}/action"
    for _ in range(times):
        requests.put(url, json={"alert": "select"}, timeout=10, verify=False)
        time.sleep(1.2)


def countdown(seconds, stop_at=None):
    stop_label = f" (stopping at {stop_at.strftime('%H:%M')})" if stop_at else ""
    try:
        for remaining in range(seconds, 0, -1):
            elapsed = seconds - remaining
            if elapsed == 120:
                print("\n✓ 2 minutes — you showed up. That counts!")
            mins, secs = divmod(elapsed, 60)
            print(
                f"\rBlue light active — {mins:02d}:{secs:02d} elapsed{stop_label}  ",
                end="",
                flush=True,
            )
            time.sleep(1)
        total_mins, total_secs = divmod(seconds, 60)
        print(f"\rBlue light active — {total_mins:02d}:{total_secs:02d} elapsed{stop_label}  ", flush=True)
    except KeyboardInterrupt:
        raise


def fmt_duration(minutes):
    if minutes >= 60:
        h, m = divmod(minutes, 60)
        return f"{h}h {m}min" if m else f"{h}h"
    return f"{minutes}min"


def append_to_logseq(config, start_time, end_time, completed):
    vault = config.get("logseq_vault")
    if not vault:
        return

    pages_dir = os.path.join(vault, "pages")
    if not os.path.isdir(pages_dir):
        return

    pomodoro_path = os.path.join(pages_dir, "pomodoro.md")
    date_ref = f"[[{start_time.strftime('%Y-%m-%d')}]]"
    duration = int((end_time - start_time).total_seconds() / 60)

    lines = []
    if os.path.exists(pomodoro_path):
        with open(pomodoro_path) as f:
            lines = f.readlines()

    updated = False
    for i, line in enumerate(lines):
        if date_ref in line:
            m = re.search(r"(\d+)h(?:\s(\d+)min)?|(\d+)min", line)
            if m:
                if m.group(3):
                    existing_min = int(m.group(3))
                else:
                    existing_min = int(m.group(1)) * 60 + int(m.group(2) or 0)
                total_min = existing_min + duration
                lines[i] = re.sub(
                    r"\d+h(?:\s\d+min)?|\d+min",
                    fmt_duration(total_min),
                    line,
                )
            updated = True
            break

    if not updated:
        lines.append(f"- {date_ref} – {fmt_duration(duration)}\n")

    with open(pomodoro_path, "w") as f:
        f.writelines(lines)


def log_session(config, start_time, end_time, completed):
    sessions = []
    if os.path.exists(SESSION_LOG_PATH):
        with open(SESSION_LOG_PATH) as f:
            sessions = json.load(f)

    duration = int((end_time - start_time).total_seconds() / 60)
    sessions.append(
        {
            "start": start_time.isoformat(timespec="seconds"),
            "end": end_time.isoformat(timespec="seconds"),
            "duration_minutes": duration,
            "completed": completed,
        }
    )

    with open(SESSION_LOG_PATH, "w") as f:
        json.dump(sessions, f, indent=2)

    append_to_logseq(config, start_time, end_time, completed)


def get_streak():
    if not os.path.exists(SESSION_LOG_PATH):
        return 0
    with open(SESSION_LOG_PATH) as f:
        sessions = json.load(f)
    completed_days = {s["start"][:10] for s in sessions if s.get("completed")}
    today = datetime.date.today()
    streak = 0
    d = today
    while d.isoformat() in completed_days:
        streak += 1
        d -= datetime.timedelta(days=1)
    return streak


def get_today_minutes():
    if not os.path.exists(SESSION_LOG_PATH):
        return 0
    with open(SESSION_LOG_PATH) as f:
        sessions = json.load(f)
    today = datetime.date.today().isoformat()
    return sum(s["duration_minutes"] for s in sessions if s["start"][:10] == today)


def format_today_time(minutes):
    if minutes <= 0:
        return None
    h, m = divmod(minutes, 60)
    return f"{h}h {m}min" if h else f"{m}min"


def show_stats():
    if not os.path.exists(SESSION_LOG_PATH):
        print("No sessions logged yet.")
        return

    with open(SESSION_LOG_PATH) as f:
        sessions = json.load(f)

    if not sessions:
        print("No sessions logged yet.")
        return

    by_date = defaultdict(int)
    minutes_by_date = defaultdict(int)
    for s in sessions:
        date = datetime.date.fromisoformat(s["start"][:10])
        by_date[date] += 1
        minutes_by_date[date] += s.get("duration_minutes", 0)

    today = datetime.date.today()
    streak = get_streak()
    today_minutes = get_today_minutes()
    print("\n=== Sleepy Hue sessions ===\n")
    if streak > 0:
        time_str = format_today_time(today_minutes) or "0min"
        print(f"🔥 Streak: {streak} days — today: {time_str}\n")

    print(f"{'Date':<15}{'Sessions':>9}{'Time':>10}")
    print("-" * 34)
    for i in range(13, -1, -1):
        d = today - datetime.timedelta(days=i)
        marker = " <" if d == today else ""
        sessions_count = by_date[d]
        time_col = fmt_duration(minutes_by_date[d]) if minutes_by_date[d] else ""
        print(f"{d.isoformat():<15}{sessions_count:>9}{time_col:>10}{marker}")

    monday_this = today - datetime.timedelta(days=today.weekday())
    monday_last = monday_this - datetime.timedelta(weeks=1)

    this_week = sum(v for d, v in by_date.items() if monday_this <= d <= today)
    last_week = sum(v for d, v in by_date.items() if monday_last <= d < monday_this)

    mth = ENGLISH_MONTHS[monday_this.month - 1]
    week_label = (
        f"{monday_this.day} {mth}–{today.day} {ENGLISH_MONTHS[today.month - 1]}"
    )
    print(f"\nThis week ({week_label}):  {this_week} sessions")
    print(f"Last week:                    {last_week} sessions")

    this_month = sum(
        v for d, v in by_date.items() if d.year == today.year and d.month == today.month
    )
    last_month_date = today.replace(day=1) - datetime.timedelta(days=1)
    last_month = sum(
        v
        for d, v in by_date.items()
        if d.year == last_month_date.year and d.month == last_month_date.month
    )

    print(
        f"\nThis month ({ENGLISH_MONTHS[today.month - 1]}):              {this_month} sessions"
    )
    print(
        f"Last month ({ENGLISH_MONTHS[last_month_date.month - 1]}):            {last_month} sessions"
    )

    print(f"\nTotal:                          {sum(by_date.values())} sessions\n")


def main():
    parser = argparse.ArgumentParser(
        description="Sleepy Hue — activate sleep scene with timer"
    )
    parser.add_argument(
        "--stats", action="store_true", help="Show session statistics"
    )
    parser.add_argument(
        "--until", metavar="HH:MM", help="Stop automatically at this time (e.g. 15:30)"
    )
    args = parser.parse_args()

    if args.stats:
        show_stats()
        return

    streak = get_streak()
    today_minutes = get_today_minutes()
    time_str = format_today_time(today_minutes)

    if streak > 0 and time_str:
        print(f"🔥 Streak: {streak} days — today: {time_str} — even 2 min counts!")
    elif streak > 0:
        print(f"🔥 Streak: {streak} days — even 2 min counts today")
    elif time_str:
        print(f"Today: {time_str} — even 2 min counts!")
    else:
        print("Ready to start — even 2 min counts!")

    config = load_config()

    if "bridge_ip" not in config or "username" not in config:
        bridge_ip = discover_bridge()
        username = create_api_key(bridge_ip)
        config = {"bridge_ip": bridge_ip, "username": username}
        save_config(config)
    else:
        bridge_ip = config["bridge_ip"]
        username = config["username"]

    config = ensure_logseq_vault(config)
    config = ensure_spotify_playlist(config)
    config = ensure_dnd(config)
    config = ensure_rooms(config, bridge_ip, username)

    try:
        scenes = get_scenes(bridge_ip, username)
        smart_scenes = get_smart_scenes(bridge_ip, username)
    except requests.exceptions.ConnectionError:
        sys.exit(
            f"Cannot reach Hue Bridge at {bridge_ip} — are you on the right network?"
        )

    room_names = config.get("rooms", [])
    groups = get_groups(bridge_ip, username) if room_names else {}

    starts = resolve_scenes(scenes, smart_scenes, SCENE_NAME, groups, room_names)
    if not starts:
        sys.exit(f"Scene '{SCENE_NAME}' not found in your Hue app.")

    ends = resolve_scenes(scenes, smart_scenes, END_SCENE_NAME, groups, room_names)
    if not ends:
        sys.exit(f"Scene '{END_SCENE_NAME}' not found in your Hue app.")

    stop_at = None
    duration = DURATION_SECONDS
    if args.until:
        try:
            stop_time = datetime.datetime.strptime(args.until, "%H:%M").replace(
                year=datetime.datetime.now().year,
                month=datetime.datetime.now().month,
                day=datetime.datetime.now().day,
            )
            secs_until = int((stop_time - datetime.datetime.now()).total_seconds())
            if secs_until <= 0:
                sys.exit(f"--until {args.until} is already in the past.")
            duration = min(secs_until, DURATION_SECONDS)
            stop_at = stop_time
        except ValueError:
            sys.exit("Invalid time format — use HH:MM, e.g. --until 15:30")

    if stop_at:
        print(f"\nActivating scene '{SCENE_NAME}' — stopping at {stop_at.strftime('%H:%M')}...")
    else:
        print(f"\nActivating scene '{SCENE_NAME}' for 25 minutes...")
    for scene in starts:
        activate_resolved_scene(bridge_ip, username, scene)
    spotify_play(config.get("spotify_playlist"))
    set_dnd(config, on=True)

    group_ids = [s[2] for s in starts if s[0] == "v1" and s[2]]
    if not group_ids:
        group_ids = [s[2] for s in ends if s[0] == "v1" and s[2]]

    session_start = datetime.datetime.now()
    full_session = False
    try:
        countdown(duration, stop_at=stop_at)
        full_session = True
    except KeyboardInterrupt:
        pass
    finally:
        session_end = datetime.datetime.now()
        elapsed_seconds = (session_end - session_start).total_seconds()
        completed = full_session or elapsed_seconds >= 120
        if group_ids:
            end_blinks = 3 if full_session else 1
            print("\nBlinking...")
            for gid in group_ids:
                blink_group(bridge_ip, username, gid, times=end_blinks)
        spotify_pause()
        set_dnd(config, on=False)
        print(f"Activating scene '{END_SCENE_NAME}'...")
        for scene in ends:
            activate_resolved_scene(bridge_ip, username, scene)
        log_session(config, session_start, session_end, completed)
        if completed:
            streak = get_streak()
            today_minutes = get_today_minutes()
            time_str = format_today_time(today_minutes) or "0min"
            print(f"\nYou showed up. Streak: {streak} days 🔥 — Today: {time_str}")


if __name__ == "__main__":
    main()
