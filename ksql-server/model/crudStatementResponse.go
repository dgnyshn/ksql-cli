package model

type CrudStatementResponse struct {
	Type          string `json:"@type"`
	StatementText string `json:"statementText"`
	CommandID     string `json:"commandId"`
	CommandStatus struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	} `json:"commandStatus"`
	CommandSequenceNumber int           `json:"commandSequenceNumber"`
	Warnings              []interface{} `json:"warnings"`
}