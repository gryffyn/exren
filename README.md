# exren

The EXif RENamer. Renames image files based on EXIF tags and a format string.

## Installation

`go get git.neveris.one/gryffyn/exren`

## Usage

```
NAME:
   exren - the exif renamer

USAGE:
   exren [global options] command [command options] [arguments...]

VERSION:
   v0.1.3

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --format value, -f value  Output format, including extension
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)

```  

The format string takes exif tags in the format `%TagName%`.

## Screenshot

![usage](img/1.png)

Copyright 2021 gryffyn, see LICENSE for details.  
EXIF parsing library copyright [rwcarlsen](https://github.com/rwcarlsen).
