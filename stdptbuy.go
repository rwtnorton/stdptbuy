package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	NumAbilities = 6
)

type ScoreCost struct {
	Score int
	Cost  int
}

type CompIndex [NumAbilities]int

var ScoreCosts []ScoreCost
var poolSize int

func init() {
	ScoreCosts = []ScoreCost{
		{7, -4},
		{8, -2},
		{9, -1},
		{10, 0},
		{11, 1},
		{12, 2},
		{13, 3},
		{14, 5},
		{15, 7},
		{16, 10},
		{17, 13},
		{18, 17},
	}
	flag.IntVar(&poolSize, "poolsize", 15, "total points for purchasing scores")
}

func main() {
	flag.Parse()

	seen := map[[NumAbilities]int]bool{}
	var valCnt = len(ScoreCosts)
	var n = 1
	for i := 0; i < NumAbilities; i++ {
		n *= valCnt
	}
	for i := 0; i < n; i++ {
		ci := decompose(i, valCnt)
		cost := ci.TotalCost()
		scores := normalizeScores(ci.Scores())
		if cost == poolSize && !seen[scores] {
			seen[scores] = true
			// fmt.Printf("%d: %v %v\n", i, ci.Scores(), ci.Costs())
			fmt.Println(toString(scores))
		}
	}
}

func decompose(n, base int) CompIndex {
	var result = CompIndex{}
	for i := range result {
		r := n % base
		n /= base
		result[i] = r
		if n == 0 {
			break
		}
	}
	return result
}

func (ci CompIndex) ScoreCosts() [NumAbilities]ScoreCost {
	var result [NumAbilities]ScoreCost
	for i, j := range ci {
		result[i] = ScoreCosts[j]
	}
	return result
}

func toString(vs [NumAbilities]int) string {
	s := make([]string, NumAbilities)
	for i, v := range vs {
		s[i] = strconv.Itoa(v)
	}
	return strings.Join(s, " ")
}

func (ci CompIndex) Scores() [NumAbilities]int {
	var result [NumAbilities]int
	for i, v := range ci.ScoreCosts() {
		result[i] = v.Score
	}
	return result
}

func (ci CompIndex) Costs() [NumAbilities]int {
	var result [NumAbilities]int
	for i, v := range ci.ScoreCosts() {
		result[i] = v.Cost
	}
	return result
}

func (ci CompIndex) TotalCost() int {
	s := 0
	for _, j := range ci {
		s += ScoreCosts[j].Cost
	}
	return s
}

func normalizeScores(s [NumAbilities]int) [NumAbilities]int {
	x := s[:]
	sort.Ints(x)
	for i, v := range x {
		s[i] = v
	}
	return s
}
