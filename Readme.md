# MTD-Go

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
