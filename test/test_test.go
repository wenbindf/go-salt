package test

import (
	"fmt"
	"testing"

	salt "github.com/xuguruogu/go-salt"
)

var (
	addr     = "192.168.88.101:8000"
	username = "salt"
	password = "salt"
)

func TestPing(t *testing.T) {
	test := New(salt.NewClient(addr, username, password, true))
	ping, err := test.Ping("*")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(ping)
}

func TestEcho(t *testing.T) {
	test := New(salt.NewClient(addr, username, password, true))
	echo, err := test.Echo("*", "Hello World")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(echo)
}
