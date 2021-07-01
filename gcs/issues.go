package gcs

type IssueState struct {
	EventName string
	DateTime  string
	User      string
}

type Issue struct {
	KeyID  string
	States []IssueState
}

type IssuesMetric struct {
	Issues []Issue
}
