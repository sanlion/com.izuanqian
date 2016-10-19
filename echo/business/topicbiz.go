package biz

import (
	"fmt"
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
