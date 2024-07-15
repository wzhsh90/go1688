package client

import (
	"errors"
	"github.com/wzhsh90/go1688/common"
	"github.com/wzhsh90/go1688/consts"
	"github.com/wzhsh90/go1688/models"
	"github.com/wzhsh90/go1688/utils"
	"strconv"
	"time"
)

type Client struct {
	appKey      string
	appSecret   []byte
	accessToken string
}

func NewClient(appKey, appSecret, accessToken string) *Client {
	return &Client{
		appKey:      appKey,
		appSecret:   []byte(appSecret),
		accessToken: accessToken,
	}
}
func (c *Client) Do(req models.Request, respPt models.Response) error {
	reqPath := c.requestPath(req)
	reqParams := req.Params()
	reqParams["_aop_timestamp"] = strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	reqParams["access_token"] = c.accessToken
	reqParams["_aop_signature"] = utils.Sign(reqPath, c.appSecret, reqParams)
	requestUri := c.requestUri(reqPath)
	respData, err := common.PostForm(requestUri, reqParams, nil)
	if err != nil {
		return err
	}
	println(string(respData))
	unErr := utils.Unmarshal(respData, respPt)
	if unErr != nil {
		return unErr
	}
	if respPt == nil {
		return errors.New("response nil")
	}
	if respPt.IsError() {
		return errors.New(respPt.GetErrorMsg())
	}
	return nil
}

func (c *Client) requestUri(reqPath string) string {
	builder := utils.GetStringsBuilder()
	defer utils.PutStringsBuilder(builder)
	builder.WriteString(consts.GATEWAY)
	builder.WriteString("/")
	builder.WriteString(reqPath)
	return builder.String()
}

func (c *Client) requestPath(req models.Request) string {
	builder := utils.GetStringsBuilder()
	defer utils.PutStringsBuilder(builder)
	builder.WriteString(consts.PROTOCOL)
	builder.WriteString("/")
	builder.WriteString(req.Path())
	builder.WriteString("/")
	builder.WriteString(c.appKey)
	return builder.String()
}
