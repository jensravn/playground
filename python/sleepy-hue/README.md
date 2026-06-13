# Sleepy Hue

Activates a Philips Hue scene for 25 minutes, then switches to a different scene when the timer ends. Designed to build a daily focus habit — even 2 minutes counts.

## Installation

```bash
pip install requests
```

## Usage

```bash
python sleepy_hue.py                # Start a session
python sleepy_hue.py --until 15:30  # Start a session, stop automatically at 15:30
python sleepy_hue.py --stats        # Show session statistics
```

**First run** automatically discovers your Hue Bridge and asks you to press the button on it to create an API key. Configuration is saved to `~/.sleepy_hue.json`, so subsequent runs start immediately.

Interrupt at any time with **Ctrl+C** — the end scene activates and the session is logged.

## Habit tracking

- **Streak** — shown at startup and after each session. Any session of 2+ minutes counts toward the streak, so the bar to "show up" is intentionally low.
- **2-minute milestone** — the lights blink and a message appears once you've sat for 2 minutes, giving immediate positive feedback before the 25-minute timer ends.
- **Stats** — `--stats` shows a 14-day calendar, weekly and monthly comparisons, and your current streak.

## Configuration

Edit the constants at the top of `sleepy_hue.py`:

- `SCENE_NAME` — scene to activate at the start (default: `"work"`)
- `END_SCENE_NAME` — scene to activate when the timer ends (default: `"Natural light"`)
- `DURATION_SECONDS` — timer duration in seconds (default: 25 minutes)

Both regular scenes and Hue intelligent scenes are supported.

## Requirements

- Python 3
- `requests` (`pip install requests`)
- Philips Hue Bridge on the same network
