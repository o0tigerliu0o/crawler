package config

const (

	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	// elasticSearch 存储 rpc服务端口
	ItemSaverPort = 1234
	WorkerPort0   = 9000

	// elasticSearch存储 路径
	ElasticIndex = "dating_profile_distributed"

	// elasticSearch存储 rpc服务
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"
)
