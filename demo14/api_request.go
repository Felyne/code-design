package main

// 接口请求不一定是URL,所以叫ApiRequest更通用
type ApiRequest struct {
	baseUrl   string
	token     string
	appId     string
	timestamp int64
}

func BuildApiRequestFromFullUrl(url string) ApiRequest {
	// TODO
	return NewApiRequest("", "", "", 0)
}

func NewApiRequest(baseUrl, token, appId string, timestamp int64) ApiRequest {
	return ApiRequest{
		baseUrl:   baseUrl,
		token:     token,
		appId:     appId,
		timestamp: timestamp,
	}
}

func (a *ApiRequest) GetBaseUrl() string {
	return a.baseUrl
}

func (a *ApiRequest) GetToken() string {
	return a.token
}

func (a *ApiRequest) GetAppId() string {
	return a.appId
}

func (a *ApiRequest) GetTimestamp() int64 {
	return a.timestamp
}

func (a *ApiRequest) GetOriginUrl() string {
	return ""
}
