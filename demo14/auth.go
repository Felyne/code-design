package main

import (
	"errors"
)

type ApiAuthenticator interface {
	Auth(request ApiRequest) error
	AuthUrl(url string) error
}

func NewApiAuthenticator(cs CredentialStorage) ApiAuthenticator {
	return &DefaultApiAuthenticator{cs}
}

var _ ApiAuthenticator = &DefaultApiAuthenticator{}

type DefaultApiAuthenticator struct {
	credentialStorage CredentialStorage
}

func (a *DefaultApiAuthenticator) Auth(request ApiRequest) error {
	appId := request.GetAppId()
	token := request.GetToken()
	timestamp := request.GetTimestamp()
	baseUrl := request.GetBaseUrl()
	clientAuthToken := NewAuthToken(token, timestamp)
	if clientAuthToken.IsExpired() {
		return errors.New("token is expired")
	}
	password, err := a.credentialStorage.GetPasswordByAppId(appId)
	if err != nil {
		return err
	}
	serverAuthToken := GenerateAuthToken(baseUrl, appId, password, timestamp)
	if !serverAuthToken.Match(clientAuthToken) {
		return errors.New("token verification failed")
	}
	return nil
}

func (a *DefaultApiAuthenticator) AuthUrl(url string) error {
	apiRequest := BuildApiRequestFromFullUrl(url)
	return a.Auth(apiRequest)
}
