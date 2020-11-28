package statements

import (
	rest ".."
	config "../../configManagement"
	visualize "../../visuailize"
	model "../model"
	"bytes"
	"encoding/json"
	"fmt"
)

func ListQueries() {
	r := model.Request{
		Ksql: "SHOW QUERIES;",
	}
	byteQuery, _ := json.Marshal(r)
	body := rest.NewRequest("POST", config.DefaultKsqlServerURL+"/ksql", bytes.NewBuffer(bytes.ToLower(byteQuery)))

	var response []model.KsqlQueryListResponse
	err := json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Exception occurred while list queries response unmarshal")
	}

	visualize.VisualizeQueryList(response)
}
