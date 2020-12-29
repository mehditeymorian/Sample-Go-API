package model

type Msg struct {
	Msg string `json:"msg"`
}

func MsgResp(msg string) Msg {
	return Msg{
		Msg: msg,
	}
}
