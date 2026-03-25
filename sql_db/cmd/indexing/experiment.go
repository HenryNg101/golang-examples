package main

type QueryExperiment struct {
	Name        string
	SetupSQL    []string
	QuerySQL    string
	TeardownSQL []string
}
