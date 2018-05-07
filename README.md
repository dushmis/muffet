# Muffet

[![Circle CI](https://img.shields.io/circleci/project/github/raviqqe/muffet/master.svg?style=flat-square)](https://circleci.com/gh/raviqqe/muffet)
[![Codecov](https://img.shields.io/codecov/c/github/raviqqe/muffet.svg?style=flat-square)](https://codecov.io/gh/raviqqe/muffet)
[![Go Report Card](https://goreportcard.com/badge/github.com/raviqqe/muffet?style=flat-square)](https://goreportcard.com/report/github.com/raviqqe/muffet)
[![Docker Build Status](https://img.shields.io/docker/build/jrottenberg/ffmpeg.svg?style=flat-square)](https://hub.docker.com/r/raviqqe/muffet)
[![License](https://img.shields.io/github/license/raviqqe/muffet.svg?style=flat-square)](LICENSE)

[![asciicast](https://asciinema.org/a/178887.png)](https://asciinema.org/a/178887)

Muffet is a website link checker which scrapes and inspects all pages in a
website recursively.

## Features

- Massive speed
- Colored outputs
- Different tags support (`a`, `img`, `link`, `script`, etc)

## Installation

```
go get -u github.com/raviqqe/muffet
```

## Usage

```
muffet https://shady.bakery.hotland
```

For more information, see `muffet --help`.

## License

[MIT](LICENSE)
