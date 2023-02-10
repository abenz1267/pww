# PWW = Playerctl Waybar Wrapper

This is a simple program that emits a json with artist or title info you can use in waybar. Additionally you can let it auto-pause media players you provide, f.e. starting a video in your browser will pause spotify.

## Install

`go install github.com/abenz1267/pww@latest`

## Usage

`pww -a firefox,spotify` will auto-pause/resume watching both Firefox and Spotify.
`pww -w spotify:title` will emit a json with the current spotify title and the Spotify status as a class. so `{"class": "Playing", "text": "Some Title"}`

Additionally a `-p` flag can be provided to specify a placeholder for empty text. F.e. when the player isn't running.

You can use `-t <player>` to toggle the playing state of the given player. If not present, the player will get launched.

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
