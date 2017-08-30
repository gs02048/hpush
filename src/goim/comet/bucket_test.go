package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestToken(t *testing.T) {
	token := "e53cb4d7c7bcfcd640ba43513dfa9921"
	userId, _ := strconv.ParseInt(token, 10, 64)
	fmt.Printf("userid:%d", userId)
}
