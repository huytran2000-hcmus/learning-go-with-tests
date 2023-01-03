package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	type resultRecord struct {
		string
		bool
	}
	results := make(map[string]bool)
	resultChannel := make(chan resultRecord)

	for _, url := range urls {
		go func(url string) {
			resultChannel <- resultRecord{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		resRecord := <-resultChannel
		results[resRecord.string] = resRecord.bool
	}

	return results
}
