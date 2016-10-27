/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package salt

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
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
	AuthToken     AuthToken
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

// RestClientTokenWrapper is a wrapper to authenticate if received 401 with a token.
// this usually due to server side loss the token when restart.
func (c *ClientImpl) RestClientTokenWrapper(callback func(rest *gorest.RestClient) (code int, err error)) error {
	if int64(c.AuthToken.Expire) < time.Now().Unix() {
		if err := c.Authenticate(); err != nil {
			return err
		}
	}
	code, err := callback(c.RestClient().Set("X-Auth-Token", c.AuthToken.Token))
	if code == 401 {
		if err := c.Authenticate(); err != nil {
			return err
		}
		code, err = callback(c.RestClient().Set("X-Auth-Token", c.AuthToken.Token))
	}
	if err != nil {
		return err
	}
	return nil
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

// ReturnResponse ...
type ReturnResponse struct {
	Return []interface{} `json:"return"`
}

func (r *ReturnResponse) parse(value interface{}) error {
	if r == nil {
		return errors.New("nil pointer return response")
	}
	if len(r.Return) != 1 {
		return fmt.Errorf("return array len is %d, this is expected to be 1", len(r.Return))
	}
	ret := r.Return[0]

	// string means error
	v := reflect.ValueOf(ret)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.String {
		return errors.New(v.String())
	}

	// marshal interface
	body, err := json.Marshal(ret)
	if err != nil {
		return err
	}

	// if want string do not unmarshal
	if val := reflect.ValueOf(value); val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.String {
		val.Elem().SetString(string(body))
		return nil
	}

	// unmarshal
	d := json.NewDecoder(strings.NewReader(string(body)))
	d.UseNumber()
	return d.Decode(value)
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
	response := ReturnResponse{}
	err := c.RestClientWithPassWord().Post("/login").Receive(&response)

	if err != nil {
		return err
	}
	return response.parse(&c.AuthToken)
}

// RunCmd ...
func (c *ClientImpl) RunCmd(target, fun string, arg []string, result interface{}) (err error) {
	response := ReturnResponse{}
	err = c.RestClientTokenWrapper(func(rest *gorest.RestClient) (code int, err error) {
		return code, rest.
			ParamStruct(struct {
				Client string   `json:"client,omitempty"`
				Fun    string   `json:"fun,omitempty"`
				Arg    []string `json:"arg,omitempty"`
				Target string   `json:"tgt,omitempty"`
			}{
				Client: "local",
				Target: target,
				Fun:    fun,
				Arg:    arg,
			}).Post("/run").
			Receive(&response, &code)
	})
	if err != nil {
		return err
	}
	return response.parse(result)
}
