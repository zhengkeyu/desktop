package persist

import (
	"testing"
	"imooc.com/ccmouse/u2pppw/crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
	"imooc.com/ccmouse/u2pppw/crawler/engine"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1314495053",
		Type: "zhenai",
		Id:   "1314495053",
		Payload: model.Profile{
			Name:       "风中的蒲公英",
			Gender:     "女",
			Marriage:   "离异",
			Age:        41,
			Height:     158,
			Weight:     48,
			Income:     "3001-5000元",
			Education:  "中专",
			Occupation: "公务员",
			Hukou:      "四川阿坝",
			Xingzuo:    "处女座",
			House:      "已购房",
			Car:        "未购车",
		},
	}
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false),
	)

	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// Save expected item
	err = Save(client, index, expected)

	if err != nil {
		t.Error(err)
	} else {
		// TODO: Try to start up elastic search
		// here using docker go client.
		client, err := elastic.NewClient(
			elastic.SetSniff(false),
		)
		if err != nil {
			panic(err)
		}

		resp, err := client.Get().
			Index(index).
			Type(expected.Type).
			Id(expected.Id).
			Do(context.Background())
		if err != nil {
			panic(err)
		}

		t.Logf("%s", *resp.Source)

		var actual engine.Item
		json.Unmarshal([]byte(*resp.Source), &actual)

		actualProfile, _ := model.FromJsonObj(actual.Payload)
		actual.Payload = actualProfile

		if actual != expected {
			t.Errorf("got %+v; expected %+v", actual, expected)
		}
	}
}
