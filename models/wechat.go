package models

type Wechat struct {
	Signature string `json:"signature"`
	Timestamp string `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Echostr   string `json:"echostr"`
}
