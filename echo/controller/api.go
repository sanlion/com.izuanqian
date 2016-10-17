package controller

import (
	"com.izuanqian/base"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func Api(m *martini.ClassicMartini) {

	base.SayHello("system")
	api_account(m)
	api_group(m)

}

func api_group(m *martini.ClassicMartini) {
	m.Get("", ListCategories) // 分类列表
	m.Group("/:categoryId", func(r martini.Router) {
		r.Get("", ListGroups) // 分类：群组列表
		m.Group("/:groupId", func(r martini.Router) {
			r.Get("", ListTopics) // 分类：群组：主题列表
			m.Group("/:topicId", func(r martini.Router) {
				r.Get("", GetTopicById)
				r.Post("/:id/reply", func() {})
				r.Get("/:id/reply", QueryTopicReplies)
				r.Post("", CreateNewTopic)
			})
		})
		r.Post("/follow", FollowTopic)
		r.Post("/unfollow", UnFollowTopic)
	})

	//m.Get("/orm", Ｏrm_demo)
	m.NotFound(func(r render.Render) {
		r.JSON(http.StatusNotFound, baseVo{Msg: "opoos~, the page was lost."})
	})
}

type baseVo struct {
	Msg string `json:"msg"`
}

func api_account(m *martini.ClassicMartini) {

	m.Group("/account", func(r martini.Router) {

		r.Post("/login", func(r render.Render, req *http.Request) {
			defer func() {
				err := recover()
				if err != nil {
					r.JSON(http.StatusNonAuthoritativeInfo, err)
				}
			}()
			var login LoginPo
			CalRequestPo(req, &login)
			Login_check(login)
			r.JSON(http.StatusOK, login)
		})
	})
}
