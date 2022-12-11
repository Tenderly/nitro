package sharedmetrics

import (
	"github.com/tenderly/nitro/go-ethereum/metrics"
	"github.com/tenderly/nitro/arbutil"
)

var (
	latestSequenceNumberGauge  = metrics.NewRegisteredGauge("arb/sequencenumber/latest", nil)
	sequenceNumberInBlockGauge = metrics.NewRegisteredGauge("arb/sequencenumber/inblock", nil)
)

func UpdateSequenceNumberGauge(sequenceNumber arbutil.MessageIndex) {
	latestSequenceNumberGauge.Update(int64(sequenceNumber))
}
func UpdateSequenceNumberInBlockGauge(sequenceNumber arbutil.MessageIndex) {
	sequenceNumberInBlockGauge.Update(int64(sequenceNumber))
}
