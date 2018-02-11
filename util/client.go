package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fhyx/Wechat/models"
)

const (
	url = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential"
)

var client = &http.Client{}

func GetAccessToken(appID, appsecret string) ([]byte, error) {
	geturl := fmt.Sprintf("%s&appid=%s&secret=%s", url, appID, appsecret)
	req, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		return []byte("get request err"), err
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte("get response err"), err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("get read resp.body err"), err
	}
	return body, nil
}

func PostAccessToken(access string) ([]byte, error) {
	button := models.Button{
		Name:      "扫码",
		Type:      "scancode_push",
		Key:       "rselfmenu_0_1",
		SubButton: nil,
	}
	var buttons []models.Button
	buttons = append(buttons, button)
	menu := models.Menu{
		Button: buttons,
	}

	buf, err := json.Marshal(menu)
	if err != nil {
		return nil, err
	}

	resp, err := client.Post("https://api.weixin.qq.com/cgi-bin/menu/create?access_token="+access, "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return []byte("response post err"), err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("read resp.body post err"), err
	}
	return body, nil
}
