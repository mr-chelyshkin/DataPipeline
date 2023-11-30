package elastic

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
)

// TODO: Commit func.

// Output to elasticsearch.
type Output struct {
	client *elasticsearch.Client
}

// NewOutput initializes elasticsearch as Output.
func NewOutput(options ...OptionFunc) (*Output, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	output := &Output{
		client: client,
	}

	for _, option := range options {
		option(output)
	}
	return output, nil
}

func (o *Output) Push(ctx context.Context, index string, buffer bytes.Buffer) error {
	type msg struct {
		Message string `json:"message"`
	}
	m := msg{Message: buffer.String()}

	bb, err := json.Marshal(m)
	if err != nil {
		return err
	}

	var bbuf bytes.Buffer
	_, err = bbuf.Write(bb)
	if err != nil {
		return err
	}

	res, err := o.client.Index(index, &bbuf, o.client.Index.WithContext(context.Background()))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		fmt.Println("Error indexing document")
	} else {
		fmt.Println("Document indexed successfully")
	}
	return err
}
