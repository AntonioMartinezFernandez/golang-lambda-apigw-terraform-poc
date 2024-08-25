package healthcheck_application

type GetHealthcheckQuery struct {
}

func NewGetHealthcheckQuery() *GetHealthcheckQuery {
	return &GetHealthcheckQuery{}
}

func (hq GetHealthcheckQuery) Data() map[string]interface{} {
	return map[string]interface{}{}
}
