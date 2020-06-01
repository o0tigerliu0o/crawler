#分布式版爬虫
##启动方式
###1 修改配置文件
crawler/distributed/config/config.go
###2 启动itemSaver  用于存储数据到elasticSearch
go run crawler/distributed/persist/server/itemSaver_server.go -port=1234
###3 启动worker 用于解析url(可起多个)
go run crawler/distributed/worker/server/woker_server.go -port=9000
go run crawler/distributed/worker/server/woker_server.go -port=9001
###4 主程序
go run crawler/distributed/main.go -itemsaver_host=":1234" -worker_host=":9000,:9001"
###5 启动页面查看结果
go run crawler/frontend/main.go
###6 登录页面搜索结果
http://127.0.0.1:8888/
![Image text](https://https://github.com/o0tigerliu0o/crawler/tree/master/frontend/view/image/readme.png)

##TODO
单机url去重效率低，无法保存之前去重的结果(之前放到map中) => 基于Kep-Value Store(如Redis)进行分布式去重