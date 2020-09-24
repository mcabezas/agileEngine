package internal

import (
	uuid "github.com/satori/go.uuid"
)

func UUID() string {
	v4 := uuid.NewV4()
	return v4.String()
}

