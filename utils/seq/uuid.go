package seq

import (
	"strings"

	"github.com/google/uuid"
)

func UUID() string {
	return uuid.New().String()
}

func UUIDShort() string {
	return strings.Replace(UUID(), "-", "", -1)
}
