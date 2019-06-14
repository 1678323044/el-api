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

type AddressData struct {
	Addr string      `json:"address" bson:"address"` 
	City string      `json:"city"`
}

type Address struct {
	Code  int         `json:"code"`
	Data  AddressData `json:"address"`
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