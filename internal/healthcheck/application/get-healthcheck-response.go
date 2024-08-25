package healthcheck_application

type GetHealthcheckResponse struct {
	Id      string `json:"id" jsonapi:"primary,healthcheck"`
	Status  string `json:"status" validate:"required" jsonapi:"attr,status"`
	Service string `json:"service" validate:"required" jsonapi:"attr,service"`
}

func NewGetHealthcheckResponse(id string, status string, service string) GetHealthcheckResponse {
	return GetHealthcheckResponse{Id: id, Status: status, Service: service}
}
