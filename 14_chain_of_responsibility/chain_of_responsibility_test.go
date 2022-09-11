package chainofresponsibility

// ExampleChainOfResponsibility 请求者(client)
func ExampleChainOfResponsibility() {
	var alice *Support = NewNoSupport("Alice")
	var bob *Support = NewLimitSupport("Bob", 100)
	var charlie *Support = NewSpecialSupport("Charlie", 429)
	var diana *Support = NewLimitSupport("Diana", 200)
	var elmo *Support = NewOddSupport("Elmo")
	var fred *Support = NewLimitSupport("Fred", 300)
	// 形成责任链
	alice.SetNext(bob).SetNext(charlie).SetNext(diana).SetNext(elmo).SetNext(fred)
	// 制造各种问题
	for i := 0; i < 500; i += 33 {
		alice.Support(NewTrouble(i))
	}

	// Output:
	// [Trouble 0] is resolved by [Bob].
	// [Trouble 33] is resolved by [Bob].
	// [Trouble 66] is resolved by [Bob].
	// [Trouble 99] is resolved by [Bob].
	// [Trouble 132] is resolved by [Diana].
	// [Trouble 165] is resolved by [Diana].
	// [Trouble 198] is resolved by [Diana].
	// [Trouble 231] is resolved by [Elmo].
	// [Trouble 264] is resolved by [Fred].
	// [Trouble 297] is resolved by [Elmo].
	// [Trouble 330] cannot be resolved.
	// [Trouble 363] is resolved by [Elmo].
	// [Trouble 396] cannot be resolved.
	// [Trouble 429] is resolved by [Charlie].
	// [Trouble 462] cannot be resolved.
	// [Trouble 495] is resolved by [Elmo].
}
