package biz

import (
	"com.izuanqian/echo/dba"
	"strconv"
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

func QueryReplies(topicId string, content string) []Reply {
	return []Reply{
		{Content: "路由匹配可以通过正则表达式或者glob的形式",
			//Time:   time.Now().Format(time.Kitchen),
			Author: "sanlion.do"},
		{Content: "路由处理器可以被相互叠加使用, 例如很有用的地方可以是在验证和授权的时候:",
			//Time:   time.Now().Format(time.Kitchen),
			Author: "sanlion.do"}}
}

func PopTopicsByUser(token string, latestTopicId string, view int) []SimpleTopicItem {
	var topics = []SimpleTopicItem{
		{
			Id:      "001",
			Title:   "【话题】自己发起的",
			Content: "contenta",
			Tags: []Tag{
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"}}},

		{
			Id:      "002",
			Title:   "【话题】自己参与过回复的",
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
			Title:   "【话题】有人@的",
			Content: "contenta",
			Tags: []Tag{
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"}},
			Images: []string{}},
		{
			Id:    "005",
			Title: "【话题】根据个人信息画像由系统推送的",
			Content: "推荐规则：\n" +
				"1.根据地理位置推荐附近的话题\n " +
				"2.根据以往关注的兴趣点推荐相似的话题\n " +
				"3.根据好友中关注的话题推荐\n" +
				"4.系统配置的运营点",
			Tags: []Tag{
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"},
				{Type: 1, Description: "description", Value: "value"}},
			Images: []string{}}}
	return topics
}

// 通过topic.id查询详细信息
func GetTopicInfoById(topicId string) Topic {
	topicDBO := dba.FetchTopicDBOById(topicId)
	return Topic{
		Author:  topicDBO.Author,
		Date:    topicDBO.Date,
		Content: topicDBO.Content,
		Id:      topicDBO.Id,
		Images:  topicDBO.Images,
		Title:   topicDBO.Title,
		Tags: []Tag{
			{
				Description: "关注人数",
				Value:       strconv.Itoa(topicDBO.ReplyCounts),
				Type:        TAG_GROUP_FOLLOW_COUNTS}},
		LatestReply: Reply{
			Author:  topicDBO.LatestReply.Author,
			Content: topicDBO.LatestReply.Content,
			Time:    topicDBO.LatestReply.Time}}
}
