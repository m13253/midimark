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
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/beevik/etree"
)

func (ev *MetaEventSequenceNumber) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventSequenceNumber) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventSequenceNumber) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventSequenceNumber) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	if ev.SequenceNumber != nil {
		el.CreateAttr("sequence-number", fmt.Sprintf("%d", *ev.SequenceNumber))
	}
	if len(ev.Undecoded) != 0 {
		el.CreateAttr("undecoded", fmt.Sprintf("% x", ev.Undecoded))
	}
	return el
}

func (ev *MetaEventSequenceNumber) MetaData() ([]byte, error) {
	if ev.SequenceNumber == nil {
		return ev.Undecoded, nil
	}
	data := make([]byte, 2+len(ev.Undecoded))
	binary.BigEndian.PutUint16(data[:2], *ev.SequenceNumber)
	copy(data[2:], ev.Undecoded)
	return data, nil
}

func (ev *MetaEventSequenceNumber) MetaLen() (VLQ, error) {
	length := len(ev.Undecoded)
	if ev.SequenceNumber != nil {
		length += 2
	}
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("text event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventSequenceNumber) Status() uint8 {
	return 0xff
}

func (ev *MetaEventSequenceNumber) MetaType() uint8 {
	return 0x00
}

func (ev *MetaEventTextEvent) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventTextEvent) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventTextEvent) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventTextEvent) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("text", ev.Text)
	return el
}

func (ev *MetaEventTextEvent) MetaData() ([]byte, error) {
	return []byte(ev.Text), nil
}

func (ev *MetaEventTextEvent) MetaLen() (VLQ, error) {
	length := len(ev.Text)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("text event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventTextEvent) Status() uint8 {
	return 0xff
}

func (ev *MetaEventTextEvent) MetaType() uint8 {
	return 0x01
}

func (ev *MetaEventCopyrightNotice) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventCopyrightNotice) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventCopyrightNotice) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventCopyrightNotice) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("text", ev.Text)
	return el
}

func (ev *MetaEventCopyrightNotice) MetaData() ([]byte, error) {
	return []byte(ev.Text), nil
}

func (ev *MetaEventCopyrightNotice) MetaLen() (VLQ, error) {
	length := len(ev.Text)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("text event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventCopyrightNotice) Status() uint8 {
	return 0xff
}

func (ev *MetaEventCopyrightNotice) MetaType() uint8 {
	return 0x02
}

func (ev *MetaEventSequenceTrackName) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventSequenceTrackName) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventSequenceTrackName) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventSequenceTrackName) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("text", ev.Text)
	return el
}

func (ev *MetaEventSequenceTrackName) MetaData() ([]byte, error) {
	return []byte(ev.Text), nil
}

func (ev *MetaEventSequenceTrackName) MetaLen() (VLQ, error) {
	length := len(ev.Text)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("text event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventSequenceTrackName) Status() uint8 {
	return 0xff
}

func (ev *MetaEventSequenceTrackName) MetaType() uint8 {
	return 0x03
}

func (ev *MetaEventInstrumentName) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventInstrumentName) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventInstrumentName) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventInstrumentName) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("text", ev.Text)
	return el
}

func (ev *MetaEventInstrumentName) MetaData() ([]byte, error) {
	return []byte(ev.Text), nil
}

func (ev *MetaEventInstrumentName) MetaLen() (VLQ, error) {
	length := len(ev.Text)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("text event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventInstrumentName) Status() uint8 {
	return 0xff
}

func (ev *MetaEventInstrumentName) MetaType() uint8 {
	return 0x04
}

func (ev *MetaEventLyric) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventLyric) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventLyric) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventLyric) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("text", ev.Text)
	return el
}

func (ev *MetaEventLyric) MetaData() ([]byte, error) {
	return []byte(ev.Text), nil
}

func (ev *MetaEventLyric) MetaLen() (VLQ, error) {
	length := len(ev.Text)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("text event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventLyric) Status() uint8 {
	return 0xff
}

func (ev *MetaEventLyric) MetaType() uint8 {
	return 0x05
}

func (ev *MetaEventMarker) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventMarker) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventMarker) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventMarker) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("text", ev.Text)
	return el
}

func (ev *MetaEventMarker) MetaData() ([]byte, error) {
	return []byte(ev.Text), nil
}

func (ev *MetaEventMarker) MetaLen() (VLQ, error) {
	length := len(ev.Text)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("text event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventMarker) Status() uint8 {
	return 0xff
}

func (ev *MetaEventMarker) MetaType() uint8 {
	return 0x06
}

func (ev *MetaEventCuePoint) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventCuePoint) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventCuePoint) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventCuePoint) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("text", ev.Text)
	return el
}

func (ev *MetaEventCuePoint) MetaData() ([]byte, error) {
	return []byte(ev.Text), nil
}

func (ev *MetaEventCuePoint) MetaLen() (VLQ, error) {
	length := len(ev.Text)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("text event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventCuePoint) Status() uint8 {
	return 0xff
}

func (ev *MetaEventCuePoint) MetaType() uint8 {
	return 0x07
}

func (ev *MetaEventProgramName) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventProgramName) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventProgramName) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventProgramName) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("text", ev.Text)
	return el
}

func (ev *MetaEventProgramName) MetaData() ([]byte, error) {
	return []byte(ev.Text), nil
}

func (ev *MetaEventProgramName) MetaLen() (VLQ, error) {
	length := len(ev.Text)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("text event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventProgramName) Status() uint8 {
	return 0xff
}

func (ev *MetaEventProgramName) MetaType() uint8 {
	return 0x08
}

func (ev *MetaEventDeviceName) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventDeviceName) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventDeviceName) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventDeviceName) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("text", ev.Text)
	return el
}

func (ev *MetaEventDeviceName) MetaData() ([]byte, error) {
	return []byte(ev.Text), nil
}

func (ev *MetaEventDeviceName) MetaLen() (VLQ, error) {
	length := len(ev.Text)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("text event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventDeviceName) Status() uint8 {
	return 0xff
}

func (ev *MetaEventDeviceName) MetaType() uint8 {
	return 0x09
}

func (ev *MetaEventMIDIChannelPrefix) EncodeSMF(w io.Writer, status, channel *uint8) error {
	ev.Channel = ev.ChannelPrefix
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventMIDIChannelPrefix) EncodeSMFLen(status, channel *uint8) (int64, error) {
	ev.Channel = ev.ChannelPrefix
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventMIDIChannelPrefix) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventMIDIChannelPrefix) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("channel-prefix", fmt.Sprintf("%d", ev.ChannelPrefix))
	if len(ev.Undecoded) != 0 {
		el.CreateAttr("undecoded", fmt.Sprintf("% x", ev.Undecoded))
	}
	return el
}

func (ev *MetaEventMIDIChannelPrefix) MetaData() ([]byte, error) {
	data := make([]byte, 1+len(ev.Undecoded))
	data[0] = ev.ChannelPrefix
	copy(data[1:], ev.Undecoded)
	return data, nil
}

func (ev *MetaEventMIDIChannelPrefix) MetaLen() (VLQ, error) {
	length := 1 + len(ev.Undecoded)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("meta event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventMIDIChannelPrefix) Status() uint8 {
	return 0xff
}

func (ev *MetaEventMIDIChannelPrefix) MetaType() uint8 {
	return 0x20
}

func (ev *MetaEventEndOfTrack) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventEndOfTrack) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventEndOfTrack) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventEndOfTrack) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	if len(ev.Undecoded) != 0 {
		el.CreateAttr("undecoded", fmt.Sprintf("% x", ev.Undecoded))
	}
	return el
}

func (ev *MetaEventEndOfTrack) MetaData() ([]byte, error) {
	return ev.Undecoded, nil
}

func (ev *MetaEventEndOfTrack) MetaLen() (VLQ, error) {
	length := len(ev.Undecoded)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("meta event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventEndOfTrack) Status() uint8 {
	return 0xff
}

func (ev *MetaEventEndOfTrack) MetaType() uint8 {
	return 0x2f
}

func (ev *MetaEventSetTempo) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventSetTempo) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventSetTempo) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventSetTempo) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("us-per-quarter", fmt.Sprintf("%d", ev.UsPerQuarter))
	if len(ev.Undecoded) != 0 {
		el.CreateAttr("undecoded", fmt.Sprintf("% x", ev.Undecoded))
	}
	return el
}

func (ev *MetaEventSetTempo) MetaData() ([]byte, error) {
	if ev.UsPerQuarter > 0x1000000 {
		return nil, newSMFEncodeError(ev, fmt.Errorf("invalid tempo value: %d\u03bcs per quarter note", ev.UsPerQuarter))
	}
	data := make([]byte, 4+len(ev.Undecoded))
	binary.BigEndian.PutUint32(data[:4], ev.UsPerQuarter)
	copy(data[4:], ev.Undecoded)
	return data[1:], nil
}

func (ev *MetaEventSetTempo) MetaLen() (VLQ, error) {
	length := 3 + len(ev.Undecoded)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("meta event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventSetTempo) Status() uint8 {
	return 0xff
}

func (ev *MetaEventSetTempo) MetaType() uint8 {
	return 0x51
}

func (ev *MetaEventSMPTEOffset) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventSMPTEOffset) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventSMPTEOffset) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventSMPTEOffset) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	if ev.ColorFrame {
		el.CreateAttr("color-frame", "yes")
	}
	if !ev.Negative {
		el.CreateAttr("timecode", fmt.Sprintf("%02d:%02d:%02d:%02d.%02d", ev.Hours, ev.Minutes, ev.Seconds, ev.Frames, ev.Fractional))
	} else {
		el.CreateAttr("timecode", fmt.Sprintf("-%02d:%02d:%02d:%02d.%02d", ev.Hours, ev.Minutes, ev.Seconds, ev.Frames, ev.Fractional))
	}
	if len(ev.Undecoded) != 0 {
		el.CreateAttr("undecoded", fmt.Sprintf("% x", ev.Undecoded))
	}
	return el
}

func (ev *MetaEventSMPTEOffset) MetaData() ([]byte, error) {
	timecodeType, ok := map[uint8]uint8{24: 0x00, 25: 0x20, 29: 0x40, 30: 0x60}[ev.Framerate]
	if !ok {
		return nil, newSMFEncodeError(ev, fmt.Errorf("invalid SMPTE framerate %d", ev.Framerate))
	}
	if ev.Hours >= 32 || ev.Minutes >= 64 || ev.Frames >= 64 {
		if !ev.Negative {
			return nil, newSMFEncodeError(ev, fmt.Errorf("invalid SMPTE time code %02d:%02d:%02d:%02d.%02d", ev.Hours, ev.Minutes, ev.Seconds, ev.Frames, ev.Fractional))
		} else {
			return nil, newSMFEncodeError(ev, fmt.Errorf("invalid SMPTE time code -%02d:%02d:%02d:%02d.%02d", ev.Hours, ev.Minutes, ev.Seconds, ev.Frames, ev.Fractional))
		}
	}
	data := make([]byte, 5+len(ev.Undecoded))
	data[0] = timecodeType | ev.Hours
	data[1] = ev.Minutes
	if ev.ColorFrame {
		data[1] |= 0x40
	}
	data[2] = ev.Seconds
	data[3] = ev.Frames
	if ev.Negative {
		data[3] |= 0x40
	}
	data[4] = ev.Fractional
	copy(data[5:], ev.Undecoded)
	return data, nil
}

func (ev *MetaEventSMPTEOffset) MetaLen() (VLQ, error) {
	length := 5 + len(ev.Undecoded)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("meta event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventSMPTEOffset) Status() uint8 {
	return 0xff
}

func (ev *MetaEventSMPTEOffset) MetaType() uint8 {
	return 0x54
}

func (ev *MetaEventTimeSignature) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventTimeSignature) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventTimeSignature) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventTimeSignature) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("numerator", fmt.Sprintf("%d", ev.Numerator))
	el.CreateAttr("denominator", fmt.Sprintf("%d", ev.Numerator))
	el.CreateAttr("midi-clocks-per-metronome", fmt.Sprintf("%d", ev.MIDIClocksPerMetronome))
	el.CreateAttr("thirty-second-notes-per-24-midi-clocks", fmt.Sprintf("%d", ev.ThirtySecondNotesPer24MIDIClocks))
	if len(ev.Undecoded) != 0 {
		el.CreateAttr("undecoded", fmt.Sprintf("% x", ev.Undecoded))
	}
	return el
}

func (ev *MetaEventTimeSignature) MetaData() ([]byte, error) {
	data := make([]byte, 4+len(ev.Undecoded))
	data[0] = ev.Numerator
	data[1] = ev.Denominator
	data[2] = ev.MIDIClocksPerMetronome
	data[3] = ev.ThirtySecondNotesPer24MIDIClocks
	copy(data[4:], ev.Undecoded)
	return data, nil
}

func (ev *MetaEventTimeSignature) MetaLen() (VLQ, error) {
	length := 4 + len(ev.Undecoded)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("meta event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventTimeSignature) Status() uint8 {
	return 0xff
}

func (ev *MetaEventTimeSignature) MetaType() uint8 {
	return 0x58
}

func (ev *MetaEventKeySignature) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventKeySignature) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventKeySignature) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventKeySignature) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("key-signature", ev.KeySignature.String())
	if len(ev.Undecoded) != 0 {
		el.CreateAttr("undecoded", fmt.Sprintf("% x", ev.Undecoded))
	}
	return el
}

func (ev *MetaEventKeySignature) MetaData() ([]byte, error) {
	data := make([]byte, 2+len(ev.Undecoded))
	binary.BigEndian.PutUint16(data[:2], uint16(ev.KeySignature))
	copy(data[2:], ev.Undecoded)
	return data, nil
}

func (ev *MetaEventKeySignature) MetaLen() (VLQ, error) {
	length := 2 + len(ev.Undecoded)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("meta event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventKeySignature) Status() uint8 {
	return 0xff
}

func (ev *MetaEventKeySignature) MetaType() uint8 {
	return 0x59
}

func (ev *MetaEventXMFPatchTypePrefix) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventXMFPatchTypePrefix) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventXMFPatchTypePrefix) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventXMFPatchTypePrefix) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("param", fmt.Sprintf("%d", ev.Param))
	if len(ev.Undecoded) != 0 {
		el.CreateAttr("undecoded", fmt.Sprintf("% x", ev.Undecoded))
	}
	return el
}

func (ev *MetaEventXMFPatchTypePrefix) MetaData() ([]byte, error) {
	data := make([]byte, 1+len(ev.Undecoded))
	data[0] = ev.Param
	copy(data[1:], ev.Undecoded)
	return data, nil
}

func (ev *MetaEventXMFPatchTypePrefix) MetaLen() (VLQ, error) {
	length := 1 + len(ev.Undecoded)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("meta event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventXMFPatchTypePrefix) Status() uint8 {
	return 0xff
}

func (ev *MetaEventXMFPatchTypePrefix) MetaType() uint8 {
	return 0x60
}

func (ev *MetaEventSequencerSpecific) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventSequencerSpecific) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventSequencerSpecific) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventSequencerSpecific) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("data", fmt.Sprintf("% x", ev.Data))
	return el
}

func (ev *MetaEventSequencerSpecific) MetaData() ([]byte, error) {
	return ev.Data, nil
}

func (ev *MetaEventSequencerSpecific) MetaLen() (VLQ, error) {
	length := len(ev.Data)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("meta event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventSequencerSpecific) Status() uint8 {
	return 0xff
}

func (ev *MetaEventSequencerSpecific) MetaType() uint8 {
	return 0x7f
}

func (ev *MetaEventUnknown) EncodeSMF(w io.Writer, status, channel *uint8) error {
	return encodeMetaEventSMF(ev, w, status, channel)
}

func (ev *MetaEventUnknown) EncodeSMFLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventSMFLen(ev, status, channel)
}

func (ev *MetaEventUnknown) EncodeRealtime() ([]byte, error) {
	return nil, nil
}

func (ev *MetaEventUnknown) EncodeXML() *etree.Element {
	el := etree.NewElement("Meta")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("type", fmt.Sprintf("%#02x", ev.MetaType()))
	el.CreateAttr("undecoded", fmt.Sprintf("% x", ev.RawData))
	return el
}

func (ev *MetaEventUnknown) MetaData() ([]byte, error) {
	return ev.RawData, nil
}

func (ev *MetaEventUnknown) MetaLen() (VLQ, error) {
	length := len(ev.RawData)
	if length > MaxVLQ {
		return 0, newSMFEncodeError(ev, errors.New("meta event too long"))
	}
	return VLQ(length), nil
}

func (ev *MetaEventUnknown) Status() uint8 {
	return 0xff
}

func (ev *MetaEventUnknown) MetaType() uint8 {
	return ev.RawType
}

func decodeMetaEvent(raw *MetaEventUnknown, warningCallback WarningCallback) Event {
	switch raw.MetaType() {
	case 0x00:
		if len(raw.RawData) == 0 {
			return &MetaEventSequenceNumber{
				EventCommon:    raw.EventCommon,
				SequenceNumber: nil,
			}
		}
		if len(raw.RawData) < 2 {
			warningCallback(newSMFDecodeError(raw.FilePosition, errors.New("incomplete meta event")))
			return raw
		}
		sequenceNumber := binary.BigEndian.Uint16(raw.RawData[:2])
		return &MetaEventSequenceNumber{
			EventCommon:    raw.EventCommon,
			SequenceNumber: &sequenceNumber,
			Undecoded:      raw.RawData[2:],
		}
	case 0x01:
		return &MetaEventTextEvent{
			EventCommon: raw.EventCommon,
			Text:        string(raw.RawData),
		}
	case 0x02:
		return &MetaEventCopyrightNotice{
			EventCommon: raw.EventCommon,
			Text:        string(raw.RawData),
		}
	case 0x03:
		return &MetaEventInstrumentName{
			EventCommon: raw.EventCommon,
			Text:        string(raw.RawData),
		}
	case 0x04:
		return &MetaEventTextEvent{
			EventCommon: raw.EventCommon,
			Text:        string(raw.RawData),
		}
	case 0x05:
		return &MetaEventLyric{
			EventCommon: raw.EventCommon,
			Text:        string(raw.RawData),
		}
	case 0x06:
		return &MetaEventMarker{
			EventCommon: raw.EventCommon,
			Text:        string(raw.RawData),
		}
	case 0x07:
		return &MetaEventCuePoint{
			EventCommon: raw.EventCommon,
			Text:        string(raw.RawData),
		}
	case 0x08:
		return &MetaEventProgramName{
			EventCommon: raw.EventCommon,
			Text:        string(raw.RawData),
		}
	case 0x09:
		return &MetaEventDeviceName{
			EventCommon: raw.EventCommon,
			Text:        string(raw.RawData),
		}
	case 0x20:
		if len(raw.RawData) < 1 {
			warningCallback(newSMFDecodeError(raw.FilePosition, errors.New("incomplete meta event")))
			return raw
		}
		raw.EventCommon.Channel = raw.RawData[0] + 1
		return &MetaEventMIDIChannelPrefix{
			EventCommon:   raw.EventCommon,
			ChannelPrefix: raw.EventCommon.Channel,
			Undecoded:     raw.RawData[1:],
		}
	case 0x2f:
		return &MetaEventEndOfTrack{
			EventCommon: raw.EventCommon,
			Undecoded:   raw.RawData,
		}
	case 0x51:
		if len(raw.RawData) < 3 {
			warningCallback(newSMFDecodeError(raw.FilePosition, errors.New("incomplete meta event")))
			return raw
		}
		return &MetaEventSetTempo{
			EventCommon:  raw.EventCommon,
			UsPerQuarter: uint32(raw.RawData[0])<<16 | uint32(raw.RawData[1])<<8 | uint32(raw.RawData[2]),
			Undecoded:    raw.RawData[3:],
		}
	case 0x54:
		if len(raw.RawData) < 5 {
			warningCallback(newSMFDecodeError(raw.FilePosition, errors.New("incomplete meta event")))
			return raw
		}
		event := &MetaEventSMPTEOffset{
			EventCommon: raw.EventCommon,
			Framerate:   map[uint8]uint8{0x00: 24, 0x20: 25, 0x40: 29, 0x60: 30}[raw.RawData[0]&0x60],
			Hours:       raw.RawData[0] & 0x1f,
			ColorFrame:  (raw.RawData[1] & 0x40) != 0,
			Minutes:     raw.RawData[1] & 0x3f,
			Seconds:     raw.RawData[2] & 0x7f,
			Negative:    (raw.RawData[3] & 0x40) != 0,
			Frames:      raw.RawData[3] & 0x3f,
			Fractional:  raw.RawData[4] & 0x7f,
			Undecoded:   raw.RawData[5:],
		}
		return event
	case 0x58:
		if len(raw.RawData) < 4 {
			warningCallback(newSMFDecodeError(raw.FilePosition, errors.New("incomplete meta event")))
			return raw
		}
		return &MetaEventTimeSignature{
			EventCommon:                      raw.EventCommon,
			Numerator:                        raw.RawData[0],
			Denominator:                      raw.RawData[1],
			MIDIClocksPerMetronome:           raw.RawData[2],
			ThirtySecondNotesPer24MIDIClocks: raw.RawData[3],
			Undecoded:                        raw.RawData[4:],
		}
	case 0x59:
		if len(raw.RawData) < 2 {
			warningCallback(newSMFDecodeError(raw.FilePosition, errors.New("incomplete meta event")))
			return raw
		}
		keySignature := KeySignature(binary.BigEndian.Uint16(raw.RawData[:2]))
		return &MetaEventKeySignature{
			EventCommon:  raw.EventCommon,
			KeySignature: keySignature,
			Undecoded:    raw.RawData[2:],
		}
	case 0x60:
		if len(raw.RawData) < 1 {
			warningCallback(newSMFDecodeError(raw.FilePosition, errors.New("incomplete meta event")))
			return raw
		}
		return &MetaEventXMFPatchTypePrefix{
			EventCommon: raw.EventCommon,
			Param:       raw.RawData[0],
			Undecoded:   raw.RawData[1:],
		}
	case 0x7f:
		return &MetaEventSequencerSpecific{
			EventCommon: raw.EventCommon,
			Data:        raw.RawData,
		}
	default:
		return raw
	}
}

func encodeMetaEventSMF(ev MetaEvent, w io.Writer, status, channel *uint8) error {
	evCommon := ev.Common()
	err := evCommon.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	*status = ev.Status()
	if evCommon.Channel-1 < 16 && *channel != evCommon.Channel {
		*channel = evCommon.Channel
		_, err = w.Write([]byte{0xff, 0x20, 0x01, evCommon.Channel, 0x00})
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte{ev.Status(), ev.MetaType()})
	if err != nil {
		return err
	}
	metaLen, err := ev.MetaLen()
	if err != nil {
		return err
	}
	err = metaLen.Encode(w)
	if err != nil {
		return err
	}
	metaData, err := ev.MetaData()
	if err != nil {
		return err
	}
	_, err = w.Write(metaData)
	return err
}

func encodeMetaEventSMFLen(ev MetaEvent, status, channel *uint8) (int64, error) {
	evCommon := ev.Common()
	length, err := evCommon.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	*status = ev.Status()
	if evCommon.Channel-1 < 16 && *channel != evCommon.Channel {
		*channel = evCommon.Channel
		length += 5
	}
	length += 2
	metaLen, err := ev.MetaLen()
	if err != nil {
		return 0, err
	}
	length1, err := metaLen.EncodeLen()
	if err != nil {
		return 0, err
	}
	length += length1
	length += int64(metaLen)
	return length, nil
}
