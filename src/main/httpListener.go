package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	Latitude string  `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Address struct {
	Code  int         `json:"code"`
	Data  AddressData `json:"address"`
}

func (p *HttpListener)handlePosition (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	query := r.URL.Query()
	if len(query) == 0 {
		fmt.Fprintf(os.Stdout,"经纬度参数不存在")
		return
	}
	latitude := query.Get("latitude")
	longitude := query.Get("longitude")
	address,err := p.dbInfo.findAddress(latitude,longitude)
	if err != nil {
		fmt.Fprintf(os.Stdout,"查询地址信息错误,%v\n",err)
		return
	}
	buf,err01 := json.Marshal(address)
	if err01 != nil {
		fmt.Fprintf(os.Stdout,"地址信息解析为json格式错误,%v\n",err01)
		return
	}
	w.Write(buf)
}