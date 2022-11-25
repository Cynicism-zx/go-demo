package main

import (
	"time"
)

type Post struct {
	IsDraft     bool      // 1 byte
	Title       string    // 16 bytes
	ID          int64     // 8 bytes
	Description string    // 16 bytes
	IsDeleted   bool      // 1 byte
	Author      string    // 16 bytes
	CreatedAt   time.Time // 24 bytes
}
