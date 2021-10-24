package main

import "sync"

/**
Custom "lazyLoad" vs "init"
-- "init": should be reserved for initialization of effectively immutable package-level state
-- "lazyLoad": call some initialization code exactly once after program launch time.
This is usually because the initialization is relatively slow and may not even
be needed every time your program runs
 */

type SlowComplicatedParser interface {
	ParseTest(string) string
}

var parser SlowComplicatedParser
/** like "sync.WaitGroup"
must make sure not to make a copy of an instance of sync.One
Because each copy has its own state to indicate whether or not it has already been used
=> Declaring a "sync.Once" instance inside a function is usually the wrong thing to do,
as a new instance will be created on every function call and there will be no memory of
previous invocations
*/
var once sync.Once


func Parse(dataToParse string) string  {
	once.Do(func() {//closure
		/**
		make sure that parser is only initialized once,
		=> we set the value of "parser" from within a closure that's
		passed to the "Do" method on "once".
		If "Parse" is called more than once.
		"once.Do" will not execute the closure again.
		 */
		parser = initParser()
	})
	return parser.ParseTest(dataToParse)
}

func initParser() SlowComplicatedParser {
	//do all sorts of setup and loading here
	return nil
}

func main() {

}
