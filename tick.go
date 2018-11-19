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

import "errors"

var ErrEventsNotSorted = errors.New("mark2midi: can not convert absolute time to delta time: events not sorted in advance")
var ErrDeltaToBig = errors.New("mark2midi: can not convert absolute time to delta time: value too big")

func (mthd *MThd) ConvertDeltaToAbsTick() {
	for _, mtrk := range mthd.Tracks {
		mtrk.ConvertDeltaToAbsTick()
	}
}

func (mthd *MThd) ConvertAbsToDeltaTick() error {
	var err error
	for _, mtrk := range mthd.Tracks {
		err = mtrk.ConvertAbsToDeltaTick()
		if err != nil {
			return err
		}
	}
	return nil
}

func (mtrk *MTrk) ConvertDeltaToAbsTick() {
	tick := uint64(0)
	for _, event := range mtrk.Events {
		evCommon := event.Common()
		delta := evCommon.DeltaTick
		tick += uint64(delta)
		evCommon.AbsTick = tick
	}
}

func (mtrk *MTrk) ConvertAbsToDeltaTick() error {
	tick := uint64(0)
	for _, event := range mtrk.Events {
		evCommon := event.Common()
		if evCommon.AbsTick < tick {
			return ErrEventsNotSorted
		}
		delta := evCommon.AbsTick - tick
		if delta > MaxVLQ {
			return ErrDeltaToBig
		}
		evCommon.DeltaTick = VLQ(delta)
	}
	return nil
}
