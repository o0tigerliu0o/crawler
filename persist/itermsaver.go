/* Copyright 2018 Inc. All Rights Reserved. */

/* File : itermsaver */
/*
modification history
--------------------
2018-05-22 09:38 , by o0TigerLiu0o, create
*/
/*
DESCRIPTION
*/

package persist

import (
	"context"
	"crawler/engine"
	"log"

	"github.com/pkg/errors"

	"github.com/olivere/elastic"
)

func ItemSaver(index string) (item chan engine.Item, err error) {

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) (err error) {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().Index(index).
		Type(item.Type).BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
