# PWW = Playerctl Waybar Wrapper

This is a simple program that emits a json with artist or title info you can use in waybar. Additionally you can let it auto-pause media players you provide, f.e. starting a video in your browser will pause spotify.

## Usage

`pww -a firefox,spotify` will autopause watching both Firefox and Spotify.
`pww -w spotify:title` will emit a json with the current spotify title and the Spotify status as a class. so `{"class": "Playing", "text": "Some Title"}`

Additionally a `-p` flag can be provided to specify a placeholder for empty text. F.e. when the player isn't running.

You can use `-t <player>` to toggle the playing state of the given player. If not present, the player will get launched.
