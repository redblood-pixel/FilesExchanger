package service

import "errors"

var errInternalServer = errors.New("Internal server error")
var errFileNotFound = errors.New("File not found")
var errFileBadName = errors.New("Bad file name")
