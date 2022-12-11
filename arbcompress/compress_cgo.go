// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package arbcompress

func Decompress(input []byte, maxSize int) ([]byte, error) {
	return []byte{}, nil
}

func compressLevel(input []byte, level int) ([]byte, error) {
	return []byte{}, nil
}

func CompressWell(input []byte) ([]byte, error) {
	return compressLevel(input, LEVEL_WELL)
}
