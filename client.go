/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package salt

import (
	"crypto/tls"
	"errors"
	"net/http"
	"time"

	"github.com/xuguruogu/gorest"
)

// Client ...
type Client interface {
	Authenticate() error
	RunCmd(target, fun string, arg []string, result interface{}) error
	Jobs() (map[string]*Job, error)
	Job(id string) (*Job, error)
	Execute(target, fun string, arg []string) (id string, er error)
	Minions() (map[string]*Minion, error)
	Minion(id string) (*Minion, error)
}

// ClientImpl ...
type ClientImpl struct {
	Addr          string
	Username      string
	Password      string
	Eauth         string
	AuthToken     *AuthToken
	SSLSkipVerify bool
}

// NewClient ...
func NewClient(addr, username, password string, SSLSkipVerify bool, eauth ...string) Client {
	return &ClientImpl{
		Addr:     addr,
		Username: username,
		Password: password,
		Eauth: func() string {
			if len(eauth) == 0 {
				return "pam"
			}
			return eauth[0]
		}(),
		SSLSkipVerify: SSLSkipVerify,
	}
}

// RestClient ...
func (c *ClientImpl) RestClient() *gorest.RestClient {
	return gorest.New().Base("https://" + c.Addr).Client(&http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: c.SSLSkipVerify,
		},
	}})
}

// RestClientWithToken ...
func (c *ClientImpl) RestClientWithToken() *gorest.RestClient {
	if c.AuthToken == nil || int64(c.AuthToken.Expire) < time.Now().Unix() {
		c.Authenticate()
	}

	if c.AuthToken == nil {
		return c.RestClient()
	}

	return c.RestClient().Set("X-Auth-Token", c.AuthToken.Token)
}

// RestClientWithPassWord ...
func (c *ClientImpl) RestClientWithPassWord() *gorest.RestClient {
	return c.RestClient().ParamStruct(struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
		Eauth    string `json:"eauth,omitempty"`
	}{
		Username: c.Username,
		Password: c.Password,
		Eauth:    c.Eauth,
	})
}

// AuthResponse ...
type AuthResponse struct {
	Tokens []*AuthToken `json:"return"`
}

// AuthToken ...
type AuthToken struct {
	Token  string   `json:"token"`
	Expire float32  `json:"expire"`
	Start  float32  `json:"start"`
	User   string   `json:"user"`
	Eauth  string   `json:"eauth"`
	Perms  []string `json:"perms"`
}

// Authenticate ...
func (c *ClientImpl) Authenticate() error {
	response := AuthResponse{}
	err := c.RestClientWithPassWord().Post("/login").Receive(&response)

	if err != nil {
		return err
	}

	if len(response.Tokens) == 0 {
		return errors.New("response token array length is 0, this should never happen")
	}
	c.AuthToken = response.Tokens[0]
	return nil
}

// RunCmd ...
func (c *ClientImpl) RunCmd(target, fun string, arg []string, result interface{}) error {
	return c.RestClientWithPassWord().ParamStruct(struct {
		Client string   `json:"client,omitempty"`
		Fun    string   `json:"fun,omitempty"`
		Arg    []string `json:"arg,omitempty"`
		Target string   `json:"tgt,omitempty"`
	}{
		Client: "local",
		Target: target,
		Fun:    fun,
		Arg:    arg,
	}).Post("/run").Receive(result)
}
