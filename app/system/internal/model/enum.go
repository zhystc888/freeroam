package model

type GetEnumListDto struct {
	Code      string `description:"枚举编码"`
	TableName string `description:"表名"`
	Module    string `description:"模块名"`
}

type GetEnumListRes struct {
	Name string
	List []*EnumListItem
}

type EnumListItem struct {
	Id   string `json:"id" orm:"value"`
	Name string `json:"name" orm:"value_desc"`
}
