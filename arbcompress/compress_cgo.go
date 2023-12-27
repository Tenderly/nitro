// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package arbcompress

import "C"
import (
	"bytes"
	"github.com/andybalholm/brotli"
)

func Decompress(_ []byte, _ int) ([]byte, error) {
	return nil, nil
}

func compressLevel(input []byte, level int) ([]byte, error) {
	options := brotli.WriterOptions{Quality: level}
	var buf bytes.Buffer
	writer := brotli.NewWriterOptions(&buf, options)
	_, err := writer.Write(input)
	if closeErr := writer.Close(); err == nil {
		err = closeErr
	}
	return buf.Bytes(), err
}

func CompressWell(input []byte) ([]byte, error) {
	return compressLevel(input, LEVEL_WELL)
}
