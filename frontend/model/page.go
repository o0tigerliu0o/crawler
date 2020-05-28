package model

type SearchResult struct {
	Hits        int64
	Start       int
	Query       string // 用于存储输入的查询条件
	PrevFrom    int    // 用于上一页功能
	NextFrom    int    // 用于下一页功能
	CurrentPage int
	TotalPage   int64
	Items       []interface{}
	//Items []engine.Item
}
