package utils

import (
	"testing"
)

func TestJudgeHostSuccess(t *testing.T) {
	hostA := "hoge"
	hostB := "hoge"

	if !JudgeHost(hostA, hostB) {
		t.Fatal("Error: hostA != hostB")
	}
	t.Log("Success!!")
}

func TestJudgeHostError(t *testing.T) {
	hostA := "hoge"
	hostB := "fuga"

	if JudgeHost(hostA, hostB) {
		t.Fatal("Error: hostA == hostB")
	}
	t.Log("Success!!")
}
