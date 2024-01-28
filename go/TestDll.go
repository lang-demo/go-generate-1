package main

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output generated.TestDll.go TestDll.go

//sys Greeting(s *uint16) (ret *uint16) = TestDll.greeting
//sys Greeting8(s *byte) (ret *byte) = TestDll.greeting8
