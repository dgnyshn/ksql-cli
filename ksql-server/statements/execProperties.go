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

func ListProperties() {
	r := model.Request{
		Ksql: "SHOW PROPERTIES;",
	}
	byteQuery, _ := json.Marshal(r)
	body := rest.NewRequest("POST", config.DefaultKsqlServerURL+"/ksql", bytes.NewBuffer(bytes.ToLower(byteQuery)))

	var response []model.ListPropertiesResponse

	err := json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Exception occurred while list properties response unmarshal")
	}

	visualize.VisualizePropertyList(response)

}
