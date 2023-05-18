# fip

Fip is a command-line interface that fetches data from FIP Radio and provides information about the currently playing track, next track, and previous track.

## Installation
### Download
Download the [corresponding binary](https://github.com/piqoni/fip/releases) based on your OS and execute it (you may need to give permissions to the file by running `chmod +x fip-binary`)

### Gophers
If you have installed the [Go programming](https://go.dev) you can also install it by running: 

```shell
go install github.com/piqoni/fip
```
This will install it in your GOPATH/bin. To see your GOPATH you can execute `go env GOPATH`

## Usage
To get the current playing track just execute the binary: 

`fip`

You can also add options, for example if you want also the spotify link you can add --spotify to the command. List of all options:
```
fip --help 

Usage of fip:
  -next
    	Specify --next to show the next track
  -now
    	Specify --now to show the current track
  -prev
    	Specify --prev to show the previous track
  -spotify
    	Specify --spotify to add spotify link at the end of the track
```

## Ideas on how to use it
One potential usage is to create an alias that will append the currently playing track to a FIP bookmarks file. So whenever you hear a great track on FIP, you just write fav on terminal:

`alias fav='fip --spotify | tee -a "/Users/user/Library/Mobile Documents/iCloud~md~obsidian/Documents/FIP.md"'`


