package factory

type Product interface {
	Use()
	getOwener() string
}

type FactoryInterface interface {
	createProduct(owner string) Product
	registerProduct(Product)
	Create(owner string) Product
}

type Factory struct {
	factory FactoryInterface
}

//Create func
func (f *Factory) Create(owner string) Product {
	p := f.factory.createProduct(owner)
	f.factory.registerProduct(p)
	return p
}
