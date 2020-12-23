# Shell Game Launcher

Shell Game Launcher is a modern dgamelaunch alternative. It is a network based shell where anyone can sign up for an account and start playing any game hosted there. Nethack is the main target audience, but any shell game should be able to be hosted this way.

## Content

- [Dependencies](#dependencies)
- [Building](#building)
- [Usage](#usage)

## Dependencies

go is required. Only go version >= 1.15.6 on linux amd64 (Gentoo and Ubuntu 20.04) and on OpenBSD amd64 has been tested.

## Building

To run tests, use :
```
go test -cover ./...
```

For a debug build, use :
```
go build
```

For a release build, use :
```
go build -ldflags="-s -w"
```

## Usage

TODO
