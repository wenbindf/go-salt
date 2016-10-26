package salt

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMinion(t *testing.T) {
	client := NewClient(addr, username, password, true)
	minions, err := client.Minions()
	if err != nil {
		t.Fatal(err)
	}
	// b, _ := json.Marshal(minions)
	// fmt.Println(string(b))

	for _, minion := range minions {
		fmt.Println(minion.ID)
		m, err := client.Minion(minion.ID)
		if err != nil {
			t.Fatal(err)
		}
		b, _ := json.Marshal(m)
		fmt.Println(string(b))
	}
}
