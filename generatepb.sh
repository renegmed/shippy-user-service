#!/bin/bash
 
protoc -I .  proto/auth/auth.proto \
   --go_out=plugins=micro:${GOPATH}/src/github.com/renegmed/shippy-user-service