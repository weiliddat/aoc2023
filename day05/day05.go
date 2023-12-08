package day05

import (
	"aoc2023/aoc_util"
	_ "embed"
	"math"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

//go:embed input.txt
var Input string

func Solve(input string) (string, string, error) {
	part01, err := Part01(input)

	if err != nil {
		return "", "", err
	}

	part02, err := Part02(input)

	if err != nil {
		return part01, "", err
	}

	return part01, part02, nil
}

func Part01(input string) (string, error) {
	almanac, err := ParseInput(input)
	if err != nil {
		return "", err
	}

	lowest := uint64(math.MaxUint64)
	for _, seed := range almanac.Seeds {
		destNum := almanac.LookupLocationFromSeed(seed)

		if lowest > destNum {
			lowest = destNum
		}
	}

	return strconv.FormatUint(lowest, 10), nil
}

func Part02(input string) (string, error) {
	almanac, err := ParseInput(input)
	if err != nil {
		return "", err
	}

	lowest := atomic.Uint64{}
	lowest.Store(math.MaxUint64)

	for i := 0; i < len(almanac.Seeds); i += 2 {
		wg := sync.WaitGroup{}

		numSeeds := almanac.Seeds[i+1]
		numWorkers := uint64(runtime.NumCPU())
		pieceSize := numSeeds / uint64(numWorkers)
		remainder := numSeeds % uint64(numWorkers)

		for j := uint64(0); j < numWorkers; j++ {
			wg.Add(1)

			rangeStart := almanac.Seeds[i] + (j * pieceSize)
			rangeEnd := almanac.Seeds[i] + (j+1)*pieceSize
			if j == numWorkers-1 {
				rangeEnd += remainder
			}
			seeds := []uint64{rangeStart, rangeEnd}

			go lookupWorker(&almanac, seeds, &lowest, &wg)
		}

		wg.Wait()
	}

	return strconv.FormatUint(lowest.Load(), 10), nil
}

func lookupWorker(a *Almanac, seeds []uint64, lowest *atomic.Uint64, wg *sync.WaitGroup) {
	defer wg.Done()

	for s := seeds[0]; s < seeds[1]; s++ {
		l := a.LookupLocationFromSeed(s)
		if lowest.Load() > l {
			lowest.Store(l)
		}
	}
}

type AlmanacMapRange struct {
	DestStart uint64
	SrcStart  uint64
	Length    uint64
}

type AlmanacMap struct {
	Src    string
	Dest   string
	Ranges []AlmanacMapRange
}

type Almanac struct {
	Seeds []uint64
	Maps  map[string]AlmanacMap
}

func (a *Almanac) LookupLocationFromSeed(seed uint64) uint64 {
	nextName, nextNum := "seed", seed

	for {
		nextName, nextNum = a.Lookup(nextName, nextNum)

		if nextName == "location" {
			break
		}
	}

	return nextNum
}

func (a *Almanac) Lookup(srcName string, srcNum uint64) (string, uint64) {
	destMap, exists := a.Maps[srcName]

	if !exists {
		return "", 0
	}

	destName, destNum := destMap.Dest, srcNum

	for _, r := range destMap.Ranges {
		if srcNum >= r.SrcStart && srcNum < r.SrcStart+r.Length {
			destNum = r.DestStart + srcNum - r.SrcStart
			break
		}
	}

	return destName, destNum
}

func ParseInput(input string) (Almanac, error) {
	almanac := Almanac{
		Seeds: []uint64{},
		Maps:  map[string]AlmanacMap{},
	}

	blocks := strings.Split(input, "\n\n")

	seedTexts := strings.Split(blocks[0], " ")[1:]
	for _, seedText := range seedTexts {
		seed, err := strconv.ParseUint(seedText, 10, 64)
		if err != nil {
			return almanac, err
		}
		almanac.Seeds = append(almanac.Seeds, seed)
	}

	blockMapMatcher := regexp.MustCompile(`^(\w+)-to-(\w+) map:`)
	for _, block := range blocks[1:] {
		matches := blockMapMatcher.FindAllStringSubmatch(block, -1)

		almanacMap := AlmanacMap{
			Src:    matches[0][1],
			Dest:   matches[0][2],
			Ranges: []AlmanacMapRange{},
		}

		rangeLines := aoc_util.SplitLines(block)[1:]
		for _, rangeLine := range rangeLines {
			ranges := strings.Split(rangeLine, " ")

			destStart, err := strconv.ParseUint(ranges[0], 10, 64)
			if err != nil {
				return almanac, err
			}

			srcStart, err := strconv.ParseUint(ranges[1], 10, 64)
			if err != nil {
				return almanac, err
			}

			length, err := strconv.ParseUint(ranges[2], 10, 64)
			if err != nil {
				return almanac, err
			}

			almanacMap.Ranges = append(almanacMap.Ranges, AlmanacMapRange{
				destStart, srcStart, length,
			})
		}

		almanac.Maps[almanacMap.Src] = almanacMap
	}

	return almanac, nil
}
