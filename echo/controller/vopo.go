package controller

import "time"

type ListGroupPo struct {
	CategoryId    string
	View          int
	LatestGroupId string
}

type User struct {
	Nick string
}

type CategoryItemVo struct {
	Id    string    `json:"id"`
	Name  string    `json:"name"`
	Cover string    `json:"cover"`
	Time  time.Time `json:"time"`
}

const (
	DEVICE_TYPE_ANDROID = "android"
	DEVICE_TYPE_IOS     = "ios"
)

/**
基础参数，放在每次请求的head中
*/
type BasePo struct {
	Version        string `json:"version"`
	Token          string `json:"token"`
	City           string `json:"city"`
	Longitude      string `json:"lng"`
	Latitude       string `json:"lat"`
	DeviceType     string `json:"deviceType"`
	DeviceCode     string `json:"deviceCode"`
	PushDeviceCode string `json:"pushDeviceCode"`
}
type LoginPo struct {
	Account  string
	Password string
}

type FollowPo struct {
	CategoryId string `json:"categoryId"`
}

type UnFollowPo struct {
	CategoryId string `json:"categoryId"`
}

type TopicCreatePo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ReplyCreatePo struct {
	Content string `json:"content"`
}

type ReplyVo struct {
	Content string `json:"content"`
	Time    string `json:"time"`
	Author  string `json:"author"`
}
