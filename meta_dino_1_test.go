package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"testing"
)

// Redit sample question: https://leetcode.com/discuss/post/685613/production-engineer-interview-questions-cgcxx/

// You will be supplied with two data files in CSV format.
// The first file contains statistics about various dinosaurs.
// The second file contains additional data.
//
// Given the following formula:
// 	Speed = ((STRIDE_LENGTH / LEG_LENGTH) - 1) * SQRT(LEG_LENGTH * g)
// 	Where g = 9.8 m/s^2 (gravitational constant)
//
// Write a program to read in the data files from disk,
// it must then print the names of only the bipedal
// dinosaurs from fastest to slowest.

func TestDino1(t *testing.T) {
	// Model Dino data based on two files
	// Read and parse file meta_dino_leg.csv
	// Read and parse file meta_dino_stride.csv
	// Merge into single dataset
	// Filter based on stance="bipedal"
	// Calculate Speed based on formula
	// Sort-descending based on Speed

	// Model Dino data based on two files
	type Dino struct {
		Name      string
		Stance    string
		LegLen    float64
		StrideLen float64
		Speed     float64
	}

	// Read and parse file meta_dino_leg.csv
	file, err := os.Open("meta_dino_leg.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rdr := csv.NewReader(file)
	recs, err := rdr.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	legDinos := make(map[string]*Dino)
	for idx, rec := range recs {
		if idx == 0 {
			continue
		}
		var dino Dino
		dino.Name = rec[0]
		dino.LegLen, err = strconv.ParseFloat(rec[1], 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		legDinos[dino.Name] = &dino
	}

	// Read and parse file meta_dino_stride.csv
	file, err = os.Open("meta_dino_stride.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rdr = csv.NewReader(file)
	recs, err = rdr.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	strideDinos := make(map[string]*Dino)
	for idx, rec := range recs {
		if idx == 0 {
			continue
		}
		var dino Dino
		dino.Name = rec[0]
		dino.StrideLen, err = strconv.ParseFloat(rec[1], 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dino.Stance = rec[2]
		strideDinos[dino.Name] = &dino
	}

	// Merge into single dataset
	var mergedFilteredDinos []*Dino
	for name, strideDino := range strideDinos {
		// Filter based on stance="bipedal"
		if strideDino.Stance != "bipedal" {
			continue
		}
		legDino, exists := legDinos[name]
		if !exists {
			continue
		}

		// Calculate Speed based on formula
		// Speed = ((STRIDE_LENGTH / LEG_LENGTH) - 1) * SQRT(LEG_LENGTH * g)
		strideDino.Speed = ((strideDino.StrideLen / legDino.LegLen) - 1) * math.Sqrt(legDino.LegLen*9.8)
		mergedFilteredDinos = append(mergedFilteredDinos, strideDino)
	}

	// Sort-descending based on Speed
	sort.Slice(mergedFilteredDinos, func(i, j int) bool {
		return mergedFilteredDinos[j].Speed < mergedFilteredDinos[i].Speed
	})

	for _, dino := range mergedFilteredDinos {
		fmt.Printf("%+v\n", dino)
	}
}
