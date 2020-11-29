package server

import (
	rest ".."
	config "../../initCfg"
	visualize "../../visuailize"
	model "../model"
	"encoding/json"
	"fmt"
)

func HealthCheck() {
	body := rest.NewRequest("GET", config.DefaultKsqlServerURL+"/healthcheck", nil)

	var response model.KsqlServerHealthCheckResponse
	err := json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Exception occurred while healthcheck response unmarshal")
	}
	visualize.VisualizeServerHealthCheck(response)
}
