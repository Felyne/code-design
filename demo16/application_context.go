package main

//饿汉式单例
var _instance = newApplicationContext()

func GetApplicationContextInstance() *applicationContext {
	return _instance
}
func newApplicationContext() *applicationContext {
	a := &applicationContext{}
	a.Init()
	return a
}

type applicationContext struct {
	alertRule    *AlertRule
	notification Notification
	alert        *Alert
}

func (a *applicationContext) Init() {
	a.alertRule = NewAlertRule()
	a.notification = NewNotification()
	a.alert = NewAlert()
	a.alert.AddAlertHandler(NewTpsAlertHandler(a.alertRule, a.notification))
	a.alert.AddAlertHandler(NewErrorAlertHandler(a.alertRule, a.notification))
}

func (a *applicationContext) GetAlert() *Alert {
	return a.alert
}
