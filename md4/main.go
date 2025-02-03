package main

import (
	"fmt"
	"strconv"
)

func main() {
	inputMsg := ""
	msg := toBinary(inputMsg)
	msg += "1"

	for len(msg)%512 != 448 {
		msg += "0"
	}

	msgLen := fmt.Sprintf("%064b", len(inputMsg))
	msg += fmt.Sprintf("%064b", len(inputMsg))

	A := 0x67452301
	B := 0xEFCDAB89
	C := 0x98BADCFE
	D := 0x10325476

	fmt.Println(len(msg), msgLen, A, B, C, D)
	x := make([]int, 16)
	chunckSize := 512
	wordSize := 32
	for i := 0; i < len(msg); i += chunckSize {
		for j := 0; j < 16; j++ {
		
			// grab 32 bits of data
			start := i* chunckSize + j * wordSize
			end := start + wordSize

			x[j], _ = strconv.Atoi(msg[start:end])

		}

		AA := A
		BB := B
		CC := C
		DD := D


		// ROUND 1
		A = (A + f(B, C, D)+ x[0])  << 3
		B = (D + f(A, B, C)+ x[1])  << 7
		C = (C + f(D, A, B)+ x[2])  << 11
		D = (B + f(C, D, A)+ x[3])  << 19
		A = (A + f(B, C, D)+ x[4])  << 3
		B = (D + f(A, B, C)+ x[5])  << 7
		C = (C + f(D, A, B)+ x[6])  << 11
		D = (B + f(C, D, A)+ x[7])  << 19
		A = (A + f(B, C, D)+ x[8])  << 3
		B = (D + f(A, B, C)+ x[9])  << 7
		C = (C + f(D, A, B)+ x[10]) << 11
		D = (B + f(C, D, A)+ x[11]) << 19
		A = (A + f(B, C, D)+ x[12]) << 3
		B = (D + f(A, B, C)+ x[13]) << 7
		C = (C + f(D, A, B)+ x[14]) << 11
		D = (B + f(C, D, A)+ x[15]) << 19


		// ROUND 2

		A = (A + g(B, C, D)+ x[0] + 0x5A827999) << 3
		B = (D + g(A, B, C)+ x[4] + 0x5A827999) << 5
		C = (C + g(D, A, B)+ x[8] + 0x5A827999) << 9
		D = (B + g(C, D, A)+ x[12] + 0x5A827999) << 13
		A = (A + g(B, C, D)+ x[1] + 0x5A827999) << 3
		B = (D + g(A, B, C)+ x[5] + 0x5A827999) << 5
		C = (C + g(D, A, B)+ x[9] + 0x5A827999) << 9
		D = (B + g(C, D, A)+ x[13] + 0x5A827999) << 13
		A = (A + g(B, C, D)+ x[2] + 0x5A827999) << 3
		B = (D + g(A, B, C)+ x[6] + 0x5A827999) << 5
		C = (C + g(D, A, B)+ x[10] + 0x5A827999) << 9
		D = (B + g(C, D, A)+ x[14] + 0x5A827999) << 13
		A = (A + g(B, C, D)+ x[3] + 0x5A827999) << 3
		B = (D + g(A, B, C)+ x[7] + 0x5A827999) << 5
		C = (C + g(D, A, B)+ x[11] + 0x5A827999) << 9
		D = (B + g(C, D, A)+ x[15] + 0x5A827999) << 13


		// ROUND 3

		A = (A + h(B, C, D)+ x[0] + 0x6ED9EBA1) << 3
		B = (D + h(A, B, C)+ x[8] + 0x6ED9EBA1) << 9
		C = (C + h(D, A, B)+ x[4] + 0x6ED9EBA1) << 11
		D = (B + h(C, D, A)+ x[12] + 0x6ED9EBA1) << 15
		A = (A + h(B, C, D)+ x[2] + 0x6ED9EBA1) << 3
		B = (D + h(A, B, C)+ x[10] + 0x6ED9EBA1) << 9
		C = (C + h(D, A, B)+ x[6] + 0x6ED9EBA1) << 11
		D = (B + h(C, D, A)+ x[14] + 0x6ED9EBA1) << 15
		A = (A + h(B, C, D)+ x[1] + 0x6ED9EBA1) << 3
		B = (D + h(A, B, C)+ x[9] + 0x6ED9EBA1) << 9
		C = (C + h(D, A, B)+ x[5] + 0x6ED9EBA1) << 11
		D = (B + h(C, D, A)+ x[13] + 0x6ED9EBA1) << 15
		A = (A + h(B, C, D)+ x[3] + 0x6ED9EBA1) << 3
		B = (D + h(A, B, C)+ x[11] + 0x6ED9EBA1) << 9
		C = (C + h(D, A, B)+ x[7] + 0x6ED9EBA1) << 11
		D = (B + h(C, D, A)+ x[15] + 0x6ED9EBA1) << 15
		

		A = (A + AA)
		B = (B + BB)
		C = (C + CC)
		D = (D + DD)

	}

	// v := fmt.

	// fmt.Println( A, B, C, D)
}

func toBinary(msg string) string {
	res := ""

	for _, c := range msg {
		res += fmt.Sprintf("%08b", c)
	}

	return res
}

func f(x, y, z int) int {
	return (x & y) | ((^x) & z)
}

func g(x, y, z int) int {
	return (x & y) | (x & z) | (y & z)
}

func h(x, y, z int) int {
	return x ^ y ^ z
}
