package grains

import "fmt"

func Square(number int) (uint64, error) {
	if number < 1 || 64 < number {
		return 0, fmt.Errorf("Not a valid space: %v", number)
	}

	var total uint64 = 1
	for i := 2; i <= number; i++ {
		total *= 2
	}
	return total, nil
}

func Total() uint64 {
	var total uint64
	for s := 1; s <= 64; s++ {
		val, _ := Square(s)
		total += val
	}
	return total
}
