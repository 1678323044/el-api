package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/* 定义结构体 保存数据库信息 */
type DBInfo struct {
	session *mgo.Session //数据库操作会话
	dbName  string       //数据库名
}

/* 设置集合常量 */
const (
	CollectionAddress string = "address"
)

func ConnToDB(cfg *Config) (*DBInfo, error) {
	session, err := mgo.Dial(cfg.DBUrl)
	if err != nil {
		fmt.Printf("连接数据库失败,%v\n",err)
	}
	//设置连接缓冲池的最大值
	session.SetPoolLimit(100)
	//给结构体初始化赋值 保存会话对象和数据库名
	dbo := DBInfo{session, cfg.DBName}
	return &dbo, nil
}

/* 关闭会话 */
func (p *DBInfo) Close() {
	p.session.Close()
}

func (p *DBInfo) findAddress(latitude,longitude string)(Address,error) {
	s := p.session.Copy()
	defer s.Close()
	c := s.DB(p.dbName).C(CollectionAddress)
	var address Address
	err := c.Find(bson.M{"data.latitude":latitude,"data.longitude":longitude}).One(&address)
	if err != nil {
		fmt.Printf("查询地址错误,err:%v\n",err)
		return address,err
	}
	return address,err
}