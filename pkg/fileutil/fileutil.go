package fileutil

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"path/filepath"
	"slices"
	"strings"
	"unicode/utf8"
)

var types = []string{".jpg", ".png", ".jpeg"}

func Validate(filename string) error {

	ext := filepath.Ext(filename)
	if pos := slices.Index(types, ext); pos == -1 {
		return errors.New("wrong ext")
	} else if utf8.RuneCountInString(filename) > 255 {
		return errors.New("filename is too large")
	}
	return nil
}

func GetName(filename string) string {
	ext := filepath.Ext(filename)
	return strings.TrimSuffix(filename, ext)
}

func HashFilename(filename string) string {
	h := sha256.New()
	ext := filepath.Ext(filename)
	h.Write([]byte(strings.TrimSuffix(filename, ext)))
	hashBytes := h.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString + ext
}
