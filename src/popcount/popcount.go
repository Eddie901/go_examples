package popcount

// pc[i] is the population count of i (number of set bits)
var pc [256]byte

func init() {
	for i := range pc {
		// i = (i / 2) * 2 + i % 2
		// to double a binary number shift left (last bit is 0)
		// so the popcount of i is popcount of i / 2 plus 1 if i % 2 == 1
		pc[i] = pc[i/2] + byte(i&1)
		//			fmt.Printf("pc[%d] is %d\n", i, pc[i])
	}
}

// Popcount returns the population count of x

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
