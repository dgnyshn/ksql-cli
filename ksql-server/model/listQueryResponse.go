package model

type KsqlQueryListResponse struct {
	Queries []QueryDetails `json:"queries"`
}

type QueryDetails struct {
	QueryString string `json:"queryString"`
	Id          string `json:"id"`
}
