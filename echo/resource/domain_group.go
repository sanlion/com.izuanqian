package res

type GroupItemVo struct {
	Id     string
	ImgUrl string
	Tags   []TagVo
}

type TagVo struct {
	Type        int
	Value       string
	Description string
}

type TopicItemVo struct {
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []TagVo  `json:"tags"`
	Images  []string `json:"images"`
}

type TopicInfoVo struct {
}
