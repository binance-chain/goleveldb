package leveldb

import (
	"testing"

	"github.com/binance-chain/goleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
