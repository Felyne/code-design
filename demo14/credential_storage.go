package main

type CredentialStorage interface {
	GetPasswordByAppId(appId string) (string, error)
}

type MemoryCredentialStorage struct {
}

func (m *MemoryCredentialStorage) GetPasswordByAppId(appId string) (string, error) {
	return appId, nil
}
