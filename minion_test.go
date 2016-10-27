package salt

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMinion(t *testing.T) {
	minions, err := client.Minions()
	if err != nil {
		t.Fatal(err)
	}

	for _, minion := range minions {
		m, err := client.Minion(minion.ID)
		if err != nil {
			t.Fatal(err)
		}
		if m.ID != minion.ID {
			t.Fatalf("id not match %s, %s", minion.ID, m.ID)
		}
		b, _ := json.Marshal(m)
		fmt.Println(string(b))
	}
}
