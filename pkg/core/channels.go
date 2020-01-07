// 2 types of channels
// 1. unbuffered  and 2. buffered
// Unbuffered, can only work when there is a received to accept the message
// So, sender has to wait until received in place. Forces a kinda syncrhonous behaviour
// Example: 	myChannel := make(chan int)

// Buffered, example, myChannel := make(chan int, 5). Pass in the number of buffers
// Does not block the sender. more async behaviour UNTIL it gets full.
// If you try to get data off a channel and there is no data on it, it will lock

package main

import (

)

