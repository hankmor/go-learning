package google

import (
	"context"
	"context-demo/userip"
	"encoding/json"
	"net/http"
)

// Results 搜索结果
type Results []Result

// Result 一条搜索结果，包含 Title 和 URL
type Result struct {
	Title, URL string
}

func Search(ctx context.Context, query string) (Results, error) {
	// Prepare the Google Search API request.
	// google 自定义搜索 api

	// https://www.googleapis.com/customsearch/v1?key=AIzaSyDDc33B_3zWY3wRtPNS-8ikt0MD1Dl5Tis&cx=053d09c819fa04043&q=golang
	key := "AIzaSyDDc33B_3zWY3wRtPNS-8ikt0MD1Dl5Tis"
	req, err := http.NewRequest("GET", "https://www.googleapis.com/customsearch/v1?key="+key+"&cx=053d09c819fa04043", nil)
	// req, err := http.NewRequest("GET", "https://ajax.googleapis.com/ajax/services/search/web?v=1.0&key="+key+"&cx=053d09c819fa04043", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query() // 查询参数
	q.Set("q", query)    // 设置搜索关键字

	// If ctx is carrying the user IP address, forward it to the server.
	// Google APIs use the user IP to distinguish server-initiated requests
	// from end-user requests.
	if userIP, ok := userip.FromContext(ctx); ok {
		q.Set("userip", userIP.String())
	}
	req.URL.RawQuery = q.Encode()
	var results Results
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Parse the JSON search result.
		// https://developers.google.com/web-search/docs/#fonje
		var data struct {
			Results []struct {
				Title string `json:"title"`
				Link  string `json:"link"`
			} `json:"items"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return err
		}
		for _, res := range data.Results {
			results = append(results, Result{Title: res.Title, URL: res.Link})
		}
		return nil
	})
	// httpDo waits for the closure we provided to return, so it's safe to
	// read results here.
	return results, err
}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	c := make(chan error, 1)
	req = req.WithContext(ctx) // 赋值一个request，使用新的context
	go func() {                // 开启单独的goroutine执行f方法
		c <- f(http.DefaultClient.Do(req)) // 执行查询，将结果传递给f方法
	}()
	select {
	case <-ctx.Done(): // 表示请求被取消或超时
		<-c              // Wait for f to return.
		return ctx.Err() // 返回取消原因，或者超时说明
	case err := <-c: // f成功执行,err 为nil，或者 f 执行返回了 error
		return err
	}
}
