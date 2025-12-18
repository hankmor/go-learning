package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func fetchURL(url string) error {
	time.Sleep(time.Second)
	if url == "bad-url" {
		return fmt.Errorf("failed to fetch %s", url)
	}
	fmt.Printf("Fetched %s\n", url)
	return nil
}

func main() {
	fmt.Println("=== errgroup Demo ===")

	g, ctx := errgroup.WithContext(context.Background())

	urls := []string{"url1", "url2", "bad-url", "url3"}

	for _, url := range urls {
		url := url // 避免闭包问题
		g.Go(func() error {
			// 如果 context 被取消，立即返回
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return fetchURL(url)
			}
		})
	}

	// 等待所有任务完成，返回第一个错误
	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("All URLs fetched successfully")
	}
}
