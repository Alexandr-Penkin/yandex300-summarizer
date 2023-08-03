package main

import (
	"context"
	"fmt"
	"log"
	"time"

	summarizer "github.com/Alexandr-Penkin/yandex300-summarizer/cmd"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s := summarizer.New("<token>")
	result, err := s.GetSummary(ctx, "https://habr.com/ru/articles/752368/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
