package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"strconv"
	"time"
)

const DefaultExpiredTimeInterval = 1 * 60 * 1000

type AuthToken struct {
	token               string
	createTime          int64 //毫秒时间戳
	expiredTimeInterval int64
}

type Option func(token *AuthToken)

//func CreateAuthToken(baseUrl string, createTime int64, params map[string]string) *AuthToken {
//	url := baseUrl
//	if len(params) > 0 {
//		var list []string
//		for k, v := range params {
//			list = append(list, k+"="+v)
//		}
//		url += "?" + strings.Join(list, "&")
//	}
//	return NewAuthToken(genToken(url), createTime)
//}

func GenerateAuthToken(baseUrl, appId, password string, timestamp int64) *AuthToken {
	url := baseUrl + "?appid=" + appId + "&pwd=" + password + "&ts=" + strconv.FormatInt(timestamp, 10)
	return NewAuthToken(genToken(url), timestamp)
}

func NewAuthToken(token string, createTime int64, opts ...Option) *AuthToken {
	a := &AuthToken{
		token:               token,
		createTime:          createTime,
		expiredTimeInterval: DefaultExpiredTimeInterval,
	}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func (a *AuthToken) GetToken() string {
	return a.token
}

func (a *AuthToken) IsExpired() bool {
	if time.Now().UnixNano()/1e6-a.createTime > a.expiredTimeInterval {
		return true
	}
	return false
}

func (a *AuthToken) Match(authToken *AuthToken) bool {
	return a.token == authToken.GetToken()
}

func genToken(data string) string {
	t := sha1.New()
	_, _ = io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
