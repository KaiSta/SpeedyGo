package wdp

import (
	"../../util"
	algos "../analysis"
	"../traceReplay"
)

type ListenerAsyncSnd struct{}
type ListenerAsyncRcv struct{}
type ListenerSync struct{}
type ListenerDataAccess struct{}
type ListenerGoFork struct{}
type ListenerGoWait struct{}
type ListenerNT struct{}
type ListenerNTWT struct{}
type ListenerPostProcess struct{}

type EventCollector struct {
	listeners []traceReplay.EventListener
}

func (l *EventCollector) Put(p *util.SyncPair) {
	//	syncPairTrace = append(syncPairTrace, p)
	for _, l := range l.listeners {
		l.Put(p)
	}
}

func Init() {
	threads = make(map[uint32]*thread)
	locks = make(map[uint32]*lock)
	signalList = make(map[uint32]signal)
	variables = make(map[uint32]variable)
	volatiles = make(map[uint32]vcepoch)
	notifies = make(map[uint32]vcepoch)

	listeners := []traceReplay.EventListener{
		&ListenerAsyncSnd{},
		&ListenerAsyncRcv{},
		&ListenerSync{},
		&ListenerDataAccess{},
		&ListenerGoFork{},
		&ListenerGoWait{},
		&ListenerNT{},
		&ListenerNTWT{},
		&ListenerPostProcess{},
	}
	algos.RegisterDetector("wdp", &EventCollector{listeners})
}

var threads map[uint32]*thread
var locks map[uint32]*lock
var signalList map[uint32]signal
var variables map[uint32]variable
var volatiles map[uint32]vcepoch
var notifies map[uint32]vcepoch

var lockClues map[uint32][][]int

func (l *ListenerDataAccess) Put(p *util.SyncPair) {
	if !p.DataAccess {
		return
	}
}
func (l *ListenerAsyncSnd) Put(p *util.SyncPair) {
	if !p.AsyncSend {
		return
	}
	if !p.Lock {
		return
	}

	t1, ok := threads[p.T1]
	if !ok {
		t1 = newT(p.T1)
	}

	lock, ok := locks[p.T2]
	if !ok {
		lock = newL()
	}

	t1.ls = append(t1.ls, p.T2)

	t1.vc = t1.vc.add(p.T1, 1)
	threads[p.T1] = t1
	locks[p.T2] = lock

}
func (l *ListenerAsyncRcv) Put(p *util.SyncPair) {
	if !p.AsyncRcv {
		return
	}
	if !p.Unlock {
		return
	}

	t1, ok := threads[p.T1]
	if !ok {
		t1 = newT(p.T1)
	}

	lock, ok := locks[p.T2]
	if !ok {
		lock = newL()
	}

	if len(t1.ls) > 0 { //remove lock from lockset
		t1.ls = t1.ls[:len(t1.ls)-1]
	}

	t1.vc = t1.vc.add(p.T1, 1)
	threads[p.T1] = t1
	locks[p.T2] = lock

}
func (l *ListenerSync) Put(p *util.SyncPair) {
}
func (l *ListenerGoFork) Put(p *util.SyncPair) {
}
func (l *ListenerGoWait) Put(p *util.SyncPair) {
}
func (l *ListenerNT) Put(p *util.SyncPair) {
}
func (l *ListenerNTWT) Put(p *util.SyncPair) {
}
func (l *ListenerPostProcess) Put(p *util.SyncPair) {
}
