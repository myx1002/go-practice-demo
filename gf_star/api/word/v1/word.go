package v1

import "github.com/gogf/gf/v2/frame/g"

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
