package config

type HttpsConfig struct {
	CertFilePath    string `json:"certFilePath"`
	CertKeyFilePath string `json:"certKeyFilePath"`
}
