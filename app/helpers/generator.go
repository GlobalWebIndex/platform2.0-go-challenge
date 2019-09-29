package helpers

import (
	"math/rand"
)

func GenerateRandomBool() bool {
	return rand.Intn(2) == 0
}

func NormalDistributionFactor() float64 {
	res := rand.NormFloat64()/8 + 0.5
	if res > 1.0 || res < 0 {
		res = NormalDistributionFactor()
	}
	return res
}

func GenerateNormalDistribution(desiredPeak float64, min float64, max float64) float64 {
	rangeMinMax := max - min
	res := NormalDistributionFactor()*rangeMinMax + min + (desiredPeak - rangeMinMax/2)
	if res < min || res > max {
		res = GenerateNormalDistribution(desiredPeak, min, max)
	}
	return res
}
