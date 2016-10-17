package biz

import (
	"fmt"
	"time"
)

// 根据用户信息，获取推荐的话题
func QueryRecommendTopicsByUserToken(userId string, latestViewId string, view int) []SimpleTopicItem {

	var topics = []SimpleTopicItem{
		{
			Id:      "001",
			Title:   "【第一期】黄山一日游",
			Content: "contenta",
			Tags: []Tag{
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"}}},

		{
			Id:      "【第一期】黄山一日游",
			Title:   "titlea",
			Content: "contenta",
			Tags: []Tag{
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"}},
			Images: []string{
				"1.jpg",
				"2.jpg",
				"3.jpg"}}}
	fmt.Println(topics)
	return topics
}

// 通过id获取详细的主题信息
func GetTopicById1(topicId string) Topic {
	return Topic{
		Id:      "0001",
		Author:  "sanlion.do",
		Content: "【第一期】黄山一日游",
		Date:    time.Now(), Images: []string{"1.jpg", "2.jpg"},
		LatestReply: Reply{
			Content: "黄山真不是能去的地方",
			Author:  "jack"}}
}

// 根据用户信息，获取推荐的话题
func QueryTopicsByUser(userId string, latestViewId string, view int) []SimpleTopicItem {

	var topics = []SimpleTopicItem{
		{
			Id:      "001",
			Title:   "【第一期】黄山一日游",
			Content: "contenta",
			Tags: []Tag{
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"}}},

		{
			Id:      "【第一期】黄山一日游",
			Title:   "titlea",
			Content: "contenta",
			Tags: []Tag{
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"}},
			Images: []string{
				"1.jpg",
				"2.jpg",
				"3.jpg"}}}
	fmt.Println(topics)
	return topics
}

func Follow(topicId string, userId string) {

}

func UnFollow(topicId string, userId string) {

}
