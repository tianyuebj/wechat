package models

import (
	"encoding/xml"
)

type Wechat struct {
	Signature string `json:"signature"`
	Timestamp string `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Echostr   string `json:"echostr"`
}

type XmlTextSchema struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        int      `xml:"MsgId"`
	// Xml          string `xml:",innerxml"`
}

type CDATAFormat struct {
	CDATA string `xml:",innerxml"`
}

func ToCDATAFormat(data string) CDATAFormat {
	return CDATAFormat{"<![CDATA[" + data + "]]>"}
}

type XmlResSchema struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   CDATAFormat `xml:"ToUserName"`
	FromUserName CDATAFormat `xml:"FromUserName"`
	CreateTime   int64       `xml:"CreateTime"`
	MsgType      CDATAFormat `xml:"MsgType"`
	Content      CDATAFormat `xml:"Content"`
	MsgId        CDATAFormat `xml:"MsgId"`
	// Xml          string `xml:",innerxml"`
}
