package main

import "fmt"

func placeQueen(queenNum int, board [][]int) bool {
	for i := 0; i<8; i++ {
		for j := 0; j<8; j++ {
			if allowedCell(i, j, board) {

				board[i][j] = 1
				queenNum = queenNum + 1

				result := true

				if queenNum < 8 {
					result = placeQueen(queenNum, board)
				}

				if !result {
					fmt.Printf("caught result false in %v %v \n", i, j)
					board[i][j] = 0
					queenNum = queenNum - 1
				}

				return result
			}
		}
	}
	return false
}

func allowedCell(row, col int, board [][]int) bool {
	/*
	 * To be implemented ...
	 */
	return board[row][col] == 0
}

func main(){
	board := [][]int{}
	for j := 0; j<8; j++ {
		board = append(board, []int{0,0,0,0,0,0,0,0})
	}

	placeQueen(0, board)

	fmt.Println("Solution:")

	for i := 0; i<8; i++ {
		for j := 0; j<8; j++ {
			fmt.Print(board[i][j], " ")
		}
		fmt.Print("\n")
	}

}