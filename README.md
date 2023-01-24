[![go-ubuntu](https://github.com/bridgewwater/drone-plugin-temple/workflows/go-ubuntu/badge.svg?branch=main)](https://github.com/bridgewwater/drone-plugin-temple/actions)
[![GoDoc](https://godoc.org/github.com/bridgewwater/drone-plugin-temple?status.png)](https://godoc.org/github.com/bridgewwater/drone-plugin-temple/)
[![GoReportCard](https://goreportcard.com/badge/github.com/bridgewwater/drone-plugin-temple)](https://goreportcard.com/report/github.com/bridgewwater/drone-plugin-temple)
[![codecov](https://codecov.io/gh/bridgewwater/drone-plugin-temple/branch/main/graph/badge.svg)](https://codecov.io/gh/bridgewwater/drone-plugin-temple)

## for what

- this project used to github golang

## depends

in go mod project

```bash
# warning use privte git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "http://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q http://github.com/bridgewwater/drone-plugin-temple.git

# test depends see full version
$ go list -v -m -versions github.com/bridgewwater/drone-plugin-temple
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -m -versions github.com/bridgewwater/drone-plugin-temple | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## evn

- golang sdk 1.17+

# dev

```bash
make init dep
```

- test code

```bash
make test
```

add main.go file and run

```bash
make run
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# see how to use
$ golang-project-temple-base -h
```
