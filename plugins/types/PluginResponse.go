package types

// PluginResponse is the response an external program will get when it has completed a task.
type PluginResponse struct {
	Handin      string       `json:"handin"`      // Which hand-in the results belong to
	Course      string       `json:"course"`      // Which course the results belong to
	Source      string       `json:"source"`      // Which tool that have generated the tests.
	Category    string       `json:"category"`    // Explanation on what type of test this is.
	TestResults []PluginTest `json:"testresults"` // List of test results from the plugin.
	TotalScore  int          `json:"severity"`    // Some information on how bad the result is. I think it is similar to your Score. Can also be binary for pass and fail.
}

// PluginTest explains results from one test.
type PluginTest struct {
	Title    string   `json:"title"`    // Title of the test. Might be overlap with comment.
	Comments []string `json:"comments"` // Output from the test. Unit tests often assert something, such as: Sent inn 1 and 3, expected 4, got 5.
	Score    int      `json:"severity"` // Some information on how bad the result is. I think it is similar to your Score. Can also be binary for pass and fail.
	File     []string `json:"files"`    // Filename(s). This can probably be extracted from Code.
	Offset   int      `json:"offset"`   // Position of the assertion. Omitted if not available.
	Length   int      `json:"length"`   // Length of the assertion. Omitted if not available.
	Code     []string `json:"codes"`    // Some reference to the file(s) tested so that the GUI can present the file and for example part of the improvements yellow.
}
