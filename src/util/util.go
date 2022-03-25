package util

import (
	"math/rand"
	"time"
)

// 工具包

func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmMNBVCXZLKJHGFDSAPOIUYTREWQ")
	res := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}
	return string(res)
}
