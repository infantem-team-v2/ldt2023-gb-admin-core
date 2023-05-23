package server

type RouteMapping struct {
	HttpMethod string `yaml:"httpMethod"`
	Route      string `yaml:"route"`
}
