package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/segmentio/events"
)

func readJSONBody(r io.ReadCloser, v interface{}) error {
	defer r.Close()
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.Wrap(err, "could not read body")
	}
	if err := json.Unmarshal(body, &v); err != nil {
		return errors.Wrapf(err, "could not parse json: %s", body)
	}
	return nil
}

func GetWebhookMessages(bucket, basicAuthUsername string) ([]map[string]interface{}, error) {
	req, err := http.NewRequest("GET", "https://webhook-e2e.segment.com/buckets/"+bucket+"?limit=100", nil)
	if err != nil {
		return nil, errors.Wrap(err, "webhook: could not create request")
	}
	req.SetBasicAuth(basicAuthUsername, "")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "webhook: could not do request")
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("webhook: received non-200 status code %d", resp.StatusCode)
	}

	var rawMsgs []string
	if err := readJSONBody(resp.Body, &rawMsgs); err != nil {
		events.Log("could not read json with error: %{err}v", err)
		return nil, errors.Wrap(err, "webhook: could not read request")
	}

	msgs := make([]map[string]interface{}, len(rawMsgs))
	for _, rawMsg := range rawMsgs {
		var msg map[string]interface{}
		rawJSON := strings.Trim(rawMsg, "\"")
		err := json.Unmarshal([]byte(rawJSON), &msg)
		if err != nil {
			events.Log("could not parse json with error: %{err}v", err, events.Arg{Name: "raw_json", Value: string(rawJSON)})
			return nil, errors.Wrap(err, "webhook: could not read msg")
		}
		msgs = append(msgs, msg)
	}

	return msgs, nil
}
