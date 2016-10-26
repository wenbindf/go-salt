/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package salt

import (
	"errors"
	"fmt"
)

// MinionsResponse ...
type MinionsResponse struct {
	Minions []map[string]*Minion `json:"return"`
}

// Minion ...
type Minion struct {
	ID            string   `json:"id"`
	Name          string   `json:"nodename"`
	Host          string   `json:"host"`
	Domain        string   `json:"domain"`
	OS            string   `json:"os"`
	OSRelease     string   `json:"osrelease"`
	OSName        string   `json:"osfullname"`
	Kernel        string   `json:"kernel"`
	KernelRelease string   `json:"kernelrelease"`
	Shell         string   `json:"shell"`
	ARCH          string   `json:"osarch"`
	CPUS          int      `json:"num_cpus"`
	RAM           int      `json:"mem_total"`
	CPUModel      string   `json:"cpu_model"`
	CPUFlags      []string `json:"cpu_flags"`
	Virtual       string   `json:"virtual"`
	IPv4          []string `json:"ipv4"`
	IPv6          []string `json:"ipv6"`
	Path          string   `json:"path"`
	ServerID      int      `json:"server_id"`
}

// Minions ...
func (c *ClientImpl) Minions() (map[string]*Minion, error) {
	response := MinionsResponse{}
	err := c.RestClientWithToken().Get("/minions").Receive(&response)
	if err != nil {
		return nil, err
	}
	if len(response.Minions) == 0 {
		return nil, errors.New("response minions array length is 0, this should never happen")
	}
	return response.Minions[0], nil
}

// Minion ...
func (c *ClientImpl) Minion(id string) (*Minion, error) {
	response := MinionsResponse{}
	err := c.RestClientWithToken().
		Get(fmt.Sprintf("/minions/%s", id)).
		Receive(&response)
	if err != nil {
		return nil, err
	}
	if len(response.Minions) == 0 || response.Minions[0][id] == nil {
		return nil, errors.New("No minions matched the target. No command was sent, no jid was assigned.")
	}
	return response.Minions[0][id], nil
}
