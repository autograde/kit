package score

import "testing"

var theSecret = "my secret code"

func init() {
	GlobalSecret = theSecret
}

var nonJSONLog = []string{
	"heere is some output",
	"some other output",
	"line contains " + theSecret,
	theSecret + " should not be revealed",
}

func TestParseNonJSONStrings(t *testing.T) {
	for _, s := range nonJSONLog {
		sc, err := Parse(s, GlobalSecret)
		if err == nil {
			t.Errorf("Expected '%v', got '<nil>'", ErrScoreNotFound.Error())
		}
		if sc != nil {
			t.Errorf("Got unexpected score object '%v', wanted '<nil>'", sc)
		}
	}
}

var jsonLog = []struct {
	in  string
	out *Score
	err error
}{
	{`{"Secret":"` + theSecret + `","TestName":"init","Score":0,"MaxScore":10,"Weight":10}`,
		NewScore(10, 10),
		nil,
	},
	{`{"Secret":"the wrong secret","TestName":"init","Score":0,"MaxScore":10,"Weight":10}`,
		nil,
		ErrScoreNotFound,
	},
}

// Equal returns true if sc equals other. Ignores the Secret field.
func (sc *Score) Equal(other *Score) bool {
	return other != nil &&
		sc.TestName == other.TestName &&
		sc.Score == other.Score &&
		sc.MaxScore == other.MaxScore &&
		sc.Weight == other.Weight
}

func TestParseJSONStrings(t *testing.T) {
	for _, s := range jsonLog {
		sc, err := Parse(s.in, GlobalSecret)
		if sc != s.out || err != s.err {
			if !s.out.Equal(sc) || err != s.err {
				t.Errorf("Failed to parse:\n%v\nGot: '%v', '%v'\nExp: '%v', '%v'", s.in, sc, err, s.out, s.err)
			}
			if sc != nil && sc.Secret == GlobalSecret {
				t.Errorf("Parse function failed to hide global secret: %v", sc.Secret)
			}
		}
	}
}
