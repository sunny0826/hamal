# Hamal

[![Build Status](https://travis-ci.org/sunny0826/hamal.svg?branch=master)](https://travis-ci.org/sunny0826/hamal)
[![Go Report Card](https://goreportcard.com/badge/github.com/sunny0826/hamal)](https://goreportcard.com/report/github.com/sunny0826/hamal)
![GitHub](https://img.shields.io/github/license/sunny0826/hamal.svg)

`Hamal` is a tool for synchronizing images between two mirrored repositories.

```
 _   _                       _ 
| | | | __ _ _ __ ___   __ _| |
| |_| |/ _\ | '_ \ _ \ / _\ | |
|  _  | (_| | | | | | | (_| | |
|_| |_|\__,_|_| |_| |_|\__,_|_|

Hamal is a tool for synchronizing images between two mirrored repositories. 
You can synchronize mirrors between two private image repositories.

WARN:The docker must be installed locally.
Currently only Linux and MacOS are supported.

Usage:
  hamal [flags]
  hamal [command]

Available Commands:
  help        Help about any command
  run         Start syncing mirror

Flags:
      --config string   config file (default is $HOME/.hamal/config.yaml)
  -h, --help            help for hamal

Use "hamal [command] --help" for more information about a command.

----------------------------------------------------------------------------

For details, please see: https://github.com/sunny0826/hamal.git

example:
hamal run -n drone-dingtalk:latest

Usage:
  hamal run [flags]

Flags:
  -h, --help          help for run
  -n, --name string   docker name:tag

Global Flags:
      --config string   config file (default is $HOME/.hamal/config.yaml)

```

#### configuration file

`$HOME/.hamal/config.yaml`

```yaml
author: <your-name>
license: MIT
dinput:
#  registry: <your-registry-input>    # if used dockerhub ,do not need registry
  repo: <your-repo-input>
  user: <your-user-input>
  pass: <your-pass-input>
  isdockerhub: true                   # use dockerhub
doutput:
  registry: <your-registry-input>
  repo: <your-repo-output>
  user: <your-user-output>
  pass: <your-pass-input>
  isdockerhub: false
```