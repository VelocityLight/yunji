package feishu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const msgApi = "https://open.feishu.cn/open-apis/im/v1/messages"

type MsgIDType string

const (
	MsgIDTypeEmail = MsgIDType("email")
)

type MsgType string

const (
	MsgTypeText = MsgType("text")
	MsgTypePost = MsgType("post")
)

type MsgReceiver struct {
	IDType MsgIDType `json:"receive_id_type"`
	ID     string    `json:"receive_id"`
}

type MsgReq struct {
	ReceiveID string `json:"receive_id"`
	MsgType   string `json:"msg_type"`
	Msg       string `json:"content"`
}

type MsgResp struct {
	Code    int                    `json:"code,omitempty"`
	Message string                 `json:"msg"`
	Data    map[string]interface{} `json:"data"`
}

func SendMsgCard(receiver MsgReceiver, msg MsgWrapper, token string) error {
	requestUrl := fmt.Sprintf("%s?receive_id_type=%s", msgApi, receiver.IDType)

	reqBody := MsgReq{
		ReceiveID: receiver.ID,
		MsgType:   msg.GetMsgType(),
		Msg:       msg.GetMsgJson(),
	}

	jsonContent, err := json.Marshal(reqBody)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonContent))
	if err != nil {
		return err
	}
	authString := fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Authorization", authString)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	result := MsgResp{}
	json.Unmarshal(body, &result)

	if result.Code != 0 {
		return fmt.Errorf("Send Feishu message error: error code %d, error msg: %s", result.Code, result.Message)
	}

	return nil
}
