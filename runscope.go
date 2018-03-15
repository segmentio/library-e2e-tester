package tester

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// RunscopeMessagesResponse is the type for the runscope response for multiple messages.
type RunscopeMessagesResponse struct {
	Data []struct {
		UUID string
	}
}

// RunscopeMessageResponse is the type for the runscope response for a single message.
type RunscopeMessageResponse struct {
	Data struct {
		Request struct {
			Body string
		}
	}
}

func readJSONBody(r io.ReadCloser, v interface{}) error {
	defer r.Close()
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.Wrap(err, "could not read body")
	}
	if err := json.Unmarshal(body, &v); err != nil {
		return errors.Wrap(err, "could not parse json")
	}
	return nil
}

func runscopeMessages(bucket, token string) ([]map[string]interface{}, error) {
	req, err := http.NewRequest("GET", "https://api.runscope.com/buckets/"+bucket+"/messages", nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not create request")
	}
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch messages")
	}

	var runscopeMessagesResponse RunscopeMessagesResponse
	if err := readJSONBody(resp.Body, &runscopeMessagesResponse); err != nil {
		return nil, errors.Wrap(err, "messages: could not read json")
	}

	msgs := make([]map[string]interface{}, 0)

	for _, data := range runscopeMessagesResponse.Data {
		req, err := http.NewRequest("GET", "https://api.runscope.com/buckets/"+bucket+"/messages/"+data.UUID, nil)
		if err != nil {
			return nil, errors.Wrap(err, "could not create request")
		}
		req.Header.Add("Authorization", "Bearer "+token)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, errors.Wrap(err, "could not fetch message")
		}

		var runscopeMessageResponse RunscopeMessageResponse
		if err := readJSONBody(resp.Body, &runscopeMessageResponse); err != nil {
			return nil, errors.Wrap(err, "message: could not read json")
		}

		var msg map[string]interface{}
		if err := json.Unmarshal([]byte(runscopeMessageResponse.Data.Request.Body), &msg); err != nil {
			return nil, errors.Wrap(err, "message data: could not parse json")
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}
