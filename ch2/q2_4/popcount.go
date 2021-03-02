// 基准测试查看q2_5
package q2_4

func PopCount(x uint64) (num int) {
	for ; x > 0; x = x >> 1 {
		num += int(x & 1)
	}

	return
}
