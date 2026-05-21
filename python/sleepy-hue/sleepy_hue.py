#!/usr/bin/env python3
import json
import os
import re
import subprocess
import sys
import time

import requests

requests.packages.urllib3.disable_warnings(
    requests.packages.urllib3.exceptions.InsecureRequestWarning
)

CONFIG_PATH = os.path.expanduser("~/.sleepy_hue.json")
DISCOVERY_URL = "https://discovery.meethue.com/"
DURATION_SECONDS = 25 * 60
SCENE_NAME = "work"
END_SCENE_NAME = "Natural light"


def load_config():
    if os.path.exists(CONFIG_PATH):
        with open(CONFIG_PATH) as f:
            return json.load(f)
    return {}


def save_config(config):
    with open(CONFIG_PATH, "w") as f:
        json.dump(config, f, indent=2)


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


def resolve_scene(scenes, smart_scenes, name):
    for scene_id, scene in scenes.items():
        if scene.get("name", "").lower() == name.lower():
            return ("v1", scene_id, scene.get("group", "0"))
    for s in smart_scenes:
        if s.get("metadata", {}).get("name", "").lower() == name.lower():
            return ("v2", s["id"])
    return None


def activate_scene(bridge_ip, username, scene_id, group_id):
    url = f"https://{bridge_ip}/api/{username}/groups/{group_id}/action"
    resp = requests.put(url, json={"scene": scene_id}, timeout=10, verify=False)
    print(f"  [v1] status={resp.status_code} body={resp.text[:200]}")


def activate_smart_scene(bridge_ip, username, smart_scene_id):
    url = f"https://{bridge_ip}/clip/v2/resource/smart_scene/{smart_scene_id}"
    headers = {"hue-application-key": username}
    resp = requests.put(
        url,
        headers=headers,
        json={"recall": {"action": "activate"}},
        timeout=10,
        verify=False,
    )
    print(f"  [v2] status={resp.status_code} body={resp.text[:200]}")


def activate_resolved_scene(bridge_ip, username, resolved):
    print(f"  type={resolved[0]} id={resolved[1]}")
    if resolved[0] == "v1":
        activate_scene(bridge_ip, username, resolved[1], resolved[2])
    else:
        activate_smart_scene(bridge_ip, username, resolved[1])


def blink_group(bridge_ip, username, group_id, times=3):
    url = f"https://{bridge_ip}/api/{username}/groups/{group_id}/action"
    for _ in range(times):
        requests.put(url, json={"alert": "select"}, timeout=10, verify=False)
        time.sleep(1.2)


def countdown(seconds):
    try:
        for remaining in range(seconds, 0, -1):
            mins, secs = divmod(remaining, 60)
            print(
                f"\rBlue light active — {mins:02d}:{secs:02d} remaining...  ",
                end="",
                flush=True,
            )
            time.sleep(1)
        print("\rBlue light active — 00:00 remaining...  ", flush=True)
    except KeyboardInterrupt:
        raise


def main():
    config = load_config()

    if "bridge_ip" not in config or "username" not in config:
        bridge_ip = discover_bridge()
        username = create_api_key(bridge_ip)
        config = {"bridge_ip": bridge_ip, "username": username}
        save_config(config)
    else:
        bridge_ip = config["bridge_ip"]
        username = config["username"]

    scenes = get_scenes(bridge_ip, username)
    smart_scenes = get_smart_scenes(bridge_ip, username)

    start = resolve_scene(scenes, smart_scenes, SCENE_NAME)
    if start is None:
        sys.exit(f"Scene '{SCENE_NAME}' not found in your Hue app.")

    end = resolve_scene(scenes, smart_scenes, END_SCENE_NAME)
    if end is None:
        sys.exit(f"Scene '{END_SCENE_NAME}' not found in your Hue app.")

    print(f"\nActivating scene '{SCENE_NAME}' for 25 minutes...")
    activate_resolved_scene(bridge_ip, username, start)

    group_id = start[2] if start[0] == "v1" else (end[2] if end[0] == "v1" else None)

    try:
        countdown(DURATION_SECONDS)
    except KeyboardInterrupt:
        pass
    finally:
        if group_id:
            print("\nBlinking...")
            blink_group(bridge_ip, username, group_id)
        print(f"Activating scene '{END_SCENE_NAME}'...")
        activate_resolved_scene(bridge_ip, username, end)


if __name__ == "__main__":
    main()
