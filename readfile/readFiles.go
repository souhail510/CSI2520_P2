/*
--------------------------------------------------------------

	student name: Souhail Daoudi
	student number: 300135458

--------------------------------------------------------------
*/
package main

import (
	"io/ioutil"
	"log"
	"strings"
)

var slice []string

func names() [][]string {

	files, err := ioutil.ReadDir(dataimagePath)
	if err != nil {
		log.Fatal(err)
	}

	// Create an array to store filenames
	var filenames []string

	// get the list of jpg files
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".jpg") {
			filenames = append(filenames, dataimagePath+file.Name())
		}
	}
	//dataimagePath+
	//create slices
	k := 1048
	p := len(filenames) / k
	numSlices := (len(filenames) + p - 1) / p
	slices := make([][]string, numSlices)

	for i := 0; i < len(filenames); i++ {
		sliceIndex := i / p
		slices[sliceIndex] = append(slices[sliceIndex], filenames[i])
	}
	return slices

}
