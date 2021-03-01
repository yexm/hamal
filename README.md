# This tool stores passwords explicitly and is not being maintained at the moment, so please use with caution.

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
  hamal [command]

Available Commands:
  help        Help about any command
  run         Start syncing mirror
  version     Prints the huamal version

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
## Quick Start

### Install hamal

Homebrew:

```bash
brew install sunny0826/tap/hamal
```

Download the binary:

```bash
# linux x86_64
curl -Lo kubecm.tar.gz https://github.com/sunny0826/hamal/releases/download/v${VERSION}/kubecm_${VERSION}_Linux_x86_64.tar.gz
# macos
curl -Lo kubecm.tar.gz https://github.com/sunny0826/hamal/releases/download/v${VERSION}/kubecm_${VERSION}_Darwin_x86_64.tar.gz

tar -zxvf hamal.tar.gz hamal
sudo mv hamal /usr/local/bin/
```

### Init hamal
```bash
hamal init
```

### configuration file

```bash
# set config dinput
hamal init set-input -u test -p test123 -n name -r registry.test.com

# set dockerhub
hamal init set-input -u test -p test123 -n name -d

# set config dinput
hamal set-output -u gxd -p 123 -n name -r registry-vpc.gxd.com

# set dockerhub
hamal set-output -u gxd -p 123 -n name -r registry-vpc.gxd.com -d
```

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

### Run

```bash
# sync docker image
hamal run -n drone-dingtalk:latest

# sync docker image and rename that image name or tag
hamal run -n drone-dingtalk:latest -r drone-test:v1.0
```
