package main

import "net/http"

/* 判断用户提交的地址并调用对应的处理函数 */
func (p *HttpListener) Router(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/position"{
		p.handlePosition(w,r)
	}
	if r.URL.Path == "/login" {
		p.handleLogin(w,r)
	}
	if r.URL.Path == "/shops" {
		p.handleShops(w,r)
	}
	if r.URL.Path == "/search" {
		p.handleSearchShops(w,r)
	}
	if r.URL.Path == "/category" {
		p.handleGoodsClass(w,r)
	}
}