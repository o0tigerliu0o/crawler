package worker

import (
	"crawler/engine"
)

type CrawlService struct {
}

func (CrawlService) Process(req Request,
	result *ParseResult) error {

	engineReq, err := DeserializeRequest(req)
	if nil != err {
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if nil != err {
		return err
	}

	*result = SerializeResult(engineResult)
	return nil
}
