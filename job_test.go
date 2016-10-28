package salt

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {
	execute, err := client.Execute("*", "test.ping", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(execute)

	// get job
	job, err := client.Job(execute.ID)
	if err != nil {
		t.Fatal(err)
	}
	if job.ID != execute.ID {
		t.Fatal("job.ID!=execute.ID", job.ID, execute.ID)
	}

	b, _ := json.Marshal(job)
	fmt.Println(string(b))
}
