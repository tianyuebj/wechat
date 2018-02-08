package web

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"

	"github.com/gin-gonic/gin"
)

const wx_token = "testToken"

func handlerWeChatGet(c *gin.Context) {
	var (
		token    string
		strs     string
		hashcode string
	)

	signature := c.DefaultQuery("signature", "")
	timestamp := c.DefaultQuery("timestamp", "")
	nonce := c.DefaultQuery("nonce", "")
	echostr := c.DefaultQuery("echostr", "")

	token = wx_token
	strSlice := sort.StringSlice{token, timestamp, nonce}
	sort.Sort(strSlice)
	for _, s := range strSlice {
		strs += s
	}
	hashcode = sha1String(strs)
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
