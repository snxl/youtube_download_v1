package entity

import "time"

type Video struct {
	Quality      string
	QualityLabel string
	AudioChannel bool
	Title        string
	Author       string
	PublishDate  time.Time
	Duration     time.Duration
}
