// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`["daniel","daniel","daniel","luis","luis","luis","luis"]`, `["10:00","10:40","11:00","09:00","11:00","13:00","15:00"]`, 
			`["daniel"]`,
		},
		{
			`["alice","alice","alice","bob","bob","bob","bob"]`, `["12:01","12:00","18:00","21:00","21:20","21:30","23:00"]`, 
			`["bob"]`,
		},
		{
			`["john","john","john"]`, `["23:58","23:59","00:01"]`, 
			`[]`,
		},
		{
			`["leslie","leslie","leslie","clare","clare","clare","clare"]`, `["13:00","13:20","14:00","18:00","18:51","19:30","19:49"]`, 
			`["clare","leslie"]`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, alertNames, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-36/problems/alert-using-same-key-card-three-or-more-times-in-a-one-hour-period/
