package mongon_db

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

/**
 * 插入一条数据person对象
 */
func InsertOnePerson(p Person) string {
	p.Id = bson.NewObjectId()
	query := func(c *mgo.Collection) error {
		return c.Insert(p)
	}
	err := getCollection("person", query)
	if err != nil {
		return "false"
	}
	return p.Id.Hex()
}

/**
 * 插入数组数据
 */
func InsertArrayPerson(persons []Person) string {

	query := func(c *mgo.Collection) error {
		// 获取bulk
		var bulk = c.Bulk()

		// 循环插入
		for index := 0; index < len(persons); index ++ {
			bulk.Insert(persons[index])
		}

		// 得到返回值
		result, err := bulk.Run()

		// 有error返回err，没有error返回nil
		if result == nil {
			return err
		} else {
			return nil
		}
	}
	// 执行query
	err := getCollection("person", query)
	if err != nil {
		return "false"
	}
	return "success"
}

/**
 * 插入多条数据
 */
func InsertMutiPerson() string {

	var person1 *Person = &Person{
		Id:    bson.NewObjectId(),
		Name:  "小吴",
		Phone: "18225254454",
	}

	var person2 *Person = &Person{
		Id:    bson.NewObjectId(),
		Name:  "小王",
		Phone: "18225254451",
	}

	query := func(c *mgo.Collection) error {
		return c.Insert(&person1, &person2)
	}

	err := getCollection("person", query)
	if err != nil {
		return "false"
	}
	return "success"
}
