package leetcode_go

import (
	"fmt"
	"testing"
)

func TestMaxTurbulenceSize(t *testing.T) {
	fmt.Println(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))
}

func TestLongestCommonSubsequence(t *testing.T) {
	fmt.Println(longestCommonSubsequence("oxcpqrsvwf", "shmtulqrypy"))
}

func TestSubSeq(t *testing.T) {
	subSeq("abcd", "", 0, func(seq string) {
		fmt.Println(seq)
	})

}

func TestIsSubSeq(t *testing.T) {
	fmt.Println(isSubSeq("shmtulqrypy", "qr"))
}

func TestSolveNQueens(t *testing.T) {
	fmt.Println(solveNQueens(4))
}

func TestSolveSudoku(t *testing.T) {
	tmp := [][]string{
		{"5","3",".",".","7",".",".",".","."},
		{"6",".",".","1","9","5",".",".","."},
		{".","9","8",".",".",".",".","6","."},
		{"8",".",".",".","6",".",".",".","3"},
		{"4",".",".","8",".","3",".",".","1"},
		{"7",".",".",".","2",".",".",".","6"},
		{".","6",".",".",".",".","2","8","."},
		{".",".",".","4","1","9",".",".","5"},
		{".",".",".",".","8",".",".","7","9"},
	}
	board := make([][]byte, 9)
	for i, ss := range tmp {
		board[i] = make([]byte, 9)
		for j, s := range ss {
			board[i][j] = s[0]
		}
	}
	solveSudoku(board)
	for _, bytes := range board {
		fmt.Println(string(bytes))
	}
}
