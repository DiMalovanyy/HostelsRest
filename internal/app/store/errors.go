package store

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEmptyData      = errors.New("empty data")
	ErrNoData         = errors.New("no data founded")
)
