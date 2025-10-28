package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"unicode"
)

var isProblem2 bool = true

var colLength int = 0
var rowLength int = 0
var matrix [][]byte
var foundTimes int = 0
var duplicates int = 0

type Pair struct {
	I int
	J int
}

var p []Pair

type Pair_2 struct {
	I1 int
	J1 int
	I2 int
	J2 int
	I3 int
	J3 int
}

var p_problem2 []Pair_2

func addPair(i int, j int) {
	p = append(p, Pair{I: i, J: j})
}

func addPair_2(i1 int, j1 int, i2 int, j2 int, i3 int, j3 int) {
	p_problem2 = append(p_problem2, Pair_2{I1: i1, J1: j1, I2: i2, J2: j2, I3: i3, J3: j3})
}

func find_file_col_row(file *os.File) {
	rows, cols := 0, 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if rows == 0 {
			for _, r := range line {
				if unicode.IsLetter(r) || r == '.' {
					cols++
				}
			}
		}
		rows++
	}

	colLength = cols
	rowLength = rows
	fmt.Println("colLength", colLength, rowLength)
}

// XMAS 88 77 65 83
func main() {
	file, err := os.Open("day4.txt")
	if err != nil {
		fmt.Println("file open failed: ", err)
		return
	}
	defer file.Close()

	find_file_col_row(file)

	file.Seek(0, io.SeekStart)

	rows, cols := colLength, rowLength
	matrix = make([][]byte, rows)
	for i := range matrix {
		matrix[i] = make([]byte, cols)
	}

	scanner := bufio.NewScanner(file)
	index_i := 0

	for scanner.Scan() {
		line := scanner.Text()
		bytes := []byte(line)
		matrix[index_i] = bytes
		index_i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("read file failed: ", err)
	}

	if isProblem2 {
		for i := 0; i < rowLength; i++ {
			for j := 0; j < colLength; j++ {
				fmt.Println("main", i, j)
				if matrix[i][j] == 77 {
					fmt.Println("found M: ", i, j)
					p = nil
					if a1, _, _ := check(i, j, 65); a1 {
						fmt.Println("found MA: ", i, j, "*", "*", p)
						for _, p_one := range p {
							fmt.Println("p_one_start:", p_one.I, p_one.J)
							i_next2, j_next2 := find1of4(i, j, p_one.I, p_one.J)
							if i_next2 < 0 || j_next2 < 0 || i_next2 == math.MaxInt {
								continue
							}
							fmt.Println("p_one:", p_one.I, p_one.J)
							if a2, _, _ := check(i_next2, j_next2, 65); a2 {
								foundTimes++
								fmt.Println("found MAS: ", i, j, p_one.I, p_one.J, i_next2, j_next2)
								addPair_2(i, j, p_one.I, p_one.J, i_next2, j_next2)
							}
						}
					}
				}
			}
		}
	} else {

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
	}

	fmt.Println("foundTimes: ", foundTimes)
	countDuplicate()
	fmt.Println(duplicates)

}

func check(i, j int, m byte) (ifFound bool, i_found, j_found int) {
	if i > 0 && i < rowLength-1 && j > 0 && j < colLength-1 {
		//don't lie at four borders
		if a, b, c := search(i, i-1, i+1, j, j-1, j+1, m); a {
			return true, b, c
		} else {
			return false, 0, 0
		}

	} else {

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
	if isProblem2 {
		return search_problem2(i, i_top, i_down, j, j_left, j_right, m)
	}

	fmt.Println("search_1", i, j)

	if m != 88 && m != 77 {
		if matrix[i][j] == m {
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

func search_problem2(i, i_top, i_down, j, j_left, j_right int, m byte) (isFound bool, i_found, j_found int) {
	fmt.Println("search_1", i, j)

	if m != 77 && m != 65 {
		if matrix[i][j] == m {
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
			fmt.Println("search_2", i, j, ii, jj, m, matrix[ii][jj], matrix[ii][jj] == m)
			if matrix[ii][jj] == m && m != 65 {
				return true, ii, jj
			}
			if matrix[ii][jj] == m && m == 65 {
				fmt.Println("addPair: ", ii, jj)
				addPair(ii, jj)
				fmt.Println(p)
			}
		}
	}

	if m == 65 && len(p) != 0 {
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

// 1   3
//
// 7   9
func find1of4(i_prev, j_prev, i_now, j_now int) (i_next, j_next int) {
	if i_now == i_prev-1 && j_now == j_prev-1 {
		i_next = i_now - 1
		j_next = j_now - 1
	} else if i_now == i_prev-1 && j_now == j_prev+1 {
		i_next = i_now - 1
		j_next = j_now + 1
	} else if i_now == i_prev+1 && j_now == j_prev-1 {
		i_next = i_now + 1
		j_next = j_now - 1
	} else if i_now == i_prev+1 && j_now == j_prev+1 {
		i_next = i_now + 1
		j_next = j_now + 1
	} else {
		return math.MaxInt, 1
	}

	if i_next == 0 && j_next == 0 {
		fmt.Println("i_next == 0 ", i_prev, j_prev, i_now, j_now)
	}
	return i_next, j_next
}

func countDuplicate() {
	countMap := make(map[[2]int]int)

	for _, p := range p_problem2 {
		key := [2]int{p.I2, p.J2}
		countMap[key]++
	}

	for _, count := range countMap {
		if count > 1 {
			duplicates++
		}
	}
}
