package cmd

import (
	"fmt"
	"testing"
)

func TestScript(t *testing.T) {
	s := "Hello World"
	r, err := client.Script("https://raw.githubusercontent.com/xuguruogu/go-salt/master/echo.sh", s, &Param{Runas: "ubuntu"})
	if err != nil {
		t.Fatal(err)
	}
	if r[target].Stdout != s {
		t.Fatalf("responese expect stand output %s, get %s", s, r[target])
	}
	fmt.Println(r)
}

func TestScriptRetcode(t *testing.T) {
	s := "Hello World"
	r, err := client.ScriptRetcode("https://raw.githubusercontent.com/xuguruogu/go-salt/master/echo.sh", s, &Param{Runas: "ubuntu"})
	if err != nil {
		t.Fatal(err)
	}
	if r[target] != 0 {
		t.Fatalf("responese expect ret code %d, get %d", 0, r[target])
	}
	fmt.Println(r)
}
