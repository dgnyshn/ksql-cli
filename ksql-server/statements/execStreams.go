package statements

import (
	rest ".."
	config "../../init-config"
	visualize "../../visuailize"
	model "../model"
	"bytes"
	"encoding/json"
	"fmt"
)

func ListStreams() {
	r := model.Request{
		Ksql: "SHOW STREAMS;",
	}
	byteQuery, _ := json.Marshal(r)

	body := rest.NewRequest("POST", config.DefaultKsqlServerURL+"/ksql", bytes.NewBuffer(bytes.ToLower(byteQuery)))

	var response []model.KsqlStreamResponse
	err := json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Exception occurred while list streams response unmarshal")
	}

	visualize.VisualizeStream(response)
}
