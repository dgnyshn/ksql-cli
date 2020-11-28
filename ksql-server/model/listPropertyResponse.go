package model

type ListPropertiesResponse struct {
	Type          string `json:"@type"`
	StatementText string `json:"statementText"`
	Properties    []struct {
		Name  string `json:"name"`
		Scope string `json:"scope"`
		Value string `json:"value"`
	} `json:"properties"`
	OverwrittenProperties []interface{} `json:"overwrittenProperties"`
	DefaultProperties     []string      `json:"defaultProperties"`
	Warnings              []interface{} `json:"warnings"`
}
