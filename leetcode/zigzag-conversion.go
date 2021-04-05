func convert(s string, numRows int) string {
    if numRows <= 1 {
        return s
    }
    row := make([][]byte, numRows)
	r, dir := 0, -1
    for i := range row {
        row[i] = []byte{}
    }
	for i := 0; i < len(s); i++ {
		row[r] = append(row[r], s[i])
		if r == numRows-1 || r == 0 {
            dir = dir*(-1)
		}
        r += dir
	}
    return string(bytes.Join(row,[]byte{}))
}