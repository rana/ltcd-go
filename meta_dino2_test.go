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

type Dino2 struct {
	Name      string
	LegLen    float64
	StrideLen float64
	Speed     float64
	Stance    string
	Diet      string
}

func readDinoLeg() (res map[string]Dino2) {
	fle, err := os.Open("meta_dino_leg.csv")
	if err != nil {
		fmt.Println("open dino leg", err)
		os.Exit(1)
	}
	defer fle.Close()

	rdr := csv.NewReader(fle)
	recs, err := rdr.ReadAll()
	if err != nil {
		fmt.Println("csv read", err)
		os.Exit(1)
	}

	res = make(map[string]Dino2)
	for idx, rec := range recs {
		if idx == 0 {
			continue
		}
		var dino Dino2
		dino.Name = rec[0]
		dino.LegLen, err = strconv.ParseFloat(rec[1], 64)
		if err != nil {
			fmt.Println("parse leg len", rec[1], err)
			continue
		}
		dino.Diet = rec[2]
		res[dino.Name] = dino
	}
	return res
}

func readDinoStride() (res map[string]Dino2) {
	fle, err := os.Open("meta_dino_stride.csv")
	if err != nil {
		fmt.Println("open dino stride", err)
		os.Exit(1)
	}
	defer fle.Close()

	rdr := csv.NewReader(fle)
	recs, err := rdr.ReadAll()
	if err != nil {
		fmt.Println("csv parse", err)
		os.Exit(1)
	}

	res = make(map[string]Dino2)
	for idx, rec := range recs {
		if idx == 0 {
			continue
		}
		// NAME,STRIDE_LENGTH,STANCE
		var dino Dino2
		dino.Name = rec[0]
		dino.StrideLen, err = strconv.ParseFloat(rec[1], 64)
		if err != nil {
			fmt.Println("parse stride", err)
			continue
		}
		dino.Stance = rec[2]
		res[dino.Name] = dino
	}
	return res
}

func TestDino2(t *testing.T) {
	// Read two files
	// Merge files
	// Filter based on "bipedal"
	// Calculate speed
	// Sort based on speed (descending)
	// Print Names

	// Read two files
	dinoLegs := readDinoLeg()
	fmt.Println("dinoLegs", dinoLegs)
	dinoStrides := readDinoStride()
	fmt.Println("dinoStrides", dinoStrides)

	// Merge files
	var mergedDinos []Dino2
	for key, dinoStride := range dinoStrides {
		dinoLeg, exists := dinoLegs[key]
		if !exists {
			continue
		}
		dinoStride.LegLen = dinoLeg.LegLen
		mergedDinos = append(mergedDinos, dinoStride)
	}
	fmt.Println("mergedDinos", mergedDinos)

	// Filter based on "bipedal"
	var filteredDinos []Dino2
	for _, dino := range mergedDinos {
		if dino.Stance == "bipedal" {
			filteredDinos = append(filteredDinos, dino)
		}
	}
	fmt.Println("filteredDinos", mergedDinos)

	// Calculate speed
	for idx, dino := range filteredDinos {
		// ((STRIDE_LENGTH / LEG_LENGTH) - 1) * SQRT(LEG_LENGTH * 9.8)
		dino.Speed = ((dino.StrideLen / dino.LegLen) - 1) * math.Sqrt(dino.LegLen*9.8)
		fmt.Println("dino", dino)
		filteredDinos[idx] = dino
	}
	fmt.Println("Speed", filteredDinos)

	// Sort based on speed (descending)
	sort.Slice(filteredDinos, func(i, j int) bool {
		return filteredDinos[i].Speed > filteredDinos[j].Speed
	})
	fmt.Println("Sort", filteredDinos)

	// Print Names
	for _, dino := range filteredDinos {
		fmt.Println(dino.Name)
	}
}
