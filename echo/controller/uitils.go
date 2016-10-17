package controller

import (
	"encoding/json"
	"net/http"
)

func CalRequestPo(req *http.Request, po interface{}) {
	defer req.Body.Close()
	json.NewDecoder(req.Body).Decode(po)
}

func WrapBasePoFromHead(req *http.Request) BasePo {
	var po BasePo
	po.Token = req.Header.Get("token")
	po.Version = req.Header.Get("version")
	po.City = req.Header.Get("city")
	po.DeviceCode = req.Header.Get("deviceCode")
	po.DeviceType = req.Header.Get("deviceType")
	po.Latitude = req.Header.Get("lat")
	po.Longitude = req.Header.Get("lng")
	po.PushDeviceCode = req.Header.Get("pushDeviceCode")
	return po
}
