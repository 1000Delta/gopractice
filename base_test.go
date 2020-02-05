package main

import "testing"

func TestBase(t *testing.T) {
	println(longestPalindromeC2B("123434333"))
}

func longestPalindromeForce(s string) string {
	// 记录：最大值、起始位置、长度
	// 查找方法：顺序查找。两个游标：回文起始游标（st）、偏移游标（of），每次匹配失败则右移偏移游标，否则记录长度，直到到达边界。从 0 到 of/2 ，比较串堆成字符是否相等，不相等直接偏移递增，比较全部通过则增加响应数据
	// 并发优化：goroutine
	start, mlen := 0, 0 // [start]len
	slen := len(s)
	countch := make(chan map[int]int, slen)

	for i := range s {
		go func(now int) {
			offset, count := 0, 1
			// 传输
			defer func() {
				tmax := make(map[int]int, 1)
				tmax[now] = count
				countch <- tmax
			}()
			for true {
				for i := offset; i >= offset/2; i-- {
					// 判否，则匹配下一组
					if s[now+i] != s[now+offset-i] {
						break
					}
					switch offset - i<<1 {
					case -1, 0, 1:
						count = offset + 1
						break
					}
				}
				// 递增
				offset++
				// 边界条件
				if now+offset >= slen {
					return
				}
			}
		}(i)
	}

	// 统计
	for range s {
		tmax := <-countch
		for ts, tl := range tmax {
			if tl > mlen {
				start, mlen = ts, tl
			}
		}
	}

	return s[start : start+mlen]
}

func longestPalindromeC2B(s string) string {
	// 记录：最大值、起始位置、长度
	// 查找方法：从回文中心开始，通过填充包围字符将回文全部扩充到奇数长度，然后判断从中间开始两边是否相等，返回最大的相等长度
	// 并发优化：goroutine
	center, mofst := 0, 0 // [start]len
	slen := len(s)
	countch := make(chan map[int]int, slen)

	for i := range s {
		go func(now int) {
			// 偏移量和长度
			ofst := 1
			// 传输
			defer func() {
				tmax := make(map[int]int, 1)
				tmax[now] = ofst - 1
				countch <- tmax
			}()
			// 查找算法
			// 奇数长度
			if now-1 >= 0 && now+1 < slen && s[now-1] == s[now+1] {
				for true {
					// 边界条件
					if now-ofst < 0 || now+ofst > slen-1 {
						return
					}
					// 回文判断
					if s[now-ofst] != s[now+ofst] {
						return
					}
					ofst++
				}
			} else if now+1 < slen && s[now] == s[now+1] {
				// 判左，可以省略判右
				for true {
					// 边界条件
					if now-ofst < 0 || now+1+ofst > slen-1 {
						return
					}
					// 回文判断
					if s[now-ofst] != s[now+1+ofst] {
						return
					}
					ofst++
				}
			}

		}(i)
	}

	// 统计
	for range s {
		tmax := <-countch
		for ts, tl := range tmax {
			if tl > mofst {
				center, mofst = ts, tl
			}
		}
	}

	start, end := center-mofst, 0
	if s[center] == s[center+1] {
		end = center + mofst + 1
	} else {
		end = center + mofst
	}

	return s[start : end+1]
}
