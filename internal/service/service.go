package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

const morseChar = ".- "

var ErrEmptyString = errors.New("string must not be empty")

// StringDetection на основе проверки строки, преобразует строку в соответствующий тип
func StringDetection(s string) (string, error) {
	if s == "" {
		return "", ErrEmptyString
	}
	if isMorse(s) {
		return morse.ToText(s), nil
	}
	return morse.ToMorse(s), nil
}

// isMorse проверяет, состоит ли строка только из знаков Морзе.
func isMorse(str string) bool {
	return str != "" && strings.Trim(str, morseChar) == ""
}
