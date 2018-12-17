# staticserver
static file server in Golang.

[![Build Status](https://travis-ci.org/schoeu/webhooks.svg?branch=master)](https://travis-ci.org/schoeu/staticserver)
[![Go Report Card](https://goreportcard.com/badge/github.com/schoeu/webhooks)](https://goreportcard.com/report/github.com/schoeu/staticserver)
[![GoDoc](https://godoc.org/github.com/schoeu/webhooks?status.svg)](https://godoc.org/github.com/schoeu/staticserver)


## Getting started

### Download

Choose the version of your computer system and download it, then copy it to the server which you want to control.

- [linux-32bit](http://qiniucdn.schoeu.com/staticserver_32bit)
- [linux-64bit](http://qiniucdn.schoeu.com/staticserver_64bit)
- [MAC](http://qiniucdn.schoeu.com/staticserver_mac)
- [windows-32bit](http://qiniucdn.schoeu.com/staticserver_32bit.exe)
- [windows-64bit](http://qiniucdn.schoeu.com/staticserver_64bit.exe)

## Help

Run `./staticserver --help` command you can get full information for use the staticserver.

```
./staticserver. --help

Output:

Usage of ./staticserver_mac:
  -help string
        help info for staticserver
  -path string
        Path to files direction.
  -port string
        Static server port. (default "8910")
  -prefix string
        Prefix of url path. (default "/static/")

```



## Example
```
./staticserver --path /Users/memee/Downloads/svn/
```
now require url `http://localhost:8910/static`

```
./staticserver --path /Users/memee/Downloads/svn  --prefix tmpfiles
```
now require url `http://localhost:8910/tmpfiles`

```
./staticserver --path ../somepath   --port 8888
```
now require url `http://localhost:8888/static`

then it can show you the files in this path.


## MIT License

Copyright (c) 2018 Schoeu

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
