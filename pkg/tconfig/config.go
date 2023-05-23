package tconfig

// Config Interface for all local packages' configs
type Config interface {
	ParseConfig() error
}

type BaseConfig struct {
	Service struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"Service"`
	System struct {
		MaxGoroutines uint64 `json:"maxGoroutines"`
		Host          string `json:"host"`
		Port          string `json:"port"`
	} `json:"System"`
}
