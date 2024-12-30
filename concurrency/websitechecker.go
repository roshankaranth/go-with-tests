package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
	//both values are anonymous
}

//operation is blocking : when it makes us wait for it to finish.
//An operation that does not block, is going to run in a goroutine(which is esentially a seperate thread)

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	//channel is a data structure to avoid race condition

	for _, url := range urls {
		//results[url] = wc(url)
		go func() {
			resultChannel <- result{url, wc(url)}
			//sending result struct to result channel using send statement
			//the proceess of sending the result to the result channel is still happending
			//concurrently.
		}()

		//if we want to start a goroutine, we need to put the keyword go in frint of a function call, and we often use
		//anonymous function for that.
	}

	for i := 0; i < len(urls); i++ {
		//recieve operation, channel on the right.
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

//we have used DI where, the function for checking the website is the dependency, and we are using mocking
//to test the function.

//anonymous functions can be executed at the same time that they are declared. they maintan access to lexical scope
//that they are defined in.

//1. first issue was that our main thread(checkwebsite) finished before any of the gorotines could put the result into the map,
//therefore we got an empty map.

//2. second issue we face is the multiple goroutine, try to write to the same map, at the same time, (concurrent map writes) causing
//the error. this is a race condition.

//race condition occurs when two or more goroutine access shared resource concurrently and at least one of the access is a write,
//leading to unpredictable behaviour.
