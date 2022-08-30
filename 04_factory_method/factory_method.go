package factorymethod

import "fmt"

// Factory 创建者(creator)
type Factory struct {
	FactoryImpl
}

func (f *Factory) Create(owner string) Product {
	p := f.CreateProduct(owner)
	f.RegisterProduct(p)
	return p
}

type FactoryImpl interface {
	CreateProduct(owner string) Product
	RegisterProduct(product Product)
}

// IDCardFactory 具体创建者(concrete factory)
type IDCardFactory struct {
	owners []string
}

func NewIDCardFactory() *IDCardFactory {
	return new(IDCardFactory)
}

func (i *IDCardFactory) CreateProduct(owner string) Product {
	return NewIDCard(owner)
}

func (i *IDCardFactory) RegisterProduct(product Product) {
	i.owners = append(i.owners, product.(*IDCard).GetOwner())
}

func (i *IDCardFactory) GetOwners() []string {
	return i.owners
}

// Product 产品
type Product interface {
	Use()
}

// IDCard 具体产品(concrete product)
type IDCard struct {
	owner string
}

func NewIDCard(owner string) *IDCard {
	fmt.Printf("制作%s的ID卡。\n", owner)
	return &IDCard{
		owner: owner,
	}
}

func (i *IDCard) Use() {
	fmt.Printf("使用%s的ID卡。\n", i.owner)
}

func (i *IDCard) GetOwner() string {
	return i.owner
}
