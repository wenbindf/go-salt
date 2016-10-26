/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package salt

import (
	"errors"
	"fmt"
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

// JobsResponse ...
type JobsResponse struct {
	Jobs []map[string]*Job `json:"return"`
}

// Jobs ...
func (c *ClientImpl) Jobs() (jobs map[string]*Job, err error) {
	response := JobsResponse{}
	err = c.RestClientWithToken().Get("/jobs").Receive(&response)
	if err != nil {
		return nil, err
	}

	if len(response.Jobs) == 0 {
		return nil, errors.New("response jobs array length is 0, this should never happen")
	}
	jobs = response.Jobs[0]
	for k, v := range jobs {
		v.ID = k
	}

	return jobs, nil
}

// JobResponse ...
type JobResponse struct {
	Job []*Job `json:"return"`
}

// Job ...
func (c *ClientImpl) Job(id string) (*Job, error) {
	response := JobResponse{}
	err := c.RestClientWithToken().Get(fmt.Sprintf("/jobs/%s", id)).Receive(&response)
	if err != nil {
		return nil, err
	}
	if len(response.Job) == 0 {
		return nil, errors.New("response job array length is 0, this should never happen")
	}
	return response.Job[0], nil
}

// ExecutionResponse ...
type ExecutionResponse struct {
	Job []*Job `json:"return"`
}

// Execute ...
func (c *ClientImpl) Execute(target, fun string, arg []string) (id string, er error) {
	response := &ExecutionResponse{}
	err := c.RestClientWithToken().ParamStruct(struct {
		Fun    string   `json:"fun,omitempty"`
		Arg    []string `json:"arg,omitempty"`
		Target string   `json:"tgt,omitempty"`
	}{
		Target: target,
		Fun:    fun,
		Arg:    arg,
	}).Post("/minions").Receive(&response)

	if err != nil {
		return "", err
	}

	if len(response.Job) == 0 {
		return "", errors.New("response execute job array length is 0, this should never happen")
	}

	return response.Job[0].ID, nil
}

// Running ...
func (j *Job) Running() bool {
	if len(j.Minions) != len(j.Result) {
		return false
	}
	return true
}
