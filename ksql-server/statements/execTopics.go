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

func ListTopics() {
	r := model.Request{
		Ksql: "SHOW TOPICS;",
	}
	byteQuery, _ := json.Marshal(r)
	buffer := bytes.NewBuffer(bytes.ToLower(byteQuery))
	body := rest.NewRequest("POST", config.DefaultKsqlServerURL+"/ksql", buffer)

	var response []model.KsqlTopicResponse
	err := json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Exception occurred while list topic response unmarshal")
	}

	visualize.VisualizeTopic(response)
}
