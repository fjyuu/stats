package lib

import (
	"errors"
	"math"
)

// Calculator calculates sample statistics.
type Calculator struct {
	count uint
	mean  float64
	m     float64
	min   float64
	max   float64
	sum   float64
}

// NewCalculator returns a new Calculator.
func NewCalculator() *Calculator {
	return &Calculator{}
}

// Input inputs a sample value.
func (c *Calculator) Input(value float64) error {
	if c.count == 0 || value < c.min {
		c.min = value
	}
	if c.count == 0 || value > c.max {
		c.max = value
	}

	c.count++
	delta := value - c.mean
	c.mean += delta / float64(c.count)
	c.m += delta * (value - c.mean)
	c.sum += value

	return nil
}

// GetResult returns the result of statistics.
func (c *Calculator) GetResult() (*Result, error) {
	if c.count <= 0 {
		return nil, errors.New("requires at least one input value")
	}

	variance := c.m / float64(c.count-1)
	std := math.Sqrt(variance)
	return &Result{
		Count: c.count,
		Mean:  c.mean,
		Std:   std,
		Min:   c.min,
		Max:   c.max,
		Sum:   c.sum,
	}, nil
}
