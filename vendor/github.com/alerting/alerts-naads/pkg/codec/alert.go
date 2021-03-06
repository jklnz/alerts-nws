package codec

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"

	"github.com/alerting/alerts/pkg/cap"
	capxml "github.com/alerting/alerts/pkg/cap/xml"
	"github.com/golang/protobuf/jsonpb"
)

type Alert struct{}

func (c *Alert) Encode(value interface{}) ([]byte, error) {
	var alert *cap.Alert

	// Handle incoming type
	switch v := value.(type) {
	case *capxml.Alert:
		return json.Marshal(v)
	case *cap.Alert:
		alert = v
	case cap.Alert:
		alert = &v
	default:
		return nil, errors.New("Unknown type provided")
	}

	// Encode to JSON
	enc := jsonpb.Marshaler{}
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	err := enc.Marshal(w, alert)
	w.Flush()
	return buf.Bytes(), err
}

func (c *Alert) Decode(data []byte) (interface{}, error) {
	var alert capxml.Alert
	err := json.Unmarshal(data, &alert)
	return alert, err
}
