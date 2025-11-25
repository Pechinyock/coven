package config

import "fmt"

type HttpsConfig struct {
	CertFilePath    string `json:"certFilePath"`
	CertKeyFilePath string `json:"certKeyFilePath"`
}

func (h *HttpsConfig) Validate() error {
	fmt.Println("https config validation is not implemented")
	return nil
}
