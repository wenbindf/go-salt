package cmd

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

func TestRun(t *testing.T) {
	r, err := client.Run("whoami", &Param{Runas: "ubuntu"})
	if err != nil {
		t.Fatal(err)
	}
	if r[target] != "ubuntu" {
		t.Fatalf("responese expect stand output %s, get %s", "ubuntu", r[target])
	}
	fmt.Print(r)
}
