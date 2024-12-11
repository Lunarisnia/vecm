package vector3f

import (
	"testing"

	"github.com/Lunarisnia/vecm"
	"github.com/stretchr/testify/assert"
)

func Test_Reflect(t *testing.T) {
	t.Run("Reflect #1", func(t *testing.T) {
		rayDirection := vecm.Vector3f{
			X: 1.0,
			Y: -1.0,
			Z: 0.0,
		}
		normal := vecm.Vector3f{
			X: 0.0,
			Y: 1.0,
			Z: 0.0,
		}
		dotProduct := Dot(rayDirection, normal)
		assert.Equal(t, float64(-1.0), dotProduct)

		reflected := Reflect(rayDirection, normal)
		assert.Equal(t, vecm.Vector3f{X: 1, Y: 1, Z: 0}, reflected)
	})
}
