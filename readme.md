# PWW = Playerctl Waybar Wrapper

This is a simple program that emits a json with metadata info you can use in f.e. [Waybar](https://github.com/Alexays/Waybar). Additionally you can let it auto-pause media players you provide, f.e. starting a video in your browser will pause spotify.

## Install

`go install github.com/abenz1267/pww@latest`

Don't forget to add the `bin` folder of your `GOPATH` to your `PATH`.

`export PATH=$PATH:$(go env GOPATH)/bin`

## Usage

| command                           | description                                                                      | example                                                                            |
| --------------------------------- | -------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `pww -a <players>`                | will launch a daemon monitoring given players in order to auto-pause/resume them | `pww -a spotify,firefox`                                                           |
| `pww -w <player>:<metadata info>` | will watch for changes and print them to stdout                                  | `pww -w spotify:title` will emit f.e. `{"class": "Playing", "text": "Some Title"}` |
| `pww -t <player>`                 | will play/pause given player, launch it if it's not running at all               | `pww -t spotify`                                                                   |

Additionally a `-p` flag can be provided to specify a placeholder for empty text. F.e. when the player isn't running.

### Waybar Example

This is how you'd use it as a custom module in [Waybar](https://github.com/Alexays/Waybar).

```
//config
"custom/spotifytitle": {
    "format": "{}",
    "max-length": 40,
    "return-type": "json",
    "exec": "pww -w spotify:title -p None 2> /dev/null",
    "on-click": "playerctl --player=spotify play-pause 2> /dev/null",
    "on-scroll-up": "playerctl --player=spotify next 2> /dev/null",
    "on-scroll-down": "playerctl --player=spotify previous 2> /dev/null"
}

//style
#custom-spotifytitle {
  background: #98bb6c;
  color: #1f1f28;
  padding: 0 10px;
  margin-right: 10px;
  opacity: 1;
  transition-property: opacity;
  transition-duration: 0.25s;
}

#custom-spotifytitle.Paused,
#custom-spotifytitle.Inactive {
  opacity: 0.5;
}
```
