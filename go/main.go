package main

import (
	"github.com/lang-library/go-global"
	"github.com/lang-library/go-winlib"
)

func main() {
	global.Echo("main() started")
	global.Echo(callGreeting("トム"))
	global.Echo(callGreeting8("マリー"))
}

func callGreeting(name string) string {
	ptr1 := winlib.StringToUTF16Ptr(name)
	ptr2 := Greeting(ptr1)
	return winlib.UTF16PtrToString(ptr2)
}

func callGreeting8(name string) string {
	ptr1 := winlib.StringToUTF8Ptr(name)
	ptr2 := Greeting8(ptr1)
	return winlib.UTF8PtrToString(ptr2)
}
