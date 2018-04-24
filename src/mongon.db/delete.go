package mongon_db

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

/*
删除person数据,只更新第一条数据
 */
func DeletePerson(query bson.M) string {
	exop := func(c *mgo.Collection) error {
		return c.Remove(query)
	}
	err := getCollection("person", exop)
	if err != nil {
		return "true"
	}
	return "false"
}

/**
删除person数据通过Id更新
 */
func DeletePersonById(id string) string {
	objId := bson.ObjectIdHex(id)

	exop := func(c *mgo.Collection) error {
		return c.RemoveId(objId)
	}
	err := getCollection("person", exop)
	if err != nil {
		return "true"
	}
	return "false"
}

/*
删除查找到的所有person数据
 */
func DeletePersonAll(query bson.M) string {
	exop := func(c *mgo.Collection) error {
		info, err := c.RemoveAll(query)
		if info == nil {
			return err
		} else {
			return nil
		}
	}
	err := getCollection("person", exop)
	if err != nil {
		return "true"
	}
	return "false"
}
