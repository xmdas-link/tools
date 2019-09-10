package common_model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Property struct {
	Name  string
	Value string
}

type Properties []Property

func (p *Properties) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, p)
	case string:
		if v != "" {
			return p.Scan([]byte(v))
		}
	default:
		return errors.New("not supported")
	}
	return nil
}

func (p Properties) Value() (driver.Value, error) {
	if len(p) == 0 {
		return nil, nil
	}
	return json.Marshal(p)
}

func (p *Properties) HasValue(name string) bool {
	if len(*p) == 0 {
		return false
	}

	for _, c := range *p {
		if c.Name == name && len(c.Value) > 0 {
			return true
		}
	}

	return false
}

func (p *Properties) GetValue(name string) string {
	for _, c := range *p {
		if c.Name == name {
			return c.Value
		}
	}
	return ""
}

func (p Properties) SetValue(key string, value string) Properties {
	for i, c := range p {
		if c.Name == key {
			p[i].Value = value
			break
		}
	}
	return p
}
