package model

type (
	RequestBodyAuthRegister struct {

		// Username 用户名
		Username bool `json:"username"`

		// Password 密码
		Password string `json:"password"`
	}
	ResponseAuthRegister struct {

		// OK，是否查询成功
		Ok bool `json:"ok"`

		// 消息，返回的消息
		Msg string `json:"msg"`
	}
)
