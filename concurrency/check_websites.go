package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)} // `[channel] <- [value]` known as "send statement"
		}()
	}

	for range urls {
		r := <-resultChannel       // `[var] := <-[channel]` known as "receive expression"
		results[r.string] = r.bool // results[url] = wc(url)
	}

	return results
}
