/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package salt

import (
	"fmt"

	"github.com/xuguruogu/gorest"
)

// Job ...
type Job struct {
	ID         string                 `json:"jid"`
	Function   string                 `json:"Function"`
	Target     string                 `json:"Target"`
	User       string                 `json:"User"`
	StartTime  string                 `json:"StartTime"`
	TargetType string                 `json:"Target-Type"`
	Arguments  []string               `json:"Arguments"`
	Minions    []string               `json:"Minions"`
	Result     map[string]interface{} `json:"Result"`
}

// Job ...
func (c *ClientImpl) Job(id string) (job *Job, err error) {
	response := ReturnResponse{}
	err = c.RestClientTokenWrapper(func(rest *gorest.RestClient) (code int, err error) {
		return code, rest.
			Get(fmt.Sprintf("/jobs/%s", id)).
			Receive(&response, &code)
	})
	if err != nil {
		return nil, err
	}
	job = &Job{}
	return job, response.parse("info", job)
}

// Execute ...
type Execute struct {
	ID      string   `json:"jid"`
	Minions []string `json:"minions"`
}

// Execute ...
func (c *ClientImpl) Execute(target, fun string, arg interface{}) (execute *Execute, err error) {
	response := ReturnResponse{}
	err = c.RestClientTokenWrapper(func(rest *gorest.RestClient) (code int, err error) {
		return code, rest.ParamStruct(struct {
			Fun    string      `json:"fun,omitempty"`
			Arg    interface{} `json:"arg,omitempty"`
			Target string      `json:"tgt,omitempty"`
		}{
			Target: target,
			Fun:    fun,
			Arg:    arg,
		}).
			Post("/minions").
			Receive(&response, &code)
	})
	if err != nil {
		return nil, err
	}
	execute = &Execute{}
	return execute, response.parse("return", execute)
}
