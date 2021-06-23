package entity

type JiraAttachmentRequest struct {
	ID       string
	Filename string
	Created  string
	Content  string
}

type JiraUserRequest struct {
	Name string `json:"displayName"`
}

type JiraIssueRequest struct {
	ID     string
	Key    string
	Fields JiraFieldsRequest
}

type ChangeLogEventRequest struct {
	Field string
	From  string `json:"fromString"`
	To    string `json:"toString"`
}

type JiraChangeLogRequest struct {
	Changes []ChangeLogEventRequest `json:"items"`
}

type JiraPriorityRequest struct {
	Name string
}

type JiraResolutionRequest struct {
	Name        string
	Description string
}

type JiraAssigneeRequest struct {
	User string `json:"displayName"`
}

type JiraStatusRequest struct {
	Name string
}

type JiraCreatorRequest struct {
	Name string `json:"displayName"`
}

type JiraReporterRequest struct {
	Name string `json:"displayName"`
}

type JiraTypeRequest struct {
	Name string
}

type JiraProjectRequest struct {
	Name string
}

type JiraFieldsRequest struct {
	ChangeDate   string `json:"statuscategorychangedate"`
	User         string `json:"displayName"`
	Priority     JiraPriorityRequest
	Resolution   JiraResolutionRequest
	Assignee     JiraAssigneeRequest
	Creator      JiraCreatorRequest
	Reporter     JiraReporterRequest
	IssueType    JiraTypeRequest
	Project      JiraProjectRequest
	Created      string
	Updated      string
	Summary      string
	Status       JiraStatusRequest
	Description  string
	Attachment   []JiraAttachmentRequest
	CustomFields []CustomField
}

type JiraRequest struct {
	EventName string `json:"webhookEvent"`
	User      JiraUserRequest
	Issue     JiraIssueRequest
	ChangeLog JiraChangeLogRequest
}
