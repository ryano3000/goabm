#!/bin/bash
protoc ./agent.proto --go_out=plugins=grpc:.