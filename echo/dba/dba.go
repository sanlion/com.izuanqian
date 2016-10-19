package dba

import "time"

func FetchTopicDBOById(topicId string) TopicDBO {
	return TopicDBO{
		Id:          topicId,
		Images:      []string{"1.jpg", "2.jpg", "3.jpg"},
		ReplyCounts: 12,
		Title:       "【话题】推荐规则说明",
		Author:      "sanlion.do",
		Content: "推荐规则：\n" +
			"1.根据地理位置推荐附近的话题\n " +
			"2.根据以往关注的兴趣点推荐相似的话题\n " +
			"3.根据好友中关注的话题推荐\n" +
			"4.系统配置的运营点",
		LatestReply: ReplyDBO{
			Author:  "sanlion.do",
			Content: "大学生活好",
			Time:    time.Now()},
		Date: time.Now()}
}
