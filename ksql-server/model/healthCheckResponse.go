package model

type KsqlServerHealthCheckResponse struct {
	IsHealthy bool `json:"isHealthy"`
	Details   struct {
		Metastore struct {
			IsHealthy bool `json:"isHealthy"`
		} `json:"metastore"`
		Kafka struct {
			IsHealthy bool `json:"isHealthy"`
		} `json:"kafka"`
	} `json:"details"`
}
