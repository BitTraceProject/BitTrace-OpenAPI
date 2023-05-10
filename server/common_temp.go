package server

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/BitTraceProject/BitTrace-Types/pkg/common"
	"time"
)

// 这里的代码需要进一步优化，目前时间不够只能这么写了

type Token struct {
	uid        string
	createTime time.Time

	token string
}

var (
	//	priKey = `MIIBVQIBADANBgkqhkiG9w0BAQEFAASCAT8wggE7AgEAAkEA7fpfUahbRrbBGW7E
	//QkIFwRTtofxR4HR/zdJMIbEya1PzbZggpwAVMBrmFTSpFY4OB5rwEtLA69NRE/Wd
	//I+EvDQIDAQABAkEAtu3ToYhhiWzn6C3eDPSFSdNVi0PSiE/lAgiaXve/wwFMzl6u
	//o1iilAVBY7pki661lhc24gj5JaAaq+LumxGEfQIhAPjw6VDdIg9JMtakxkRMClFu
	//R30I9YIgoydRQHUgLu9fAiEA9LnhDJXuWyDvYYTE7FiFpWRTNKEM1uP6+1bliqZh
	//dRMCIG34RqlhNoWYKWwmmCtdiAVW+R+kEIhukIRy4U2cbP9zAiEAjiQce8PlspZZ
	//k4mbGy97SIoR7eKQ44t9LljVcAedWXMCIFTX8WCxg3c/7TLtrpZWmAZ0isdRdRQD
	//oDXtdkV1YqUa`
	//	pubKey = `MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAO36X1GoW0a2wRluxEJCBcEU7aH8UeB0
	//f83STCGxMmtT822YIKcAFTAa5hU0qRWODgea8BLSwOvTURP1nSPhLw0CAwEAAQ==`

	initUid = map[string]string{
		"jianhuizh": "",
		"swu":       "",
		"rise":      "",
	}
	tokenKV = map[string]Token{}

	expireDuration = 3 * time.Hour
)

func genHash(s string) string {
	c := sha1.New()
	c.Write([]byte(s))
	bytes := c.Sum(nil)
	return hex.EncodeToString(bytes)
}

func newToken(uid string) (token string) {
	now := time.Now()
	if _t, ok := initUid[uid]; ok {
		if _, ok := tokenKV[_t]; ok {
			// 已存在，则删掉
			removeToken(_t)
		}
	}
	// 不存在，创建
	token = genHash(uid + "-" + common.FromTime(now).String())
	tokenKV[token] = Token{
		uid:        uid,
		createTime: now,

		token: token,
	}
	initUid[uid] = token
	return token
}

func removeToken(token string) {
	t, ok := tokenKV[token]
	if ok {
		delete(tokenKV, token)
		initUid[t.uid] = ""
	}
}

func checkToken(token string) bool {
	t, ok := tokenKV[token]
	if !ok {
		return false
	}
	// 再检查一下 uid 是否还存在
	if !checkUid(t.uid) {
		return false
	}
	if t.expired() {
		removeToken(token)
		return false
	}
	return true
}

func (t *Token) expired() bool {
	now := time.Now()
	if now.After(t.createTime.Add(expireDuration)) {
		removeToken(t.token)
		return true
	}
	return false
}

func checkUid(uid string) bool {
	_, ok := initUid[uid]
	if !ok {
		return false
	}
	return true
}
