package maps

import (
	"errors"
)

func Search(d map[string]string, key string) string {
	v, ok := d[key]
	if !ok {
		return ""
	}
	return v
}

type Dictionary map[string]string

var (
	ErrorNotFound  = errors.New("could not found key")
	ErrorWordExist = errors.New("word has exists")
)

func (self Dictionary) Search(key string) (string, error) {
	v, ok := self[key]
	if !ok {
		return "", ErrorNotFound
	}
	return v, nil
}

func (self Dictionary) Add(key string, value string) error {
	_, err := self.Search(key)
	switch err {
	case ErrorNotFound:
		self[key] = value
	case nil:
		return ErrorWordExist
	default:
		return err
	}

	return nil
}
