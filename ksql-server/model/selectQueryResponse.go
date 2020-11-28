package model

type SelectStatementHeaderResponse struct {
	Header struct {
		QueryID string `json:"queryId"`
		Schema  string `json:"schema"`
	} `json:"header"`
}

type SelectStatementBodyResponse struct {
	Row struct {
		Columns []interface{} `json:"columns"`
	} `json:"row"`
}

func (response SelectStatementBodyResponse) GetColumns() interface{} {
	return response.Row.Columns
}
