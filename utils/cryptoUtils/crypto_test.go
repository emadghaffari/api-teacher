package cryptoutils

import "fmt"

func ExampleGetMD5() {
	res := GetMD5("ABCD")
	fmt.Println(res)
	// OutPut: cb08ca4a7bb5f9683c19133a84872ca7
}
