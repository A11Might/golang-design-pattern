package factorymethod

func ExampleFactoryMethod() {
	factory := Factory{
		FactoryIface: NewIDCardFactory(),
	}
	var card1, card2, card3 Product
	card1 = factory.Create("小明")
	card2 = factory.Create("小红")
	card3 = factory.Create("小刚")
	card1.Use()
	card2.Use()
	card3.Use()

	// Output:
	// 制作小明的ID卡。
	// 制作小红的ID卡。
	// 制作小刚的ID卡。
	// 使用小明的ID卡。
	// 使用小红的ID卡。
	// 使用小刚的ID卡。
}
