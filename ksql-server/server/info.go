package server

import (
	rest ".."
	config "../../init-config"
	visualize "../../visuailize"
	model "../model"
	"encoding/json"
	"fmt"
)

func Info() {

	body := rest.NewRequest("GET", config.DefaultKsqlServerURL+"/info", nil)

	var response model.KsqlServerInfoResponse
	err := json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Exception occurred while info response unmarshal")
	}

	visualize.VisualizeServerInfo(response)
}
