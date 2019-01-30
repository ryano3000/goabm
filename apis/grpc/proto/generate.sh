#!/bin/bash
protoc ./agent.proto --go-out=plugins=grpc:.