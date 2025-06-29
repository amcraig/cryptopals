package matrix

func Transpose[T any](matrix [][]T) [][]T {
	result := make([][]T, len(matrix[0]))
	for i := range result {
		result[i] = make([]T, len(matrix))
	}

	for i := range matrix {
		for j := range matrix[i] {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}

// Postive values shift left, negative values shift right
func RotateVector[T any](vector []T, pos int) []T {
	if len(vector) == 0 {
		return vector
	}

	if pos < 0 {
		return RotateVector(vector, len(vector)+pos)
	}

	pos = pos % len(vector)

	vector = append(vector[pos:], vector[:pos]...)
	return vector

}
