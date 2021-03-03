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

// Flag interface implemented by application flags
type Flag interface {
	Parsed() bool
	IsUnique() bool
	IsRequired() bool
	Name() string
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

// AddFlag method stores a new flag in FlagSet's collection
// The flag can be accessible by using Get method
func (fs *FlagSet) AddFlag(flg Flag) {
	fs.flags[flg.Name()] = flg
}

// Get method returns the Flag interface associated to his name from the FlagSet collection.
//
// If a key doesn't exists, the returned interface will be nil
func (fs *FlagSet) Get(flagName string) Flag {
	return fs.flags[flagName]
}