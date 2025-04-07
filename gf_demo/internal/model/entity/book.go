// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Book is the golang structure for table book.
type Book struct {
	Id          uint        `json:"id"           orm:"id"           description:"ID"`   // ID
	Name        string      `json:"name"         orm:"name"         description:"书名"`   // 书名
	Author      string      `json:"author"       orm:"author"       description:"作者"`   // 作者
	Price       float64     `json:"price"        orm:"price"        description:"价格"`   // 价格
	PublishTime *gtime.Time `json:"publish_time" orm:"publish_time" description:"出版时间"` // 出版时间
}
