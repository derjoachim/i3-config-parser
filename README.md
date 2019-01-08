# i3-config-parser
Find and parse an i3 configuration file and display in a nice way.

## Introduction

This is a reimplementation of [my earlier i3wm confiiguration parser](https://github.com/derjoachim/joatools/blob/master/i3wm_help.py). My main goal is to get more familiar with the go language and is solely intended for personal use. However, if anyone is happy with this tool, they are free to use it. See the LICENCE file.

## Configuration

### Prerequisites
- Make sure that you have golang installed on your system and your go workspace is properly configured.
- Obviously, you should have i3wm installed of at least a i3wm config file in one of the appropriate locations.

### Installation
- Clone (or fork and clone) in your go development environment.
- Run `go build i3-confiig-parser.go`.
- Run `go install i3-config-parser.go`.

## How it works

As per the [excellent i3wm documentation](https://i3wm.org/docs/userguide.html), there are several possible places to find a configuration file for i3. The config parser will try to find a configuration file and parse its settings and keybindings. 

The most important setting is the configured modifier key, typically Left Alt. Mod keys can be overridden through Xmodmap, so a nice feature would be to search a `.Xmodmap` file for overrides. Since I do not use one, I do not have the intention to add this functionality just yet. :-)

Variables in key bindings are substituted: `$mod+Return` => `WinL+Return`

## TODO

- Display in a nice window. For now, the file is parsed into a struct that is displayed on the CLI. Still working on that one.
- Support for `.xmodmap` files . I do not use one myself, so this one has a low priority.
- Automated testing. Just to get used to golang testing.
- Group key bindings by number (e.g. `Mod+1-9` => Switch to workspace 1-9)
- Display of modes, e.g. resize modes.
