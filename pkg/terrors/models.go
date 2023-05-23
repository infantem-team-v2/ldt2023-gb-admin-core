package terrors

type externalMessage struct {
	Description  string `json:"description"`
	InternalCode uint32 `json:"internal_code"`
}

type serializedExternalMessage struct {
	Message    string `yaml:"message"`
	StatusCode int    `yaml:"statusCode"`
}
