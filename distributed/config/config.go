package config

const (

	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	// elasticSearch存储 路径
	ElasticIndex = "dating_profile_distributed"

	// elasticSearch存储 rpc服务
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	// 单个woker爬取速度
	Qps = 20
)
