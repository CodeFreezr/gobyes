package toy

func GCD(smaller, larger uint64) uint64 {
	if larger < smaller {
		larger, smaller = smaller, larger
	}

	for {
		remainder := larger % smaller
		if remainder == 0 {
			return smaller
		}
		larger = smaller
		smaller = remainder
	}
}

func LCM(smaller, larger uint64) uint64 {
	return (larger / GCD(smaller, larger)) * smaller
}
