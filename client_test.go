package salt

import (
	"fmt"
	"testing"
	"time"
)

var (
	addr     = "192.168.88.101:8000"
	username = "salt"
	password = "salt"
)

func TestAuthenticate(t *testing.T) {
	client := NewClient(addr, username, password, true)
	err := client.Authenticate()
	if err != nil {
		t.Fatal(err)
	}
	start := int64(client.(*ClientImpl).AuthToken.Start)
	expire := int64(client.(*ClientImpl).AuthToken.Expire)
	fmt.Println("start at ", time.Unix(start, 0).String())
	fmt.Println("expire at ", time.Unix(expire, 0).String())
}

func TestCmd(t *testing.T) {
	client := NewClient(addr, username, password, true)
	var s string
	err := client.RunCmd("*", "test.ping", nil, &s)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(s)
}
