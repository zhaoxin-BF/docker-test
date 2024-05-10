package docker_log

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

func GetLog() {
	// 创建 Elasticsearch 客户端
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.31.213:9200"), elastic.SetSniff(false), elastic.SetBasicAuth("elastic", "u6Xgdk5yvLnSXtvdO3kA"))
	if err != nil {
		log.Fatalf("创建 Elasticsearch 客户端时出错：%v", err)
	}
	fmt.Println(client)

	// 定义查询条件
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("container.id", "0917672ac32362f0c36fc0f4e8fe478532c16c7ad0f58290655af858ea86a35b"))
}
