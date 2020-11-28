package model

type DescribeResponse struct {
	Type              string `json:"@type"`
	StatementText     string `json:"statementText"`
	SourceDescription struct {
		Name         string        `json:"name"`
		WindowType   interface{}   `json:"windowType"`
		ReadQueries  []interface{} `json:"readQueries"`
		WriteQueries []interface{} `json:"writeQueries"`
		Fields       []struct {
			Name   string `json:"name"`
			Schema struct {
				Type         string      `json:"type"`
				Fields       interface{} `json:"fields"`
				MemberSchema interface{} `json:"memberSchema"`
			} `json:"schema"`
		} `json:"fields"`
		Type        string `json:"type"`
		Timestamp   string `json:"timestamp"`
		Statistics  string `json:"statistics"`
		ErrorStats  string `json:"errorStats"`
		Extended    bool   `json:"extended"`
		KeyFormat   string `json:"keyFormat"`
		ValueFormat string `json:"valueFormat"`
		Topic       string `json:"topic"`
		Partitions  int    `json:"partitions"`
		Replication int    `json:"replication"`
		Statement   string `json:"statement"`
	} `json:"sourceDescription"`
	Warnings []interface{} `json:"warnings"`
}
