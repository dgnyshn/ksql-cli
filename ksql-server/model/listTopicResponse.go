package model


type KsqlTopicResponse struct {
	Topics []TopicDetails `json:"topics"`
}

type TopicDetails struct {
	Name        string `json:"name"`
	ReplicaInfo []int  `json:"replicaInfo"`
}
