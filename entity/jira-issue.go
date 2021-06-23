package entity

type JiraTransition struct {
	LastState    string
	CurrentState string
}

type JiraEvent struct {
	EventUser string
	DateTime  string
	EventName string //webhook event
}

type JiraIssue struct {
	Event        JiraEvent
	Transitions  []JiraTransition
	Key          string
	Assignee     string
	Creator      string
	Reporter     string
	ChangeDate   string
	Priority     string
	Status       string
	Type         string
	CreatedDate  string
	UpdatedDate  string
	Summary      string
	Description  string
	Project      string
	CustomFields []CustomField
	Attachment   []JiraAttachmentRequest
}
