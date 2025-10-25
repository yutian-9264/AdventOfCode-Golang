package main

import (
	"bufio"
	"fmt"
	"os"
)

var colLength int = 10
var rowLength int = 10
var matrix [][]byte
var foundTimes int = 0

type Pair struct {
	I int
	J int
}

var p []Pair

func addPair(i int, j int) {
	p = append(p, Pair{I: i, J: j})
}

// XMAS 88 77 65 83
func main() {

	rows, cols := colLength, rowLength
	matrix = make([][]byte, rows)
	for i := range matrix {
		matrix[i] = make([]byte, cols)
	}
	//
	file, err := os.Open("day4.txt")
	if err != nil {
		fmt.Println("file open failed: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	index_i := 0

	// if scanner.Scan() {
	// 	line := scanner.Text()

	// 	for _, r := range line {
	// 		if unicode.IsLetter(r) {
	// 			col_count++
	// 		}
	// 	}
	// }

	// colLength = col_count

	// file.Seek(0, 0)

	for scanner.Scan() {
		line := scanner.Text()
		bytes := []byte(line)
		matrix[index_i] = bytes
		// matrix = append(matrix, []byte(bytes))

		index_i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("read file failed: ", err)
	}

	// rowLength = index_i

	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {
			fmt.Println("main", i, j)
			if matrix[i][j] == 88 {
				fmt.Println("found X: ", i, j)
				p = nil
				if a1, b1, c1 := check(i, j, 77); a1 {
					fmt.Println("found XM: ", i, j, "*", "*")
					for _, p_one := range p {
						i_next2, j_next2 := find1of8(i, j, p_one.I, p_one.J)
						if i_next2 < 0 || j_next2 < 0 {
							continue
						}
						fmt.Println("p_one:", p_one.I, p_one.J)
						if a2, b2, c2 := check(i_next2, j_next2, 65); a2 {
							fmt.Println("found XMA: ", i, j, b1, c1, b2, c2)
							fmt.Println("i_next2:", i_next2, j_next2, b2, c2)
							i_next3, j_next3 := find1of8(p_one.I, p_one.J, b2, c2)
							fmt.Println("i_next3:", i_next3, j_next3)
							if i_next3 < 0 || j_next3 < 0 {
								continue
							}
							if a3, b3, c3 := check(i_next3, j_next3, 83); a3 {
								foundTimes++
								fmt.Println("found XMAS: ", i, j, b1, c1, b2, c2, b3, c3, i_next2, j_next2)
								continue
							}
						}
					}

				}
			}
		}
	}

	// check(1, 2, 88)
	// check(0, 2, 77)
	// check(3, 0, 77)

	fmt.Println("foundTimes: ", foundTimes)

	fmt.Println(matrix)
}

func check(i, j int, m byte) (ifFound bool, i_found, j_found int) {
	// if m == 65 {
	// 	fmt.Println("begin seach A", i, j)
	// }
	// var i_top, i_down, j_left, j_right int
	// fmt.Println("check: ", i, j, m)
	// fmt.Println("i == rowLength-1:", i == rowLength-1)
	// fmt.Println("j == colLength-1:", j == colLength-1)
	if i > 0 && i < rowLength-1 && j > 0 && j < colLength-1 {
		//don't lie at four borders
		if a, b, c := search(i, i-1, i+1, j, j-1, j+1, m); a {
			return true, b, c
		} else {
			return false, 0, 0
		}

	} else {
		// if i == 0 {
		// 	i_top = 0
		// }else if i == rowLength {
		// 	i_down =
		// }

		// lie at four corners
		if i == 0 && j == 0 {
			if a, b, c := search(i, i, i+1, j, j, j+1, m); a {
				return true, b, c
			} else {
				return false, 0, 0
			}
		}
		if i == rowLength-1 && j == 0 {
			if a, b, c := search(i, i-1, i, j, j, j+1, m); a {
				return true, b, c
			} else {
				return false, 0, 0
			}
		}
		if i == 0 && j == colLength-1 {
			if a, b, c := search(i, i, i+1, j, j-1, j, m); a {
				return true, b, c
			} else {
				return false, 0, 0
			}
		}
		if i == rowLength-1 && j == colLength-1 {
			if a, b, c := search(i, i-1, i, j, j-1, j, m); a {
				return true, b, c
			} else {
				return false, 0, 0
			}
		}

		//lie at four borders
		if i == 0 {
			if a, b, c := search(i, i, i+1, j, j-1, j+1, m); a {
				return true, b, c
			} else {
				return false, 0, 0
			}
		} else if i == rowLength-1 {
			if a, b, c := search(i, i-1, i, j, j-1, j+1, m); a {
				return true, b, c
			} else {
				return false, 0, 0
			}
		}
		if j == 0 {
			if a, b, c := search(i, i-1, i+1, j, j, j+1, m); a {
				return true, b, c
			} else {
				return false, 0, 0
			}
		} else if j == colLength-1 {
			fmt.Println("j == colLength-1 ")
			if a, b, c := search(i, i-1, i+1, j, j-1, j, m); a {
				return true, b, c
			} else {
				return false, 0, 0
			}
		}
	}

	return false, 0, 0
}

func search(i, i_top, i_down, j, j_left, j_right int, m byte) (isFound bool, i_found, j_found int) {
	fmt.Println("search_1", i, j)
	if m != 88 && m != 77 {
		if matrix[i][j] == m {
			// fmt.Println(ii, jj)
			// foundTimes++
			return true, i, j
		} else {
			return false, 0, 0
		}
	}

	for ii := i_top; ii < i_down+1; ii++ {
		for jj := j_left; jj < j_right+1; jj++ {
			//don't check itself
			if ii == i && jj == j {
				continue
			}
			fmt.Println("search_2", i, j, ii, jj, m)
			if matrix[ii][jj] == m && m != 77 {
				// fmt.Println(ii, jj)
				// foundTimes++
				return true, ii, jj
			}
			if matrix[ii][jj] == m && m == 77 {
				addPair(ii, jj)
			}
		}
	}

	if m == 77 && len(p) != 0 {
		return true, 0, 0
	}

	return false, 0, 0
}

// 1 2 3
// 4   6
// 7 8 9
func find1of8(i_prev, j_prev, i_now, j_now int) (i_next, j_next int) {
	if i_now == i_prev-1 && j_now == j_prev-1 {
		i_next = i_now - 1
		j_next = j_now - 1
	} else if i_now == i_prev-1 && j_now == j_prev {
		i_next = i_now - 1
		j_next = j_now
	} else if i_now == i_prev-1 && j_now == j_prev+1 {
		i_next = i_now - 1
		j_next = j_now + 1
	} else if i_now == i_prev && j_now == j_prev-1 {
		i_next = i_now
		j_next = j_now - 1
	} else if i_now == i_prev && j_now == j_prev+1 {
		i_next = i_now
		j_next = j_now + 1
	} else if i_now == i_prev+1 && j_now == j_prev-1 {
		i_next = i_now + 1
		j_next = j_now - 1
	} else if i_now == i_prev+1 && j_now == j_prev {
		i_next = i_now + 1
		j_next = j_now
	} else if i_now == i_prev+1 && j_now == j_prev+1 {
		i_next = i_now + 1
		j_next = j_now + 1
	}
	return i_next, j_next
}
