package param

type GetFromStorageRequest struct {
	Key string `json:"key"`
}

type GetFromStorageResponse struct {
	Value string `json:"value"`
}
