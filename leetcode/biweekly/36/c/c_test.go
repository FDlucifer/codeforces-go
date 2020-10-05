// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`[3,8]`, `[4,7]`, 
			`[[3,0],[1,7]]`,
		},
		{
			`[5,7,10]`, `[8,6,8]`, 
			`[[0,5,0],[6,1,0],[2,0,8]]`,
		},
		{
			`[14,9]`, `[6,9,8]`, 
			`[[0,9,5],[6,0,3]]`,
		},
		{
			`[1,0]`, `[1]`, 
			`[[1],[0]]`,
		},
		{
			`[0]`, `[0]`, 
			`[[0]]`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, restoreMatrix, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-36/problems/find-valid-matrix-given-row-and-column-sums/
