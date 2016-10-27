package salt

import (
	"fmt"
	"testing"
	"time"
)

var client = NewClient(addr, username, password, true)

var (
	addr     = "192.168.88.101:8000"
	username = "salt"
	password = "salt"
)

// var (
// 	addr     = "10.94.99.95:8000"
// 	username = "hyxc"
// 	password = "123456aa"
// )

func TestAuthenticate(t *testing.T) {
	err := client.Authenticate()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(client.(*ClientImpl).AuthToken.Token)
	start := int64(client.(*ClientImpl).AuthToken.Start)
	expire := int64(client.(*ClientImpl).AuthToken.Expire)
	fmt.Println("start at ", time.Unix(start, 0).String())
	fmt.Println("expire at ", time.Unix(expire, 0).String())
}

func TestEcho(t *testing.T) {
	echo := map[string]string{}
	err := client.RunCmd("*", "test.echo", []string{"hello world"}, &echo)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(echo)
}

func TestPing(t *testing.T) {
	ping := map[string]bool{}
	err := client.RunCmd("*", "test.ping", nil, &ping)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(ping)
}

func TestArg(t *testing.T) {
	arg := map[string]interface{}{}
	err := client.RunCmd("*", "test.arg", []string{"hello world", "shell='dd'", "shell3='dd'"}, &arg)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(arg)
}
