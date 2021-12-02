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
			`[1, 0]`, `[2, 3]`, `[5, 4, 3]`, `[8, 2, 6, 7]`, 
			`18`,
		},
		{
			`[0, 0]`, `[0, 0]`, `[5]`, `[26]`, 
			`0`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minCost, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-66/problems/minimum-cost-homecoming-of-a-robot-in-a-grid/