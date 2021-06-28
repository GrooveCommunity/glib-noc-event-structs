package entity

type Turno struct {
	Inicio string
	Fim    string
}

type JiraNocUser struct {
	Name  string
	Turno Turno
}

type NocUsers struct {
	JiraUsers []JiraNocUser `json:"jira-users"`
}
