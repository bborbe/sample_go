package main

import (
	"context"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"os"
	"runtime"
	"strconv"
)

const (
	indexName = "test"
)

var (
	urlPtr      = flag.String("url", "", "elasticsearch url")
	usernamePtr = flag.String("user", "", "elasticsearch username")
	passwordPtr = flag.String("pass", "", "elasticsearch password")
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := do(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "failure: %v\n", err)
		os.Exit(1)
	}
}

type tweet struct {
	User     string
	Message  string
	Retweets int
}

func do(ctx context.Context) error {
	username := *usernamePtr
	if len(username) == 0 {
		return fmt.Errorf("parameter username missing")
	}
	password := *passwordPtr
	if len(password) == 0 {
		return fmt.Errorf("parameter password missing")
	}
	url := *urlPtr
	if len(url) == 0 {
		return fmt.Errorf("parameter url missing")
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(url),
		elastic.SetBasicAuth(username, password),
	)
	if err != nil {
		return fmt.Errorf("create client failed: %v", err)
	}
	if err := createOrOpenIndex(ctx, client, indexName); err != nil {
		return fmt.Errorf("open index %v failed: %v", indexName, err)
	}

	for i := 0; i < 100; i++ {
		t := tweet{
			User:     "olivere",
			Message:  "Take Five",
			Retweets: 0,
		}
		put1, err := client.Index().Index(indexName).Type("tweet").Id(strconv.Itoa(i)).BodyJson(t).Do(context.Background())
		if err != nil {
			return fmt.Errorf("create failed: %v", err)
		}
		if !put1.Created {
			fmt.Printf("create entry %d in index failed\n", i)
		}
	}
	fmt.Printf("create\n")
	return nil
}

func createOrOpenIndex(ctx context.Context, client *elastic.Client, index string) error {
	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		return fmt.Errorf("index %v not found: %v", index, err)
	}
	if !exists {
		result, err := client.CreateIndex(index).Do(ctx)
		if err != nil {
			return fmt.Errorf("create index %v failed: %v", index, err)
		}
		if !result.Acknowledged {
			return fmt.Errorf("create not acknowledged")
		}
		return nil
	}
	result, err := client.OpenIndex(index).Do(ctx)
	if err != nil {
		return fmt.Errorf("open index %v failed: %v", index, err)
	}
	if !result.Acknowledged {
		return fmt.Errorf("open not acknowledged")
	}
	return nil
}
