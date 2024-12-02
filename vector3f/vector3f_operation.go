package vector3f

import (
	"math"

	"github.com/Lunarisnia/vecm"
)

// Dot product of two vector
func Dot(a vecm.Vector3f, b vecm.Vector3f) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// LengthSquared get the squared length of a vector is equivalent to Dot(a, a)
func LengthSquared(a vecm.Vector3f) float64 {
	return Dot(a, a)
}

// Length get the length of a vector
func Length(a vecm.Vector3f) float64 {
	return math.Sqrt(LengthSquared(a))
}

// Add add vector a + b
func Add(a vecm.Vector3f, b vecm.Vector3f) vecm.Vector3f {
	return vecm.Vector3f{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

// Subtract a - b
func Subtract(a vecm.Vector3f, b vecm.Vector3f) vecm.Vector3f {
	return vecm.Vector3f{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

// Multiply a * b
func Multiply(a vecm.Vector3f, b vecm.Vector3f) vecm.Vector3f {
	return vecm.Vector3f{
		X: a.X * b.X,
		Y: a.Y * b.Y,
		Z: a.Z * b.Z,
	}
}

// Scale a * scalar
func Scale(a vecm.Vector3f, scalar float64) vecm.Vector3f {
	return vecm.Vector3f{
		X: a.X * scalar,
		Y: a.Y * scalar,
		Z: a.Z * scalar,
	}
}

// Divide a / scalar
func Divide(a vecm.Vector3f, scalar float64) vecm.Vector3f {
	t := 1.0 / scalar
	return Scale(a, t)
}

// Normalize also known as Unit Vector
func Normalize(a vecm.Vector3f) vecm.Vector3f {
	length := Length(a)
	return Divide(a, length)
}
