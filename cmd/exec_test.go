package cmd

import (
	"fmt"
	"testing"
)

func TestExecCode(t *testing.T) {
	r, err := client.ExecCode("python", "print 1+2")
	if err != nil {
		t.Fatal(err)
	}
	if r[target] != "3" {
		t.Fatalf("responese expect %s, get %s", "3", r[target])
	}
	fmt.Print(r)
}

func TestExecCodeAll(t *testing.T) {
	r, err := client.ExecCodeAll("python", "print 1+2")
	if err != nil {
		t.Fatal(err)
	}
	if r[target].Stdout != "3" {
		t.Fatalf("responese expect stand output %s, get %s", "3", r[target])
	}
	fmt.Print(r)
}
