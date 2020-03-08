package client

import (
	"ch15/series"
	"testing"
)

func TestSeries(t *testing.T) {
	t.Log(series.GetFibonacci(5))
	//t.Log(series.square(3))
}
