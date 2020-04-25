package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"imooc.com/ccmouse/u2pppw/crawler/engine"
	"github.com/pkg/errors"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		//elastic.SetURL("http://127.0.0.1:19200"),
		elastic.SetSniff(false),
	)

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	itemCount := 0
	go func() {
		item := <-out
		log.Printf("Item Saver: got item #%d: %v", itemCount, item)
		itemCount++

		err := Save(client, index, item)
		if err != nil {
			log.Printf("Item Saver: error saving item %v: %v", item, err)
		}

	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	//fmt.Printf("%+v", resp)
	return nil
}
