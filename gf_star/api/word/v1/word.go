package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ProficiencyLevel uint

const (
	ProficiencyLevel1 ProficiencyLevel = iota + 1
	ProficiencyLevel2
	ProficiencyLevel3
	ProficiencyLevel4
	ProficiencyLevel5
)

type WordCreateReq struct {
	g.Meta             `path:"words" method:"post" sm:"创建" tags:"单词"`
	Word               string           `v:"required|length:1,20#请输入单词|单词长度为1-20位" dc:"单词"`
	Definition         string           `v:"required|length:1,100#请输入单词释义|单词释义长度为1-100位" dc:"单词定义"`
	ExampleSentence    string           `v:"required|length:1,100#请输入例句|例句长度为1-100位" dc:"例句"`
	ChineseTranslation string           `v:"required|length:1,100#请输入中文翻译|中文翻译长度为1-100位" dc:"中文翻译"`
	Pronunciation      string           `v:"required|length:1,100#请输入音标|音标长度为1-100位" dc:"发音"`
	ProficiencyLevel   ProficiencyLevel `v:"required|in:1,2,3,4,5#请输入熟练度|熟练度只能为1-5" dc:"熟练度，1最低，5最高"`
}

type WordCreateRes struct {
}

type WordUpdateReq struct {
	g.Meta             `path:"words/{id}" method:"put" sm:"更新" tags:"单词"`
	Id                 uint             `v:"required|length:1,20#请输入单词id|单词id长度为1-20位" dc:"单词id"`
	Word               string           `v:"required|length:1,20#请输入单词|单词长度为1-20位" dc:"单词"`
	Definition         string           `v:"required|length:1,100#请输入单词释义|单词释义长度为1-100位" dc:"单词定义"`
	ExampleSentence    string           `v:"required|length:1,100#请输入例句|例句长度为1-100位" dc:"例句"`
	ChineseTranslation string           `v:"required|length:1,100#请输入中文翻译|中文翻译长度为1-100位" dc:"中文翻译"`
	Pronunciation      string           `v:"required|length:1,100#请输入音标|音标长度为1-100位" dc:"发音"`
	ProficiencyLevel   ProficiencyLevel `v:"required|in:1,2,3,4,5#请输入熟练度|熟练度只能为1-5" dc:"熟练度，1最低，5最高"`
}

type WordUpdateRes struct {
}

type WordDeleteReq struct {
	g.Meta `path:"words/{id}" method:"delete" sm:"删除" tags:"单词"`
	Id     uint `v:"required|length:1,20#请输入单词id|单词id长度为1-20位" dc:"单词id"`
}

type WordDeleteRes struct {
}

type WordDetailReq struct {
	g.Meta `path:"words/{id}" method:"get" sm:"详情" tags:"单词"`
	Id     uint `v:"required|length:1,20#请输入单词id|单词id长度为1-20位" dc:"单词id"`
}

type WordDetailRes struct {
	Id                 uint             `json:"id"                 dc:"单词id"`
	Word               string           `json:"word"               dc:"单词"`
	Definition         string           `json:"definition"         dc:"单词定义"`
	ExampleSentence    string           `json:"exampleSentence"    dc:"例句"`
	ChineseTranslation string           `json:"chineseTranslation" dc:"中文翻译"`
	Pronunciation      string           `json:"pronunciation"      dc:"发音"`
	ProficiencyLevel   ProficiencyLevel `json:"proficiencyLevel"   dc:"熟练度，1最低，5最高"`
	CreatedAt          *gtime.Time      `json:"createdAt"          dc:"创建时间"`
	UpdatedAt          *gtime.Time      `json:"updatedAt"          dc:"更新时间"`
}

type WordListReq struct {
	g.Meta `path:"words" method:"get" sm:"列表" tags:"单词"`
	Page   int    `v:"min:1#请输入正确的页码" df:"1" dc:"页码"`
	Size   int    `v:"min:1#请输入正确的页大小" df:"10" dc:"页大小"`
	Word   string `dc:"单词"`
}

type WordListRes struct {
	List  []WordDetailRes `json:"list" dc:"列表"`
	Total int             `json:"total" dc:"总数"`
}

type WordRandomListReq struct {
	g.Meta `path:"words/random" method:"get" sm:"随机列表" tags:"单词"`
	Size   int `v:"min:1#请输入正确的页大小" df:"10" dc:"页大小"`
}

type WordRandomListRes struct {
	List []WordDetailRes `json:"list" dc:"列表"`
}

type SetLevelReq struct {
	g.Meta `path:"words/{id}/level" method:"patch"`
	Id     uint             `json:"id" v:"required"`
	Level  ProficiencyLevel `json:"level" v:"required|between:1,5"`
}

type SetLevelRes struct {
}
