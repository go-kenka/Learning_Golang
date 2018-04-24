package mongon_db

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

/*
更新person数据,只更新第一条数据
 */
func UpdatePerson(query bson.M, change bson.M) string {
	exop := func(c *mgo.Collection) error {
		return c.Update(query, change)
	}
	err := getCollection("person", exop)
	if err != nil {
		return "true"
	}
	return "false"
}

/**
更新person数据通过Id更新
 */
func UpdatePersonById(id string, change bson.M) string {
	objId := bson.ObjectIdHex(id)

	exop := func(c *mgo.Collection) error {
		return c.UpdateId(objId, change)
	}
	err := getCollection("person", exop)
	if err != nil {
		return "true"
	}
	return "false"
}

/*
更新查找到的所有person数据
 */
func UpdatePersonAll(query bson.M, change bson.M) string {
	exop := func(c *mgo.Collection) error {
		info, err := c.UpdateAll(query, change)
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
