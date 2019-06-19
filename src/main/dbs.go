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
	CollectionUser string = "user"
	CollectionAddress string = "address"
	CollectionShops string = "shops"
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

//查询用户登录信息
func (p *DBInfo) findUser(username,password string) (UserData,error){
	s := p.session.Copy()
	defer s.Close()
	c := s.DB(p.dbName).C(CollectionUser)
	var userData UserData
	err := c.Find(bson.M{"username": username,"password": password}).One(&userData)
	if err != nil {
		debugLog.Printf("查询用户信息错误,err:%v\n",err)
		return userData,err
	}
	return userData,err
}

//查询地址信息
func (p *DBInfo) findAddress(latitude,longitude string)(AddressData,error) {
	s := p.session.Copy()
	defer s.Close()
	c := s.DB(p.dbName).C(CollectionAddress)
	var addressData AddressData
	err := c.Find(bson.M{"latitude": latitude,"longitude": longitude}).One(&addressData)
	if err != nil {
		debugLog.Printf("查询地址错误，err：%v\n",err)
		return addressData,err
	}
	return addressData,err
}

//查询商铺信息
func (p *DBInfo) findShops() ([]ShopsData,error){
	s := p.session.Copy()
	defer s.Close()
	c := s.DB(p.dbName).C(CollectionShops)
	var shopsData []ShopsData
	err := c.Find(nil).All(&shopsData)
	if err != nil {
		debugLog.Printf("查询商铺列表错误,err:%v\n",err)
		return nil,err
	}
	return shopsData,err
}

//查询搜索商铺信息
func (p *DBInfo) findSearchShops(val string) ([]ShopsData,error) {
	s := p.session.Copy()
	defer s.Close()
	c := s.DB(p.dbName).C(CollectionShops)
	var shopsData []ShopsData
	err := c.Find(bson.M{"name": val}).All(&shopsData)
	if err != nil {
		debugLog.Printf("查询搜索商铺列表错误,err:%v\n",err)
		return nil,err
	}
	return shopsData,err
}