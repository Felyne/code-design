package ocp

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

type Rule struct {
}

func (r Rule) getMaxTps() int64 {
	return 0
}

func (r Rule) getMaxErrorCount() int64 {
	return 0
}

type AlertRule struct {
}

func (a *AlertRule) getMatchedRule(api string) Rule {
	return Rule{}
}

func NewAlert(rule AlertRule, notification Notification) *Alert {
	return &Alert{
		rule:         rule,
		notification: notification,
	}
}

type Alert struct {
	rule         AlertRule
	notification Notification
}

func (a *Alert) check(api string, requestCount, errorCount, durationOfSeconds int64) {
	tps := requestCount / durationOfSeconds
	if tps > a.rule.getMatchedRule(api).getMaxTps() {
		a.notification.Notify(URGENCY, "...")
	}
	if errorCount > a.rule.getMatchedRule(api).getMaxErrorCount() {
		a.notification.Notify(SEVERE, "...")
	}
}
