from unittest.mock import patch

import sleepy_hue


@patch("subprocess.run")
def test_notify_uses_macos_notification(run_mock):
    sleepy_hue.notify("Focus", "Time is up")

    run_mock.assert_called_once()
    args = run_mock.call_args.args[0]
    assert args[0] == "osascript"
    assert args[1] == "-e"
    assert "display notification" in args[2]
    assert "Focus" in args[2]
