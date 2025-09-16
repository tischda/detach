[![Build Status](https://github.com/tischda/detach/actions/workflows/build.yml/badge.svg)](https://github.com/tischda/detach/actions/workflows/build.yml)
[![Test Status](https://github.com/tischda/detach/actions/workflows/test.yml/badge.svg)](https://github.com/tischda/detach/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/tischda/detach/badge.svg)](https://coveralls.io/r/tischda/detach)
[![Go Report Card](https://goreportcard.com/badge/github.com/tischda/detach)](https://goreportcard.com/report/github.com/tischda/detach)

# detach

detach starts a new process specified by cmdPath and cmdArgs in a detached state on Windows.
The new process will not be attached to the current console and will run independently.

The process will inherit the current environment and update USER and SYSTEM environment
variables from the Windows registry.

I wrote this utility because [refresh](https://github.com/tischda/refresh) can't update the environment when
starting Alacritty from [`whkd`](https://github.com/LGUG2Z/whkd).

### Install

~~~
go install github.com/tischda/detach@latest
~~~

### Usage

~~~
Usage: detach [OPTIONS] cmdPath cmdArgs...

OPTIONS:

  -h, --help
        display this help message
  -v, --version
        print version and exit
~~~

### Examples

From terminal
~~~
$ detach alacritty.exe
  2025/08/30 22:32:16 Started detached process 'alacritty.exe' with PID 34568
~~~

whkd configuration (~\.config\whkdrc)
~~~
.shell cmd

# Custom keys
alt + return            : detach "C:\Program Files\Alacritty\alacritty.exe"
~~~
