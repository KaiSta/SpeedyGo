package tsanwrd

import (
	"../../util"
	"../report"
	"github.com/xojoc/bitset"
)

type ListenerDataAccessWCPHBC struct{}
type ListenerAsyncRcvWCPHBC struct{}
type ListenerAsyncSndWCPHBC struct{}

func (l *ListenerDataAccessWCPHBC) Put(p *util.SyncPair) {
	if !p.DataAccess {
		return
	}

	t1, ok := threads[p.T1]
	if !ok {
		t1 = newT(p.T1)
	}

	varstate, ok := variables[p.T2]
	if !ok {
		varstate = newVar()
	}

	varstate.current++
	newFE := &dot{v: t1.vc.clone(), int: varstate.current, t: uint16(p.T1),
		sourceRef: p.Ev.Ops[0].SourceRef, line: uint16(p.Ev.Ops[0].Line),
		write: p.Write, ls: make(map[uint32]struct{}), lsCount: make(map[uint32]int), pos: p.Ev.LocalIdx}
	for _, k := range t1.ls { //copy lockset
		//lk := locks[k]
		//newFE.lsCount[k] = lk.count
		newFE.ls[k] = struct{}{}
	}

	//var conflictingCritSections vcepoch = newvc2()
	if p.Write {
		newFrontier := make([]*dot, 0, len(varstate.frontier))
		//connectTo := make([]*dot, 0)

		for _, f := range varstate.frontier {
			k := f.v.get(uint32(f.t))
			thi_at_j := t1.vc.get(uint32(f.t))

			if k > thi_at_j {
				newFrontier = append(newFrontier, f) // RW(x) =  {j#k | j#k ∈ RW(x) ∧ k > Th(i)[j]}

				if !intersect(newFE.ls, f.ls) {
					report.ReportRace(report.Location{File: uint32(f.sourceRef), Line: uint32(f.line), W: f.write},
						report.Location{File: p.Ev.Ops[0].SourceRef, Line: p.Ev.Ops[0].Line, W: true}, false, 0)
				} else { //interesting case, two events are concurrent but protected by the same lock, if the previous one is a read and the current a write, we need to synchronize the threads! zu konvervativ! wer sagt dass das echt last write nicht zwishcen die crits geschoben werden kann?
					// if !f.write { //prev is a read
					// 	// should be release -> write sync, dont have the release event at the moment needs to be fixed in code
					// 	for k := range newFE.ls {
					// 		if _, ok := f.ls[k]; ok {
					// 			xx := f.lsCount[k]
					// 			csHist := t1.csHistory[k]
					// 			conflictingCritSections = conflictingCritSections.ssync(csHist[xx].rel)
					// 		}
					// 	}
					// }
				}

			} else if k < thi_at_j {
				//connectTo = append(connectTo, f)
			}
		}

		//	t1.vc = t1.vc.ssync(conflictingCritSections)

		//varstate.updateGraph3(newFE, connectTo)

		newFrontier = append(newFrontier, newFE) // ∪{i#Th(i)[i]}
		varstate.frontier = newFrontier

		varstate.lastWrite = t1.vc.clone().ssync(t1.hb)
		varstate.lwDot = newFE
	} else if p.Read {
		if varstate.lwDot != nil { // != nil if a last write exists!
			newFE.v = newFE.v.ssync(varstate.lastWrite) //sync with last write in advance, necessary for the graph analysis in the following loop!
		}

		newFrontier := make([]*dot, 0, len(varstate.frontier))

		for _, f := range varstate.frontier {
			k := f.v.get(uint32(f.t))          //j#k
			thi_at_j := t1.vc.get(uint32(f.t)) //Th(i)[j]

			if k > thi_at_j {

				newFrontier = append(newFrontier, f) // RW(x) =  {j]k | j]k ∈ RW(x) ∧ k > Th(i)[j]}

				if f.write {
					if !intersect(newFE.ls, f.ls) {
						report.ReportRace(report.Location{File: uint32(f.sourceRef), Line: uint32(f.line), W: f.write},
							report.Location{File: p.Ev.Ops[0].SourceRef, Line: p.Ev.Ops[0].Line, W: newFE.write}, false, 0)
					}

				}
			} else {
				if f.write {
					newFrontier = append(newFrontier, f)
				}
			}
		}

		//write-read sync
		if varstate.lwDot != nil {
			t1.vc = t1.vc.ssync(varstate.lastWrite) //Th(i) = max(Th(i),L W (x))
			for _, l := range t1.ls {
				lk := locks[l]
				lk.rel = lk.rel.ssync(varstate.lastWrite)
				locks[l] = lk
			}
		}

		for _, l := range t1.ls {
			lkHistory := t1.csHistory[l]
			nhistory := make([]vcPair, 0, len(lkHistory))

			for _, h := range lkHistory {
				localTimeForLastOwner := t1.vc.get(h.owner)
				relTimeForLastOwner := h.rel.get(h.owner)

				//read event is ordered within previous critical section
				if h.acq.v < localTimeForLastOwner && localTimeForLastOwner < relTimeForLastOwner {
					t1.vc = t1.vc.ssync(h.relhb)
				}

				if !(relTimeForLastOwner < localTimeForLastOwner) {
					nhistory = append(nhistory, h)
				}
			}
			t1.csHistory[l] = nhistory
		}

		newFrontier = append(newFrontier, newFE) // ∪{i#Th(i)[i]}
		varstate.frontier = newFrontier

	} else { //volatile synchronize
		vol, ok := volatiles[p.T2]
		if !ok {
			vol = newvc2()
		}
		t1.vc = t1.vc.ssync(vol)
		vol = t1.vc.clone()

		volatiles[p.T2] = vol
	}

	t1.vc = t1.vc.add(p.T1, 1)
	threads[p.T1] = t1
	variables[p.T2] = varstate
}

func (l *ListenerAsyncRcvWCPHBC) Put(p *util.SyncPair) {
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

	nPair := vcPair{owner: p.T1, acq: lock.acq, rel: t1.vc.clone(),
		relhb: t1.vc.clone().ssync(t1.hb)}
	//lock.history = append(lock.history, nPair)
	lock.history = lock.history.add(nPair)
	for _, t := range threads {
		if t.id != p.T1 {
			// cshist := t.csHistory[p.T2]
			// cshist = append(cshist, nPair)
			// t.csHistory[p.T2] = cshist
			cshist := t.csHistory[p.T2]
			cshist = cshist.add(nPair)
			//cshist = append(cshist, nPair)
			t.csHistory[p.T2] = cshist
		}
	}

	lock.count++
	lock.hb = lock.hb.ssync(t1.vc)

	t1.vc = t1.vc.add(p.T1, 1)
	threads[p.T1] = t1
	locks[p.T2] = lock
}

func (l *ListenerAsyncSndWCPHBC) Put(p *util.SyncPair) {
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

	//t1.ls[p.T2] = struct{}{} //inlcude lock to lockset
	t1.ls = append(t1.ls, p.T2)
	t1.vc = t1.vc.ssync(lock.rel)

	lock.acq = newEpoch(p.T1, t1.vc.get(p.T1))

	if len(t1.ls) > 1 {
		multiLock++
	} else {
		singleLock++
	}

	if len(t1.ls) > maxMulti {
		maxMulti = len(t1.ls)
	}

	t1.hb = t1.hb.ssync(lock.hb)

	t1.vc = t1.vc.add(p.T1, 1)
	threads[p.T1] = t1
	locks[p.T2] = lock
}

type ListenerDataAccessWCPEEHBC struct{}

func (l *ListenerDataAccessWCPEEHBC) Put(p *util.SyncPair) {
	if !p.DataAccess {
		return
	}

	t1, ok := threads[p.T1]
	if !ok {
		t1 = newT(p.T1)
	}

	varstate, ok := variables[p.T2]
	if !ok {
		varstate = newVar()
	}

	varstate.current++
	newFE := &dot{v: t1.vc.clone(), int: varstate.current, t: uint16(p.T1),
		sourceRef: p.Ev.Ops[0].SourceRef, line: uint16(p.Ev.Ops[0].Line),
		write: p.Write, ls: make(map[uint32]struct{}), pos: p.Ev.LocalIdx}
	for _, k := range t1.ls { //copy lockset
		newFE.ls[k] = struct{}{}
	}

	newFrontier := make([]*dot, 0, len(varstate.frontier))
	connectTo := make([]*dot, 0)

	if p.Write {

		for _, f := range varstate.frontier {
			k := f.v.get(uint32(f.t))
			thi_at_j := t1.vc.get(uint32(f.t))

			if k > thi_at_j {
				newFrontier = append(newFrontier, f) // RW(x) =  {j#k | j#k ∈ RW(x) ∧ k > Th(i)[j]}

				if !intersect(newFE.ls, f.ls) {
					report.ReportRace(report.Location{File: uint32(f.sourceRef), Line: uint32(f.line), W: f.write},
						report.Location{File: p.Ev.Ops[0].SourceRef, Line: p.Ev.Ops[0].Line, W: true}, false, 0)
				}

				varstate.findRaces(newFE, f, &bitset.BitSet{}, 0)

			} else {
				connectTo = append(connectTo, f)
			}
		}

		varstate.updateGraph3(newFE, connectTo)

		newFrontier = append(newFrontier, newFE) // ∪{i#Th(i)[i]}
		varstate.frontier = newFrontier

		varstate.lastWrite = t1.vc.clone()
		varstate.lwDot = newFE

		// if len(lockClues) > 0 {
		// 	for _, l := range t1.ls {
		// 		lk := locks[l]
		// 		clues := lockClues[l]
		// 		if clues.Get(lk.count) {
		// 			varstate.lastWrite = varstate.lastWrite.ssync(lk.hb)
		// 		}
		// 	}
		// }

		list, ok := varstate.graph.get(newFE.int)
		if ok && len(list) == 0 {
			list = append(list, &startDot)
			varstate.graph.add(newFE, list)
		}
	} else if p.Read {
		if varstate.lwDot != nil {
			newFE.v = newFE.v.ssync(varstate.lastWrite) //sync with last write in advance, necessary for the graph analysis in the following loop!
		}

		for _, f := range varstate.frontier {
			k := f.v.get(uint32(f.t))          //j#k
			thi_at_j := t1.vc.get(uint32(f.t)) //Th(i)[j]

			if k > thi_at_j {

				newFrontier = append(newFrontier, f) // RW(x) =  {j]k | j]k ∈ RW(x) ∧ k > Th(i)[j]}

				if f.write {
					if !intersect(newFE.ls, f.ls) {
						report.ReportRace(report.Location{File: uint32(f.sourceRef), Line: uint32(f.line), W: f.write},
							report.Location{File: p.Ev.Ops[0].SourceRef, Line: p.Ev.Ops[0].Line, W: newFE.write}, false, 0)
					}

					varstate.findRaces(newFE, f, &bitset.BitSet{}, 0)
				}
			} else {
				if f.int > 0 {
					connectTo = append(connectTo, f)
				}
				if f.write {
					newFrontier = append(newFrontier, f)
				}
			}
		}

		//write-read sync
		if varstate.lwDot != nil {
			t1.vc = t1.vc.ssync(varstate.lastWrite) //Th(i) = max(Th(i),L W (x))
			for _, l := range t1.ls {
				lk := locks[l]
				lk.rel = lk.rel.ssync(varstate.lastWrite)
				locks[l] = lk
			}
		}

		for _, l := range t1.ls {
			lkHistory := t1.csHistory[l]
			nhistory := make([]vcPair, 0, len(lkHistory))

			for _, h := range lkHistory {
				localTimeForLastOwner := t1.vc.get(h.owner)
				relTimeForLastOwner := h.rel.get(h.owner)

				//read event is ordered within previous critical section
				if h.acq.v < localTimeForLastOwner && localTimeForLastOwner < relTimeForLastOwner {
					t1.vc = t1.vc.ssync(h.relhb)
				}

				if !(relTimeForLastOwner < localTimeForLastOwner) {
					nhistory = append(nhistory, h)
				}
			}
			t1.csHistory[l] = nhistory
		}

		varstate.updateGraph3(newFE, connectTo)

		newFrontier = append(newFrontier, newFE) // ∪{i#Th(i)[i]}
		varstate.frontier = newFrontier

		//connect to artifical start dot if no connection exists

		list, ok := varstate.graph.get(newFE.int)
		if ok && len(list) == 0 {
			list = append(list, &startDot)
			varstate.graph.add(newFE, list)
		}
	} else { //volatile synchronize
		vol, ok := volatiles[p.T2]
		if !ok {
			vol = newvc2()
		}
		t1.vc = t1.vc.ssync(vol)
		vol = t1.vc.clone()

		volatiles[p.T2] = vol
	}

	t1.vc = t1.vc.add(p.T1, 1)
	threads[p.T1] = t1
	variables[p.T2] = varstate
}
