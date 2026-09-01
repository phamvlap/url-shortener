package utility

import (
	"strings"

	"github.com/google/uuid"
)

func EnforceHTTPSProtocol(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "https://" + url
	}
	return url
}

func GenerateShortID() string {
	id := uuid.New().String()
	id = strings.ReplaceAll(id, "-", "")[:16]
	return id
}
