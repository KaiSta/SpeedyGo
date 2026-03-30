package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	race "github.com/kai/SpeedyGo/traceReplay/detector"
	traceParser "github.com/kai/SpeedyGo/traceReplay/parser"
	"github.com/kai/SpeedyGo/traceReplay/util"

	"runtime/pprof"
)

func main() {
	trace := flag.String("trace", "", "path to trace")
	json := flag.Bool("json", false, "output as json")
	plain := flag.Bool("plain", true, "output as plain text")
	bench := flag.Bool("bench", false, "used for benchmarks only")
	analysis := flag.String("mode", "fasttrack", "Use -ls to list available race predictors")
	parser := flag.String("parser", "javainc", "java,javainc")
	repeat := flag.Int("repeat", 2, "multiple runs of the same trace")
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")
	validateTrace := flag.Bool("validate", false, "set if the trace should be validated for correctness")
	n := flag.Uint64("n", ^uint64(0), "number of events to verify")
	s := flag.Uint64("s", 0, "start verification at s events")
	filter := flag.Uint64("filter", 0, "filter level")
	postProcess := flag.Bool("pp", true, "PostProcess")
	ls := flag.Bool("ls", false, "List available race detectors")

	flag.Parse()

	if *ls {
		race.LSDetectors()
		return
	}

	// Optimize GOMAXPROCS for better garbage collection performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Validate flags
	if err := validateFlags(trace, json, plain, bench, analysis); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Start CPU profiling if requested
	var profileFile *os.File
	if *cpuprofile != "" {
		var err error
		profileFile, err = os.Create(*cpuprofile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating profile file: %v\n", err)
			os.Exit(1)
		}
		pprof.StartCPUProfile(profileFile)
		defer func() {
			pprof.StopCPUProfile()
			profileFile.Close()
		}()
	}

	fmt.Println("Parsing trace...")
	tracePath, err := filepath.Abs(*trace)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting absolute path: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Trace path:", tracePath)

	var items []util.Item
	switch *parser {
	case "go":
		items = traceParser.ParseTracev2(*trace)
	case "java", "java2":
		items = traceParser.ParseJTracev2(*trace, *n)
	default:
		// Handle incremental parsing cases
	}

	fmt.Println("Parsing trace complete!")

	if *validateTrace {
		fmt.Println("Validating trace...")
		race.RunAPIV2(items, *analysis, *filter)
		fmt.Println("Validating successful")
	}

	// Run the analysis
	t1 := time.Now()
	runAnalysis(*parser, *analysis, *postProcess, *trace, items, *n, *s, *repeat, *filter)
	t2 := time.Now()
	fmt.Println("Complete Time:", t2.Sub(t1))
}

func validateFlags(trace *string, json, plain, bench *bool, analysis *string) error {
	if !*json && !*plain && !*bench {
		return fmt.Errorf("no output format defined")
	}
	if trace == nil || *trace == "" {
		return fmt.Errorf("no valid trace file")
	}
	if *analysis == "" {
		return fmt.Errorf("no analysis mode chosen")
	}
	return nil
}

func runAnalysis(parser, analysis string, postProcess bool, trace string, items []util.Item, n, s uint64, repeat int, filter uint64) {
	switch parser {
	case "javainc2":
		race.RunAPISingleIncDouble(analysis, postProcess, trace, n, s, repeat)
	case "java2":
		race.RunAPISingleDouble(items, analysis, filter, postProcess)
	case "javainc":
		race.RunAPISingleInc(analysis, postProcess, trace, n, s, filter)
	default:
		race.RunAPISingle(items, analysis, filter, postProcess)
	}
}
