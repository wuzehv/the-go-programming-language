// 基准测试查看q2_5
package q2_3

var Pc [256]byte = func() (Pc [256]byte) {
	for i := range Pc {
		Pc[i] = Pc[i/2] + byte(i&1)
	}
	return
}()

func PopCount(x uint64) int {
	return int(Pc[byte(x>>(0*8))] +
		Pc[byte(x>>(1*8))] +
		Pc[byte(x>>(2*8))] +
		Pc[byte(x>>(3*8))] +
		Pc[byte(x>>(4*8))] +
		Pc[byte(x>>(5*8))] +
		Pc[byte(x>>(6*8))] +
		Pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var r byte
	for i := 0; i < 8; i++ {
		r += Pc[byte(x>>(i*8))]
	}
	return int(r)
}
