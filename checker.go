package tester

import (
	"github.com/go-test/deep"
)

// SegmentEqual returns true if the two Segment messages can be considered the same.
func SegmentEqual(a, b map[string]interface{}) bool {
	return deep.Equal(cleanMsg(a), cleanMsg(b)) == nil
}

func cleanMsg(m map[string]interface{}) map[string]interface{} {
	ignoredKeys := []string{"messageId", "timestamp", "receivedAt", "sentAt", "originalTimestamp", "channel", "version", "projectId", "writeKey"}
	m = delete(m, ignoredKeys...)
	if _, ok := m["context"].(map[string]interface{}); ok {
		m["context"] = delete(m["context"].(map[string]interface{}), "library")
	}
	// delete empty fields
	for k, v := range m {
		field, ok := v.(map[string]interface{})
		if ok {
			empty := true
			for _ = range field {
				empty = false
			}
			if empty {
				m = delete(m, k)
			}
		}
	}
	return m
}

// delete returns a copy of the input map with the given keys deleted.
func delete(m map[string]interface{}, keys ...string) map[string]interface{} {
	out := make(map[string]interface{})
	for k, v := range m {
		found := false
		for _, ignoredKey := range keys {
			if k == ignoredKey {
				found = true
				break
			}
		}
		if !found {
			out[k] = v
		}
	}
	return out
}
