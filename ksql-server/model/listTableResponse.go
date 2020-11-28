package model

type KsqlTableResponse struct {
	Tables []TableDetails `json:"tables"`
}

type TableDetails struct {
	Name   string `json:"name"`
	Topic  string `json:"topic"`
	Format string `json:"format"`
}
