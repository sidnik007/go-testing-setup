package main

import (
	"flag"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"os"
	"testing"
)

var opt = godog.Options{
	Output: colors.Colored(os.Stdout),
	Paths:  []string{"features"},
}

func thereIsAWhichCosts(arg1 string, arg2 int) error {
	fmt.Print(arg1, arg2)
	return nil
}

func iAddTheToTheBasket(arg1 string) error {
	fmt.Print(arg1)
	return nil
}

func iShouldHaveProductsInTheBasket(arg1 int) error {
	fmt.Print(arg1)
	return nil
}

func theOverallBasketPriceShouldBe(arg1 int) error {
	fmt.Print(arg1)
	return nil
}

func featureContext(s *godog.Suite) {
	s.Step(`^there is a "([^"]*)", which costs £(\d+)$`, thereIsAWhichCosts)
	s.Step(`^I add the "([^"]*)" to the basket$`, iAddTheToTheBasket)
	s.Step(`^I should have (\d+) products in the basket$`, iShouldHaveProductsInTheBasket)
	s.Step(`^the overall basket price should be £(\d+)$`, theOverallBasketPriceShouldBe)
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	opt.Paths = flag.Args()
	status := godog.RunWithOptions("acceptance", func(s *godog.Suite) {
		featureContext(s)
	}, opt)
	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
