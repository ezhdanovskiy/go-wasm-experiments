package calculator

// Calculator представляет собой структуру для базовых математических операций
type Calculator struct{}

// Add выполняет сложение двух чисел
func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract выполняет вычитание
func (c *Calculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply выполняет умножение
func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}

// Divide выполняет деление
func (c *Calculator) Divide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
