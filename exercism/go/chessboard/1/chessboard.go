package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    if file < "A" || file > "H" || len(file) != 1 {
        return 0
    }
    
	occupied := 0

    for _, f := range cb[file] {
        if f {
            occupied++
        }
    }

    return occupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
        return 0
    }

    occupied := 0

    for _, r := range cb {
        if r[rank-1] {
            occupied++
        }
    }

    return occupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	sq := 0
    for _, r := range cb {
        sq += len(r)
    }
    return sq
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupied := 0
    for f, _ := range cb {
        occupied += CountInFile(cb, f)
    }
    return occupied
}
