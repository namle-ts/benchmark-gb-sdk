package main

import (
	"github.com/growthbook/growthbook-golang"
	"time"
)

func main() {

}

func getGBInstance() *growthbook.GrowthBook {
	context := growthbook.NewContext().
		WithClientKey("sdk-z6H20L1y8f0zGfll").
		WithCacheTTL(10 * time.Second)

	gb := growthbook.New(context)
	gb.LoadFeatures(&growthbook.FeatureRepoOptions{
		AutoRefresh: true,
	})

	return gb
}

func evaluateFeature(gb *growthbook.GrowthBook, attVal string) bool {
	gb.WithAttributes(growthbook.Attributes{"id": attVal})
	rs := gb.EvalFeature("enable-dark-mode").On
	return rs
}
