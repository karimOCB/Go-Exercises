package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	counter := 0
    for _, value := range cb[file] {
        if value {
            counter ++
        }
    }
    return counter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank <= 0 || rank >= 9 {
        return 0
    }
    counter := 0
	for _, value := range cb {
        if value[rank - 1] {
            counter++;
        }
    }
    return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	counter := 0
    for key := range cb {
        for range cb[key] {
            counter++
        }
    }
	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	counter := 0
    for key := range cb {
        for _, value := range cb[key] {
            if value {
                counter++
            }
        }
    }
	return counter
}
