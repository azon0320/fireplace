package fireplacehook

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestHook(t *testing.T) {
	t.Error("nothing to test...")

	// log := logrus.New()
	// log.AddHook(NewFireplaceHook(&FireplaceHookConfig{
	// 	Application:  "test_app",
	// 	FireplaceURL: "http://localhost:8999",
	// }))
	// log.WithFields(logrus.Fields{
	// 	"error":   fmt.Errorf("no error"),
	// 	"status":  "ok",
	// 	"done":    1,
	// 	"right":   true,
	// 	"Unicode": "九",
	// }).Info("testInfo")
	// t.Error()
}

func TestHookBeforeLog(t *testing.T) {
	flag := false
	log := logrus.New()
	log.AddHook(NewFireplaceHook(&FireplaceHookConfig{
		Application:  "test_app",
		FireplaceURL: "http://localhost:8999",
		BeforeLog: func(entry *logrus.Entry) bool {
			// if logrus.Fields contains model and its value is 'flag', then set flag to true
			if entry.Data["model"] == "flag" {
				flag = true
			}
			return true
		},
	}))
	log.WithField("model", "notflag").Info("model not flag")
	if flag != false {
		t.Errorf("when mode is not flag, flag must be %t but got %t", false, flag)
		return
	}
	log.WithField("model", "flag").Info("model flag")
	if flag != true {
		t.Errorf("when mode is flag, flag must be %t but got %t", true, flag)
		return
	}
}
