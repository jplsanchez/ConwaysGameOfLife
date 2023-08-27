package main

type Matrix [][]Cell

type Cell struct {
	alive      bool
	neighbours int
}

func NewMatrix(seedMatrix [][]int) *Matrix {
	var matrix Matrix = make([][]Cell, len(seedMatrix))
	for i := range matrix {
		matrix[i] = make([]Cell, len(seedMatrix[i]))
		for j := range matrix[i] {
			matrix[i][j].alive = seedMatrix[i][j] == 1
			matrix[i][j].neighbours = 0
		}
	}
	return &matrix
}

func (m *Matrix) Update() {
	for i := 0; i < len(*m); i++ {
		for j := 0; j < len((*m)[i]); j++ {
			m.setCellNeighbours(i, j)
		}
	}

	for i := 0; i < len(*m); i++ {
		for j := 0; j < len((*m)[i]); j++ {
			m.setNextState(i, j)
		}
	}
}

func (m *Matrix) setCellNeighbours(i int, j int) {
	cell := &(*m)[i][j]

	var kernel = [][]bool{
		{true, true, true},
		{true, false, true},
		{true, true, true}}

	cell.neighbours = 0
	for k := 0; k < len(kernel); k++ {
		for l := 0; l < len(kernel[k]); l++ {
			if !kernel[k][l] {
				continue
			}
			x := i + k - 1
			y := j + l - 1
			isOnBoundaries := x >= 0 && x < len(*m) && y >= 0 && y < len((*m)[x])
			if isOnBoundaries && (*m)[x][y].alive {
				cell.neighbours++
			}
		}
	}
}

func (m *Matrix) setNextState(i, j int) {
	cell := &(*m)[i][j]

	// 1. Any live cell with two or three live neighbours survives.
	if cell.alive && cell.neighbours == 2 || cell.neighbours == 3 {
		cell.alive = true
		return
	}
	// 2. Any dead cell with three live neighbours becomes a live cell.
	if !cell.alive && cell.neighbours == 3 {
		cell.alive = true
		return
	}
	// 3. All other live cells die in the next generation. Similarly, all other dead cells stay dead.
	cell.alive = false
}
