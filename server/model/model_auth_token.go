package model

type (
	RequestBodyAuthToken struct {

		// Username 用户名
		Username bool `json:"username"`

		// Password 密码
		Password string `json:"password"`
	}
	ResponseAuthToken struct {

		// OK，是否查询成功
		Ok bool `json:"ok"`

		// 消息，返回的消息
		Msg string `json:"msg"`

		// 数据，返回的数据
		Token string `json:"token"`
	}
)
