package main

import ("fmt"
		"strings"
		"strconv")

func placeQueen(queenNum, row, col int, board [][]int) bool {

	if queenNum >= 8 {
		return true
	}

	if queenNum < 8 && row >=8 && col >= 8 {
		return false
	}

	result := false

	for i := 0; i<8; i++ {
		for j := 0; j<8; j++ {

			if allowedCell(i, j, board) {

				board[i][j] = 1
				queenNum = queenNum + 1

				if result = placeQueen(queenNum, i, j, board);result == false {
					board[i][j] = 0
					queenNum = queenNum - 1
				}

			}
		}
	}

	return result
}

func allowedCell(row, col int, board [][]int) bool {

	if board[row][col] != 0 {
		return false
	}

	// walk vertical and horizontal trajectories
	for i:=0; i<8; i++ {
		if board[row][i]==1 || board[i][col]==1 {
			return false
		}
	}

	i, j := row+1, col+1
	SouthEast:
	for ;; i++ {
		for ;; j++ {
			if i<8 && j<8 {
				if board[i][j]==1 {
					return false
				}
				i++
				j++
			} else {
				break SouthEast
			}

		}
	}

	h, k := row-1, col-1
	// walk diagonal North-West
	NorthWest:
	for ;; {
		for ;; {
			if h>=0 && k>=0 {
				if board[h][k]==1 {
					return false
				}
				h--
				k--
			} else {
				break NorthWest
			}
		}
	}


	l, m := row-1, col+1
	// walk diagonal North-East
	NorthEast:
	for ;; {
		for ;; {
			if l>=0 && m<8 {
				if board[l][m]==1 {
					return false
				}
				l--
				m++
			} else {
				break NorthEast
			}
		}
	}


	s, p := row+1, col-1
	SouthWest:
	for ;; {
		for ;; {
			if s < 8 && p >=0 {
				if board[s][p]==1 {
					return false
				}
				s++
				p--
			} else {
				break SouthWest
			}
		}
	}

	return true
}


func main(){
	board := [][]int{}
	for j := 0; j<8; j++ {
		board = append(board, []int{0,0,0,0,0,0,0,0})
	}

	placeQueen(0, 0, 0, board)

	fmt.Println("Solution:")

	printSolution(board)

}

func printSolution(board [][]int){
	s := []string{}
	s = append(s, "")
	for i := 0; i<8; i++ {
		for j := 0; j<8; j++ {
			s = append(s, strconv.Itoa(board[i][j]))
		}
		s = append(s, "\n")
	}
	fmt.Print(strings.Join(s, " "))
}