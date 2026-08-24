#!/bin/sh
set -e
go mod init genhash >/dev/null 2>&1 || true
go get golang.org/x/crypto/bcrypt >/dev/null 2>&1
go run genhash.go -- "Admin@123"
