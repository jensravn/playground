#!/usr/bin/python

import os

from dotenv import load_dotenv
from phue import Bridge

# Load environment variables from .env file
load_dotenv()

# See IP in Hue app - go to Settings > Hue Bridges
ip = os.getenv("HUE_BRIDGE_IP")

if not ip:
    raise ValueError("HUE_BRIDGE_IP environment variable is not set")

b = Bridge(ip)

# If the app is not registered and the button is not pressed, press the button and call connect() (this only needs to be run a single time)
b.connect()

# Get the bridge state (This returns the full dictionary that you can explore)
b.get_api()

lights = b.lights

# Print light names
for l in lights:
    print(l.name)

# Set brightness of each light
for l in lights:
    l.brightness = 255

print("Set brightness of all lights to 255")
