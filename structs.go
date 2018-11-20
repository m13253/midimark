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
	"io"

	"github.com/beevik/etree"
)

type Event interface {
	Common() *EventCommon
	EncodeSMF(w io.Writer, status, channel *uint8) error
	EncodeSMFLen(status, channel *uint8) (int64, error)
	EncodeRealtime() ([]byte, error)
	EncodeXML() *etree.Element
	Status() uint8
}

type MetaEvent interface {
	Event
	MetaData() ([]byte, error)
	MetaLen() (VLQ, error)
	MetaType() uint8
}

type Sequence struct {
	Header    *MThd
	Tracks    []*MTrk
	Undecoded []byte
}

type MThd struct {
	FilePosition int64
	Format       uint16
	NTrks        uint16
	Framerate    uint8
	Division     uint16
	Undecoded    []byte
}

type MTrk struct {
	FilePosition int64
	TempoTable   *TempoTable
	Events       []Event
}

type EventCommon struct {
	FilePosition int64
	AbsTick      uint64
	DeltaTick    VLQ
	Channel      uint8
}

// 8n
type EventNoteOff struct {
	EventCommon
	Key           Key
	Velocity      uint8
	RelatedNoteOn *EventNoteOff
}

// 9n
type EventNoteOn struct {
	EventCommon
	Key            Key
	Velocity       uint8
	RelatedNoteOff *EventNoteOn
}

// An
type EventPolyphonicKeyPressure struct {
	EventCommon
	Key           Key
	Velocity      uint8
	RelatedNoteOn []*EventNoteOn
}

// Bn
type EventControlChange struct {
	EventCommon
	Control uint8
	Value   uint8
}

// Cn
type EventProgramChange struct {
	EventCommon
	Program uint8
}

// Dn
type EventChannelPressure struct {
	EventCommon
	Velocity uint8
}

// En
type EventPitchWheelChange struct {
	EventCommon
	Pitch int16
}

// F0
type EventSystemExclusive struct {
	EventCommon
	Data []byte
}

// F1
type EventTimeCodeQuarterFrame struct {
	EventCommon
	MessageType uint8
	Values      uint8
}

// F2
type EventSongPositionPointer struct {
	EventCommon
	SongPosition uint16
}

// F3
type EventSongSelect struct {
	EventCommon
	SongNumber uint8
}

// F6
type EventTuneRequest struct {
	EventCommon
}

// F7
type EventEscape struct {
	EventCommon
	Data []byte
}

// F8
type EventTimingClock struct {
	EventCommon
}

// FA
type EventStart struct {
	EventCommon
}

// FB
type EventContinue struct {
	EventCommon
}

// FC
type EventStop struct {
	EventCommon
}

// FE
type EventActiveSensing struct {
	EventCommon
}

// FF 00
type MetaEventSequenceNumber struct {
	EventCommon
	SequenceNumber *uint16
	Undecoded      []byte
}

// FF 01
type MetaEventTextEvent struct {
	EventCommon
	Text string
}

// FF 02
type MetaEventCopyrightNotice struct {
	EventCommon
	Text string
}

// FF 03
type MetaEventSequenceTrackName struct {
	EventCommon
	Text string
}

// FF 04
type MetaEventInstrumentName struct {
	EventCommon
	Text string
}

// FF 05
type MetaEventLyric struct {
	EventCommon
	Text string
}

// FF 06
type MetaEventMarker struct {
	EventCommon
	Text string
}

// FF 07
type MetaEventCuePoint struct {
	EventCommon
	Text string
}

// FF 08
type MetaEventProgramName struct {
	EventCommon
	Text string
}

// FF 09
type MetaEventDeviceName struct {
	EventCommon
	Text string
}

// FF 20
type MetaEventMIDIChannelPrefix struct {
	EventCommon
	ChannelPrefix uint8
	Undecoded     []byte
}

// FF 2F
type MetaEventEndOfTrack struct {
	EventCommon
	Undecoded []byte
}

// FF 51
type MetaEventSetTempo struct {
	EventCommon
	UsPerQuarter uint32
	Undecoded    []byte
}

// FF 54
type MetaEventSMPTEOffset struct {
	EventCommon
	Framerate  uint8
	ColorFrame bool
	Negative   bool
	Hours      uint8
	Minutes    uint8
	Seconds    uint8
	Frames     uint8
	Fractional uint8
	Undecoded  []byte
}

// FF 58
type MetaEventTimeSignature struct {
	EventCommon
	Numerator                        uint8
	Denominator                      uint8
	MIDIClocksPerMetronome           uint8
	ThirtySecondNotesPer24MIDIClocks uint8
	Undecoded                        []byte
}

// FF 59
type MetaEventKeySignature struct {
	EventCommon
	KeySignature KeySignature
	Undecoded    []byte
}

// FF 60
type MetaEventXMFPatchTypePrefix struct {
	EventCommon
	Param     uint8
	Undecoded []byte
}

// FF 7F
type MetaEventSequencerSpecific struct {
	EventCommon
	Data []byte
}

type MetaEventUnknown struct {
	EventCommon
	Type    uint8
	Unknown []byte
}

type EventUnknown struct {
	EventCommon
	Unknown []byte
}

type TempoTable struct {
	Framerate uint8
	Division  uint16
	Changes   []TempoChange
}

type TempoChange struct {
	AbsTick      uint64
	FilePosition int64
	UsPerQuarter uint32
}

func (ev *EventCommon) Common() *EventCommon {
	return ev
}
