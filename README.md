# gotmoji

> Badges

> A interactive command line tool for using emojis on commits without dependencies.

## About

This project provides an cli tool for usign [gitmoji](https://gitmoji.carloscuesta.me/) from your command line.

## Alternatvies

- official [gitmoji-cli](https://gitmoji.carloscuesta.me/)

## Install

## Usage

First add files to the stage as usual with "git add" files. Then "gotmoji" instead of "git commit" can be used and an interactive dialog starts.

With the parameter "-u" the UTF8 version of the emoji is used instead of the code (for example ": fire:").

```
gotmoji -u
```

```bash
gotmoji --help
A interactive command line tool for using emojis on commits
      without dependencies.
      Gotmoji: https://github.com/pkuebler/gotmoji/

      Gitmoji Overview: https://gitmoji.carloscuesta.me/
      NodeJS - Orginal: https://github.com/carloscuesta/gitmoji-cli

Usage:
  gotmoji [flags]
  gotmoji [command]

Available Commands:
  commit      commit
  help        Help about any command
  list        Print all emojis with description
  version     Print the version number of gotmoji

Flags:
  -h, --help   help for gotmoji
  -u, --utf8   utf8 emoji (default :code:)

Use "gotmoji [command] --help" for more information about a command.
```
