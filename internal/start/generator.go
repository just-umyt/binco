package start

import (
	"slices"

	"math/rand"
)

func (det *Det) GeneratorAndRemove() string {
	if len(det.StartBalls) == 1 {
		generatedBall := det.StartBalls[0]
		det.StartBalls = []string{}
		return generatedBall
	}

	randomInt := rand.Intn(len(det.StartBalls) - 1)
	generatedBall := det.StartBalls[randomInt]
	det.StartBalls = slices.Delete(det.StartBalls, randomInt, randomInt+1)

	return generatedBall
}
