package exercise

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/autograde/kit/score"
)

type Choices []struct {
	Number int
	Want   rune
}

// MultipleChoice computes the score of a multiple choice exercise
// with student answers provided in fileName, and the answers provided
// in the answerKey object. The function requires a Score object, and
// will produce both string output and JSON output.
func MultipleChoice(t *testing.T, sc *score.Score, fileName string, answers Choices) {
	defer sc.WriteString(os.Stdout)
	defer sc.WriteJSON(os.Stdout)

	// Read the whole file
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		sc.Score = 0
		t.Fatalf(fmt.Sprintf("%v: error reading the file: %v", fileName, err))
		return
	}

	for i := range answers {
		// Find the user's answer to the corresponding question number
		regexStr := "\n" + strconv.Itoa(answers[i].Number) + "[.)]*[ \t\v\r\n\f]*[A-Za-z]*"
		regex := regexp.MustCompile(regexStr)
		userAnswer := regex.Find(bytes)

		if userAnswer == nil {
			t.Errorf("%v %d: Answer not found.\n", sc.TestName, answers[i].Number)
			sc.Dec()
		} else {
			r, _ := utf8.DecodeLastRune(userAnswer)
			got, _ := utf8.DecodeLastRuneInString(strings.ToUpper(string(r)))
			if got != answers[i].Want {
				t.Errorf("%v %d: %q is incorrect.\n", sc.TestName, answers[i].Number, got)
				sc.Dec()
			}
		}
	}
}
