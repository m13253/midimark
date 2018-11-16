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
)

func (ev *MetaEventSequenceNumber) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventSequenceNumber) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventSequenceNumber) Type() uint8 {
	return 0x00
}

func (ev *MetaEventTextEvent) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventTextEvent) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventTextEvent) Type() uint8 {
	return 0x01
}

func (ev *MetaEventCopyrightNotice) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventCopyrightNotice) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventCopyrightNotice) Type() uint8 {
	return 0x02
}

func (ev *MetaEventSequenceTrackName) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventSequenceTrackName) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventSequenceTrackName) Type() uint8 {
	return 0x03
}

func (ev *MetaEventInstrumentName) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventInstrumentName) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventInstrumentName) Type() uint8 {
	return 0x04
}

func (ev *MetaEventLyric) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventLyric) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventLyric) Type() uint8 {
	return 0x05
}

func (ev *MetaEventMarker) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventMarker) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventMarker) Type() uint8 {
	return 0x06
}

func (ev *MetaEventCuePoint) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventCuePoint) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventCuePoint) Type() uint8 {
	return 0x07
}

func (ev *MetaEventProgramName) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventProgramName) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventProgramName) Type() uint8 {
	return 0x08
}

func (ev *MetaEventDeviceName) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventDeviceName) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventDeviceName) Type() uint8 {
	return 0x09
}

func (ev *MetaEventMIDIChannelPrefix) Encode(w io.Writer, status, channel *uint8) error {
	ev.Channel = ev.ChannelPrefix
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventMIDIChannelPrefix) EncodeLen(status, channel *uint8) (int64, error) {
	ev.Channel = ev.ChannelPrefix
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventMIDIChannelPrefix) Type() uint8 {
	return 0x20
}

func (ev *MetaEventEndOfTrack) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventEndOfTrack) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventEndOfTrack) Type() uint8 {
	return 0x2f
}

func (ev *MetaEventSetTempo) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventSetTempo) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventSetTempo) Type() uint8 {
	return 0x51
}

func (ev *MetaEventSMPTEOffset) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventSMPTEOffset) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
}

func (ev *MetaEventSMPTEOffset) MetaData() ([]byte, error) {
	data := make([]byte, 5+len(ev.Undecoded))
	copy(data[:5], ev.Timecode[:])
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

func (ev *MetaEventSMPTEOffset) Type() uint8 {
	return 0x54
}

func (ev *MetaEventTimeSignature) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventTimeSignature) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventTimeSignature) Type() uint8 {
	return 0x58
}

func (ev *MetaEventKeySignature) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventKeySignature) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventKeySignature) Type() uint8 {
	return 0x59
}

func (ev *MetaEventXMFPatchTypePrefix) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventXMFPatchTypePrefix) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventXMFPatchTypePrefix) Type() uint8 {
	return 0x60
}

func (ev *MetaEventSequencerSpecific) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventSequencerSpecific) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventSequencerSpecific) Type() uint8 {
	return 0x7f
}

func (ev *MetaEventUnknown) Encode(w io.Writer, status, channel *uint8) error {
	return encodeMetaEvent(ev, w, status, channel)
}

func (ev *MetaEventUnknown) EncodeLen(status, channel *uint8) (int64, error) {
	return encodeMetaEventLen(ev, status, channel)
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

func (ev *MetaEventUnknown) Type() uint8 {
	return ev.RawType
}

func decodeMetaEvent(raw *MetaEventUnknown, warningCallback WarningCallback) Event {
	switch raw.Type() {
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
		raw.EventCommon.Channel = raw.RawData[0]
		return &MetaEventMIDIChannelPrefix{
			EventCommon:   raw.EventCommon,
			ChannelPrefix: raw.RawData[0],
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
			Undecoded:   raw.RawData[5:],
		}
		copy(event.Timecode[:], raw.RawData[:5])
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

func encodeMetaEvent(ev MetaEvent, w io.Writer, status, channel *uint8) error {
	evCommon := ev.Common()
	err := evCommon.DeltaTime.Encode(w)
	if err != nil {
		return err
	}
	*status = ev.Status()
	if evCommon.Channel < 16 && *channel != evCommon.Channel {
		*channel = evCommon.Channel
		_, err = w.Write([]byte{0xff, 0x20, 0x01, evCommon.Channel, 0x00})
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte{ev.Status(), ev.Type()})
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

func encodeMetaEventLen(ev MetaEvent, status, channel *uint8) (int64, error) {
	evCommon := ev.Common()
	length, err := evCommon.DeltaTime.EncodeLen()
	if err != nil {
		return 0, err
	}
	*status = ev.Status()
	if evCommon.Channel < 16 && *channel != evCommon.Channel {
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
