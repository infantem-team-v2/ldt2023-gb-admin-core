package config

type TLoggerConfig struct {
	Loki struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Batch    struct {
			Number int `json:"number"`
			Wait   int `json:"wait"`
		} `json:"Batch"`
	} `json:"Loki"`
}

func (tlc *TLoggerConfig) ParseConfig() error {
	return nil
}
