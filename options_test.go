package option_test

import (
	"fmt"
	"github.com/debudda/option"
	"strings"
)

type Writer struct {
	Name  string
	Age   int
	Alive bool
}

var options = option.Options[Writer]{
	option.O(Writer{Name: "Douglas Adams", Age: 49}),
	option.O[Writer](),
	option.O(Writer{Name: "Neil Gaiman", Age: 61, Alive: true}),
	option.O(Writer{Name: "Neal Stephenson", Age: 62, Alive: true}),
}

func ExampleMap() {
	sentences := option.Map(options, func(w Writer) string {
		tobe := "is"
		if !w.Alive {
			tobe = "was"
		}
		return fmt.Sprintf("%s %s %d", w.Name, tobe, w.Age)
	})
	for _, s := range sentences {
		fmt.Println(s)
	}
	// Output: Douglas Adams was 49
	// Neil Gaiman is 61
	// Neal Stephenson is 62
}

func ExampleFilter() {
	alive := option.Filter(options, func(w Writer) bool {
		return w.Alive
	})
	for _, w := range alive {
		fmt.Printf("%s is %d\n", w.Name, w.Age)
	}
	// Output: Neil Gaiman is 61
	// Neal Stephenson is 62
}

func ExampleOFilter() {
	options.OFilter(func(w Writer) bool {
		return w.Alive
	}).Each(func(w Writer) {
		fmt.Printf("%s is %d\n", w.Name, w.Age)
	})
	// Output: Neil Gaiman is 61
	// Neal Stephenson is 62
}

func ExampleFoldl() {
	b := option.Foldl(options, func(res strings.Builder, next Writer) strings.Builder {
		res.WriteString(next.Name)
		res.WriteString(" and ")
		return res
	}, strings.Builder{})
	fmt.Println(b.String())
	// Output: Douglas Adams and Neil Gaiman and Neal Stephenson and
}

func ExampleFoldr() {
	b := option.Foldr(options, func(res strings.Builder, next Writer) strings.Builder {
		res.WriteString(next.Name)
		res.WriteString(" and ")
		return res
	}, strings.Builder{})
	fmt.Println(b.String())
	// Output: Neal Stephenson and Neil Gaiman and Douglas Adams and
}

func ExampleFoldli() {
	b := option.Foldli(options, func(i int, res strings.Builder, next Writer) strings.Builder {
		res.WriteString(next.Name)
		if i != len(options)-1 {
			res.WriteString(" and ")
		}
		return res
	}, strings.Builder{})
	fmt.Println(b.String())
	// Output: Douglas Adams and Neil Gaiman and Neal Stephenson
}

func ExampleFoldri() {
	b := option.Foldri(options, func(i int, res strings.Builder, next Writer) strings.Builder {
		res.WriteString(next.Name)
		if i != len(options)-1 {
			res.WriteString(" and ")
		}
		return res
	}, strings.Builder{})
	fmt.Println(b.String())
	// Output: Neal Stephenson and Neil Gaiman and Douglas Adams
}
