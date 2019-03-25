package tester

import (
	"testing"
)

func TestSegmentEqual(t *testing.T) {
	testCases := []struct {
		Name  string
		A     map[string]interface{}
		B     map[string]interface{}
		Equal bool
	}{
		{
			Name: "ignored fields can differ",
			A: map[string]interface{}{
				"messageId":         "x",
				"timestamp":         "x",
				"receivedAt":        "x",
				"sentAt":            "x",
				"originalTimestamp": "x",
				"channel":           "x",
				"version":           "x",
				"projectId":         "x",
				"writeKey":          "x",
				"_metadata":         "x",
				"context": map[string]interface{}{
					"library": "x",
				},
			},
			B: map[string]interface{}{
				"messageId":         "y",
				"timestamp":         "y",
				"receivedAt":        "y",
				"sentAt":            "y",
				"originalTimestamp": "y",
				"channel":           "y",
				"version":           "y",
				"projectId":         "y",
				"writeKey":          "y",
				"_metadata":         "y",
				"context": map[string]interface{}{
					"library": "y",
				},
			},
			Equal: true,
		},
		{
			Name: ".context should not differ",
			A: map[string]interface{}{
				"context": map[string]interface{}{
					"language": "English",
				},
			},
			B: map[string]interface{}{
				"context": map[string]interface{}{
					"language": "English",
					"location": "US",
				},
			},
			Equal: false,
		},
		{
			Name: ".context can differ by library",
			A: map[string]interface{}{
				"context": map[string]interface{}{
					"language": "English",
					"library":  "x",
				},
			},
			B: map[string]interface{}{
				"context": map[string]interface{}{
					"language": "English",
					"library":  "y",
				},
			},
			Equal: true,
		},
		{
			Name: ".integrations cannot differ",
			A: map[string]interface{}{
				"integrations": map[string]interface{}{
					"Amplitude": true,
				},
			},
			B: map[string]interface{}{
				"integrations": map[string]interface{}{
					"Amplitude": false,
				},
			},
			Equal: false,
		},
		{
			Name: "empty values are equivalent to non-existent",
			A: map[string]interface{}{
				"type":         "track",
				"integrations": map[string]interface{}{},
				"context":      map[string]interface{}{},
			},
			B: map[string]interface{}{
				"type": "track",
			},
			Equal: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			equal, _ := SegmentEqual(tc.A, tc.B)
			if equal != tc.Equal {
				t.Errorf("Expected equality %v but was %v: \nA: %+v\nB: %+v\n", tc.Equal, equal, tc.A, tc.B)
			}
		})
	}
}
