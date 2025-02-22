package main

import (
	"flag"
	"github.com/cucumber/godog"
	s "playwraight-golang/stepdefinitions"
	"testing"
)

var opts = godog.Options{}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestFeatures(t *testing.T) {
	opts := godog.Options{
		Format:   "pretty,cucumber:./report/cucumber.json",
		Paths:    []string{"features"},
		TestingT: t,
		// `Concurrency: 2` sets the number of concurrent goroutines to be used for running the tests. In this
		// case, it is set to 2, which means that two tests can be run concurrently. This can help speed up the
		// test execution time.
		Concurrency: 2,
	}

	flag.Parse()

	suite := godog.TestSuite{
		ScenarioInitializer: s.InitializeScenario,
		Options:             &opts,
	}

	if suite.Run() != 0 {
		t.Fatal("Error")
	}

}
