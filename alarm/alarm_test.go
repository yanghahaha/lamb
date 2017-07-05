package alarm

import "testing"

func Test_alarm(t *testing.T) {
	GetAlarm().Config("http://user-console.ejudata.com/lsa/send", "id", "key")
	msg, err := GetAlarm().SendAlarm("Alarm Test From AMS")
	if err != nil || msg != "ok" {
		t.Error(msg)
	}
	msg, err = GetAlarm().SendWarning("Warning Test From AMS")
	if err != nil || msg != "ok" {
		t.Error(msg)
	}
}
