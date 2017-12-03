# Travelex Pre Interview Application #
## About ##
This repository contains a simple microservice application written in [go](http://golang.com).

## Requirements ##
You'll need docker and a git client to run this application.

## Running ##
```bash
git clone git@github.com:pennywisdom-other/travelex.git
cd travelex
docker build -t travelex-test .
docker run -it -p 8080:8080 travelex-test
```