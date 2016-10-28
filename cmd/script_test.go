package cmd

import (
	"fmt"
	"testing"
)

func TestScript(t *testing.T) {
	s := "Hello World"
	r, err := client.Script("https://github.com/xuguruogu/go-salt/echo.sh", s, &Param{Runas: "ubuntu"})
	if err != nil {
		t.Fatal(err)
	}
	if r[target] != s {
		t.Fatalf("responese expect output %s, get %s", s, r[target])
	}
	fmt.Println(r)
}
