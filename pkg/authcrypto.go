package pkg

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

const (
	AccessSecret = "uOvKLmVfztaXGpNYd4Z0I1SiT7MweJhl"         // 私钥
	AccessExpire = 86400                                      // 持有期限
	Header       = `{"alg": "HS256","typ": "JWT"}`            // JWT的header部分 固定
	BaseHeader   = "eyJhbGciOiAiSFMyNTYiLCJ0eXAiOiAiSldUIn0=" // JWT的header部分 base64编码
)

type Payload struct {
	Exp int64 // 过期时间
	Iat int64 // 发布时间
	Id  int64 // 用户或视频id
}

// 传入用户的id 即可返回token
func GetAuthToken(id int64) (string, error) {
	payload := Payload{
		Exp: time.Now().Unix() + AccessExpire,
		Iat: time.Now().Unix(),
		Id:  id,
	}
	// HS256加密部分
	data, err := json.Marshal(payload)
	BaseData := base64.URLEncoding.EncodeToString(data)
	message := BaseHeader + BaseData
	if err != nil {
		return "", err
	}
	mac := hmac.New(sha256.New, []byte(AccessSecret))
	_, err = mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	secret := base64.URLEncoding.EncodeToString(mac.Sum(nil))
	return BaseHeader + "." + BaseData + "." + secret, nil
}

// 传入token判断是否有效
func AuthToken(token string) error {
	iter := strings.Split(token, ".")
	data, err := base64.StdEncoding.DecodeString(iter[1])
	if err != nil {
		return err
	}
	var payload Payload
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return err
	}
	// token是否过期
	if time.Now().Unix() > payload.Exp {
		return errors.New("token out of data")
	}
	message := iter[0] + iter[1]
	mac := hmac.New(sha256.New, []byte(AccessSecret))
	_, err = mac.Write([]byte(message))
	if err != nil {
		return err
	}
	secret := base64.URLEncoding.EncodeToString(mac.Sum(nil))
	// token是否有效
	if secret != iter[2] {
		return errors.New("token is wrong")
	} else {
		return nil
	}
}
