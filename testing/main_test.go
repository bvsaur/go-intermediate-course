package main

import "testing"

func TestSum(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{4, 5, 9},
		{10, 15, 25},
	}

	for _, table := range tables {
		if sum := Sum(table.a, table.b); sum != table.n {
			t.Errorf("Sum was incorrect, expected %d got %d", table.n, sum)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{5, 10, 10},
		{20, 45, 45},
		{30, 2, 30},
	}

	for _, table := range tables {
		if max := GetMax(table.a, table.b); max != table.n {
			t.Errorf("Incorrect max value, expected %d got %d", table.n, max)
		}
	}
}
