package model

type KsqlServerInfoResponse struct {
	KsqlServerInfo struct {
		Version        string `json:"version"`
		KafkaClusterID string `json:"kafkaClusterId"`
		KsqlServiceId   string `json:"ksqlServiceId"`
	} `json:"KsqlServerInfo"`
}
