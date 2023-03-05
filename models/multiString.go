package models

import (
	"database/sql/driver"
	"errors"
	"strings"
)

type MultiString []string

func (s *MultiString) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("MultiString src does not hold byte value")
	}

	*s = strings.Split(string(bytes), ",")
	return nil
}

func (s MultiString) Value() (driver.Value, error) {
	if len(s) == 0 {
		return nil, nil
	}

	return strings.Join(s, ","), nil
}
