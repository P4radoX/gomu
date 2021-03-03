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
	"os"
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

// FlagSet struct represents the flags handler.
//
// It must be initialized with the NewFlagSet constructor.
type FlagSet struct {
	MinArgs int
	usage func()
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

// Add method stores a new flag in FlagSet's collection
// The flag can be accessible by using Get method
func (fs *FlagSet) Add(flg Flag) {
	fs.flags[flg.Who()] = flg
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
	flag.Usage = fs.usage
	
	// Declare flags
	for k,v := range fs.flags {
		switch v.Type() {
		case FLAGBOOL:
			flag.BoolVar(v.(*BoolFlag).Value, k, false, v.(*BoolFlag).Description)
		case FLAGINT:
			flag.IntVar(v.(*IntFlag).Value, k, 0, v.(*IntFlag).Description)
		case FLAGSTRING:
			flag.StringVar(v.(*StringFlag).Value, k, "", v.(*StringFlag).Description)
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