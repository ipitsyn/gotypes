package gotypes

import (
	"errors"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// DurationAsNanoseconds
// Marshal / unmarshal time.Duration as Nanoseconds
type DurationAsNanoseconds struct {
	time.Duration
}

func (d *DurationAsNanoseconds) UnmarshalJSON(b []byte) error {
	return unmarshalDurationJSON(d, b)
}

func (d DurationAsNanoseconds) MarshalJSON() ([]byte, error) {
	return marshalDurationJSON(d)
}

func (d *DurationAsNanoseconds) UnmarshalYAML(v *yaml.Node) error {
	return unmarshalDurationYAML(d, v)
}

func (d DurationAsNanoseconds) MarshalYAML() (interface{}, error) {
	return marshalDurationYAML(d)
}

// DurationAsMicroseconds
// Marshal / unmarshal time.Duration as Microseconds
type DurationAsMicroseconds struct {
	time.Duration
}

func (d *DurationAsMicroseconds) UnmarshalJSON(b []byte) error {
	return unmarshalDurationJSON(d, b)
}

func (d DurationAsMicroseconds) MarshalJSON() ([]byte, error) {
	return marshalDurationJSON(d)
}

func (d *DurationAsMicroseconds) UnmarshalYAML(v *yaml.Node) error {
	return unmarshalDurationYAML(d, v)
}

func (d DurationAsMicroseconds) MarshalYAML() (interface{}, error) {
	return marshalDurationYAML(d)
}

// DurationAsMilliseconds
// Marshal / unmarshal time.Duration as Milliseconds
type DurationAsMilliseconds struct {
	time.Duration
}

func (d *DurationAsMilliseconds) UnmarshalJSON(b []byte) error {
	return unmarshalDurationJSON(d, b)
}

func (d DurationAsMilliseconds) MarshalJSON() ([]byte, error) {
	return marshalDurationJSON(d)
}

func (d *DurationAsMilliseconds) UnmarshalYAML(v *yaml.Node) error {
	return unmarshalDurationYAML(d, v)
}

func (d DurationAsMilliseconds) MarshalYAML() (interface{}, error) {
	return marshalDurationYAML(d)
}

// DurationAsSeconds
// Marshal / unmarshal time.Duration as Seconds
type DurationAsSeconds struct {
	time.Duration
}

func (d *DurationAsSeconds) UnmarshalJSON(b []byte) error {
	return unmarshalDurationJSON(d, b)
}

func (d DurationAsSeconds) MarshalJSON() ([]byte, error) {
	return marshalDurationJSON(d)
}

func (d *DurationAsSeconds) UnmarshalYAML(v *yaml.Node) error {
	return unmarshalDurationYAML(d, v)
}

func (d DurationAsSeconds) MarshalYAML() (interface{}, error) {
	return marshalDurationYAML(d)
}

// DurationAsMinutes
// Marshal / unmarshal time.Duration as Minutes
type DurationAsMinutes struct {
	time.Duration
}

func (d *DurationAsMinutes) UnmarshalJSON(b []byte) error {
	return unmarshalDurationJSON(d, b)
}

func (d DurationAsMinutes) MarshalJSON() ([]byte, error) {
	return marshalDurationJSON(d)
}

func (d *DurationAsMinutes) UnmarshalYAML(v *yaml.Node) error {
	return unmarshalDurationYAML(d, v)
}

func (d DurationAsMinutes) MarshalYAML() (interface{}, error) {
	return marshalDurationYAML(d)
}

// DurationAsHours
// Marshal / unmarshal time.Duration as Hours
type DurationAsHours struct {
	time.Duration
}

func (d *DurationAsHours) UnmarshalJSON(b []byte) error {
	return unmarshalDurationJSON(d, b)
}

func (d DurationAsHours) MarshalJSON() ([]byte, error) {
	return marshalDurationJSON(d)
}

func (d *DurationAsHours) UnmarshalYAML(v *yaml.Node) error {
	return unmarshalDurationYAML(d, v)
}

func (d DurationAsHours) MarshalYAML() (interface{}, error) {
	return marshalDurationYAML(d)
}

// *
// JSON functions
// *
func unmarshalDurationJSON(d interface{}, b []byte) error {
	n, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	switch v := d.(type) {
	case *DurationAsNanoseconds:
		v.Duration = time.Duration(n) * time.Nanosecond
	case *DurationAsMicroseconds:
		v.Duration = time.Duration(n) * time.Microsecond
	case *DurationAsMilliseconds:
		v.Duration = time.Duration(n) * time.Millisecond
	case *DurationAsSeconds:
		v.Duration = time.Duration(n) * time.Second
	case *DurationAsMinutes:
		v.Duration = time.Duration(n) * time.Minute
	case *DurationAsHours:
		v.Duration = time.Duration(n) * time.Hour
	default:
		panic("unsupported type passed in unmarshalDurationJSON")
	}
	return nil
}

func marshalDurationJSON(d interface{}) ([]byte, error) {
	var i int64
	switch v := d.(type) {
	case DurationAsNanoseconds:
		i = int64(v.Duration / time.Nanosecond)
	case DurationAsMicroseconds:
		i = int64(v.Duration / time.Microsecond)
	case DurationAsMilliseconds:
		i = int64(v.Duration / time.Millisecond)
	case DurationAsSeconds:
		i = int64(v.Duration / time.Second)
	case DurationAsMinutes:
		i = int64(v.Duration / time.Minute)
	case DurationAsHours:
		i = int64(v.Duration / time.Hour)
	default:
		panic("unsupported type passed in marshalDurationJSON")
	}
	return []byte(strconv.FormatInt(i, 10)), nil
}

// *
// YAML functions
// *
func unmarshalDurationYAML(d interface{}, node *yaml.Node) error {
	if node.Kind != yaml.ScalarNode {
		return errors.New("invalid YAML node kind: non-scalar")
	}
	n, err := strconv.ParseInt(node.Value, 10, 64)
	if err != nil {
		return err
	}
	switch v := d.(type) {
	case *DurationAsNanoseconds:
		v.Duration = time.Duration(n) * time.Nanosecond
	case *DurationAsMicroseconds:
		v.Duration = time.Duration(n) * time.Microsecond
	case *DurationAsMilliseconds:
		v.Duration = time.Duration(n) * time.Millisecond
	case *DurationAsSeconds:
		v.Duration = time.Duration(n) * time.Second
	case *DurationAsMinutes:
		v.Duration = time.Duration(n) * time.Minute
	case *DurationAsHours:
		v.Duration = time.Duration(n) * time.Hour
	default:
		panic("unsupported type passed in unmarshalDurationYAML")
	}
	return nil
}

func marshalDurationYAML(d interface{}) (interface{}, error) {
	var i int64
	switch v := d.(type) {
	case DurationAsNanoseconds:
		i = int64(v.Duration / time.Nanosecond)
	case DurationAsMicroseconds:
		i = int64(v.Duration / time.Microsecond)
	case DurationAsMilliseconds:
		i = int64(v.Duration / time.Millisecond)
	case DurationAsSeconds:
		i = int64(v.Duration / time.Second)
	case DurationAsMinutes:
		i = int64(v.Duration / time.Minute)
	case DurationAsHours:
		i = int64(v.Duration / time.Hour)
	default:
		panic("unsupported type passed in marshalDurationYAML")
	}
	return i, nil
}
