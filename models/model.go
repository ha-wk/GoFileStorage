package models

import "time"

type User struct {
	Username string
	Password string
	Quota    int64
	Used     int64
}

type FileMetadata struct {
	Filename    string
	Size        int64
	UploadedAt  time.Time
	OriginalName string
}