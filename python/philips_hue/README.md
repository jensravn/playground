# Philips Hue

https://github.com/studioimaginaire/phue

1. Create virtual environment
```
python3 -m venv path/to/venv
```

2. Activate virtual enviroment

```
source path/to/venv/bin/activate
```

3. Install dependencies

```
python3 -m pip install -r requirements.txt
```

4. Create `.env` file (copy `.env.example`) and set `HUE_BRIDGE_IP`.

5. Press button on Hue Bridge to connect first time.

6. Run the program

```
python main.py
```