[![Build Status](https://github.com/tamada/uniq2/workflows/build/badge.svg?branch=master)](https://github.com/tamada/uniq2/actions?workflow=build)
[![Coverage Status](https://coveralls.io/repos/github/tamada/uniq2/badge.svg?branch=master)](https://coveralls.io/github/tamada/uniq2?branch=master)
[![codebeat badge](https://codebeat.co/badges/855266ea-99d4-4d80-ac43-81a1712f0f90)](https://codebeat.co/projects/github-com-tamada-uniq2-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/tamada/uniq2)](https://goreportcard.com/report/github.com/tamada/uniq2)
[![License](https://img.shields.io/badge/License-WTFPL-blue.svg)](https://github.com/tamada/uniq2/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-1.1.1-yellowgreen.svg)](https://github.com/tamada/uniq2/releases/tag/v1.1.1)

# uniq2

Delete duplicated lines.

[GNU core utilities](https://www.gnu.org/software/coreutils/) has `uniq` command for deleting duplicate lines.
However, `uniq` command deletes only continuous duplicate lines.
When deleting not continuous duplicate lines, we use `sort` command together, in that case, the order of the list was not kept.

We want to delete not continuous duplicated lines with remaining the order.

## Install

### Install by Homebrew

Simply type the following commands.

```sh
$ brew tap tamada/brew # only the first time.
$ brew install uniq2
```

### Install by Go

Simply type the following command.

```sh
$ go get github.com/tamada/uniq2
```

## Usage

```sh
uniq2 [OPTIONS] [INPUT [OUTPUT]]
OPTIONS
    -a, --adjacent        delete only adjacent duplicated lines.
    -d, --delete-lines    only prints deleted lines.
    -i, --ignore-case     case sensitive.
    -h, --help            print this message.

INPUT                     gives file name of input.  If argument is single dash ('-')
                          or absent, the program read strings from stdin.
OUTPUT                    represents the destination.
```

## :whale: Docker

```sh
docker run --rm -v $PWD:/home/uniq2 tamada/uniq2:latest [OPTIONS] [ARGUMENTS...]
```

The meaning of the options of above command are as follows.

* `--rm`
    * remove container after running Docker.
* `-v $PWD:/home/uniq2`
    * share volumen `$PWD` in the host OS to `/home/uniq2` in the container OS.
    * Note that `$PWD` must be the absolute path.

[![Docker](https://img.shields.io/badge/docker-tamada%2Funiq2%3Alatest-blue?logo=docker&style=social)](https://hub.docker.com/r/tamada/uniq2)

## License

[WTFPL](https://github.com/tamada/uniq2/blob/master/LICENSE)
