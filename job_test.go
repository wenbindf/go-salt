package salt

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJobs(t *testing.T) {
	client := NewClient(addr, username, password, true)
	_, err := client.Jobs()
	if err != nil {
		t.Fatal(err)
	}
}

func TestExecute(t *testing.T) {
	client := NewClient(addr, username, password, true)

	id, err := client.Execute("*", "test.ping", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(id)

	// get job
	job, err := client.Job(id)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(job)
	fmt.Println(string(b))
}
