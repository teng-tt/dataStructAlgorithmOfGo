package main

import (
	"fmt"
	"github.com/OneOfOne/xxhash"
)

// 使用号称计算速度最快的哈希 xxhash，我们直接用该库来实现哈希

// XXHash 实现hash函数
func XXHash(key []byte) uint64 {
	h := xxhash.New64()
	h.Write(key)
	return h.Sum64()
}

// 拿到哈希值之后，我们要对结果取余，方便定位到数组下标 index。
// 如果数组的长度为 len，那么 index = xxhash(key) % len
// 选择 2^x 作为数组长度有一个很好的优点，就是计算速度变快了
// 这样取余 % 操作将变成按位 & 操作
// 选择 2^x 长度会使得计算速度更快，但是相当于截断二进制后保留后面的 k 位，
// 如果存在很多哈希值的值很大，位数超过了 k 位，而二进制后 k 位都相同，那么会导致大片哈希冲突。
// 但存在很大哈希值的情况很少发生，大部分哈希值的二进制位数都不会超过 k 位，
// 因此Golang 使用了这种 2^x 长度作为哈希表的数组长度。
// 实际上 hash(key) % len 的分布是和 len 有关的，一组均匀分布的 hash(key) 在 len 是素数时才能做到均匀

func main() {
	keys := []string{"hi", "my", "friend"}
	for _, key := range keys {
		fmt.Printf("xxhash('%s') = %d\n", key, XXHash([]byte(key)))
	}
}