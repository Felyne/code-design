package main

import "fmt"

func NewUserServiceTest() *UserServiceTest {
	u := &UserServiceTest{}
	testCaseTemplate := NewTestCaseTemplate(u) // 父类需要持有子类的引用
	u.TestCaseTemplate = testCaseTemplate      // 子类需要匿名组合父类
	return u
}

type UserServiceTest struct {
	*TestCaseTemplate
}

func (u *UserServiceTest) DoTest() bool {
	//...
	fmt.Println("UserServiceTest test.")
	return true
}

//Register(&UserServiceTest{})
