package model

type OssToken struct {
	EndPoint string `json:"endpoint"`
	Bucket   string `json:"bucket"`
	Path     string `json:"path"`
	Key      string `json:"key"`
	Secret   string `json:"secret"`
	Token    string `json:"token"`
	Expire   int64  `json:"expire"`
}
