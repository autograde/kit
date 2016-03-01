package types

// PluginRequest is the expected format of an external request for building a solution.
type PluginRequest struct {
	StudentRepo     string   `json:"student_repository"` // Repository with student code.
	TestRepo        string   `json:"test_repository"`    // Repository with seperate maintained test cases.
	Commands        []string `json:"commands"`           // Extra commands to run. Omitted if no commands to run.
	Files           []string `json:"files"`              // Files to analyse. Eks: files to execute in build, files to run plagiraizm on.
	FolderStructure string   `json:"folder_structure"`   // A base folder structure used to locate the repository location.
	ServiceUsername string   `json:"service_username"`   // Username used to access repositories. Omitted if authentication is not needed.
	ServicePassword string   `json:"service_password"`   // Password used to access repositories. Omitted if authentication is not needed.
}
