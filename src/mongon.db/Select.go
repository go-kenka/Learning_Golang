package mongon_db

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

/**
 * 通过objectid获取一条记录
 */
func GetPersonById(id string) *Person {
	objid := bson.ObjectIdHex(id)
	person := new(Person)
	query := func(c *mgo.Collection) error {
		return c.FindId(objid).One(&person)
	}
	getCollection("person", query)
	return person
}

//获取所有的person数据
func PagePerson() []Person {
	var persons []Person
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&persons)
	}
	err := getCollection("person", query)
	if err != nil {
		return persons
	}
	return persons
}

/**
 * 执行查询，此方法可拆分做为公共方法
 * [SearchPerson description]
 * @param {[type]} collectionName string [description]
 * @param {[type]} query          bson.M [description]
 * @param {[type]} sort           string [description]“-name”降序 "name" 升序
 * @param {[type]} fields         bson.M [description]
 * @param {[type]} skip           int    [description]
 * @param {[type]} limit          int)   (results []Person [description], err error [description]
 */
func SearchPerson(collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int) (results []interface{}, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Sort(sort).Select(fields).Skip(skip).Limit(limit).All(&results)
	}
	err = getCollection(collectionName, exop)
	return
}