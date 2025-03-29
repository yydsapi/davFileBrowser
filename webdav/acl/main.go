package main

import (
	acl "github.com/hectane/go-acl"
	"golang.org/x/sys/windows"
)

func main() {
	if err := acl.Apply(
		"E:\\ex",
		false,
		false,
		acl.GrantName(windows.GENERIC_ALL, "xinxin"),
		//acl.DenyName(windows.GENERIC_WRITE, "Bob"),
	); err != nil {
		//fmt.println(err)
	}
}
