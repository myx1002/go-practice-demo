// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Book is the golang structure of table book for DAO operations like Where/Data.
type Book struct {
	g.Meta      `orm:"table:book, do:true"`
	Id          interface{} // ID
	Name        interface{} // 书名
	Author      interface{} // 作者
	Price       interface{} // 价格
	PublishTime *gtime.Time // 出版时间
}
