package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Color string
	Size  string
}

type ProductBuilder struct {
	name  string
	price float64
	color string
	size  string
}

func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{}
}

func (b *ProductBuilder) Name(name string) *ProductBuilder {
	b.name = name
	return b
}

func (b *ProductBuilder) Price(price float64) *ProductBuilder {
	b.price = price
	return b
}

func (b *ProductBuilder) Color(color string) *ProductBuilder {
	b.color = color
	return b
}

func (b *ProductBuilder) Size(size string) *ProductBuilder {
	b.size = size
	return b
}

func (b *ProductBuilder) Build() *Product {
	return &Product{
		Name:  b.name,
		Price: b.price,
		Color: b.color,
		Size:  b.size,
	}
}

func main() {
	product := NewProductBuilder().Name("Product 1").Price(100).Color("Red").Size("Large").Build()
	fmt.Println(product)
}
