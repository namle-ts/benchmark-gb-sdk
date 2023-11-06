package main

import (
	"github.com/growthbook/growthbook-golang"
	"os"
	"testing"
)

var gbInstance *growthbook.GrowthBook

func TestMain(m *testing.M) {
	gbInstance = getGBInstance()
	os.Exit(m.Run())
}

func BenchmarkEvaluateFeature(b *testing.B) {
	for i := 0; i < b.N; i++ {
		evaluateFeature(gbInstance, "123")
	}
}
