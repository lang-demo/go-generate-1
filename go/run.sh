#! /usr/bin/env bash
set -uvx
set -e
cwd=`pwd`
cd $cwd/../cpp
g++ -shared -o $cwd/TestDll.dll TestDll.cpp -static
cd $cwd
pexports ./TestDll.dll
./gen.sh
go run .
