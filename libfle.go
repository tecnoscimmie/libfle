package libfle

import (
	"math/rand"
	"os"
	"strings"

	"github.com/fatih/color"
)

func getfles() []func(string) string {
	return []func(string) string{
		func(d string) string { return color.New(color.FgYellow).SprintFunc()(d) },
		func(d string) string { return color.New(color.FgBlue).SprintFunc()(d) },
		func(d string) string { return color.New(color.FgRed).SprintFunc()(d) },
		func(d string) string { return color.New(color.FgYellow).SprintFunc()(d) },
		func(d string) string { return color.New(color.FgMagenta).SprintFunc()(d) },
		func(d string) string { return color.New(color.FgCyan).SprintFunc()(d) },
		func(d string) string { return color.New(color.FgWhite).SprintFunc()(d) },
	}
}

func splitargs() string {
	args := os.Args[1:]
	s := ""

	for i := 0; i < len(args); i++ {
		s += args[i] + " "
	}

	return s
}

func stringtoslice(s string) []string {
	return strings.Split(s, "")
}

func NewFle(tofle string) string {
	flist := getfles()

	// argstring := splitargs()
	s := stringtoslice(tofle)
	result := ""

	for i := 0; i < len(s); i++ {
		result += flist[rand.Intn(len(flist))](s[i])
	}

	return result
}
