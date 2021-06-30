package gcs

type IssueState struct {
	Name     string
	DateTime string
	User     string
}

type Issue struct {
	KeyID string
	State IssueState
}

type IssuesMetric struct {
	Issues []Issue
}
