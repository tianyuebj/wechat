package web

import (
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sort"
	"time"

	"github.com/gin-gonic/gin"

	"fhyx/Wechat/models"
	"fhyx/Wechat/util"
)

const wx_token = "testToken"

func handlerWeChatGet(c *gin.Context) {
	var (
		strs     string
		hashcode string
	)

	signature := c.DefaultQuery("signature", "")
	timestamp := c.DefaultQuery("timestamp", "")
	nonce := c.DefaultQuery("nonce", "")
	echostr := c.DefaultQuery("echostr", "")

	strSlice := sort.StringSlice{wx_token, timestamp, nonce}
	sort.Sort(strSlice)
	for _, s := range strSlice {
		strs += s
	}
	hashcode = sha1String(strs)

	//获取access 暂时放这
	const (
		appID     = "wxe61d7bef9bd8376a"
		appsecret = "5d6e48510b0ffcd762def2403798038a"
	)
	getData, err := util.GetAccessToken(appID, appsecret)
	if err != nil {
		log.Println("get access token err:", err)
		return
	}
	var accessResp models.AccessToken
	err = json.Unmarshal(getData, &accessResp)
	if err != nil {
		log.Println("access unmarshal err:", err)
		return
	}
	// log.Println("wa:", accessResp)
	//向微信服务器发POST
	postData, err := util.PostAccessToken(accessResp.AccessToken)
	if err != nil {
		log.Println("post access_token return err:", err)
		return
	}
	var accessReturn models.AccessToken
	err = json.Unmarshal(postData, &accessReturn)
	if err != nil {
		log.Println("access_return unmarshal err:", err)
		return
	}
	// log.Println("gua:", accessReturn)

	if hashcode == signature && echostr != "" {
		c.Writer.Write([]byte(echostr))
		return
	}

	c.Writer.WriteHeader(403)
	return
}

func sha1String(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func handlerWeChatPost(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("read xml err: %s", err)
		c.Writer.WriteHeader(400)
		return
	}
	var xmlReq models.XmlTextSchema
	xml.Unmarshal(data, &xmlReq)

	if xmlReq.MsgType != "text" {
		c.XML(403, "success")
		return
	}
	xmlRes := models.XmlResSchema{
		ToUserName:   models.ToCDATAFormat(xmlReq.FromUserName),
		FromUserName: models.ToCDATAFormat(xmlReq.ToUserName),
		CreateTime:   time.Now().Unix(),
		MsgType:      models.ToCDATAFormat(xmlReq.MsgType),
		Content:      models.ToCDATAFormat(xmlReq.Content),
	}

	c.XML(200, xmlRes)
}
