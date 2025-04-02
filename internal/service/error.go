package service

import "errors"

var errFileExists = errors.New("File already exists")
var errFileNotFound = errors.New("File not found")
var errFileInvalidName = errors.New("File can contain only letters, digits, point and underscore")
var errFileInvalidType = errors.New("File can only be a picture of type jpg, png, jpeg")
