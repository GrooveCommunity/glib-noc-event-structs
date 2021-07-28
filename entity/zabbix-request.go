package entity

type ZabbixRequestFields struct {
	EventID                   string
	Eventname                 string
	EventRecoveryValue        string
	EventSeverity             string
	EventRecoveryDate         string
	EventDate                 string
	EventTime                 string
	EventRecoveryTime         string
	EventAck                  string
	EventDuration             string
	Message                   string
	User                      string
	TriggerID                 string
	TriggerStatus             string
	TriggerExpression         string
	TriggerExpressionRecovery string
	Hostname                  string
	RequestDate               string
}

type ZabbixRequest struct {
	Fields ZabbixRequestFields
}
