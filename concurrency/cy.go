package concurrency

type WebsiteChecker func(string) bool

type Result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChan := make(chan Result)
	for _, url := range urls {
		go func(u string) {
			resultChan <- Result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChan
		results[r.string] = r.bool
	}
	return results
}
