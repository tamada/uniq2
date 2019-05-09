---
title: Utilities
---

## Demo

```
$ cat -n file
line1
line1
line3
line1
$ uniq file
line1
line3
line1
$ uniq2 file
line1
line3
```

## Delete duplicate entries in PATH

```
export PATH=$(echo $PATH | tr : '\n' | uniq2 | paste -s -d : -)
```

