package common

type Params struct {
	UserId uint64 `json:"-"`
}

type Response struct {
	Message      string `json:"description"`
	InternalCode int    `json:"internal_code"`
}
