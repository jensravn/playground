# Sleepy Hue

Activates a Philips Hue scene for 25 minutes, then switches to a different scene when the timer ends.

## Installation

```bash
pip install requests
```

## Usage

```bash
python sleepy_hue.py
```

**First run** automatically discovers your Hue Bridge and asks you to press the button on it to create an API key. Configuration is saved to `~/.sleepy_hue.json`, so subsequent runs start immediately.

Interrupt at any time with **Ctrl+C** — the end scene activates immediately.

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
