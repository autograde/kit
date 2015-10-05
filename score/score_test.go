package score

import (
	"encoding/json"
	"testing"
)

func TestScoreStrings(t *testing.T) {
	scoreLine := `{"Secret":"92bc4fd5549d8e4f14ab4e96f7a0be02","TestName":"Assignment1","Score":0.8,"MaxScore":1.0,"Weight":1.0}`
	var score Score
	err := json.Unmarshal([]byte(scoreLine), &score)
	if err != nil && err.Error() != "json: cannot unmarshal number 0.8 into Go value of type int" {
		t.Error("Unmarshal returned unexpected string:", err)
	}
	scoreLine = `{"Secret":"92bc4fd5549d8e4f14ab4e96f7a0be02","TestName":"Assignment1","Score":8,"MaxScore":1.0,"Weight":1.0}`
	err = json.Unmarshal([]byte(scoreLine), &score)
	if err != nil && err.Error() != "json: cannot unmarshal number 1.0 into Go value of type int" {
		t.Error("Unmarshal returned unexpected string:", err)
	}
	scoreLine = `{"Secret":"92bc4fd5549d8e4f14ab4e96f7a0be02","TestName":"Assignment1","Score":8,"MaxScore":10,"Weight":100}`
	err = json.Unmarshal([]byte(scoreLine), &score)
	if err != nil {
		t.Error("Unmarshal returned unexpected string:", err)
	}
}
