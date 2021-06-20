package entity

type CustomField struct {
	CustomID      string
	ID            string
	Name          string
	Value         string
	RequestType   RequestTypeCustomField
	CurrentStatus CurrentStatusCustomField
}

type RequestTypeCustomField struct {
	CustomID    string
	ID          string
	Name        string
	Description string
}

type CurrentStatusCustomField struct {
	Status         string
	StatusCategory string
}

type TypeCustomField struct {
	RequestType   RequestTypeCustomField
	CurrentStatus CurrentStatusCustomField
}
