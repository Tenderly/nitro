// Copyright 2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package das

import (
	"time"

	"github.com/tenderly/nitro/go-ethereum/log"
	"github.com/tenderly/nitro/arbstate"
	"github.com/tenderly/nitro/util/pretty"
)

func logPut(store string, data []byte, timeout uint64, reader arbstate.DataAvailabilityReader, more ...interface{}) {
	if len(more) == 0 {
		log.Trace(
			store, "message", pretty.FirstFewBytes(data), "timeout", time.Unix(int64(timeout), 0),
			"this", reader,
		)
	} else {
		log.Trace(
			store, "message", pretty.FirstFewBytes(data), "timeout", time.Unix(int64(timeout), 0),
			"this", reader, more,
		)
	}
}
