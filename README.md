# Good

Personal playground composed of CLI commands made in Go.

## Prerequisites

[Install Go](https://go.dev/doc/install)

## List of available commands

### Fun commands

* `good fun catfact`: get absolutely essential facts about cats
* `good fun kanji`: get meanings and other info about a specific kanji
* `good fun spinthewheel`: Don't know what to decide? Maybe leave it to chance...
* `good fun 1337`: Convert a string to leetspeak and shine on Internet

### Checking tools

* `good check archi`: check your architecture with the Topology function
* `good check bios`: check your BIOS
* `good check chassis`: check your chassis info
* `good check cpu`: check your **C**PU
* `good check gpu`: check your **G**PU
* `good check ip`: check your public IP with ipinfo.io
* `good check memory`: check your memory
* `good check network`: check your network
* `good check os`: check your current OS
* `good check product`: check your product info (~ the "about" of your machine)
* `good check storage`: check your storage
* `good check version`: check the current version of the command-line tool

_N.B.: Many commands are just wrappers for utils provided by the [ghw package](https://github.com/jaypipes/ghw)_

### Legit Runners

* `good run countdown`: interesting countdown...
* `good run htmlcom`: check if the web page contains HTML comments
* `good run htmlimg`: download all images you can find on a given URL in `~/good-images/`
* `good run passwordgen`: generate long & random passwords

### Hacking tools

* `good hack cipher`: encipher/decipher a secret message
* `good hack ports`: scan **open** ports on IP
* `good hack rot13`: encode/decode a rot13 message
* `good hack obfuscate`: obfuscate/deobfuscate string with hexadecimal encoding
* `good hack keylogger`: start a keylogger (as root user)

## Installation

```
go install github.com/jmau111-org/good@latest
```

## Help

```
good help
```

## Compilation

```
go get && go build
```
