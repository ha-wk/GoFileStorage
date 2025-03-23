package models

import "time"

type User struct {
	Username string
	Password string
	Quota    int64         //todo check if any more we can have in this sruct???
	Used     int64
}

//check if originalname really required or not???

type FileMetadata struct {
	Filename    string
	Size        int64
	UploadedAt  time.Time
	OriginalName string
}