package main

import "fmt"

/* 检查字段是否为空 */
func (p *HttpListener)checkFields(fields ...string) bool{
	for _,val := range fields {
		if val == "" {
			return false
		}
	}
	return true
}

/* 生成处理结果 */
func (p *HttpListener) makeResult (code int,msg string) string{
	return fmt.Sprintf(`{"code": %d,"msg": "%s"}`, code, msg)
}