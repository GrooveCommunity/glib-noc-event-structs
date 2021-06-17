package entity

type JiraTransition struct {
	LastState    string
	CurrentState string
	DateTime     string
}

type JiraEvent struct {
	EventUser JiraUser
	DateTime  string
	Name      string
}

type JiraUser struct {
	Username string
}

type JiraCustomField struct {
	ID    string
	Label string
	Value string
}

type JiraIssue struct {
	Event        JiraEvent
	Transitions  []JiraTransition
	Key          string
	Assignee     JiraUser
	Creator      JiraUser
	Reporter     JiraUser
	ChangeDate   string
	Priority     string
	Status       string
	Type         string
	CreatedDate  string
	UpdatedDate  string
	Summary      string
	Description  string
	Project      string
	CustomFields []JiraCustomField
}
