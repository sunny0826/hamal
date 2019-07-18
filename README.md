# Hamal

`Hamal` is a tool for synchronizing images between two mirrored repositories.

```shell
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
hamal run -r guoxudongdocker/drone-dingtalk:latest

Usage:
  hamal run [flags]

Flags:
  -h, --help              help for run
  -r, --repontag string   docker repo:tag

Global Flags:
      --config string   config file (default is $HOME/.hamal/config.yaml)
```

#### configuration file

`$HOME/.hamal/config.yaml`

```yaml
author: Guo Xudong <sunnydog0826@gmail.com>
license: MIT
dinput:
  registry: <your-registry-input>
  user: <your-user-input>
  pass: <your-pass-input>
doutput:
  registry: <your-registry-output>
  user: <your-user-output>
  pass: <your-pass-output>

```