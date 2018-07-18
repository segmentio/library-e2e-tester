package tester

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

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

func runscopeMessages(bucket, username string) ([]map[string]interface{}, error) {
	req, err := http.NewRequest("GET", "https://webhook-e2e.segment.com/buckets/"+bucket+"?limit=10", nil)
	if err != nil {
		return nil, errors.Wrap(err, "messages: could not create request")
	}
	auth := username + ":"
	authData := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", "Basic "+authData)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "messages: could not fetch")
	}

	var runscopeMessagesResponse []interface{}
	if err := readJSONBody(resp.Body, &runscopeMessagesResponse); err != nil {
		return nil, errors.Wrap(err, "messages: could not read json")
	}

	msgs := make([]map[string]interface{}, 0)

	for _, data := range runscopeMessagesResponse {
		var msg map[string]interface{}

		b := []byte(data.(string))
		if err := json.Unmarshal(b, &msg); err != nil {
			events.Log("message %{message}v: could not parse json %{err}v", data, err)
			continue
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}
