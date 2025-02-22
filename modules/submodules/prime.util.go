package submodules

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

var smallPrime = map[uint64]bool{
	1:  true,
	2:  true,
	3:  true,
	5:  true,
	7:  true,
	11: true,
	13: true,
	17: true,
	19: true,
	23: true,
	29: true,
	31: true,
	37: true,
	41: true,
	43: true,
	47: true,
	53: true,
	59: true,
	61: true,
	67: true,
	71: true,
	73: true,
	79: true,
	83: true,
	89: true,
	97: true,
}


func byte32ToInt64(read []byte) int64 {

	var value int32
	value |= int32(read[3])
	value |= int32(read[2]) << 8
	value |= int32(read[1]) << 16
	value |= int32(read[0]) << 24
	return int64(value)
}

func _read(fileName, tableName string, k int64) ([]uint32, error) {
	file, err := os.Open(fileName)
	table, err2 := os.Open(tableName)

	if err != nil || err2 != nil {
		return []uint32{}, errors.Join(err, err2)
	}

	ref1 := make([]byte, 4)
	ref2 := make([]byte, 4)

	i := 4 * (k - 1)

	table.ReadAt(ref1, i)
	table.ReadAt(ref2, i+4)

	start := byte32ToInt64(ref1) / 8
	end := byte32ToInt64(ref2) / 8

	leng := end - start

	val := make([]byte, leng-1)
	file.ReadAt(val, start)

	return _to32Slice(string(val), file, table)
}

func _to32Slice(s string, f, t *os.File) ([]uint32, error) {
	split := strings.Split(s, " ")

	out := make([]uint32, len(split))
	for i, s := range split {
		if strings.ContainsRune(s, '*') {
			panic(s)
		}

		i32, err := strconv.Atoi(s)
		if err != nil {
			return []uint32{}, err
		}
		out[i] = uint32(i32)
	}
	return out, nil
}
