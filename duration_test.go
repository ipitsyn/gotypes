package gotypes

import (
	"encoding/json"
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	TimeInterval DurationAsSeconds `json:"timeInterval" yaml:"timeInterval"`
}

var inYAML = `
---
timeInterval: 5025
`

var inJSON = `
{
	"timeInterval": 5025
}
`

func TestDurationJSON(t *testing.T) {
	c := new(Configuration)

	json.Unmarshal([]byte(inJSON), c)
	d := c.TimeInterval.Duration
	t.Log("JSON Unmarshaled duration:", d)

	out, err := json.Marshal(c)
	if err != nil {
		t.Error("Failed to marshal back to JSON:", err)
	}
	t.Log("JSON Marshal result:", string(out))
}

func TestDurationYAML(t *testing.T) {
	c := new(Configuration)

	yaml.Unmarshal([]byte(inYAML), c)
	d := c.TimeInterval.Duration
	t.Log("YAML Unmarshaled duration:", d)

	out, err := yaml.Marshal(c)
	if err != nil {
		t.Error("Failed to marshal back to YAML:", err)
	}
	t.Log("YAML Marshal result:", string(out))
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
