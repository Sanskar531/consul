# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: BUSL-1.1

ARG GOLANG_VERSION=1.20.12
FROM golang:${GOLANG_VERSION}

WORKDIR /consul
