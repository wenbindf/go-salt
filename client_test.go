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

func TestPing(t *testing.T) {
	client := NewClient(addr, username, password, true)
	ping := map[string]bool{}
	err := client.RunCmd("*", "test.ping", nil, &ping)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(ping)
}

func TestEcho(t *testing.T) {
	client := NewClient(addr, username, password, true)
	echo := map[string]string{}
	err := client.RunCmd("*", "test.echo", []string{"hello world"}, &echo)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(echo)
}
