package biz

import (
	"fmt"
	"time"
)

/**
查询分类下的群组
*/
func QueryGroupsByCategory(cid string, latestGid string, view int) []GroupItem {

	groups := make([]GroupItem, 0)
	for i := 1; i <= view; i++ {
		groups = append(groups, GroupItem{Id: string(i), ImgUrl: "http://img.izuanqian.com/" + string(i) + ".jpg",
			Tags: []Tag{
				{Type: TAG_GROUP_NAME, Value: "suzhou", Description: "group name"},
				{Type: TAG_GROUP_FOLLOW_COUNTS, Value: "29,100", Description: "follow counts"}}})
	}
	return groups
}

func QueryReplies(tid string, content string) []Reply {
	return []Reply{
		{Content: "路由匹配可以通过正则表达式或者glob的形式",
			Time:   time.Now().Format(time.Kitchen),
			Author: "sanlion.do"},
		{Content: "路由处理器可以被相互叠加使用, 例如很有用的地方可以是在验证和授权的时候:",
			Time:   time.Now().Format(time.Kitchen),
			Author: "sanlion.do"}}
}

func QueryRecommendTopics(token string, latestViewTopicId string, view int) []SimpleTopicItem {

	fmt.Println("根据用户习惯等信息点推送热门讨论主题topic")
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
			Id:      "002",
			Title:   "titlea",
			Content: "contenta",
			Tags: []Tag{
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"}},
			Images: []string{
				"1.jpg",
				"2.jpg",
				"3.jpg"}},

		{
			Id:      "003",
			Title:   "titlea",
			Content: "contenta",
			Tags: []Tag{
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"}},
			Images: []string{}}}
	fmt.Println(topics)
	return topics
}
