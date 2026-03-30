package parser

import (
	"bufio"
	"encoding/binary"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"

	"github.com/kai/SpeedyGo/traceReplay/util"
)

/*
TODO: types in util.Item were changed to uint32, was uint64 before. Parsing go traces it is necessary to change the memory addresses into counted objects like in
the java tracing. This way uint32 is enough and reduces the memory consumption significantly.

*/

var Parserlock = sync.Mutex{}
var FileNumToString map[uint32]string
var FileStringToNum map[string]uint32

// Operation type mapping for Go traces
var goOpTypeMap = map[uint32]struct {
	kind         util.OpKind
	mutex        int
	setChFromPPC bool
	bufSize      uint32
}{
	1:  {util.READ, 0, false, 0},
	2:  {util.WRITE, 0, false, 0},
	3:  {util.ATOMICREAD, 0, false, 0},
	4:  {util.ATOMICWRITE, 0, false, 0},
	5:  {util.SEND, 0, false, 0},
	6:  {util.RCV, 0, false, 0},
	7:  {util.SEND, util.LOCK, false, 1},  // LOCK
	8:  {util.RCV, util.UNLOCK, false, 1}, // UNLOCK
	9:  {util.SIG, 0, true, 0},
	10: {util.WAIT, 0, true, 0},
	11: {util.CLS, 0, false, 0},
	12: {util.PREBRANCH, 0, false, 0},
	13: {util.POSTBRANCH, 0, false, 0},
	14: {util.NT, 0, true, 0},
	15: {util.NTWT, 0, true, 0},
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseTracev2(s string) []util.Item {
	// Initialize with empty slice
	events := make([]util.Item, 0, 1000)

	// file table first
	FileNumToString = make(map[uint32]string)
	locfilePath := path.Join(s, "locfile.log")
	logfilePath := path.Join(s, "mylog.log")

	locfile, err := os.Open(locfilePath)
	if err != nil {
		return events
	}
	defer locfile.Close()
	scanner := bufio.NewScanner(locfile)
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, ";")
		if len(splitted) < 2 {
			continue
		}
		i, err := strconv.ParseUint(splitted[0], 10, 32)
		if err != nil {
			continue
		}
		Parserlock.Lock()
		FileNumToString[uint32(i)] = splitted[1]
		Parserlock.Unlock()
	}

	logfile, err := os.Open(logfilePath)
	if err != nil {
		return events
	}
	defer logfile.Close()

	// Get file size for capacity estimation
	fi, err := logfile.Stat()
	if err == nil {
		// Estimate events: file size / record size (52 bytes)
		estimatedEvents := fi.Size() / 52
		if estimatedEvents > 0 {
			events = make([]util.Item, 0, estimatedEvents)
		}
	}

	buff := make([]byte, 52)
	for {
		n, err := logfile.Read(buff)
		if n == 0 {
			break // EOF
		}
		if err != nil {
			break
		}
		if n != 52 {
			continue // Skip incomplete records
		}

		it := util.Item{}
		it.Partner = binary.LittleEndian.Uint32(buff[4:])
		it.Thread = binary.LittleEndian.Uint32(buff[8:])

		op := parseOp(buff)
		it.Ops = []util.Operation{op}

		events = append(events, it)
	}

	return events
}

func parseOp(buff []byte) util.Operation {
	op := util.Operation{}

	// Parse basic fields using direct binary reads for better performance
	op.Ch = uint32(binary.LittleEndian.Uint64(buff[36:]))
	op.BufSize = binary.LittleEndian.Uint32(buff[44:])
	op.BufField = binary.LittleEndian.Uint32(buff[48:])
	op.SourceRef = binary.LittleEndian.Uint32(buff[12:])
	op.Line = binary.LittleEndian.Uint32(buff[16:])
	op.PC = binary.LittleEndian.Uint64(buff[20:])
	op.PPC = binary.LittleEndian.Uint64(buff[28:])

	head := binary.LittleEndian.Uint32(buff)

	// Set PREPARE or POST based on bit 0
	if head&1 > 0 {
		op.Kind |= util.POST
	} else {
		op.Kind |= util.PREPARE
	}

	// Decode operation type from head (bits 16-31) using lookup table
	opType := head >> 16
	if opInfo, ok := goOpTypeMap[opType]; ok {
		op.Kind |= opInfo.kind
		op.Mutex = opInfo.mutex
		if opInfo.setChFromPPC {
			op.Ch = uint32(op.PPC)
		}
		if opInfo.bufSize > 0 {
			op.BufSize = opInfo.bufSize
		}
	} else {
		// Unknown operation type, default to READ
		op.Kind |= util.READ
	}

	return op
}
