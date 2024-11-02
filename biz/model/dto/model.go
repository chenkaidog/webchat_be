package dto

type Model struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ModelQueryResp struct {
	Models []*Model `json:"models"`
}
