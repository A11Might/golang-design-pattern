package decorator

func ExampleDecorator() {
	var (
		b1 *Display
		b2 *Display
		b3 *Display
		b4 *Display
	)
	b1 = NewStringDisplay("Hello, world.")
	b2 = NewSideBorder(b1, '#')
	b3 = NewFullBorder(b2)
	b1.Show()
	b2.Show()
	b3.Show()

	b4 =
		NewSideBorder(
			NewFullBorder(
				NewFullBorder(
					NewSideBorder(
						NewFullBorder(
							NewStringDisplay("你好，世界。"),
						),
						'*',
					),
				),
			),
			'/',
		)
	b4.Show()

	// Output:
	// Hello, world.
	// #Hello, world.#
	// +---------------+
	// |#Hello, world.#|
	// +---------------+
	// /+------------------------+/
	// /|+----------------------+|/
	// /||*+------------------+*||/
	// /||*|你好，世界。|*||/
	// /||*+------------------+*||/
	// /|+----------------------+|/
	// /+------------------------+/
}
