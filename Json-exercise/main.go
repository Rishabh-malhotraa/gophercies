package main

import (
	"encoding/json"
	"fmt"
)

/*
1. Marshal Unmarshal JSON data
2. Parse JSON from a URL
*/

type Person struct {
	FirstName      string
	lastName       string
	Age            int
	HeightinMeters float32
	IsMale         bool
}

var rishabh Person = Person{
	FirstName:      "Rishabh",
	lastName:       "Malhotra",
	Age:            21,
	HeightinMeters: 1.75,
	IsMale:         true,
}

type PersonMap map[string]interface{}

var rishabhMap PersonMap = PersonMap{
	"FirstName":      "Rishabh",
	"lastName":       "Malhotra",
	"Age":            21,
	"HeightinMeters": 1.75,
	"IsMale":         true,
}

// adding meta data to struct-tags

type Concepts struct {
	DynamicProgramming bool `json:"dp" yaml:"dp"`
	BinarySearch       bool `json:"bs" yaml:"bs"`
	UnionFind          bool `json:"dsu" yaml:"dsu"`
}

var topics Concepts = Concepts{true, false, true}

// ============================================

func encode() {

	var jsonData string = `
	{
		"FirstName": "John",
		"Age": 21,
		"Username": "johndoe91",
		"Grades": null,
		"Languages": [
			"English",
			"French"
		]
	}`

	jsonString := []byte(jsonData)

	isValid := json.Valid(jsonString)
	buffer, _ := json.MarshalIndent(topics, "", "  ")

	fmt.Println(string(buffer), isValid)
}

type Profile struct {
	Username string
	Follower string
}

type Student struct {
	FirstName, lastName string
	HeightinMeters      float64
	isMale              bool
	Languages           [2]string
	Subjects            []string
	Grades              map[string]string
	Profile             Profile
}

func decode() {

	var jsonString string = `
	{
		"FirstName": "John",
		"HeightInMeters": 1.75,
		"IsMale": null,
		"Languages": [ "English", "Spanish", "German" ],
		"Subjects": [ "Math", "Science" ],
		"Grades": { "Math": "A" },
		"Profile": {
			"Username": "johndoe91",
			"Followers": 1975
		}
	}`

	var Rishabh Student = Student{}

	_ = json.Unmarshal([]byte(jsonString), &Rishabh)

	fmt.Printf("%#v\n", Rishabh)
}

func main() {
	decode()
}
