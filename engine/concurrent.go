package engine

type ConcurrentEnigne struct {
	Scheduler   Scheduler
	WorkerCount int
	// 用于存储数据的队列
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

// 实现者只要实现ReadyNotifer及Scheduler中的函数即可
type Scheduler interface {
	ReadyNotifer
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifer interface {
	WorkReady(chan Request)
}

// 并发版爬虫引擎
func (e *ConcurrentEnigne) Run(seeds ...Request) {
	out := make(chan ParseResult)
	// 配置调度器通道
	e.Scheduler.Run()

	// 开启WorkerCount个工作
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// 种子首先运行
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		// out等待接受ParseResult
		result := <-out
		// 打印出接收到的数据，以及个数。
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}
		// 分配任务
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
	return
}

// 存储URL、实行去掉重复URL的操作
var URLstore = make(map[string]bool)

func isDuplicate(url string) bool {
	if URLstore[url] {
		return true
	}
	URLstore[url] = true
	return false
}

// 工作函数，逻辑是 in通道接收到request，即会调用worker函数爬每一个
// request中的网址，用对应的解析器。 解析完成后，将ParseResult返回给通道out
func (e *ConcurrentEnigne) createWorker(in chan Request,
	out chan ParseResult, ready ReadyNotifer) {
	go func() {
		for {
			// 传递到调度器，提示可以开始工作
			ready.WorkReady(in)
			// 有任务到工作中
			request := <-in
			// 开始工作，分布式要改为rpc调用
			result, err := e.RequestProcessor(request)
			if nil != err {
				continue
			}
			// 工作结果返回
			out <- result
		}
	}()
}
