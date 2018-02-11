package models

type Menu struct {
	Button []Button `json:"button"`
}

type Button struct {
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	Key       string      `json:"key"`
	SubButton []SubButton `json:"sub_button"`
}

type SubButton struct {
	Type      string   `json:"type"`
	Name      string   `json:"name"`
	Key       string   `json:"key"`
	SubButton []string `json:"sub_button"`
}

// type ButtonSuccess struct {
// 	Errcode int    `json:"errcode"`
// 	Errmsg  string `json:"errmsg"`
// }
