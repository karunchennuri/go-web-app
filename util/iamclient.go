package util

import (
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"github.com/astaxie/beego"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

type IAMClient struct {
	iamUrl string
	authorizationHeader string
	Content string
	Host string
	TimeStamp string
	httpClient *http.Client
}

func NewIAMClient(url string) (*IAMClient, error) {
	return newIAMClient(url, http.DefaultClient)
}

func newIAMClient(url string, client *http.Client) (*IAMClient, error) {
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	return &IAMClient {
		//iamUrl: fmt.Sprintf("%sv3/auth/tokens", url),
		iamUrl: fmt.Sprintf("%sv1.0/1087c3c1e90c483cb9dc4cc997aa3273/quotas/dms", url),
		httpClient: client,
	}, nil
}

func (ctx *IAMClient) createHttpRequest() (*http.Request, error) {
	var body io.Reader
	body = nil
	if len(ctx.Content) > 0 {
		body = strings.NewReader(ctx.Content)
	}

	req, err := http.NewRequest("GET", ctx.iamUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("host", ctx.Host)
	req.Header.Add("x-sdk-date", ctx.TimeStamp)
	req.Header.Add("authorization", ctx.authorizationHeader)

	return req, nil
}

func (ctx *IAMClient) ValidateIAMToken(token string) ([]byte, error) {
	beego.Info("*** Inside ValidateIAMToken ****")
	req, _ := ctx.createHttpRequest()
	beego.Info("token=", token)
	var ret []byte
	var err error

	if len(token) > 0 {
		req.Header.Add("x-auth-token", token)
	}

	resp, err := ctx.httpClient.Do(req)
	beego.Info("resp=", resp)
	beego.Info("resp status Code=", resp.StatusCode)
	beego.Info("err=", err)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case 200:
		beego.Info("**** VALID TOKEN ****")
		ret, err = ioutil.ReadAll(resp.Body)
		beego.Info("ret=", string(ret))
		beego.Info("err=", err)
	case 401:
		beego.Info("**** Unauthorized TOKEN ****")
		err = errors.New("Unauthorized Token")
	default:
		beego.Info("**** INVALID TOKEN ****")
		err = errors.New("Invalid Token")
	}

	defer resp.Body.Close()

	beego.Info("*** End ValidateIAMToken ****")
	return ret, err
}