package dba

import "time"

type TopicDBO struct {
	Id          string
	Title       string
	Content     string
	Images      []string
	Author      string
	Date        time.Time
	LatestReply ReplyDBO
	ReplyCounts int
}

type ReplyDBO struct {
	Content string
	Time    time.Time
	Author  string
}
