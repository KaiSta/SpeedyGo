package tsanwrd

import (
	"fmt"

	"../../util"
	algos "../analysis"
	"../traceReplay"
)

// WRDs + WRD transfer over locks only

type ListenerAsyncSnd2 struct{}
type ListenerAsyncRcv2 struct{}

type ListenerDataAccessOneRunSound struct{}
type ListenerDataAccessWRDW1 struct{}
type ListenerGoFork struct{}
type ListenerGoWait struct{}
type ListenerNT struct{}
type ListenerNTWT struct{}
type ListenerPostProcess struct{}
type ListenerPostProcess2 struct{}

type EventCollector struct {
	listeners     []traceReplay.EventListener
	preProcessing []traceReplay.EventListener
	phase         uint
}

//var lockClues map[uint32]bitset.BitSet

func (l *EventCollector) Put(p *util.SyncPair) {
	if l.phase == 2 {
		for _, l := range l.listeners {
			l.Put(p)
		}
	} else {
		for _, l := range l.preProcessing {
			l.Put(p)
		}

		if p.PostProcess {
			l.phase++
		}
	}
}

func Init() {
	threads = make(map[uint32]*thread)
	locks = make(map[uint32]*lock)
	signalList = make(map[uint32]vcepoch)
	variables = make(map[uint32]*variable)
	volatiles = make(map[uint32]vcepoch)
	notifies = make(map[uint32]vcepoch)
	csHistory = make(map[uint32][]vcPair)
	lockClues = make(map[uint32][][]int)

	algos.RegisterDetector("polockset", &EventCollector{[]traceReplay.EventListener{
		&ListenerAsyncSnd{},
		&ListenerAsyncRcv{},
		&ListenerDataAccess{},
		&ListenerGoFork{},
		&ListenerGoWait{},
		&ListenerNT{},
		&ListenerNTWT{},
		&ListenerPostProcess{},
	}, nil, 2})

	algos.RegisterDetector("powrdlockset", &EventCollector{[]traceReplay.EventListener{
		&ListenerAsyncSnd{},
		&ListenerAsyncRcv{},
		//	&ListenerDataAccess{},
		&ListenerDataAccessWRD{},
		&ListenerGoFork{},
		&ListenerGoWait{},
		&ListenerNT{},
		&ListenerNTWT{},
		&ListenerPostProcess{},
	}, nil, 2})

	algos.RegisterDetector("w3po", &EventCollector{[]traceReplay.EventListener{
		&ListenerAsyncSndWCP{},
		&ListenerAsyncRcvWCP{},
		&ListenerDataAccessWCP{},
		&ListenerGoFork{},
		&ListenerGoWait{},
		&ListenerNT{},
		&ListenerNTWT{},
		&ListenerPostProcess{},
	}, nil, 2})

	algos.RegisterDetector("w3hbc", &EventCollector{[]traceReplay.EventListener{
		&ListenerAsyncSndWCPHBC{},
		&ListenerAsyncRcvWCPHBC{},
		&ListenerDataAccessWCPHBC{},
		&ListenerGoFork{},
		&ListenerGoWait{},
		&ListenerNT{},
		&ListenerNTWT{},
		&ListenerPostProcess{},
	}, nil, 2})

	algos.RegisterDetector("w3s", &EventCollector{[]traceReplay.EventListener{
		&ListenerAsyncSndW1P2{},
		&ListenerAsyncRcvW1P2{},
		//	&ListenerDataAccess{},
		&ListenerDataAccessWRD{},
		&ListenerGoFork{},
		&ListenerGoWait{},
		&ListenerNT{},
		&ListenerNTWT{},
		&ListenerPostProcess{},
	}, []traceReplay.EventListener{
		&ListenerAsyncSndW1P1{},
		&ListenerAsyncRcvW1P1{},
		&ListenerDataAccessW1P1{},
		&ListenerGoFork{},
		&ListenerGoWait{},
		&ListenerNT{},
		&ListenerNTWT{},
		&ListenerPostProcessW1P1{},
	}, 1})

	algos.RegisterDetector("polocksetEE", &EventCollector{[]traceReplay.EventListener{
		&ListenerAsyncSnd{},
		&ListenerAsyncRcv{},
		&ListenerDataAccessEE{},
		&ListenerGoFork{},
		&ListenerGoWait{},
		&ListenerNT{},
		&ListenerNTWT{},
		&ListenerPostProcess{},
	}, nil, 2})

	algos.RegisterDetector("powrdlocksetEE", &EventCollector{[]traceReplay.EventListener{
		&ListenerAsyncSnd{},
		&ListenerAsyncRcv{},
		&ListenerDataAccessWRDEE{},
		&ListenerGoFork{},
		&ListenerGoWait{},
		&ListenerNT{},
		&ListenerNTWT{},
		&ListenerPostProcess{},
	}, nil, 2})

	algos.RegisterDetector("w3poEE", &EventCollector{[]traceReplay.EventListener{
		&ListenerAsyncSndWCP{},
		&ListenerAsyncRcvWCP{},
		&ListenerDataAccessWCPEE{},
		&ListenerGoFork{},
		&ListenerGoWait{},
		&ListenerNT{},
		&ListenerNTWT{},
		&ListenerPostProcess{},
	}, nil, 2})

	algos.RegisterDetector("w3sEE", &EventCollector{[]traceReplay.EventListener{
		&ListenerAsyncSndW1P2{},
		&ListenerAsyncRcvW1P2{},
		&ListenerDataAccessWRDEE{},
		&ListenerGoFork{},
		&ListenerGoWait{},
		&ListenerNT{},
		&ListenerNTWT{},
		&ListenerPostProcess{},
	}, []traceReplay.EventListener{
		&ListenerAsyncSndW1P1{},
		&ListenerAsyncRcvW1P1{},
		&ListenerDataAccessW1P1{},
		&ListenerGoFork{},
		&ListenerGoWait{},
		&ListenerNT{},
		&ListenerNTWT{},
		&ListenerPostProcessW1P1{},
	}, 1})

	algos.RegisterDetector("w3hbcEE", &EventCollector{[]traceReplay.EventListener{
		&ListenerAsyncSndWCPHBC{},
		&ListenerAsyncRcvWCPHBC{},
		&ListenerDataAccessWCPEEHBC{},
		&ListenerGoFork{},
		&ListenerGoWait{},
		&ListenerNT{},
		&ListenerNTWT{},
		&ListenerPostProcess{},
	}, nil, 2})

}

var threads map[uint32]*thread
var locks map[uint32]*lock
var signalList map[uint32]vcepoch
var variables map[uint32]*variable
var volatiles map[uint32]vcepoch
var notifies map[uint32]vcepoch

var csHistory map[uint32][]vcPair

var multiLock uint
var singleLock uint
var maxMulti int
var numCrits uint

func (l *ListenerGoFork) Put(p *util.SyncPair) {
	if !p.IsFork {
		return
	}

	t1, ok := threads[p.T1]

	if !ok {
		t1 = newT(p.T1)
	}

	signalList[p.T2] = t1.vc.clone()

	t1.vc = t1.vc.add(p.T1, 1)
	threads[p.T1] = t1
}

func (l *ListenerGoWait) Put(p *util.SyncPair) {
	if !p.IsWait {
		return
	}

	t2, ok := signalList[p.T2]

	if ok {
		t1, ok := threads[p.T1]
		if !ok {
			t1 = newT(p.T1)
		}
		t1.vc = t1.vc.ssync(t2)
		t1.vc = t1.vc.add(p.T1, 1)

		threads[p.T1] = t1
	}

}

func (l *ListenerNT) Put(p *util.SyncPair) {
	if !p.IsNT {
		return
	}

	t1, ok := threads[p.T1]
	if !ok {
		t1 = newT(p.T1)
	}
	vc, ok := notifies[p.T2]
	if !ok {
		vc = newvc2()
	}

	vc = vc.ssync(t1.vc)
	t1.vc = t1.vc.ssync(vc)

	t1.vc = t1.vc.add(p.T1, 1)
	notifies[p.T2] = vc
	threads[p.T1] = t1
}

func (l *ListenerNTWT) Put(p *util.SyncPair) {
	if !p.IsNTWT {
		return
	}

	//post wait event, so notify is already synchronized
	if vc, ok := notifies[p.T2]; ok {
		t1, ok := threads[p.T1]
		if !ok {
			t1 = newT(p.T1)
		}

		t1.vc = t1.vc.ssync(vc)
		t1.vc = t1.vc.add(p.T1, 1)

		vc = t1.vc.clone()

		threads[p.T1] = t1
		notifies[p.T2] = vc
	}

}

func (l *ListenerPostProcess) Put(p *util.SyncPair) {
	if !p.PostProcess {
		return
	}
	fmt.Println("Multilock:", multiLock)
	fmt.Println("SingleLock:", singleLock)
	fmt.Println("MaxMulti:", maxMulti)
	fmt.Println("NumCrits:", numCrits)
	fmt.Println("NumHistories:", numHistories, len(locks))
	fmt.Println("MaxHistory:", maxSizeHistories)
	fmt.Println("avg histories:", float64(avghistories)/float64(countSamples))

}
