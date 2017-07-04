package main

import "fmt"

const MAX int = 3

type imag struct {
	width  int
	height int
}

func main() {
	/*
		a := []int{10, 100, 200}
		var i int
		var ptr [MAX]*int
		for i = 0; i < MAX; i++ {
			ptr[i] = &a[i] // 整数地址赋值给指针数组
		}
		for i = 0; i < MAX; i++ {
			fmt.Printf("a[%d] = %d\n", i, *ptr[i])
		}

		b := []string{"123", "345", "567"}
		var ptrstr [MAX]*string
		for i = 0; i < MAX; i++ {
			ptrstr[i] = &b[i] // 字符地址赋值给指针数组
		}

		for i = 0; i < MAX; i++ {
			fmt.Printf("b[%d] = %s\n", i, *ptrstr[i])
		}
	*/
	/*
		var b [][]imag
		b = [][]imag{{{10, 11}, {100, 111}},
			{{200, 777}, {300, 666}}}
		var i, j int

		var ptr **imag
		var ptr1 *imag
		for i = 0; i < 2; i++ {
			for j = 0; j < 2; j++ {
				ptr1 = &b[i][j]
				ptr = &ptr1
			}
		}
		for i = 0; i < 2; i++ {
			for j = 0; j < 2; j++ {
				fmt.Println(**ptr)
			}
		}
			for i = 0; i < 2; i++ {
				for j = 0; j < 2; j++ {
					fmt.Println(b[i][j].height)
					fmt.Println(b[i][j].width)
				}
			}


				c := []imag{{10, 11}, {100, 111}}
				var ptr *imag
				for i = 0; i < 2; i++ {
					ptr = &c[i]
				}

				for i = 0; i < 2; i++ {
					fmt.Println(*ptr)
				}
	*/

	var p *int
	p = make(int, 100)

	var t *int
	fmt.Println(*t)
}
