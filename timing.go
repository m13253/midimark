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
	"errors"
	"fmt"
	"sort"
	"time"
)

var ErrEventsNotSorted = errors.New("mark2midi: can not convert absolute time to delta time: events not sorted in advance")
var ErrDeltaToBig = errors.New("mark2midi: can not convert absolute time to delta time: value too big")

func (seq *Sequence) ConvertDeltaToAbsTick() {
	for _, mtrk := range seq.Tracks {
		mtrk.ConvertDeltaToAbsTick()
	}
}

func (seq *Sequence) ConvertAbsToDeltaTick() error {
	var err error
	for _, mtrk := range seq.Tracks {
		err = mtrk.ConvertAbsToDeltaTick()
		if err != nil {
			return err
		}
	}
	return nil
}

func (seq *Sequence) CalculateTempoTable() {
	var table *TempoTable
	for i, mtrk := range seq.Tracks {
		if i == 0 || seq.Header.Format == 2 {
			table = &TempoTable{
				Framerate: seq.Header.Framerate,
				Division:  seq.Header.Division,
			}
		}
		absTick := uint64(0)
		for _, event := range mtrk.Events {
			absTick += uint64(event.Common().DeltaTick)
			if tempoChange, ok := event.(*MetaEventSetTempo); ok {
				if len(table.Changes) == 0 || table.Changes[len(table.Changes)-1].AbsTick != absTick {
					table.Changes = append(table.Changes, TempoChange{
						AbsTick:      absTick,
						FilePosition: tempoChange.FilePosition,
						UsPerQuarter: tempoChange.UsPerQuarter,
					})
				} else {
					table.Changes[len(table.Changes)-1] = TempoChange{
						AbsTick:      absTick,
						FilePosition: tempoChange.FilePosition,
						UsPerQuarter: tempoChange.UsPerQuarter,
					}
				}
			}
		}
		mtrk.TempoTable = table
	}
	if table != nil && seq.Header.Format != 2 {
		sort.Slice(table.Changes, func(i, j int) bool {
			return table.Changes[i].AbsTick < table.Changes[j].AbsTick || (table.Changes[i].AbsTick == table.Changes[j].AbsTick && table.Changes[i].FilePosition < table.Changes[j].FilePosition)
		})
	}
}

func (seq *Sequence) CalculateNotePair() {
	if seq.Header.Format == 2 {
		for _, mtrk := range seq.Tracks {
			var noteOn [128][]*EventNoteOn
			for _, event := range mtrk.Events {
				switch ev := event.(type) {
				case *EventNoteOff:
					if ev.Key >= 0x80 || len(noteOn[ev.Key]) == 0 {
						ev.RelatedNoteOn = nil
						break
					}
					ev.RelatedNoteOn = noteOn[ev.Key][0]
					noteOn[ev.Key][0].RelatedNoteOff = ev
					copy(noteOn[ev.Key], noteOn[ev.Key][1:])
					noteOn[ev.Key] = noteOn[ev.Key][:len(noteOn[ev.Key])-1]
				case *EventNoteOn:
					ev.RelatedNoteOff = nil
					if ev.Key >= 0x80 || len(noteOn[ev.Key]) == 0 {
						break
					}
					noteOn[ev.Key] = append(noteOn[ev.Key], ev)
				case *EventPolyphonicKeyPressure:
					if ev.Key >= 0x80 || len(noteOn[ev.Key]) == 0 {
						ev.RelatedNoteOn = nil
						break
					}
					ev.RelatedNoteOn = make([]*EventNoteOn, len(noteOn[ev.Key]))
					copy(ev.RelatedNoteOn, noteOn[ev.Key])
				}
			}
		}
	} else {
		var noteOn [128][]*EventNoteOn
		i := make([]int, len(seq.Tracks))
		absTick := make([]uint64, len(seq.Tracks))
		for {
			var selectedEvent Event
			selectedTrack := 0
			selectedTick := uint64(0)
			for t, mtrk := range seq.Tracks {
				if i[t] < len(mtrk.Events) && (selectedEvent == nil || absTick[t]+uint64(mtrk.Events[i[t]].Common().DeltaTick) < selectedTick) {
					selectedEvent = mtrk.Events[i[t]]
					selectedTrack = t
					selectedTick = absTick[t] + uint64(selectedEvent.Common().DeltaTick)
					absTick[t] = selectedTick
				}
			}
			if selectedEvent == nil {
				break
			}
			i[selectedTrack]++

			switch ev := selectedEvent.(type) {
			case *EventNoteOff:
				if ev.Key >= 0x80 || len(noteOn[ev.Key]) == 0 {
					ev.RelatedNoteOn = nil
					break
				}
				ev.RelatedNoteOn = noteOn[ev.Key][0]
				noteOn[ev.Key][0].RelatedNoteOff = ev
				copy(noteOn[ev.Key], noteOn[ev.Key][1:])
				noteOn[ev.Key] = noteOn[ev.Key][:len(noteOn[ev.Key])-1]
			case *EventNoteOn:
				ev.RelatedNoteOff = nil
				if ev.Key >= 0x80 || len(noteOn[ev.Key]) == 0 {
					break
				}
				noteOn[ev.Key] = append(noteOn[ev.Key], ev)
			case *EventPolyphonicKeyPressure:
				if ev.Key >= 0x80 || len(noteOn[ev.Key]) == 0 {
					ev.RelatedNoteOn = nil
					break
				}
				ev.RelatedNoteOn = make([]*EventNoteOn, len(noteOn[ev.Key]))
				copy(ev.RelatedNoteOn, noteOn[ev.Key])
			}
		}
	}
}

func (mtrk *MTrk) ConvertDeltaToAbsTick() {
	absTick := uint64(0)
	for _, event := range mtrk.Events {
		evCommon := event.Common()
		delta := evCommon.DeltaTick
		absTick += uint64(delta)
		evCommon.AbsTick = absTick
	}
}

func (mtrk *MTrk) ConvertAbsToDeltaTick() error {
	absTick := uint64(0)
	for _, event := range mtrk.Events {
		evCommon := event.Common()
		if evCommon.AbsTick < absTick {
			return ErrEventsNotSorted
		}
		delta := evCommon.AbsTick - absTick
		if delta > MaxVLQ {
			return ErrDeltaToBig
		}
		evCommon.DeltaTick = VLQ(delta)
	}
	return nil
}

func (mtrk *MTrk) ConvertAbsTickToDuration(absTick uint64) time.Duration {
	if mtrk.TempoTable == nil {
		panic(errors.New("midimark: track does not contain a tempo table"))
	}
	if mtrk.TempoTable.Framerate != 0 {
		if mtrk.TempoTable.Framerate == 29 {
			return time.Duration(absTick) * (time.Second * 100) / (time.Duration(mtrk.TempoTable.Division) * 2997)
		} else {
			return time.Duration(absTick) * time.Second / (time.Duration(mtrk.TempoTable.Division) * time.Duration(mtrk.TempoTable.Framerate))
		}
	}
	lastChange := uint64(0)
	numerator := uint64(0)
	denominator := mtrk.TempoTable.Division
	usPerQuarter := uint32(500000)
	for i := 0; i < len(mtrk.TempoTable.Changes) && absTick < mtrk.TempoTable.Changes[i].AbsTick; i++ {
		numerator += (mtrk.TempoTable.Changes[i].AbsTick - lastChange) * uint64(usPerQuarter)
		lastChange = mtrk.TempoTable.Changes[i].AbsTick
		usPerQuarter = mtrk.TempoTable.Changes[i].UsPerQuarter
	}
	if absTick < lastChange {
		panic(fmt.Errorf("midimark: internal error: absTick - lastChange = %d - %d = %d", absTick, lastChange, absTick-lastChange))
	}
	numerator += (absTick - lastChange) * uint64(usPerQuarter)
	return time.Duration(numerator) * time.Microsecond / time.Duration(denominator)
}
