package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func Convert(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", nil
	}
	if isMorseCode(input) {
		return morse.ToText(input), nil
	} else {
		return morse.ToMorse(input), nil
	}
}

func isMorseCode(s string) bool {
	for _, r := range s {
		switch r {
		case '.', '-', ' ', '/', '\n', '\r', '\t':
			continue
		default:
			return false
		}
	}
	return true
}
