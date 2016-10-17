package controller

import (
	. "com.izuanqian/echo/business"
	"com.izuanqian/echo/resource"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"strings"
)

func ListGroups(r render.Render, req *http.Request) {
	var po ListGroupPo
	CalRequestPo(req, &po)
	var token = req.Header.Get("Authorization")
	if !strings.Contains(token, "token_") {
		r.JSON(http.StatusUnauthorized, "{}")
		return
	}
	if strings.Split(token, "_")[1] != "sanlion" {
		r.JSON(http.StatusUnauthorized, "{}")
		return
	}
	var groups = QueryGroupsByCategory(po.CategoryId, po.LatestGroupId, po.View)
	r.JSON(http.StatusOK, groups)
}

func ListCategories(r render.Render) {
	r.JSON(
		http.StatusOK,
		[]CategoryItemVo{
			{Id: "001", Name: "category_city"},
			{Id: "002", Name: "category_train"},
			{Id: "003", Name: "category_ly"}})
}

func ListTopics(r render.Render, params martini.Params, req *http.Request) {

	var id = params["id"]
	fmt.Println(id)
	r.JSON(
		http.StatusOK,
		QueryRecommendTopics("token", "latestViewTopicId", 100))
}

func GetTopicById(r render.Render, params martini.Params) {
	var id = params["id"]
	r.JSON(
		http.StatusOK,
		res.TopicItemVo{
			Title:   "#" + string(id) + "#",
			Content: "content",
			Tags: []res.TagVo{
				{Type: TAG_GROUP_NAME, Value: "suzhou", Description: "group name"},
				{Type: TAG_GROUP_FOLLOW_COUNTS, Value: "29,100", Description: "follow counts"}}})
}

func CreateNewTopic(r render.Render, params martini.Params, req *http.Request) {
	var categoryId = params["categoryId"]
	var groupId = params["groupId"]
	var topicId = params["topicId"]
	fmt.Println(categoryId)
	fmt.Println(groupId)
	fmt.Println(topicId)
	var po TopicCreatePo
	CalRequestPo(req, &po)

	res.InsertSingleTopic(res.Topic{Id: "0001", Title: po.Title, Content: po.Content})
	r.JSON(
		http.StatusOK,
		res.TopicItemVo{
			Title:   po.Title,
			Content: po.Content,
			Tags: []res.TagVo{
				{Type: TAG_GROUP_NAME, Value: "suzhou", Description: "group name"},
				{Type: TAG_GROUP_FOLLOW_COUNTS, Value: "29,100", Description: "follow counts"}}})
}

func QueryTopicReplies(r render.Render, params martini.Params, req *http.Request) {
	var tid = params["id"]
	var po ReplyCreatePo
	CalRequestPo(req, &po)
	r.JSON(http.StatusOK, QueryReplies(tid, po.Content))
}

func FollowTopic(r render.Render, params martini.Params) {
	var categoryId = params["id"]
	r.JSON(http.StatusOK, "success, followed category:"+categoryId)
}

func UnFollowTopic(r render.Render, params martini.Params) {
	var categoryId = params["id"]
	r.JSON(http.StatusOK, "success, unfollowed category:"+categoryId)
}

func Login_check(login LoginPo) {
	if login.Account != "sanlion" {
		panic("invalid account info, please check")
	}
}
