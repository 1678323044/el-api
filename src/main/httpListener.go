package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpListener struct {
	config  *Config
	dbInfo  *DBInfo
}

func ListenHttpService(cfg *Config, dbo *DBInfo) {
	listenter := HttpListener{}

	listenter.dbInfo = dbo
	listenter.config = cfg
	listenter.listen(cfg.HttpPort)
}

/* 将端口号和函数绑定 封装监听功能 */
func (p *HttpListener) listen(port int) {
	sPort := fmt.Sprintf(":%d", port)
	http.HandleFunc("/", p.Router)
	http.ListenAndServe(sPort, nil)
}

/* 处理用户登录 */
func (p *HttpListener) handleLogin(w http.ResponseWriter, r *http.Request) {
	//访问控制允许全部来源 允许跨域
	w.Header().Set("Access-Control-Allow-Origin","*")

	username := r.FormValue("username")
	password := r.FormValue("password")

	checkResult := p.checkFields(username,password)
	if !checkResult {
		sErr := p.makeResult(1000,"缺少必要字段")
		w.Write([]byte(sErr))
		return
	}
	userData,err := p.dbInfo.findUser(username,password)
	if err != nil {
		sErr := p.makeResult(1001,"用户名或密码错误")
		w.Write([]byte(sErr))
		return
	}
	user := User{
		Code: 0,
		Data: userData,
	}
	buf,_ := json.Marshal(user)
	w.Write(buf)
}

/* 处理地址信息 */
func (p *HttpListener)handlePosition (w http.ResponseWriter, r *http.Request) {
	//访问控制允许全部来源 允许跨域
	w.Header().Set("Access-Control-Allow-Origin","*")

	query := r.URL.Query()
	latitude := query.Get("latitude")
	longitude := query.Get("longitude")

	checkResult := p.checkFields(latitude,longitude)
	if !checkResult {
		sErr := p.makeResult(1000,"缺少必要字段")
		w.Write([]byte(sErr))
		return
	}
	addressData,err := p.dbInfo.findAddress(latitude,longitude)
	if err != nil {
		sErr := p.makeResult(1001,"查询地址信息失败")
		w.Write([]byte(sErr))
		return
	}
	address := Address{
		Code: 0,
		Data: addressData,
	}
	buf,_ := json.Marshal(address)
	w.Write(buf)
}

/* 处理商铺列表 */
func (p *HttpListener) handleShops(w http.ResponseWriter, r *http.Request){
	//访问控制允许全部来源 允许跨域
	w.Header().Set("Access-Control-Allow-Origin","*")

	shopsData,err := p.dbInfo.findShops()
	if err != nil {
		sErr := p.makeResult(1001,"查询商铺列表失败")
		w.Write([]byte(sErr))
		return
	}
	shops := Shops{
		Code: 0,
		Data: shopsData,
	}
	buf,_ := json.Marshal(shops)
	w.Write(buf)
}

/* 处理搜索商铺列表 */
func (p *HttpListener)handleSearchShops (w http.ResponseWriter, r *http.Request) {
	//访问控制允许全部来源 允许跨域
	w.Header().Set("Access-Control-Allow-Origin","*")

	val := r.FormValue("val")
	checkResult := p.checkFields(val)
	if !checkResult {
		sErr := p.makeResult(1000,"缺少必要字段")
		w.Write([]byte(sErr))
		return
	}
	shopsData,err := p.dbInfo.findSearchShops(val)
	if err != nil {
		sErr := p.makeResult(1001,"查询搜索商铺列表失败")
		w.Write([]byte(sErr))
		return
	}
	shops := Shops{
		Code: 0,
		Data: shopsData,
	}
	buf,_ := json.Marshal(shops)
	w.Write(buf)
}

/* 处理商品分类 */
func (p *HttpListener) handleGoodsClass(w http.ResponseWriter, r *http.Request) {
	//访问控制允许全部来源 允许跨域
	w.Header().Set("Access-Control-Allow-Origin","*")

	categoryData,err := p.dbInfo.findGoodsClass()
	if err != nil {
		sErr := p.makeResult(1001,"查询商品分类失败")
		w.Write([]byte(sErr))
		return
	}
	category := Category{
		Code: 0,
		Data: categoryData,
	}
	buf,_ := json.Marshal(category)
	w.Write(buf)
}