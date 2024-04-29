/*
--------------------------------------------------------------

	student name: Souhail Daoudi
	student number: 300135458

--------------------------------------------------------------
*/
package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"sync"
	"time"
)

type Histo2 struct {
	Name string
	N    float64
}
type histolist []Histo2

// Less compares values of pairs at two indices
func (p histolist) Less(i, j int) bool {
	return p[i].N < p[j].N
}

// Swap swaps pairs at two indices
func (p histolist) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Len returns the length of the PairList
func (s histolist) Len() int {
	return len(s)
}

var dataimagePath string
var queryimagePath string

var queryhisto Histo
var res1 = make(chan (Histo), 2485)

func main() {
	args := os.Args
	dataPath := args[2]
	queryPath := args[1]
	dataimagePath = dataPath
	queryimagePath = queryPath

	var wg sync.WaitGroup

	// get name of database files from readFiles
	slices := names()

	// start timer
	timeStart := time.Now()
	/////////////////////////////////////////////////////////////////
	// compute histogram of database images
	// send k slices to computeHistograms
	for _, slice := range slices {
		wg.Add(1)
		go func(path string) {
			computeHistograms(slice, 3, res1, &wg)
		}(queryimagePath)

	}
	wg.Wait()
	/////////////////////////////////////////////////////////////////
	// compute query image histogram
	wg.Add(1)
	go func(path string) {
		histo, err := computeHistogram(queryimagePath, 3, &wg)
		if err != nil {
			fmt.Print("erreur")
		}
		queryhisto = histo

	}(queryimagePath)
	wg.Wait()

	/////////////////////////////////////////////////////////////////
	// compare data histogram with query histogram
	wg.Add(1)

	var histos histolist

	go func() {

		defer wg.Done()

		// loop through channel
		for datahisto := range res1 {
			sum = 0
			for i := 0; i < 512; i++ {
				sum += math.Min(datahisto.H[i], queryhisto.H[i])
			}
			// add and sort histograms in histos
			histos = append(histos, Histo2{datahisto.Name, sum})
			sort.Sort(sort.Reverse(histos))
		}

		for j := 0; j < 5; j++ {
			fmt.Println(histos[j].Name)
		}
	}()
	close(res1)
	wg.Wait()

	/////////////////////////////////////////////////////////////////
	// print execution time
	time := time.Since(timeStart)
	fmt.Println("execution time ", time)
}
