[![Build Status](https://travis-ci.org/tamada/uniq2.svg?branch=master)](https://travis-ci.org/tamada/uniq2)
[![Coverage Status](https://coveralls.io/repos/github/tamada/uniq2/badge.svg?branch=master)](https://coveralls.io/github/tamada/uniq2?branch=master)
[![codebeat badge](https://codebeat.co/badges/855266ea-99d4-4d80-ac43-81a1712f0f90)](https://codebeat.co/projects/github-com-tamada-uniq2-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/tamada/uniq2)](https://goreportcard.com/report/github.com/tamada/uniq2)
[![License](https://img.shields.io/badge/License-WTFPL-blue.svg)](https://github.com/tamada/uniq2/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-0.1.0-yellowgreen.svg)](https://github.com/tamada/uniq2/releases/tag/v0.1.0)

# uniq2

Delete duplicated lines.
GNU core utilities have `uniq` command for deleting duplicate lines.
However, `uniq` command deletes only continuous duplicate lines.
When deleting not continuous duplicate lines, we use `sort` command together, in that case, the order of the list was not kept.

We want to delete not continuous duplicated lines with remaining the order.

## License

[WTFPL](https://github.com/tamada/uniq2/blob/master/LICENSE)
