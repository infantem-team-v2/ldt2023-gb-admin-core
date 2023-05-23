package config

type TStorageConfig struct {
	Relational struct {
		Postgres struct {
			Host            string `json:"host"`
			Port            string `json:"port"`
			User            string `json:"user"`
			Password        string `json:"password"`
			DBName          string `json:"dbName"`
			SSLMode         string `json:"sslMode"`
			PgDriver        string `json:"pgDriver"`
			ConnMaxIdleTime int    `json:"connMaxIdleTime"`
			MaxOpenConns    int    `json:"maxOpenConns"`
		} `json:"Postgres"`
	} `json:"Relational"`
	Cache struct {
		Redis struct {
			Host     string `json:"host"`
			Port     string `json:"port"`
			DB       string `json:"db"`
			Password string `json:"password"`
		} `json:"Redis"`
	} `json:"Cache"`
	Metrics struct {
		Prometheus struct {
			Name string `json:"name"`
			Help string `json:"help"`
		} `json:"Prometheus"`
	} `json:"Metrics"`
	NonRelational struct {
	}
}
