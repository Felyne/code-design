package main

func main() {
	info := NewApiStateInfo("/v1/user", 1000, 10, 20)
	GetApplicationContextInstance().GetAlert().Check(info)
}
