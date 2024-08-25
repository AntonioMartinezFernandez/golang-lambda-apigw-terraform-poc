package http_pkg

type CommonJsonapiResponse struct {
	Id      string `json:"id" jsonapi:"primary,common_response"`
	Message string `json:"message" validate:"required" jsonapi:"attr,message"`
}
