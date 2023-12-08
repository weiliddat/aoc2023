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

	lowest := math.MaxInt
	for _, seed := range almanac.Seeds {
		destNum := almanac.LookupLocationFromSeed(seed)

		if lowest > destNum {
			lowest = destNum
		}
	}

	return strconv.Itoa(lowest), nil
}

func Part02(input string) (string, error) {
	almanac, err := ParseInput(input)
	if err != nil {
		return "", err
	}

	lowest := math.MaxInt
	for i := 0; i < len(almanac.Seeds); i += 2 {
		locations := make(chan int)
		lowestChan := make(chan int)
		wg := sync.WaitGroup{}

		numSeeds := almanac.Seeds[i+1]
		numWorkers := runtime.NumCPU()
		pieceSize := numSeeds / numWorkers
		remainder := numSeeds % numWorkers

		for i := 0; i < numWorkers; i++ {
			wg.Add(1)

			rangeStart := almanac.Seeds[0] + (i * pieceSize)
			rangeEnd := almanac.Seeds[0] + (i+1)*pieceSize
			if i == numWorkers-1 {
				rangeEnd += remainder
			}
			seeds := []int{rangeStart, rangeEnd}

			go lookupWorker(&almanac, seeds, locations, &wg)
		}

		go lowestWorker(locations, lowestChan)
		wg.Wait()
		close(locations)

		lowestInGroup := <-lowestChan

		if lowest > lowestInGroup {
			lowest = lowestInGroup
		}
	}

	return strconv.Itoa(lowest), nil
}

func lookupWorker(a *Almanac, seeds []int, locations chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for s := seeds[0]; s < seeds[1]; s++ {
		l := a.LookupLocationFromSeed(s)
		locations <- l
	}
}

func lowestWorker(locations <-chan int, lc chan<- int) {
	lowest := math.MaxInt

	for l := range locations {
		if lowest > l {
			lowest = l
		}
	}

	lc <- lowest
}

type AlmanacMapRange struct {
	DestStart int
	SrcStart  int
	Length    int
}

type AlmanacMap struct {
	Src    string
	Dest   string
	Ranges []AlmanacMapRange
}

type Almanac struct {
	Seeds []int
	Maps  map[string]AlmanacMap
}

func (a *Almanac) LookupLocationFromSeed(seed int) int {
	nextName, nextNum := "seed", seed

	for {
		nextName, nextNum = a.Lookup(nextName, nextNum)

		if nextName == "location" {
			break
		}
	}

	return nextNum
}

func (a *Almanac) Lookup(srcName string, srcNum int) (string, int) {
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
		Seeds: []int{},
		Maps:  map[string]AlmanacMap{},
	}

	blocks := strings.Split(input, "\n\n")

	seedTexts := strings.Split(blocks[0], " ")[1:]
	for _, seedText := range seedTexts {
		seed, err := strconv.Atoi(seedText)
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

			destStart, err := strconv.Atoi(ranges[0])
			if err != nil {
				return almanac, err
			}

			srcStart, err := strconv.Atoi(ranges[1])
			if err != nil {
				return almanac, err
			}

			length, err := strconv.Atoi(ranges[2])
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
