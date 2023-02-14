package model

import "errors"

var (
	ErrNotDBExecOne  = errors.New("not exec one")
	ErrNotDBExecMany = errors.New("not exec many")
	ErrNotFoundDb    = errors.New("not found")
)
