package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// Redit sample question: https://leetcode.com/discuss/post/685613/production-engineer-interview-questions-cgcxx/
const gravitationalConstant1 = 9.8

type Dinosaur1 struct {
	nam string
	lgL float64
	stL float64
	stn string
	spd float64
}

func Dino1() {
	// Read data from CSV files
	dsLeg1 := readDsLeg1_1("meta_dino_leg.csv")
	dsLeg2 := readDsLeg2_1("meta_dino_stride.csv")

	// Join datasets and calculate speeds
	dns := join1(dsLeg1, dsLeg2)

	// Filter bipedal dinosaurs
	bpd := filterBipedal1(dns)

	// Sort by speed (fastest to slowest)
	sort.Slice(bpd, func(i, j int) bool {
		return bpd[i].spd > bpd[j].spd
	})

	// Print results
	for _, d := range bpd {
		fmt.Println(d.nam)
	}
}

func readDsLeg1_1(fln string) map[string]float64 {
	fil, err := os.Open(fln)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", fln, err)
		os.Exit(1)
	}
	defer fil.Close()

	rdr := csv.NewReader(fil)

	// Skip header row
	_, err = rdr.Read()
	if err != nil {
		fmt.Printf("Error reading header: %v\n", err)
		os.Exit(1)
	}

	res := make(map[string]float64)

	for {
		rec, err := rdr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading CSV: %v\n", err)
			continue
		}

		if len(rec) < 3 {
			fmt.Printf("Invalid record format: %v\n", rec)
			continue
		}

		nam := rec[0]
		lgL, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			fmt.Printf("Error parsing LEG_LENGTH for %s: %v\n", nam, err)
			continue
		}

		res[nam] = lgL
	}

	return res
}

func readDsLeg2_1(fln string) map[string]struct {
	stL float64
	stn string
} {
	fil, err := os.Open(fln)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", fln, err)
		os.Exit(1)
	}
	defer fil.Close()

	rdr := csv.NewReader(fil)

	// Skip header row
	_, err = rdr.Read()
	if err != nil {
		fmt.Printf("Error reading header: %v\n", err)
		os.Exit(1)
	}

	res := make(map[string]struct {
		stL float64
		stn string
	})

	for {
		rec, err := rdr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading CSV: %v\n", err)
			continue
		}

		if len(rec) < 3 {
			fmt.Printf("Invalid record format: %v\n", rec)
			continue
		}

		nam := rec[0]
		stL, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			fmt.Printf("Error parsing STRIDE_LENGTH for %s: %v\n", nam, err)
			continue
		}

		stn := rec[2]

		res[nam] = struct {
			stL float64
			stn string
		}{stL, stn}
	}

	return res
}

func join1(ds1 map[string]float64, ds2 map[string]struct {
	stL float64
	stn string
}) []Dinosaur1 {
	var dns []Dinosaur1

	for nam, data2 := range ds2 {
		lgL, ok := ds1[nam]
		if !ok {
			fmt.Printf("Warning: Dinosaur %s not found in dataset1\n", nam)
			continue
		}

		// Calculate speed
		spd := calculateSpeed1(data2.stL, lgL)

		dns = append(dns, Dinosaur1{
			nam: nam,
			lgL: lgL,
			stL: data2.stL,
			stn: data2.stn,
			spd: spd,
		})
	}

	return dns
}

func filterBipedal1(dns []Dinosaur1) []Dinosaur1 {
	var bpd []Dinosaur1

	for _, d := range dns {
		if d.stn == "bipedal" {
			bpd = append(bpd, d)
		}
	}

	return bpd
}

func calculateSpeed1(stL, lgL float64) float64 {
	return ((stL / lgL) - 1) * math.Sqrt(lgL*gravitationalConstant1)
}

func TestDino1(t *testing.T) {
	testCases := []struct {
		name               string
		legContent         string
		strideContent      string
		expectedDinosaurs  []string
		expectedErrorCheck func(err error) bool
	}{
		{
			name: "Sample data from problem statement",
			legContent: `NAME,LEG_LENGTH,DIET
Hadrosaurus,1.4,herbivore
Struthiomimus,0.72,omnivore
Velociraptor,1.8,carnivore
Triceratops,0.47,herbivore
Euoplocephalus,2.6,herbivore
Stegosaurus,1.50,herbivore
Tyrannosaurus Rex,6.5,carnivore`,
			strideContent: `NAME,STRIDE_LENGTH,STANCE
Euoplocephalus,1.97,quadrupedal
Stegosaurus,1.70,quadrupedal
Tyrannosaurus Rex,4.76,bipedal
Hadrosaurus,1.3,bipedal
Deinonychus,1.11,bipedal
Struthiomimus,1.24,bipedal
Velociraptor,2.62,bipedal`,
			// Calculate expected order manually - bipedal dinosaurs sorted by speed
			expectedDinosaurs: []string{
				"Tyrannosaurus Rex",
				"Struthiomimus",
				"Hadrosaurus",
			},
			expectedErrorCheck: func(err error) bool { return err == nil },
		},
		{
			name:               "Empty datasets",
			legContent:         `NAME,LEG_LENGTH,DIET`,
			strideContent:      `NAME,STRIDE_LENGTH,STANCE`,
			expectedDinosaurs:  []string{},
			expectedErrorCheck: func(err error) bool { return err == nil },
		},
		{
			name: "No bipedal dinosaurs",
			legContent: `NAME,LEG_LENGTH,DIET
Euoplocephalus,2.6,herbivore
Stegosaurus,1.50,herbivore`,
			strideContent: `NAME,STRIDE_LENGTH,STANCE
Euoplocephalus,1.97,quadrupedal
Stegosaurus,1.70,quadrupedal`,
			expectedDinosaurs:  []string{},
			expectedErrorCheck: func(err error) bool { return err == nil },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create temporary files
			dataset1File, err := os.CreateTemp("", "dataset1*.csv")
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer os.Remove(dataset1File.Name())

			dataset2File, err := os.CreateTemp("", "dataset2*.csv")
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer os.Remove(dataset2File.Name())

			// Write test data to temp files
			if _, err := dataset1File.WriteString(tc.legContent); err != nil {
				t.Fatalf("Failed to write to temp file: %v", err)
			}
			if _, err := dataset2File.WriteString(tc.strideContent); err != nil {
				t.Fatalf("Failed to write to temp file: %v", err)
			}

			dataset1File.Close()
			dataset2File.Close()

			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Run the test
			ds1 := readDsLeg1_1(dataset1File.Name())
			ds2 := readDsLeg2_1(dataset2File.Name())
			dns := join1(ds1, ds2)
			bpd := filterBipedal1(dns)

			sort.Slice(bpd, func(i, j int) bool {
				return bpd[i].spd > bpd[j].spd
			})

			for _, d := range bpd {
				fmt.Println(d.nam)
			}

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			var buf [1024]byte
			n, _ := r.Read(buf[:])
			output := string(buf[:n])

			// Parse output into lines
			var actualDinosaurs []string
			for _, line := range strings.Split(strings.TrimSpace(output), "\n") {
				if line != "" {
					actualDinosaurs = append(actualDinosaurs, line)
				}
			}

			// Check results
			if !reflect.DeepEqual(actualDinosaurs, tc.expectedDinosaurs) {
				t.Errorf("Expected dinosaurs %v, got %v", tc.expectedDinosaurs, actualDinosaurs)
			}
		})
	}
}
