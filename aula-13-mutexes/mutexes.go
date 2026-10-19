package main

func main() {

}

//	Mutexes allow us to lock access to data. This ensures that we can control which goroutines can access certain data at which time.
//	Go's standard library provides built-in implementation of a mutex with the "sync.Mutex" type and its two methods:
//		.Lock()
//		.Unlock()
//	We can protect a block of code by surrounding it with a call to Lock and Unlock.

// func protected() {
// 	mux.Lock()
// 	defer mux.Unlock()

// 	//	The rest of the function is protected
// 	//	Any other calls to mux.Lock() will block

// }

//	Mutexes are powerful. Like most powerful things, they can also cause many bugs if used carelessly.

//	Maps are not thread-safe.
//	Maps are not safe for concurrent use. If you have multiple goroutines accessing the same map, and at least one of them is writing
//	to the map, you must lock your maps with a mutex.

//	RW Mutex
//	In addition to the Lock() and Unlock() functions, RW Mutex also has:
//		RLock()
//		RUnlock()
//	A RW (Read-Write) Mutex allows us to have multiple readers at the same time.
//	RW Mutex can help with performance if we have a read-intensive process.
//	Many goroutines can safely read from the map at the same time(multiple RLock() calls can happen simultaneously).
//	However, only one goroutine can hold a Lock() and all RLock()'s will also be excluded.
