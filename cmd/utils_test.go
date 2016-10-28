package cmd

import (
	"fmt"
	"testing"
)

func TestTty(t *testing.T) {
	s, err := client.Tty("pts1", "This is a test")
	if err != nil {
		t.Fatal(err)
	}
	if len(s[target]) == 0 {
		t.Fatal("responese len is 0")
	}
	fmt.Println(s)
}

func TestWhich(t *testing.T) {
	s, err := client.Which("cat")
	if err != nil {
		t.Fatal(err)
	}
	if len(s[target]) == 0 {
		t.Fatal("responese len is 0")
	}
	fmt.Println(s)
}

func TestHasExec(t *testing.T) {
	s, err := client.HasExec("cat")
	if err != nil {
		t.Fatal(err)
	}
	if s[target] == false {
		t.Fatal("responese is false")
	}
	fmt.Println(s)
}

func TestShells(t *testing.T) {
	s, err := client.Shells()
	if err != nil {
		t.Fatal(err)
	}
	if len(s[target]) == 0 {
		t.Fatal("responese len is 0")
	}
	fmt.Println(s)
}
