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
	ID         string   `json:"jid"`
	Function   string   `json:"Function"`
	Target     string   `json:"Target"`
	User       string   `json:"User"`
	StartTime  string   `json:"StartTime"`
	TargetType string   `json:"Target-Type"`
	Arguments  []string `json:"Arguments"`
	Minions    []string `json:"Minions"`
	Result     map[string]struct {
		Return interface{} `json:"return"`
	} `json:"Result"`
}

// Jobs ...
func (c *ClientImpl) Jobs() (jobs map[string]*Job, err error) {
	response := ReturnResponse{}
	err = c.RestClientTokenWrapper(func(rest *gorest.RestClient) (code int, err error) {
		return code, rest.Get("/jobs").Receive(&response, &code)
	})
	if err != nil {
		return nil, err
	}

	jobs = map[string]*Job{}
	err = response.parse(&jobs)
	if err != nil {
		return nil, err
	}
	for k, v := range jobs {
		v.ID = k
	}

	return jobs, nil
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
	return job, response.parse(job)
}

// Execute ...
func (c *ClientImpl) Execute(target, fun string, arg []string) (id string, err error) {
	response := ReturnResponse{}

	err = c.RestClientTokenWrapper(func(rest *gorest.RestClient) (code int, err error) {
		return code, rest.ParamStruct(struct {
			Fun    string   `json:"fun,omitempty"`
			Arg    []string `json:"arg,omitempty"`
			Target string   `json:"tgt,omitempty"`
		}{
			Target: target,
			Fun:    fun,
			Arg:    arg,
		}).
			Post("/minions").
			Receive(&response, &code)
	})
	if err != nil {
		return "", err
	}
	job := Job{}
	return job.ID, response.parse(&job)
}

// Running ...
func (j *Job) Running() bool {
	if len(j.Minions) != len(j.Result) {
		return false
	}
	return true
}
