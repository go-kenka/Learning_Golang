package mongon_db

import "gopkg.in/mgo.v2"

// 定义全局变量  DB连接url
const URL = "localhost:27017" //mongodb连接字符串
// 定义全局变量  DB连接用户名
const USER = ""

// 定义全局变量  DB连接密码
const PASSWORD = ""

// 定义一个全局的Session
var mgoSession *mgo.Session

// 定义一个全局的Schema
var database = "Rxc"

/**
 *  获取MongoDB的session连接
 */
func getSession() *mgo.Session {
	// 在session为空的场合，创建session
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(URL)
		if err != nil {
			// 直接终止程序运行
			panic(err)
		}
	}

	// 设置模式为Monotonic
	mgoSession.SetMode(mgo.Monotonic, true)

	return mgoSession.Clone()
}

/**
 *  切换集合
 *  collectionNm ：  集合名
 *  query ： 语句
 */
func getCollection(collectionNm string, query func(collection *mgo.Collection) error) error {
	// 获取DB连接
	session := getSession()
	// 延迟关闭
	defer session.Close()
	// 获取Collection
	collection := session.DB(database).C(collectionNm)
	// 返回error
	return query(collection)
}
