package main

func main() {
	_ = NewApiAuthenticator(&MemoryCredentialStorage{})
}
