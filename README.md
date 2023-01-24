[![go-ubuntu](https://github.com/bridgewwater/drone-plugin-temple/workflows/go-ubuntu/badge.svg?branch=main)](https://github.com/bridgewwater/drone-plugin-temple/actions)
[![GoDoc](https://godoc.org/github.com/bridgewwater/drone-plugin-temple?status.png)](https://godoc.org/github.com/bridgewwater/drone-plugin-temple/)
[![GoReportCard](https://goreportcard.com/badge/github.com/bridgewwater/drone-plugin-temple)](https://goreportcard.com/report/github.com/bridgewwater/drone-plugin-temple)
[![codecov](https://codecov.io/gh/bridgewwater/drone-plugin-temple/branch/main/graph/badge.svg)](https://codecov.io/gh/bridgewwater/drone-plugin-temple)

## for what

- this project used to drone CI

## Pipeline Settings (.drone.yml)

`1.x`

```yaml
steps:
  - name: notification
    image: bridgewwater/drone-plugin-temple:latest
    pull: if-not-exists
    settings:
      debug: false
      webhook:
        # https://docs.drone.io/pipeline/environment/syntax/#from-secrets
        from_secret: webhook_token
      msg_type: your-message-type
      timeout_second: 10 # default 10
    when:
      event: # https://docs.drone.io/pipeline/exec/syntax/conditions/#by-event
        - promote
        - rollback
        - push
        - pull_request
        - tag
      status: # only support failure/success,  both open will send anything
        - failure
        # - success
```

# dev

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

```bash
make init dep
```

- test code

add env then test

```bash
export PLUGIN_MSG_TYPE=post \
  export PLUGIN_WEBHOOK=7138d7b3-abc
```

```bash
make test
```

- see help

```bash
make dev
```

update main.go file set env then and run

```bash
export PLUGIN_MSG_TYPE= \
  export PLUGIN_WEBHOOK= \
  export DRONE_REPO=bridgewwater/drone-plugin-temple \
  export DRONE_REPO_NAME=drone-plugin-temple \
  export DRONE_REPO_NAMESPACE=bridgewwater \
  export DRONE_REMOTE_URL=https://github.com/bridgewwater/drone-plugin-temple \
  export DRONE_REPO_OWNER=bridgewwater \
  export DRONE_COMMIT_AUTHOR=bridgewwater \
  export DRONE_COMMIT_AUTHOR_AVATAR=  \
  export DRONE_COMMIT_AUTHOR_EMAIL=bridgewwatergmppt@gmail.com \
  export DRONE_COMMIT_BRANCH=main \
  export DRONE_COMMIT_LINK=https://github.com/bridgewwater/drone-plugin-temple/commit/68e3d62dd69f06077a243a1db1460109377add64 \
  export DRONE_COMMIT_SHA=68e3d62dd69f06077a243a1db1460109377add64 \
  export DRONE_COMMIT_REF=refs/heads/main \
  export DRONE_COMMIT_MESSAGE="mock message commit" \
  export DRONE_STAGE_STARTED=1674531206 \
  export DRONE_STAGE_FINISHED=1674532106 \
  export DRONE_BUILD_STATUS=success \
  export DRONE_BUILD_NUMBER=1 \
  export DRONE_BUILD_LINK=https://drone.xxx.com/bridgewwater/drone-plugin-temple/1 \
  export DRONE_BUILD_EVENT=push \
  export DRONE_BUILD_STARTED=1674531206 \
  export DRONE_BUILD_FINISHED=1674532206
```

- then run

```bash
make run
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# if run error
# like this error
# err: missing webhook, please set webhook
#  fix env settings then test

# see run docker fast
$ make dockerTestRunLatest

# clean test build
$ make dockerTestPruneLatest

# see how to use
$ docker run --rm bridgewwater/drone-plugin-temple:latest -h
```
