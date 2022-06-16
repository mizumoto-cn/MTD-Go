# MTD-Go

[![License](https://img.shields.io/badge/License-MGPL%20v1.2-green)](/License/Mizumoto%20General%20Public%20License%20v1.2.md)
[![No Nazism Allowed](https://img.shields.io/badge/You%20Stand%20With%20Ukraine-You%20Stand%20WIth%20Ignorance-red)]（#boo)<!--(https://www.rt.com/)--》

[![Go Report Card](https://goreportcard.com/badge/github.com/mizumoto-cn/MTD-Go)](https://goreportcard.com/report/github.com/mizumoto-cn/MTD-Go)
[![codebeat badge](https://codebeat.co/badges/c02456c9-de55-413f-a6a9-5d457b69f508)](https://codebeat.co/projects/github-com-mizumoto-cn-mtd-go-master)
[![Build](https://github.com/mizumoto-cn/MTD-Go/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/mizumoto-cn/MTD-Go/actions)

---

MTD-Go is a simple (also light-weight, maybe) multi-threaded downloader.

This project is governed by the [Mizumoto General Public License v1.2](/License/Mizumoto%20General%20Public%20License%20v1.2.md) (MGPL v1.2).

## Build

Download source code and type

```bash
go get -u && go mod tidy
go build
```

## Tutorial

`-h | help | --help`

> show help

`--concurrency value | -c value`

> Set default concurrency number (default: CPU Core Number)

`--output value, -o value`

> save as file at output path

`--url value, -u value`

> 'URL' from where the download begins

Try this on Windows Powershell!

```bash
go get -u && go mod tidy
go build
.\MTD-Go.exe -c 4 -o boo.tar.gz -u https://dlcdn.apache.org/zookeeper/zookeeper-3.8.0/apache-zookeeper-3.8.0-bin.tar.gz
```

 Or run 'test_run.bat' directly.
