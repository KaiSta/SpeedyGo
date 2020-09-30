package wdp

type thread struct {
	vc vcepoch
	//ls        map[uint32]struct{}
	hb      vcepoch
	ls      []uint32
	posLock map[uint32]int
	//csHistory map[uint32][]vcPair
	csHistory map[uint32]limCSHist
	id        uint32
	deps      []intpair
}

func newT(id uint32) *thread {
	t := &thread{id: id, vc: newvc2().set(id, 1), hb: newvc2(),
		ls:      make([]uint32, 0),                                           /*ls: make(map[uint32]struct{})*/
		posLock: make(map[uint32]int), csHistory: make(map[uint32]limCSHist)} //csHistory: make(map[uint32][]vcPair)}

	if len(lockClues) > 0 {
		for k, l := range locks {
			list, ok := t.csHistory[k]
			if !ok {
				list = make([]vcPair, 0)
			}

			for i := range l.history { //0; i < l.count; i++ {
				list = append(list, l.history[i])
			}

			t.csHistory[k] = list
		}
	} else {
		for k, l := range locks {
			list, ok := t.csHistory[k]
			if !ok {
				list = make([]vcPair, 0)
			}
			list = append(list, l.history...)

			t.csHistory[k] = list
		}
	}
	return t
}

type lock struct {
	rel  vcepoch
	rels []vcepoch
	hb   vcepoch
	//history    []vcPair
	history    limCSHist
	acq        epoch
	nextAcq    epoch
	count      int
	strongSync bool
	fullSync   bool
}

func (l *lock) reset() {
	l.rel = newvc2()
	l.hb = newvc2()
	l.count = 0
}

func newL() *lock {
	return &lock{rel: newvc2(), hb: newvc2(), history: make([]vcPair, 0)}
}

type variable struct {
	//	races    []datarace
	history  []variableHistory
	frontier []*dot
	//	graph       *fsGraph
	lastWrite   vcepoch
	lwLocks     map[uint32]int
	lwOpenLocks []intpair
	lwDot       *dot
	current     int
}

func newVar() *variable {
	return &variable{lastWrite: newvc2(), lwDot: nil, frontier: make([]*dot, 0),
		current: 0,
		history: make([]variableHistory, 0), lwLocks: make(map[uint32]int)}
}

var csHistoryLimit = 5 //int(^uint(0) >> 1)

var numHistories = 0
var maxSizeHistories = 0
var avghistories = 0
var countSamples = 0

type limCSHist []vcPair

func newLimCSHist() limCSHist {
	numHistories++
	return make([]vcPair, 0)
}

func (lcsh limCSHist) add(p vcPair) limCSHist {
	if len(lcsh) >= csHistoryLimit {
		lcsh = lcsh[1:]
	}
	if len(lcsh) > maxSizeHistories {
		maxSizeHistories = len(lcsh)
	}
	countSamples++
	avghistories += len(lcsh)
	return append(lcsh, p)
}

type dot struct {
	int
	v         vcepoch
	ls        map[uint32]struct{}
	sourceRef uint32
	pos       int
	line      uint16
	t         uint16
	write     bool
	lsCount   map[uint32]int
}

type signal struct {
	vc vcepoch
}
