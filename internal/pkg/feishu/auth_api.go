package feishu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type FeishuAppInfo struct {
	AppID     string
	AppSecret string
}

var FeishuClient = &FeishuAppInfo{}

func SetFeishuApp(appId, appSecret string) {
	FeishuClient.AppID = appId
	FeishuClient.AppSecret = appSecret
}

const tenantAccessTokenApi = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"

type AccessTokenResponse struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int    `json:"expire"`
}

func GetAccessToken() (string, error) {
	return GetAccessTokenFromApp(FeishuClient.AppID, FeishuClient.AppSecret)
}

func GetAccessTokenFromApp(appId, appSecret string) (string, error) {
	requestUrl := tenantAccessTokenApi

	jsonContent := fmt.Sprintf(` { "app_id": "%s", "app_secret": "%s" } `, appId, appSecret)

	var jsonData = []byte(jsonContent)
	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Body

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	result := AccessTokenResponse{}
	json.Unmarshal(respBody, &result)
	if result.Code != 0 {
		return "", fmt.Errorf("Get Feishu access token error: error code %d, error msg: %s", result.Code, result.Msg)
	}

	return result.TenantAccessToken, nil
}
