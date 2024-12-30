package concurrency

type WebsiteChecker func(string) bool

//functions can be treated as varaibles, passes as arguments or returned as results. over here WebsiteChecker is alisas to
//any which takes string as parameter and returns a boolean value.
//functions can be assigned to varaibles, stored in data structure, returned and passes as parameters.

type result struct {
	string
	bool
	//both fields are anonymous. not recommended.
}

//operation is blocking : when it makes us wait for it to finish.
//An operation that does not block, is going to run in a goroutine(which is esentially a seperate thread)

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	//channel is a data structure to avoid race condition

	for _, url := range urls {
		//first value is the index and the second is the value itself
		go func() {
			resultChannel <- result{url, wc(url)}
			//sending result struct to result channel using send statement
			//the proceess of sending the result to the result channel is still happening
			//concurrently.
		}()

		//if we want to start a goroutine, we need to put the keyword go in front of a function call, and we often use
		//anonymous function for that.
	}

	for i := 0; i < len(urls); i++ {
		//recieve operation, channel on the right.
		r := <-resultChannel //r is going to be of the type struct reesult
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

//blocking flow :
//1.as the channel is unbuffered, when a goroutine sends to the channel, other goroutines block at the send operation, until the channel is empty again.
//so they wait for other thread to recieve the value inside the channel. when someone recieves the value from the channel, the goroutines which were block, send
//value to the channel.

//2.The main thread is blocked at revcieve operation, until there is a value available in the channel to receieve. when there is a value to recieve in the channel,
//the value is recieved by r varaible, when is then inserted into the map. this allows vales to be writeen in the map, in orderely fashion.
