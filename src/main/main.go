package main

import (
	"mongon.db"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"encoding/json"
)

func main() {
	//Insert()
	Search()
	//Update()
	//Delete()
}

/*
插入
 */
func Insert() {
	var person mongon_db.Person = mongon_db.Person{
		Id:    bson.NewObjectId(),
		Name:  "小王",
		Phone: "18225254451",
	}

	mongon_db.InsertOnePerson(person)

	mongon_db.InsertMutiPerson()

	var persons []mongon_db.Person = []mongon_db.Person{
		mongon_db.Person{
			Id:    bson.NewObjectId(),
			Name:  "小王",
			Phone: "18225254451",
		},
		mongon_db.Person{
			Id:    bson.NewObjectId(),
			Name:  "小李",
			Phone: "18225254425",
		},
		mongon_db.Person{
			Id:    bson.NewObjectId(),
			Name:  "小张",
			Phone: "18225254455",
		},
	}

	mongon_db.InsertArrayPerson(persons)
}

/*
检索
 */
func Search() {
	personOne := mongon_db.GetPersonById("5ade847f5f52100cf8f5fdcd")

	fmt.Printf("这个是通过ID查询到的用户名：%s，电话番号：%s\r\n", personOne.Name, personOne.Phone)

	persons := mongon_db.PagePerson()

	for index := 0; index < len(persons); index ++ {
		fmt.Printf("这是第%d个用户名：%s，电话番号：%s\r\n", index+1, persons[index].Name, persons[index].Phone)
	}

	/**
	* 执行查询，此方法可拆分做为公共方法
	* [SearchPerson description]
	* @param {[type]} collectionName string [description]
	* @param {[type]} query          bson.M [description]
	* @param {[type]} sort           bson.M [description]
	* @param {[type]} fields         bson.M [description]
	* @param {[type]} skip           int    [description]
	* @param {[type]} limit          int)   (results      []interface{}, err error [description]
	*/
	resultPersons, err := mongon_db.SearchPerson("person", bson.M{"st_name": "小王"}, "st_phone", nil, 0, 0)

	var person []mongon_db.Person

	if err != nil {
		fmt.Printf(err.Error())
	} else {
		result, err := bson.MarshalJSON(resultPersons)
		if err != nil {
			fmt.Printf(err.Error())
		} else {
			fmt.Println(string(result))

			json.Unmarshal(result, &person)
		}

		for index := 0; index < len(person); index ++ {
			fmt.Printf("这是第%d个用户名：%s，电话番号：%s\r\n", index+1, person[index].Name, person[index].Phone)
		}
	}
}

/*
更新
 */
func Update() {
	selector := bson.M{"st_name": "小王"}

	data := bson.M{"$set": bson.M{"st_phone": "182257585742"}}

	mongon_db.UpdatePerson(selector, data)

	data1 := bson.M{"$set": bson.M{"st_phone": "182257585744"}}

	mongon_db.UpdatePersonById("5ade847f5f52100cf8f5fdcd", data1)

	selector2 := bson.M{"st_name": "小王"}

	data2 := bson.M{"$set": bson.M{"st_phone": "182257585745"}}

	mongon_db.UpdatePersonAll(selector2, data2)
}

/*
删除
 */
func Delete() {
	selector := bson.M{"st_name": "小王"}

	mongon_db.DeletePerson(selector)

	mongon_db.DeletePersonById("5ade847f5f52100cf8f5fdcd")

	selector2 := bson.M{"st_name": "小王"}

	mongon_db.DeletePersonAll(selector2)
}
