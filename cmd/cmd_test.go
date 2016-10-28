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
		t.Fatalf("responese expect output %s, get %s", "ubuntu", r[target])
	}
	fmt.Println(r)
}

func TestRetcode(t *testing.T) {
	r, err := client.Retcode("whoami", &Param{Runas: "ubuntu"})
	if err != nil {
		t.Fatal(err)
	}
	if r[target] != 0 {
		t.Fatalf("responese expect ret code %d, get %d", 0, r[target])
	}
	fmt.Println(r)
}

func TestRunStderr(t *testing.T) {
	r, err := client.RunStderr("whoami", &Param{Runas: "ubuntu"})
	if err != nil {
		t.Fatal(err)
	}
	if r[target] != "" {
		t.Fatalf("responese expect stand error %s, get %s", "", r[target])
	}
	fmt.Println(r)
}

func TestRunStdout(t *testing.T) {
	r, err := client.RunStdout("whoami", &Param{Runas: "ubuntu"})
	if err != nil {
		t.Fatal(err)
	}
	if r[target] != "ubuntu" {
		t.Fatalf("responese expect stand output %s, get %s", "ubuntu", r[target])
	}
	fmt.Println(r)
}

func TestRunAll(t *testing.T) {
	r, err := client.RunAll("whoami", &Param{Runas: "ubuntu"})
	if err != nil {
		t.Fatal(err)
	}
	if r[target].Stdout != "ubuntu" {
		t.Fatalf("responese expect stand output %s, get %s", "ubuntu", r[target])
	}
	fmt.Println(r)
}

func TestRunBg(t *testing.T) {
	r, err := client.RunBg("ls -a", &Param{Runas: "ubuntu"})
	if err != nil {
		t.Fatal(err)
	}
	if r[target]["pid"] == 0 {
		t.Fatalf("responese expect pid none 0, get %d", r[target]["pid"])
	}
	fmt.Println(r)
}

func TestRunChroot(t *testing.T) {
	r, err := client.RunChroot("/home", "pwd", &Param{Runas: "ubuntu"})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(r)
}

func TestPowershell(t *testing.T) {
	r, err := client.Powershell("whoami", &Param{Runas: "ubuntu"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
}

func TestShell(t *testing.T) {
	r, err := client.Shell("whoami", &Param{Runas: "ubuntu"})
	if err != nil {
		t.Fatal(err)
	}
	if r[target] != "ubuntu" {
		t.Fatalf("responese expect stand output %s, get %s", "ubuntu", r[target])
	}
	fmt.Println(r)
}
