package main

import "fmt"

// 开闭原则的英文全称是 Open Closed Principle，简写为 OCP
// 对扩展开放、修改关闭

type NotificationEmergencyLevel string

const (
	SEVERE  NotificationEmergencyLevel = "SEVERE"
	URGENCY NotificationEmergencyLevel = "URGENCY"
	NORMAL  NotificationEmergencyLevel = "NORMAL"
	TRIVIAL NotificationEmergencyLevel = "TRIVIAL"
)

type Notification interface {
	Notify(level NotificationEmergencyLevel, msg string)
}

func NewNotification() Notification {
	return &DefaultNotification{}
}

type DefaultNotification struct {
}

func (d *DefaultNotification) Notify(level NotificationEmergencyLevel, msg string) {
	fmt.Println(level, msg)
}

type Rule struct {
}

func (r *Rule) GetMaxTps() int64 {
	return 0
}

func (r *Rule) GetMaxErrorCount() int64 {
	return 0
}

func (r *Rule) GetMaxTimeoutTps() int64 {
	return 0
}

func NewAlertRule() *AlertRule {
	return &AlertRule{
		apiRule: make(map[string]Rule),
	}
}

type AlertRule struct {
	apiRule map[string]Rule
}

func (a *AlertRule) AddApiRule(api string, rule Rule) {
	a.apiRule[api] = rule
}

func (a *AlertRule) GetMatchedRule(api string) *Rule {
	rule, ok := a.apiRule[api]
	if ok {
		return &rule
	}
	return nil
}

// 原代码
//type Alert struct {
//	rule         AlertRule
//	notification Notification
//}
//
//func (a *Alert) check(api string, requestCount, errorCount, durationOfSeconds, timeoutCount int64) {
//	tps := requestCount / durationOfSeconds
//	if tps > a.rule.GetMatchedRule(api).GetMaxTps() {
//		a.notification.Notify(URGENCY, "...")
//	}
//	if errorCount > a.rule.GetMatchedRule(api).GetMaxErrorCount() {
//		a.notification.Notify(SEVERE, "...")
//	}
//	timeoutTps := timeoutCount / durationOfSeconds
//	if timeoutTps > a.rule.GetMatchedRule(api).GetMaxTimeoutTps() {
//		a.notification.Notify(URGENCY, "...")
//	}
//
//}

func NewAlert() *Alert {
	return &Alert{}
}

type Alert struct {
	alertHandlers []AlertHandler
}

func (a *Alert) AddAlertHandler(h ...AlertHandler) {
	a.alertHandlers = append(a.alertHandlers, h...)
}

func (a *Alert) Check(info ApiStateInfo) {
	for _, h := range a.alertHandlers {
		h.Check(info)
	}
}

func NewApiStateInfo(api string, requestCount, errorCount, durationOfSeconds int64) ApiStateInfo {
	return ApiStateInfo{
		api:               api,
		requestCount:      requestCount,
		errorCount:        errorCount,
		durationOfSeconds: durationOfSeconds,
	}
}

type ApiStateInfo struct {
	api               string
	requestCount      int64
	errorCount        int64
	durationOfSeconds int64
}

func (a *ApiStateInfo) GetApi() string {
	return a.api
}

func (a *ApiStateInfo) GetRequestCount() int64 {
	return a.requestCount
}

func (a *ApiStateInfo) GetErrorCount() int64 {
	return a.errorCount
}

func (a *ApiStateInfo) GetDurationOfSeconds() int64 {
	return a.durationOfSeconds
}

type AlertHandler interface {
	Check(info ApiStateInfo)
}

func NewAlertHandlerTemplate(r *AlertRule, n Notification, h AlertHandler) *AlertHandlerTemplate {
	return &AlertHandlerTemplate{
		rule:         r,
		notification: n,
		AlertHandler: h,
	}
}

type AlertHandlerTemplate struct {
	rule         *AlertRule
	notification Notification
	AlertHandler
}

func NewTpsAlertHandler(r *AlertRule, n Notification) AlertHandler {
	tpsAlertHandler := &TpsAlertHandler{}
	ah := NewAlertHandlerTemplate(r, n, tpsAlertHandler)
	tpsAlertHandler.AlertHandler = ah
	return tpsAlertHandler
}

type TpsAlertHandler struct {
	*AlertHandlerTemplate
}

func (t *TpsAlertHandler) Check(info ApiStateInfo) {
	tps := info.GetRequestCount() / info.GetDurationOfSeconds()
	if tps > t.rule.GetMatchedRule(info.GetApi()).GetMaxTps() {
		t.notification.Notify(URGENCY, "...")
	}
}

func NewErrorAlertHandler(r *AlertRule, n Notification) AlertHandler {
	errorAlertHandler := &ErrorAlertHandler{}
	ah := NewAlertHandlerTemplate(r, n, errorAlertHandler)
	errorAlertHandler.AlertHandler = ah
	return errorAlertHandler
}

type ErrorAlertHandler struct {
	*AlertHandlerTemplate
}

func (t *ErrorAlertHandler) Check(info ApiStateInfo) {
	if info.GetErrorCount() > t.rule.GetMatchedRule(info.GetApi()).GetMaxErrorCount() {
		t.notification.Notify(SEVERE, "...")
	}
}
