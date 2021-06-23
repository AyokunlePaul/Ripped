#!/bin/zsh

protoc auth/auth.proto --go_out=plugins=grpc:.