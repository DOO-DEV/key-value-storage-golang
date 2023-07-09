package param

type PutInStoreRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PutInStorageResponse struct {
	Status Status `json:"status"`
}

type Status string

const (
	StatusCreated Status = "CREATED"
	StatusUpdated Status = "Updated"
)
