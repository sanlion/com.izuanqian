package biz

import "time"

type GroupItem struct {
	Id     string
	ImgUrl string
	Tags   []Tag
}

type Tag struct {
	Type        int
	Value       string
	Description string
}

type SimpleTopicItem struct {
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []Tag    `json:"tags"`
	Images  []string `json:"images"`
}

type Topic struct {
	Id          string
	Title       string
	Content     string
	Tags        []Tag
	Images      []string
	Author      string
	Date        time.Time
	LatestReply Reply
}

type Reply struct {
	Content string `json:"content"`
	Time    string `json:"time"`
	Author  string `json:"author"`
}
