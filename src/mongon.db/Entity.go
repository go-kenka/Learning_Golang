package mongon_db

import "gopkg.in/mgo.v2/bson"

type Person struct {
	Id    bson.ObjectId `bson:"_id"json:"_id"`
	Name  string        `bson:"st_name"json:"st_name"` //bson:"name" 表示mongodb数据库中对应的字段名称
	Phone string        `bson:"st_phone"json:"st_phone"`
}
