package test

import salt "github.com/xuguruogu/go-salt"

// Test ...
type Test interface {
	Ping() (ping map[string]bool, err error)
	Echo(param string) (echo map[string]string, err error)
}

// Impl ...
type Impl struct {
	client salt.Client
	target string
}

// New ...
func New(target string, client salt.Client) Test {
	return &Impl{client: client, target: target}
}

// Ping ...
func (ti *Impl) Ping() (ping map[string]bool, err error) {
	ping = map[string]bool{}
	return ping, ti.client.RunCmd(ti.target, "test.ping", nil, nil, &ping)
}

// Echo ...
func (ti *Impl) Echo(param string) (echo map[string]string, err error) {
	echo = map[string]string{}
	return echo, ti.client.RunCmd(ti.target, "test.echo", []interface{}{param}, nil, &echo)
}
