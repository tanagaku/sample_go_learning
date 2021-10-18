package factory

import "fmt"

type iDCardProduct struct {
	owner string
}

type IDCardFactory struct {
	*Factory
	owners []string
}

//NewIDCardFatry uc for initilizig IDCardFacory
func NewIDCardFatry() *IDCardFactory {
	idCardFactory := &IDCardFactory{
		Factory: &Factory{},
	}

	idCardFactory.factory = idCardFactory
	return idCardFactory
}

func (i *IDCardFactory) createProduct(owner string) Product {
	return newIDCartProduct(owner)
}

func (i *IDCardFactory) registerProduct(product Product) {
	i.owners = append(i.owners, product.getOwener())
}

func newIDCartProduct(owner string) *iDCardProduct {
	fmt.Printf("I'll create %s's card\n", owner)
	return &iDCardProduct{owner}
}

//User func for using card
func (i *iDCardProduct) Use() {
	fmt.Printf("I'll use %s's card\n", i.owner)
}

func (i *iDCardProduct) getOwener() string {
	return i.owner
}
