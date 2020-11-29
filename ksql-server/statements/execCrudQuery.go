package statements

import (
	rest ".."
	config "../../initCfg"
	visualize "../../visuailize"
	model "../model"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func ExecuteCrudStatement(query string) {
	r := model.Request{
		Ksql: strings.TrimSpace(query),
	}
	byteQuery, _ := json.Marshal(r)

	body := rest.NewRequest("POST", config.DefaultKsqlServerURL+"/ksql", bytes.NewBuffer(bytes.ToLower(byteQuery)))

	var response [] model.CrudStatementResponse
	err := json.Unmarshal(body, &response)

	fmt.Println(string(body))
	fmt.Println("-------------------------------------------")
	if err != nil {
		fmt.Println("Exception occurred while executeQuery response unmarshal")
	}

	visualize.VisualizeCrudStatementResponse(response)
}
