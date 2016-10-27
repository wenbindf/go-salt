package test

import salt "github.com/xuguruogu/go-salt"

// Test ...
type Test interface {
	Ping(target string) (ping map[string]bool, err error)
	Echo(target string, param string) (echo map[string]string, err error)
}

type TestImpl struct {
	client salt.Client
}

// New ...
func New(client salt.Client) Test {
	return &TestImpl{client: client}
}

// Ping ...
func (ti *TestImpl) Ping(target string) (ping map[string]bool, err error) {
	ping = map[string]bool{}
	return ping, ti.client.RunCmd(target, "test.ping", nil, &ping)
}

// Echo ...
func (ti *TestImpl) Echo(target string, param string) (echo map[string]string, err error) {
	echo = map[string]string{}
	return echo, ti.client.RunCmd(target, "test.echo", []string{param}, &echo)
}
