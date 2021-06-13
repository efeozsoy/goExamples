package main

import "fmt"

func main() {
	s := make([]int, 5)

	fmt.Println("Empty : ", s)

	s[0] = 25
	s[1] = 227
	s[2] = 9

	fmt.Println("set : ", s)
	fmt.Println("get : ", s[2])

	fmt.Println("len : ", len(s))

	s = append(s, 2)
	s = append(s, 98, 57)
	fmt.Println("appended : ", s)

	c := make([]int, len(s))
	copy(c, s)
	fmt.Println("copy : ", c)

	l := s[2:5]
	fmt.Println("slice 1 : ", l)

	l = s[:5]
	fmt.Println("slice 2 : ", l)

	l = s[:]
	fmt.Println("slice 3 ", l)

	l = s[1:]
	fmt.Println("slice 4 ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl : ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
