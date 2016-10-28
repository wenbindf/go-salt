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

var (
	target = "minion"
	client = New(target, salt.NewClient(addr, username, password, true))
)

func TestPing(t *testing.T) {
	ping, err := client.Ping()
	if err != nil {
		t.Fatal(err)
	}
	if ping[target] != true {
		t.Errorf("ping %s failed", target)
	}
	fmt.Print(ping)
}

func TestEcho(t *testing.T) {
	s := "Hello World"
	echo, err := client.Echo(s)
	if err != nil {
		t.Fatal(err)
	}
	if echo[target] != s {
		t.Errorf("echo failed, expect %s, get %s", s, echo[target])
	}
	fmt.Print(echo)
}
