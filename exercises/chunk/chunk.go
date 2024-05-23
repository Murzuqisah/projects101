/* Write a function called Chunk that receives as parameters a slice, slice []int, and a number size int.
The goal of this function is to chunk a slice into many sub slices where each sub slice has the length of size.

If the size is 0 it should print a newline ('\n').
*/

package main

import (
	"github.com/01-edu/z01"
)

func Chunk(slice []int, size int) {
	if size == 0 {
		z01.PrintRune('\n')
		return
	}
	if len(slice) == 0 {
		return
	}

	var result [][]int
	var buffer []int
	for i:= 0;i <=len(slice)-1;i++{
		if i % size == 0 && size != 0 {
			result = append(result,buffer)
			buffer = []int{}
			buffer = append(buffer,slice[i])
		}else {
			buffer = append(buffer,slice[i])
		}
		if len(buffer) > 0{
			result = append(result,buffer)
		}
		
	}

	// var i int
	// var end int
	// if size > 0 {
	// 	for i = 0; i < len(slice); i += size {
	// 		end = i + size
	// 		if end > len(slice) {
	// 			end = len(slice)
	// 		}
	// 		result = append(result, slice[i:end])
	// 	}

	// }

	z01.PrintRune('[')
	for i, chunk := range result {
		if i > 0 {
			z01.PrintRune(' ')
		}
		z01.PrintRune('[')
		for j, num := range chunk {
			if j > 0 {
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune(num + '0'))
		}
		z01.PrintRune(']')
	}
	z01.PrintRune(']')
	z01.PrintRune('\n')

}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
