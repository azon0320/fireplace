package fireplacehook

import (
	"testing"
)

func TestHook(t *testing.T) {
	t.Error("nothing to test...")
	/*
		log := logrus.New()
		log.AddHook(NewFireplaceHook(&FireplaceHookConfig{
			Application:  "test_app",
			FireplaceURL: "http://localhost:8099",
		}))
		log.WithFields(logrus.Fields{
			"error":  fmt.Errorf("no error"),
			"status": "ok",
			"done":   1,
			"right":  true,
		}).Info("testInfo")
	*/
	// t.Error()
}
