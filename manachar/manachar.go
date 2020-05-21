/*
Author:	LazyboyChen7
Date:	2020-05-21
Description:
	Manachar 是一个寻找字符串中最长子回文串的算法。
	时间复杂度：O(n)
	空间复杂度：O(n)

	算法讲解：
		https://ethsonliu.com/2018/04/manacher.html
		https://blog.crimx.com/2017/07/06/manachers-algorithm/
	算法例题：
		https://leetcode-cn.com/problems/longest-palindromic-substring/
 */
package manachar

func Manachar(s string) (string, int) {
	slc,p := make([]byte, len(s)<<1+3), make([]int, len(s)<<1+3)
	slc[0],slc[1],slc[len(slc)-1] = '$','#','@'
	for i:=range s {
		slc[(i+1)<<1], slc[(i+1)<<1+1] = s[i], '#'
	}
	mxl,idx,mx,re := -1,0,0,0
	for i:=1;i<len(slc)-1;i++ {
		if i < mx {
			p[i] = min(p[idx<<1-i], mx-i)
		} else {
			p[i] = 1
		}
		for slc[i-p[i]] == slc[i+p[i]] {
			p[i]++
		}
		if mx < i + p[i] {
			mx,idx = i+p[i],i
		}
		if mxl < p[i]-1 {
			re,mxl = idx,p[i]-1
		}
	}
	return s[(re-p[re]+2)>>1-1:(re+p[re]-2)>>1],mxl
}

func min(a, b int) int {
	if a < b {return a}
	return b
}

func max(a, b int) int {
	if a < b {return b}
	return a
}