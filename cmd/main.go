package main

import (
	"context"
	"log"

	"github.com/mr-chelyshkin/DataPipeline/stream"
)

func main() {
	st, err := stream.NewStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	st.Start()
}
