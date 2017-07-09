// +---------------+----------------+----------------+
// |timestamp(ms)42  | worker id(10) | sequence(12)  |
// +---------------+----------------+----------------+

package uuid

import (
	"errors"
	"sync"
	"time"
)

const (
	Epoch          = 1474802888000
	TimeStampShift = 22
	WorkerIdBits   = 10
	WorkerIdShift  = 12
	SequenceMask   = 0xfff
)

type ISnowflake struct {
	workerId      int64
	lastTimeStamp int64
	sequence      int64
	maxWorkerId   int64
	lock          *sync.Mutex
}

// WorkID : < 1024
func NewSnowflake(workerid int64) (iw *ISnowflake, err error) {
	iw = new(ISnowflake)
	iw.workerId = workerid
	iw.maxWorkerId = (-1 ^ -1<<WorkerIdBits)
	if workerid > iw.maxWorkerId || workerid < 0 {
		return nil, errors.New("worker not fit")
	}
	iw.sequence = 0
	iw.lastTimeStamp = -1
	iw.lock = new(sync.Mutex)
	return iw, nil
}

// return in ms
func (iw *ISnowflake) timeGen() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func (iw *ISnowflake) nextTimeMillis(last int64) int64 {
	ts := time.Now().UnixNano() / 1000 / 1000
	for {
		if ts < last {
			ts = iw.timeGen() / 1000 / 1000
		} else {
			break
		}
	}
	return ts
}

func (iw *ISnowflake) GenerateID() (ts int64, err error) {
	iw.lock.Lock()
	defer iw.lock.Unlock()

	ts = iw.timeGen()
	if ts == iw.lastTimeStamp {
		iw.sequence = (iw.sequence + 1) & SequenceMask
		// 毫秒内序列溢出, 取下一个毫秒
		if iw.sequence == 0 {
			ts = iw.nextTimeMillis(ts)
		}
	} else {
		iw.sequence = 0
	}

	if ts < iw.lastTimeStamp {
		return 0, errors.New("refused to generate id")
	}
	iw.lastTimeStamp = ts
	ts = (ts-Epoch)<<TimeStampShift | iw.workerId<<WorkerIdShift | iw.sequence
	return ts, nil
}
