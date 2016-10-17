package handler

import (
	"com.izuanqian/ObjViewer/db"
	//. "com.izuanqian/ObjViewer/domain"
	//"encoding/json"
	"fmt"
	//"github.com/garyburd/redigo/redis"
	"log"
)

type OrderCreateHandler struct {
}

//func (OrderCreateHandler) Do() {
//	fmt.Println("OrderCreateHandler doing...")
//	conn := db.GetOptConnection()
//	v, err := redis.String(conn.Do("BLPOP", "EVENT:ORD:CREATE", 10))
//	if err != nil {
//		log.Println(err)
//	}
//	if len(v) != 0 {
//
//		var obj OrderCreate
//		json.Unmarshal([]byte(v), &obj)
//		o, _ := json.Marshal(Order{Id: obj.Id, Price: obj.Price, Time: obj.Time})
//		log.Println("新建ORDER\n" + string(o))
//		conn.Do("SET", "ORDER:"+obj.Id, string(o))
//	}
//}

type OrderCompleteHandler struct {
}

//func (OrderCompleteHandler) Do() {
//	fmt.Println("OrderCompleteHandler doing...")
//	conn := db.GetOptConnection()
//	v, err := redis.String(conn.Do("BLPOP", "EVENT:ORD:COMPLETE", 10))
//	if err != nil {
//		fmt.Println(err)
//	}
//	if len(v) != 0 {
//		var obj OrderCompete
//		json.Unmarshal([]byte(v), &obj)
//		orderJson, _ := redis.String(conn.Do("GET", "ORDER:"+obj.Id))
//		if len(orderJson) == 0 {
//			fmt.Println(obj.Id + " not exists.")
//			return
//		}
//		var order Order
//		json.Unmarshal([]byte(orderJson), &order)
//		order.Success = true
//		o, _ := json.Marshal(order)
//		conn.Do("SET", "ORDER:"+obj.Id, string(o))
//	}
//}

func BLPOP() {
	fmt.Println("BLPOP doing...")
	conn := db.GetOptConnection()
	//v, err := redis.Values(conn.Do("BLPOP", "demo", 10))
	v := conn.Cmd("BLPOP", "demo", 10)
	if v.Err != nil {
		log.Println(v.Err.Error())
	}
	fmt.Println(v)
}
