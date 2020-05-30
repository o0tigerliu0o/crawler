package persist

import (
	"crawler/engine"
	"crawler/persist"
	"log"

	"github.com/olivere/elastic"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

// elsticSearch存储服务
func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.",item)
	if nil == err {
		*result = "ok"
		log.Printf("Error saving item %v:%v",item,err)
	}
	return err

}
