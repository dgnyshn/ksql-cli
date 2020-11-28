package model

type Request struct {
	Ksql string
}

type KsqlStreamResponse struct {
	StatementText string           `json:"statementText"`
	Streams       []StreamsDetails `json:"streams"`
}

type StreamsDetails struct {
	Name   string `json:"name"`
	Topic  string `json:"topic"`
	Format string `json:"format"`
}