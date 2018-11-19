/*
  MIT License

  Copyright (c) 2018 Star Brilliant

  Permission is hereby granted, free of charge, to any person obtaining a copy
  of this software and associated documentation files (the "Software"), to deal
  in the Software without restriction, including without limitation the rights
  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
  copies of the Software, and to permit persons to whom the Software is
  furnished to do so, subject to the following conditions:

  The above copyright notice and this permission notice shall be included in
  all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
  SOFTWARE.
*/

package midimark

import (
	"fmt"

	"github.com/beevik/etree"
)

type WarningCallback func(err error)

func IgnoreWarnings(err error) {}

type ErrSMFEncode struct {
	Obj interface{}
	Err error
}

type ErrSMFDecode struct {
	Pos int64
	Err error
}

type ErrXMLDecode struct {
	Obj etree.Token
	Err error
}

func newSMFEncodeError(obj interface{}, err error) *ErrSMFEncode {
	return &ErrSMFEncode{
		Obj: obj,
		Err: err,
	}
}

func newSMFDecodeError(pos int64, err error) *ErrSMFDecode {
	return &ErrSMFDecode{
		Pos: pos,
		Err: err,
	}
}

func newXMLDecodeError(Obj etree.Token, err error) *ErrXMLDecode {
	return &ErrXMLDecode{
		Obj: Obj,
		Err: err,
	}
}

func (e *ErrSMFEncode) Error() string {
	return fmt.Sprintf("MIDI encode error: %v", e.Err)
}

func (e *ErrSMFDecode) Error() string {
	if e.Pos < 0 {
		return fmt.Sprintf("MIDI decode error: %v", e.Err)
	}
	return fmt.Sprintf("MIDI decode error at %#x: %v", e.Pos, e.Err)
}

func (e *ErrXMLDecode) Error() string {
	return fmt.Sprintf("MIDI Markup decode error: %v", e.Err)
}
