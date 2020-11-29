package statements

import (
	rest ".."
	config "../../initCfg"
	visualize "../../visuailize"
	model "../model"
	"bytes"
	"encoding/json"
	"fmt"
)

func ListTables() {
	r := model.Request{
		Ksql: "SHOW TABLES;",
	}
	byteQuery, _ := json.Marshal(r)
	body := rest.NewRequest("POST", config.DefaultKsqlServerURL+"/ksql", bytes.NewBuffer(bytes.ToLower(byteQuery)))

	var response []model.KsqlTableResponse
	err := json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Exception occurred while list tables response unmarshal")
	}
	visualize.VisualizeTable(response)
}
