package config

type TSecureConfig struct {
	Fernet struct {
		Key string `json:"key"`
	} `json:"Fernet"`

	RSA struct {
		PublicKey  string `json:"public_key"`
		PrivateKey string `json:"private_key"`
	} `json:"RSA"`
}

func (tsc *TSecureConfig) ParseConfig() error {
	return nil
}
