package concurrency

type checkResult struct {
	url   string
	valid bool
}

func CheckWebsites(checker func(url string) bool, websites []string) map[string]bool {
	output := map[string]bool{}
	c := make(chan checkResult)

	for _, url := range websites {
		go func(url string) {
			c <- checkResult{
				url:   url,
				valid: checker(url),
			}
		}(url)
	}

	for i := 0; i < len(websites); i++ {
		result := <-c
		output[result.url] = result.valid
	}
	return output
}
