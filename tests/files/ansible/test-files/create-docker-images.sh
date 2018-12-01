#!/bin/bash -x

# creates the necessary docker images to run testrunner.sh locally

docker build --tag="teslafunds/cppjit-testrunner" docker-cppjit
docker build --tag="teslafunds/python-testrunner" docker-python
docker build --tag="teslafunds/go-testrunner" docker-go
