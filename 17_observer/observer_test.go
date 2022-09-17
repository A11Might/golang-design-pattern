package observer

func ExampleObserver() {
	generator := NewRandomNumberGenerator()
	observer1 := NewDigitObserver()
	observer2 := NewGraphObserver()
	generator.AddObserver(observer1)
	generator.AddObserver(observer2)
	generator.Execute()

	// Output:
	// DigitObserver:1
	// GraphObserver:*
	// DigitObserver:7
	// GraphObserver:*******
	// DigitObserver:7
	// GraphObserver:*******
	// DigitObserver:9
	// GraphObserver:*********
	// DigitObserver:1
	// GraphObserver:*
	// DigitObserver:8
	// GraphObserver:********
	// DigitObserver:5
	// GraphObserver:*****
	// DigitObserver:0
	// GraphObserver:
	// DigitObserver:6
	// GraphObserver:******
	// DigitObserver:0
	// GraphObserver:
	// DigitObserver:4
	// GraphObserver:****
	// DigitObserver:1
	// GraphObserver:*
	// DigitObserver:2
	// GraphObserver:**
	// DigitObserver:9
	// GraphObserver:*********
	// DigitObserver:8
	// GraphObserver:********
	// DigitObserver:4
	// GraphObserver:****
	// DigitObserver:1
	// GraphObserver:*
	// DigitObserver:5
	// GraphObserver:*****
	// DigitObserver:7
	// GraphObserver:*******
	// DigitObserver:6
	// GraphObserver:******
}
