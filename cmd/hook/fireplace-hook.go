package fireplacehook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/azon0320/fireplace/pkg/logentry"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// FireplaceHookConfig Fireplace logging hook config
type FireplaceHookConfig struct {
	// Application name
	Application string
	// Fireplace Server URL
	FireplaceURL string
	// before log fn, return if should log
	BeforeLog func(entry *logrus.Entry) bool
}

// FireplaceHook The hook which should be injected by logrus.New().AddHook()
type FireplaceHook struct {
	client *http.Client
	config *FireplaceHookConfig
}

// NewFireplaceHook new fireplace hook object
func NewFireplaceHook(config *FireplaceHookConfig) *FireplaceHook {
	return &FireplaceHook{
		client: &http.Client{Timeout: time.Second * 2},
		config: config,
	}
}

// Fire implemented from logrus.Hook
func (h *FireplaceHook) Fire(entry *logrus.Entry) error {
	var err error

	if h.config.BeforeLog != nil {
		canSend := h.config.BeforeLog(entry)
		if !canSend {
			return nil
		}
	}

	data := &logentry.CreateLogEntryRequest{
		Application: h.config.Application,
		Level:       entry.Level.String(),
		Time:        entry.Time.UTC().Format(time.RFC3339),
		Message:     entry.Message,
		Details:     make(logentry.LogEntryDetailItemCollection, 0, 10),
	}

	for key, value := range entry.Data {
		if errorData, isError := value.(error); logrus.ErrorKey == key && value != nil && isError {
			data.Details = append(data.Details, &logentry.LogEntryDetailItem{
				Key:   key,
				Value: errorData.Error(),
			})
		} else {
			data.Details = append(data.Details, &logentry.LogEntryDetailItem{
				Key:   key,
				Value: fmt.Sprintf("%v", value),
			})
		}
	}

	if err = h.send(data); err != nil {
		fmt.Printf("%s\n", err.Error())
		return errors.Wrapf(err, "Unable to write logrus entry in hook Fireplace Hook")
	}

	return nil
}

// Levels implemented from logrus.Hook
func (h *FireplaceHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

// send Send log data to fireplace server
func (h *FireplaceHook) send(entry *logentry.CreateLogEntryRequest) error {
	var err error
	var entryJSON []byte
	var response *http.Response

	if entryJSON, err = json.Marshal(entry); err != nil {
		return errors.Wrapf(err, "Error converting log entry to JSON")
	}

	reader := bytes.NewReader(entryJSON)

	if response, err = h.client.Post(h.config.FireplaceURL+"/logentry", "application/json", reader); err != nil {
		return errors.Wrapf(err, "Error sending log entry to Fireplace Server")
	} else if response.StatusCode != 200 {
		dat, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("err read response body: %s", err.Error())
		}
		return fmt.Errorf("err sending log entry to server: %s", string(dat))
	}

	defer response.Body.Close()
	return nil
}
