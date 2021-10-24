package main

/**
properly handle closed channels.
When that happens, we rely on something that looks like an error: reading a nil channel.
As we saw earlier, reading from or writing to a nil channel causes your code to hang forever.
While that is bad if it is triggered by a bug, you can use a nil channel to disable a case in a select.
When you detect that a channel has been closed, set the channel’s variable to nil.
The associated case will no longer run, because the read from the nil channel never returns a value:
*/

func closeChannel()  {
	in := make(chan int)
	in2 := make(chan int)
	done := make(chan struct{})
	// in and in2 are channels, done is a done channel.
	for {
		select {
		case _, ok := <-in:
			if !ok {
				in = nil // the case will never succeed again!
				continue
			}
			// process the v that was read from in
		case _, ok := <-in2:
			if !ok {
				in2 = nil // the case will never succeed again!
				continue
			}
			// process the v that was read from in2
		case <-done:
			return
		}
	}
}

func main() {

}
