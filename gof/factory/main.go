package main

import factory "github.com/tanagaku/sample_go_learning/gof/factory/framework"

func startMain(factoryObject factory.FactoryInterface) {
	card1 := factoryObject.Create("TEST Taro")
	card2 := factoryObject.Create("Jiro")
	card3 := factoryObject.Create("TEST Hanako")
	card1.Use()
	card2.Use()
	card3.Use()
}

func main() {
	startMain(factory.NewIDCardFatry())
}
