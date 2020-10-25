package main

import "fmt"

//type UserServiceTest struct {
//}
//
//func (u *UserServiceTest) DoTest() bool {
//	//...
//  fmt.Println("UserServiceTest test.")
//	return true
//}
//
//func main() {
//	u := &UserServiceTest{}
//	if u.DoTest() {
//		fmt.Println("Test succeed.")
//	} else {
//		fmt.Println("Test failed")
//	}
//}

// 下面是控制反转(ioc)

type TestCase interface {
	DoTest() bool
	Run()
}

func NewTestCaseTemplate(t TestCase) *TestCaseTemplate {
	return &TestCaseTemplate{t}
}

type TestCaseTemplate struct {
	TestCase
}

func (t *TestCaseTemplate) Run() {
	if t.DoTest() {
		fmt.Println("Test succeed.")
	} else {
		fmt.Println("Test failed")
	}
}

var testCases []TestCase

func Register(t TestCase) {
	testCases = append(testCases, t)
}

func main() {
	Register(NewUserServiceTest())
	for _, t := range testCases {
		t.Run()
	}
}
