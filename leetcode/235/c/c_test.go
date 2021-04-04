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
			`[1,7,5]`, `[2,3,5]`, 
			`3`,
		},
		{
			`[2,4,6,8,10]`, `[2,4,6,8,10]`, 
			`0`,
		},
		{
			`[1,10,4,4,2,7]`, `[9,3,5,1,7,4]`, 
			`20`,
		},
		// TODO 测试入参最小的情况
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minAbsoluteSumDiff, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-235/problems/minimum-absolute-sum-difference/
