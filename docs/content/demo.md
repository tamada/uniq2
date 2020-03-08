---
title: Demo
---

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


## Demo

```sh
$ cat -n testdata/test1.txt
1	a1
2	a1  # <- is the duplicate of the previous line.
3	a2
4	a2  # <- is the duplicate of the previous line.
5	a3
6	a4
7	a1  # <- is the duplicate of the first line.
8	A1
$ uniq2 testdata/test1.txt
a1
a2
a3
a4
A1
$ uniq2 -a testdata/test1.txt  # same result as uniq command.
a1
a2
a3
a4
a1  # <- this line is not deleted.
A1
$ uniq2 -i testdata/test1.txt # ignore case
a1
a2
a3
a4
$ uniq2 -d testdata/test1.txt # print delete lines.
a1
a2
a1
```

## Delete duplicate entries in PATH

```sh
export PATH=$(echo $PATH | tr : '\n' | uniq2 | paste -s -d : -)
```

* `tr : '\n'` replaces `:` to `\n` of data from STDIN,
* `uniq2` deletes duplicate lines from the result of `tr`, and
* `paste -s -d : -` joins given strings with `:`.
