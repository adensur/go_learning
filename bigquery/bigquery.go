package main

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

func main() {
	ts := time.Now()
	ctx := context.Background()
	project := "voisey-feed-ranking"
	client, err := bigquery.NewClient(ctx, project)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	q := client.Query("SELECT * FROM `bigquery-public-data.covid19_public_forecasts.county_14d_historical_` LIMIT 100000;")
	it, err := q.Read(ctx)
	if err != nil {
		panic(err)
	}
	for {
		var row []bigquery.Value
		err = it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", row)
	}
	println(fmt.Sprintf("Elapsed time: %v", time.Since(ts)))
}