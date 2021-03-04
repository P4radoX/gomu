// Copyright (c) 2021 P4radoX

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package flags

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// FlagType enumeration
const (
	UNKNOWN int = iota
	FLAGBOOL
	FLAGINT
	FLAGSTRING
)

// Flag interface implemented by application flags
type Flag interface {
	Parsed() bool
	IsUnique() bool
	IsRequired() bool
	Who() string
	Type() int
}

// Usage function is the default usage message set at FlagSet.Parse() time
func Usage(flags ...Flag) func() {
	msg := fmt.Sprintf("%s\n\nUsage:\n", strings.TrimPrefix(os.Args[0], "./"))

	for _, flag := range flags {
		switch flag.Type() {
		case FLAGBOOL:
			msg += fmt.Sprintf("\t-%s {%s}:\t%s\n", flag.Who(), "bool", flag.(*BoolFlag).Description)
		case FLAGINT:
			msg += fmt.Sprintf("\t-%s {%s}:\t%s\n", flag.Who(), "int", flag.(*IntFlag).Description)
		case FLAGSTRING:
			msg += fmt.Sprintf("\t-%s {%s}:\t%s\n", flag.Who(), "string", flag.(*StringFlag).Description)
		default:
		}
	}

	msg += "\t-help, --help:\tDisplays this message\n\n"

	return func() {fmt.Fprintf(os.Stdout, msg)}
}

// FlagSet struct represents the flags handler.
//
// It must be initialized with the NewFlagSet constructor.
type FlagSet struct {
	MinArgs int
	flags map[string]Flag
}

// NewFlagSet function is the FlagSet struct constructor.
//
// It returns a FlagSet struct pointer with the initialized collection
func NewFlagSet() *FlagSet {
	return &FlagSet{
		flags: make(map[string]Flag),
	}
}

// Add method stores severals new flag in FlagSet's collection.
// If a flag already exists, it will be simply discarded.
// The flag can be accessible by using Get method
func (fs *FlagSet) Add(flags ...Flag) {
	for _, flag := range flags {
		if _, ok := fs.flags[flag.Who()]; !ok {
			fs.flags[flag.Who()] = flag
		}
	}
}

// Get method returns the Flag interface associated to his name from the FlagSet collection.
//
// If a key doesn't exists, the returned interface will be nil
func (fs *FlagSet) Get(flagName string) Flag {
	return fs.flags[flagName]
}

// Parse method
func (fs *FlagSet) Parse() {
	// Set usage
	flag.Usage = Usage(fs.Flags()...)
	
	// Declare flags
	for k,v := range fs.flags {
		switch v.Type() {
		case FLAGBOOL:
			flag.BoolVar(&v.(*BoolFlag).Value, k, false, v.(*BoolFlag).Description)
		case FLAGINT:
			flag.IntVar(&v.(*IntFlag).Value, k, 0, v.(*IntFlag).Description)
		case FLAGSTRING:
			flag.StringVar(&v.(*StringFlag).Value, k, "", v.(*StringFlag).Description)
		default:
			// Handle non-typed flag
		}
	}

	// Parse
	flag.Parse()

	// Check min args
	if fs.MinArgs > 0 {
		if len(os.Args) < fs.MinArgs {
			flag.Usage()
			os.Exit(2)
		}
	}
}

// Flags method returns all registered flags as a Flag interface slice
func (fs *FlagSet) Flags() (lst []Flag) {
	for _, flag := range fs.flags {
		lst = append(lst, flag)
	}

	return
}