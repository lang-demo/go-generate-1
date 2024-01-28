#! /usr/bin/env bash
set -uvx
set -e
go generate TestDll.go
sed -i -e "s/NewLazySystemDLL/NewLazyDLL/g" generated.TestDll.go
