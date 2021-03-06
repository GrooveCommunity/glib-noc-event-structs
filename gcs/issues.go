package gcs

type IssueState struct {
	EventName  string
	EventDate  string
	EventUser  string
	Status     string
	CreateDate string
	ChangeDate string
	Priority   string
	Assignee   string
	Reporter   string
}

type Issue struct {
	KeyID  string
	States []IssueState
}

type IssuesMetric struct {
	Issues []Issue
}
