package salt

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {
	execute, err := client.Execute("*", "test.ping", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(execute)

	// get job
	job, err := client.Job(execute.ID)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(job)
	fmt.Println(string(b))
}
