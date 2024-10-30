package dto

type Model struct {
	ModelId   string `json:"model_id"`
	ModelName string `json:"model_name"`
}

type ModelQueryResp struct {
	Models []*Model `json:"models"`
}
