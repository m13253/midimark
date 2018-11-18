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
	"io"

	"github.com/beevik/etree"
)

func (ev *EventCommon) encodeCommonXMLAttr(el *etree.Element) {
	el.CreateAttr("pos", fmt.Sprintf("%#x", ev.FilePosition))
	el.CreateAttr("tick", fmt.Sprintf("%d", ev.AbsTick))
	el.CreateAttr("delta", fmt.Sprintf("%d", ev.DeltaTick))
	if ev.Channel-1 < 16 {
		el.CreateAttr("channel", fmt.Sprintf("%d", ev.Channel))
	}
}

func (ev *EventNoteOff) Encode(w io.Writer, status, channel *uint8) error {
	if ev.Channel-1 >= 16 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid MIDI channel %d", ev.Channel))
	}
	if ev.Key >= 0x80 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid key code %d", ev.Key))
	}
	if ev.Velocity >= 0x80 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid velocity value %d", ev.Velocity))
	}
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	if *status == ev.Status() {
		_, err = w.Write([]byte{uint8(ev.Key), ev.Velocity})
	} else {
		*status = ev.Status()
		*channel = ev.Channel
		_, err = w.Write([]byte{ev.Status(), uint8(ev.Key), ev.Velocity})
	}
	return err
}

func (ev *EventNoteOff) EncodeLen(status, channel *uint8) (int64, error) {
	if ev.Channel-1 >= 16 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid MIDI channel %d", ev.Channel-1))
	}
	if ev.Key >= 0x80 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid key code %d", ev.Key))
	}
	if ev.Velocity >= 0x80 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid velocity value %d", ev.Velocity))
	}
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	if *status == ev.Status() {
		length += 2
	} else {
		*status = ev.Status()
		*channel = ev.Channel
		length += 3
	}
	return length, nil
}

func (ev *EventNoteOff) EncodeXML() *etree.Element {
	el := etree.NewElement("NoteOff")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("key", ev.Key.String())
	el.CreateAttr("velocity", fmt.Sprintf("%d", ev.Velocity))
	return el
}

func (ev *EventNoteOff) Status() uint8 {
	return 0x80 | (ev.Channel - 1&0x0f)
}

func (ev *EventNoteOn) Encode(w io.Writer, status, channel *uint8) error {
	if ev.Channel-1 >= 16 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid MIDI channel %d", ev.Channel))
	}
	if ev.Key >= 0x80 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid key code %d", ev.Key))
	}
	if ev.Velocity >= 0x80 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid velocity value %d", ev.Velocity))
	}
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	if *status == ev.Status() {
		_, err = w.Write([]byte{uint8(ev.Key), ev.Velocity})
	} else {
		*status = ev.Status()
		*channel = ev.Channel
		_, err = w.Write([]byte{ev.Status(), uint8(ev.Key), ev.Velocity})
	}
	return err
}

func (ev *EventNoteOn) EncodeLen(status, channel *uint8) (int64, error) {
	if ev.Channel-1 >= 16 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid MIDI channel %d", ev.Channel-1))
	}
	if ev.Key >= 0x80 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid key code %d", ev.Key))
	}
	if ev.Velocity >= 0x80 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid velocity value %d", ev.Velocity))
	}
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	if *status == ev.Status() {
		length += 2
	} else {
		*status = ev.Status()
		*channel = ev.Channel
		length += 3
	}
	return length, nil
}

func (ev *EventNoteOn) EncodeXML() *etree.Element {
	el := etree.NewElement("NoteOn")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("key", ev.Key.String())
	el.CreateAttr("velocity", fmt.Sprintf("%d", ev.Velocity))
	return el
}

func (ev *EventNoteOn) Status() uint8 {
	return 0x90 | (ev.Channel - 1&0x0f)
}

func (ev *EventPolyphonicKeyPressure) Encode(w io.Writer, status, channel *uint8) error {
	if ev.Channel-1 >= 16 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid MIDI channel %d", ev.Channel))
	}
	if ev.Key >= 0x80 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid key code %d", ev.Key))
	}
	if ev.Velocity >= 0x80 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid velocity value %d", ev.Velocity))
	}
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	if *status == ev.Status() {
		_, err = w.Write([]byte{uint8(ev.Key), ev.Velocity})
	} else {
		*status = ev.Status()
		*channel = ev.Channel
		_, err = w.Write([]byte{ev.Status(), uint8(ev.Key), ev.Velocity})
	}
	return err
}

func (ev *EventPolyphonicKeyPressure) EncodeLen(status, channel *uint8) (int64, error) {
	if ev.Channel-1 >= 16 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid MIDI channel %d", ev.Channel-1))
	}
	if ev.Key >= 0x80 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid key code %d", ev.Key))
	}
	if ev.Velocity >= 0x80 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid velocity value %d", ev.Velocity))
	}
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	if *status == ev.Status() {
		length += 2
	} else {
		*status = ev.Status()
		*channel = ev.Channel
		length += 3
	}
	return length, nil
}

func (ev *EventPolyphonicKeyPressure) EncodeXML() *etree.Element {
	el := etree.NewElement("KeyPressure")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("key", ev.Key.String())
	el.CreateAttr("velocity", fmt.Sprintf("%d", ev.Velocity))
	return el
}

func (ev *EventPolyphonicKeyPressure) Status() uint8 {
	return 0xa0 | (ev.Channel - 1&0x0f)
}

func (ev *EventControlChange) Encode(w io.Writer, status, channel *uint8) error {
	if ev.Channel-1 >= 16 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid MIDI channel %d", ev.Channel))
	}
	if ev.Control >= 0x80 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid control type %d", ev.Control))
	}
	if ev.Value >= 0x80 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid control value %d", ev.Value))
	}
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	if *status == ev.Status() {
		_, err = w.Write([]byte{ev.Control, ev.Value})
	} else {
		*status = ev.Status()
		*channel = ev.Channel
		_, err = w.Write([]byte{ev.Status(), ev.Control, ev.Value})
	}
	return err
}

func (ev *EventControlChange) EncodeLen(status, channel *uint8) (int64, error) {
	if ev.Channel-1 >= 16 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid MIDI channel %d", ev.Channel-1))
	}
	if ev.Control >= 0x80 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid control type %d", ev.Control))
	}
	if ev.Value >= 0x80 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid control value %d", ev.Value))
	}
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	if *status == ev.Status() {
		length += 2
	} else {
		*status = ev.Status()
		*channel = ev.Channel
		length += 3
	}
	return length, nil
}

func (ev *EventControlChange) EncodeXML() *etree.Element {
	el := etree.NewElement("ControlChange")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("control", fmt.Sprintf("%d", ev.Control))
	el.CreateAttr("value", fmt.Sprintf("%d", ev.Value))
	return el
}

func (ev *EventControlChange) Status() uint8 {
	return 0xb0 | (ev.Channel - 1&0x0f)
}

func (ev *EventProgramChange) Encode(w io.Writer, status, channel *uint8) error {
	if ev.Channel-1 >= 16 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid MIDI channel %d", ev.Channel))
	}
	if ev.Program-1 >= 0x80 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid program value %d", ev.Program))
	}
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	if *status == ev.Status() {
		_, err = w.Write([]byte{ev.Program - 1})
	} else {
		*status = ev.Status()
		*channel = ev.Channel
		_, err = w.Write([]byte{ev.Status(), ev.Program - 1})
	}
	return err
}

func (ev *EventProgramChange) EncodeLen(status, channel *uint8) (int64, error) {
	if ev.Channel-1 >= 16 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid MIDI channel %d", ev.Channel-1))
	}
	if ev.Program >= 0x80 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid program value %d", ev.Program))
	}
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	if *status == ev.Status() {
		length++
	} else {
		*status = ev.Status()
		*channel = ev.Channel
		length += 2
	}
	return length, nil
}

func (ev *EventProgramChange) EncodeXML() *etree.Element {
	el := etree.NewElement("NoteOff")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("program", fmt.Sprintf("%d", ev.Program))
	return el
}

func (ev *EventProgramChange) Status() uint8 {
	return 0xc0 | (ev.Channel - 1&0x0f)
}

func (ev *EventChannelPressure) Encode(w io.Writer, status, channel *uint8) error {
	if ev.Channel-1 >= 16 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid MIDI channel %d", ev.Channel))
	}
	if ev.Velocity >= 0x80 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid velocity value %d", ev.Velocity))
	}
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	if *status == ev.Status() {
		_, err = w.Write([]byte{ev.Velocity})
	} else {
		*status = ev.Status()
		*channel = ev.Channel
		_, err = w.Write([]byte{ev.Status(), ev.Velocity})
	}
	return err
}

func (ev *EventChannelPressure) EncodeLen(status, channel *uint8) (int64, error) {
	if ev.Channel-1 >= 16 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid MIDI channel %d", ev.Channel-1))
	}
	if ev.Velocity >= 0x80 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid velocity value %d", ev.Velocity))
	}
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	if *status == ev.Status() {
		length++
	} else {
		*status = ev.Status()
		*channel = ev.Channel
		length += 2
	}
	return length, nil
}

func (ev *EventChannelPressure) EncodeXML() *etree.Element {
	el := etree.NewElement("ChannelPressure")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("velocity", fmt.Sprintf("%d", ev.Velocity))
	return el
}

func (ev *EventChannelPressure) Status() uint8 {
	return 0xd0 | (ev.Channel - 1&0x0f)
}

func (ev *EventPitchWheelChange) Encode(w io.Writer, status, channel *uint8) error {
	if ev.Channel-1 >= 16 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid MIDI channel %d", ev.Channel))
	}
	if ev.Pitch >= 0x2000 || ev.Pitch < -0x2000 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid pitch value %d", ev.Pitch))
	}
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	if *status == ev.Status() {
		_, err = w.Write([]byte{uint8((ev.Pitch+0x2000)>>7) & 0x7f, uint8(ev.Pitch+0x2000) & 0x7f})
	} else {
		*status = ev.Status()
		*channel = ev.Channel
		_, err = w.Write([]byte{ev.Status(), uint8((ev.Pitch+0x2000)>>7) & 0x7f, uint8(ev.Pitch+0x2000) & 0x7f})
	}
	return err
}

func (ev *EventPitchWheelChange) EncodeLen(status, channel *uint8) (int64, error) {
	if ev.Channel-1 >= 16 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid MIDI channel %d", ev.Channel-1))
	}
	if ev.Pitch >= 0x2000 || ev.Pitch < -0x2000 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid pitch value %d", ev.Pitch))
	}
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	if *status == ev.Status() {
		length += 2
	} else {
		*status = ev.Status()
		*channel = ev.Channel
		length += 3
	}
	return length, nil
}

func (ev *EventPitchWheelChange) EncodeXML() *etree.Element {
	el := etree.NewElement("PitchWheel")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("pitch", fmt.Sprintf("%d", ev.Pitch))
	return el
}

func (ev *EventPitchWheelChange) Status() uint8 {
	return 0xe0 | (ev.Channel - 1&0x0f)
}

func (ev *EventSystemExclusive) Encode(w io.Writer, status, channel *uint8) error {
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		_, err = w.Write([]byte{0xff, 0x20, 0x01, ev.Channel - 1, 0x00})
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte{ev.Status()})
	if err != nil {
		return err
	}
	packetLen := len(ev.Data)
	if packetLen > MaxVLQ {
		packetLen = MaxVLQ
	}
	err = VLQ(packetLen).Encode(w)
	if err != nil {
		return err
	}
	_, err = w.Write(ev.Data[:packetLen])
	if err != nil {
		return err
	}
	for i := MaxVLQ; i < len(ev.Data); i += MaxVLQ {
		*status = 0xf7
		_, err = w.Write([]byte{0x00, 0xf7})
		if err != nil {
			return err
		}
		packetLen = len(ev.Data) - i
		if packetLen > MaxVLQ {
			packetLen = MaxVLQ
		}
		err = VLQ(packetLen).Encode(w)
		if err != nil {
			return err
		}
		_, err = w.Write(ev.Data[i : i+packetLen])
		if err != nil {
			return err
		}
	}
	return nil
}

func (ev *EventSystemExclusive) EncodeLen(status, channel *uint8) (int64, error) {
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		length += 5
	}
	length++
	packetLen := len(ev.Data)
	if packetLen > MaxVLQ {
		packetLen = MaxVLQ
	}
	length1, err := VLQ(packetLen).EncodeLen()
	if err != nil {
		return 0, err
	}
	length += length1
	length += int64(packetLen)
	for i := MaxVLQ; i < len(ev.Data); i += MaxVLQ {
		*status = 0xf7
		length += 2
		packetLen = len(ev.Data) - i
		if packetLen > MaxVLQ {
			packetLen = MaxVLQ
		}
		length1, err = VLQ(packetLen).EncodeLen()
		if err != nil {
			return 0, err
		}
		length += length1
		length += int64(packetLen)
	}
	return length, nil
}

func (ev *EventSystemExclusive) EncodeXML() *etree.Element {
	el := etree.NewElement("SysEx")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("data", fmt.Sprintf("% x", ev.Data))
	return el
}

func (ev *EventSystemExclusive) Status() uint8 {
	return 0xf0
}

func (ev *EventTimeCodeQuarterFrame) Encode(w io.Writer, status, channel *uint8) error {
	if ev.MessageType >= 0x8 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid message type %d", ev.MessageType))
	}
	if ev.Values >= 0x10 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid time code quarter frame value %d", ev.Values))
	}
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		_, err = w.Write([]byte{0xff, 0x20, 0x01, ev.Channel - 1, 0x00})
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte{ev.Status(), ev.MessageType<<4 | ev.Values})
	return err
}

func (ev *EventTimeCodeQuarterFrame) EncodeLen(status, channel *uint8) (int64, error) {
	if ev.MessageType >= 0x8 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid message type %d", ev.MessageType))
	}
	if ev.Values >= 0x10 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid time code quarter frame value %d", ev.Values))
	}
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		length += 5
	}
	length += 2
	return length, nil
}

func (ev *EventTimeCodeQuarterFrame) EncodeXML() *etree.Element {
	el := etree.NewElement("TimeCodeQuarterFrame")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("message-type", fmt.Sprintf("%d", ev.MessageType))
	el.CreateAttr("values", fmt.Sprintf("%d", ev.Values))
	return el
}

func (ev *EventTimeCodeQuarterFrame) Status() uint8 {
	return 0xf1
}

func (ev *EventSongPositionPointer) Encode(w io.Writer, status, channel *uint8) error {
	if ev.SongPosition >= 0x4000 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid song position %d", ev.SongPosition))
	}
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		_, err = w.Write([]byte{0xff, 0x20, 0x01, ev.Channel - 1, 0x00})
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte{ev.Status(), uint8(ev.SongPosition>>7) & 0x7f, uint8(ev.SongPosition) & 0x7f})
	return err
}

func (ev *EventSongPositionPointer) EncodeLen(status, channel *uint8) (int64, error) {
	if ev.SongPosition >= 0x4000 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid song position %d", ev.SongPosition))
	}
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		length += 5
	}
	length += 3
	return length, nil
}

func (ev *EventSongPositionPointer) EncodeXML() *etree.Element {
	el := etree.NewElement("SongPosition")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("song-position", fmt.Sprintf("%d", ev.SongPosition))
	return el
}

func (ev *EventSongPositionPointer) Status() uint8 {
	return 0xf2
}

func (ev *EventSongSelect) Encode(w io.Writer, status, channel *uint8) error {
	if ev.SongNumber >= 0x80 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid song number %d", ev.SongNumber))
	}
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	*status = 0xf3
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		_, err = w.Write([]byte{0xff, 0x20, 0x01, ev.Channel - 1, 0x00})
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte{ev.Status(), ev.SongNumber & 0x7f})
	return err
}

func (ev *EventSongSelect) EncodeLen(status, channel *uint8) (int64, error) {
	if ev.SongNumber >= 0x80 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid song number %d", ev.SongNumber))
	}
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	*status = 0xf3
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		length += 5
	}
	length += 2
	return length, nil
}

func (ev *EventSongSelect) EncodeXML() *etree.Element {
	el := etree.NewElement("SongSelect")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("song-number", fmt.Sprintf("%d", ev.SongNumber))
	return el
}

func (ev *EventSongSelect) Status() uint8 {
	return 0xf3
}

func (ev *EventTuneRequest) Encode(w io.Writer, status, channel *uint8) error {
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		_, err = w.Write([]byte{0xff, 0x20, 0x01, ev.Channel - 1, 0x00})
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte{*status})
	return err
}

func (ev *EventTuneRequest) EncodeLen(status, channel *uint8) (int64, error) {
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	*status = 0xf6
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		length += 5
	}
	length++
	return length, nil
}

func (ev *EventTuneRequest) EncodeXML() *etree.Element {
	el := etree.NewElement("TuneRequest")
	ev.encodeCommonXMLAttr(el)
	return el
}

func (ev *EventTuneRequest) Status() uint8 {
	return 0xf6
}

func (ev *EventEscape) Encode(w io.Writer, status, channel *uint8) error {
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		_, err = w.Write([]byte{0xff, 0x20, 0x01, ev.Channel - 1, 0x00})
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte{ev.Status()})
	if err != nil {
		return err
	}
	packetLen := len(ev.Data)
	if packetLen > MaxVLQ {
		packetLen = MaxVLQ
	}
	err = VLQ(packetLen).Encode(w)
	if err != nil {
		return err
	}
	_, err = w.Write(ev.Data[:packetLen])
	if err != nil {
		return err
	}
	for i := MaxVLQ; i < len(ev.Data); i += MaxVLQ {
		_, err = w.Write([]byte{0x00, 0xf7})
		if err != nil {
			return err
		}
		packetLen = len(ev.Data) - i
		if packetLen > MaxVLQ {
			packetLen = MaxVLQ
		}
		err = VLQ(packetLen).Encode(w)
		if err != nil {
			return err
		}
		_, err = w.Write(ev.Data[i : i+packetLen])
		if err != nil {
			return err
		}
	}
	return nil
}

func (ev *EventEscape) EncodeLen(status, channel *uint8) (int64, error) {
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		length += 5
	}
	length++
	packetLen := len(ev.Data)
	if packetLen > MaxVLQ {
		packetLen = MaxVLQ
	}
	length1, err := VLQ(packetLen).EncodeLen()
	if err != nil {
		return 0, err
	}
	length += length1
	length += int64(packetLen)
	for i := MaxVLQ; i < len(ev.Data); i += MaxVLQ {
		length += 2
		packetLen = len(ev.Data) - i
		if packetLen > MaxVLQ {
			packetLen = MaxVLQ
		}
		length1, err = VLQ(packetLen).EncodeLen()
		if err != nil {
			return 0, err
		}
		length += length1
		length += int64(packetLen)
	}
	return length, nil
}

func (ev *EventEscape) EncodeXML() *etree.Element {
	el := etree.NewElement("Escape")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("data", fmt.Sprintf("% x", ev.Data))
	return el
}

func (ev *EventEscape) Status() uint8 {
	return 0xf7
}

func (ev *EventTimingClock) Encode(w io.Writer, status, channel *uint8) error {
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		_, err = w.Write([]byte{0xff, 0x20, 0x01, ev.Channel - 1, 0x00})
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte{*status})
	return err
}

func (ev *EventTimingClock) EncodeLen(status, channel *uint8) (int64, error) {
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		length += 5
	}
	length++
	return length, nil
}

func (ev *EventTimingClock) EncodeXML() *etree.Element {
	el := etree.NewElement("TimingClock")
	ev.encodeCommonXMLAttr(el)
	return el
}

func (ev *EventTimingClock) Status() uint8 {
	return 0xf8
}

func (ev *EventStart) Encode(w io.Writer, status, channel *uint8) error {
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		_, err = w.Write([]byte{0xff, 0x20, 0x01, ev.Channel - 1, 0x00})
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte{*status})
	return err
}

func (ev *EventStart) EncodeLen(status, channel *uint8) (int64, error) {
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		length += 5
	}
	length++
	return length, nil
}

func (ev *EventStart) EncodeXML() *etree.Element {
	el := etree.NewElement("Start")
	ev.encodeCommonXMLAttr(el)
	return el
}

func (ev *EventStart) Status() uint8 {
	return 0xfa
}

func (ev *EventContinue) Encode(w io.Writer, status, channel *uint8) error {
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		_, err = w.Write([]byte{0xff, 0x20, 0x01, ev.Channel - 1, 0x00})
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte{*status})
	return err
}

func (ev *EventContinue) EncodeLen(status, channel *uint8) (int64, error) {
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		length += 5
	}
	length++
	return length, nil
}

func (ev *EventContinue) EncodeXML() *etree.Element {
	el := etree.NewElement("Continue")
	ev.encodeCommonXMLAttr(el)
	return el
}

func (ev *EventContinue) Status() uint8 {
	return 0xfb
}

func (ev *EventStop) Encode(w io.Writer, status, channel *uint8) error {
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		_, err = w.Write([]byte{0xff, 0x20, 0x01, ev.Channel - 1, 0x00})
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte{*status})
	return err
}

func (ev *EventStop) EncodeLen(status, channel *uint8) (int64, error) {
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		length += 5
	}
	length++
	return length, nil
}

func (ev *EventStop) EncodeXML() *etree.Element {
	el := etree.NewElement("Stop")
	ev.encodeCommonXMLAttr(el)
	return el
}

func (ev *EventStop) Status() uint8 {
	return 0xfc
}

func (ev *EventActiveSensing) Encode(w io.Writer, status, channel *uint8) error {
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		_, err = w.Write([]byte{0xff, 0x20, 0x01, ev.Channel - 1, 0x00})
		if err != nil {
			return err
		}
	}
	_, err = w.Write([]byte{*status})
	return err
}

func (ev *EventActiveSensing) EncodeLen(status, channel *uint8) (int64, error) {
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	*status = ev.Status()
	if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
		*channel = ev.Channel
		length += 5
	}
	length++
	return length, nil
}

func (ev *EventActiveSensing) EncodeXML() *etree.Element {
	el := etree.NewElement("ActiveSensing")
	ev.encodeCommonXMLAttr(el)
	return el
}

func (ev *EventActiveSensing) Status() uint8 {
	return 0xfe
}

func (ev *EventUnknown) Encode(w io.Writer, status, channel *uint8) error {
	if len(ev.RawData) == 0 {
		return nil
	}
	if ev.RawData[0] < 0x80 {
		return newSMFEncodeError(ev, fmt.Errorf("invalid status byte %#02x", ev.RawData[0]))
	}
	err := ev.DeltaTick.Encode(w)
	if err != nil {
		return err
	}
	if ev.Status() < 0xf0 {
		ev.Channel = (ev.Status() & 0x0f) + 1
		if *status == ev.Status() {
			_, err = w.Write(ev.RawData[1:])
		} else {
			*status = ev.Status()
			*channel = ev.Channel
			_, err = w.Write(ev.RawData)
		}
	} else {
		*status = ev.Status()
		if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
			*channel = ev.Channel
			_, err = w.Write([]byte{0xff, 0x20, 0x01, ev.Channel - 1, 0x00})
			if err != nil {
				return err
			}
		}
		_, err = w.Write(ev.RawData)
	}
	return err
}

func (ev *EventUnknown) EncodeLen(status, channel *uint8) (int64, error) {
	if len(ev.RawData) == 0 {
		return 0, nil
	}
	if ev.RawData[0] < 0x80 {
		return 0, newSMFEncodeError(ev, fmt.Errorf("invalid status byte %#02x", ev.RawData[0]))
	}
	length, err := ev.DeltaTick.EncodeLen()
	if err != nil {
		return 0, err
	}
	if ev.Status() < 0xf0 {
		ev.Channel = (ev.Status() & 0x0f) + 1
		if *status == ev.Status() {
			length += int64(len(ev.RawData) - 1)
		} else {
			*status = ev.Status()
			*channel = ev.Channel
			length += int64(len(ev.RawData))
		}
	} else {
		*status = ev.Status()
		if ev.Channel-1 < 16 && *channel-1 != ev.Channel-1 {
			*channel = ev.Channel
			length += 5
		}
		length += int64(len(ev.RawData))
	}
	return length, err
}

func (ev *EventUnknown) EncodeXML() *etree.Element {
	el := etree.NewElement("Event")
	ev.encodeCommonXMLAttr(el)
	el.CreateAttr("undecoded", fmt.Sprintf("% x", ev.RawData))
	return el
}

func (ev *EventUnknown) Status() uint8 {
	if len(ev.RawData) == 0 {
		panic(newSMFEncodeError(ev, errors.New("invalid MIDI event")))
	}
	if ev.RawData[0] < 0x80 {
		panic(newSMFEncodeError(ev, fmt.Errorf("invalid status byte %#02x", ev.RawData[0])))
	}
	return ev.RawData[0]
}

func DecodeEvent(r io.ReadSeeker, status, channel *uint8, warningCallback WarningCallback) (event Event, err error) {
	pos := tell(r)

	var delta VLQ
	delta, err = DecodeVLQ(r, warningCallback)
	if err != nil {
		return
	}

	var buf [3]byte
	_, err = io.ReadFull(r, buf[:1])
	if err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return
	}
	if buf[0] < 0x80 {
		if *status < 0x80 {
			warningCallback(newSMFDecodeError(pos, fmt.Errorf("invalid running status %#02x", *status)))
		}
		buf[0] = *status
		_, err = r.Seek(-1, io.SeekCurrent)
		if err != nil {
			return
		}
	} else {
		*status = buf[0]
	}
	if *status < 0xf0 {
		*channel = (*status & 0x0f) + 1
	}

	eventCommon := EventCommon{
		FilePosition: pos,
		DeltaTick:    delta,
		Channel:      *channel,
	}

	switch buf[0] & 0xf0 {
	case 0x80:
		_, err = io.ReadFull(r, buf[1:3])
		if err != nil {
			if err == io.EOF {
				err = io.ErrUnexpectedEOF
			}
			return
		}
		if buf[1] >= 0x80 || buf[2] >= 0x80 {
			warningCallback(newSMFDecodeError(pos, fmt.Errorf("invalid MIDI event % x", buf[:3])))
		}
		event = &EventNoteOff{
			EventCommon: eventCommon,
			Key:         Key(buf[1] & 0x7f),
			Velocity:    buf[2] & 0x7f,
		}
	case 0x90:
		_, err = io.ReadFull(r, buf[1:3])
		if err != nil {
			if err == io.EOF {
				err = io.ErrUnexpectedEOF
			}
			return
		}
		if buf[1] >= 0x80 || buf[2] >= 0x80 {
			warningCallback(newSMFDecodeError(pos, fmt.Errorf("invalid MIDI event % x", buf[:3])))
		}
		event = &EventNoteOn{
			EventCommon: eventCommon,
			Key:         Key(buf[1] & 0x7f),
			Velocity:    buf[2] & 0x7f,
		}
	case 0xa0:
		_, err = io.ReadFull(r, buf[1:3])
		if err != nil {
			if err == io.EOF {
				err = io.ErrUnexpectedEOF
			}
			return
		}
		if buf[1] >= 0x80 || buf[2] >= 0x80 {
			warningCallback(newSMFDecodeError(pos, fmt.Errorf("invalid MIDI event % x", buf[:3])))
		}
		event = &EventPolyphonicKeyPressure{
			EventCommon: eventCommon,
			Key:         Key(buf[1] & 0x7f),
			Velocity:    buf[2] & 0x7f,
		}
	case 0xb0:
		_, err = io.ReadFull(r, buf[1:3])
		if err != nil {
			if err == io.EOF {
				err = io.ErrUnexpectedEOF
			}
			return
		}
		if buf[1] >= 0x80 || buf[2] >= 0x80 {
			warningCallback(newSMFDecodeError(pos, fmt.Errorf("invalid MIDI event % x", buf[:3])))
		}
		event = &EventControlChange{
			EventCommon: eventCommon,
			Control:     buf[1] & 0x7f,
			Value:       buf[2] & 0x7f,
		}
	case 0xc0:
		_, err = io.ReadFull(r, buf[1:2])
		if err != nil {
			if err == io.EOF {
				err = io.ErrUnexpectedEOF
			}
			return
		}
		if buf[1] >= 0x80 {
			warningCallback(newSMFDecodeError(pos, fmt.Errorf("invalid MIDI event % x", buf[:2])))
		}
		event = &EventProgramChange{
			EventCommon: eventCommon,
			Program:     (buf[1] & 0x7f) + 1,
		}
	case 0xd0:
		_, err = io.ReadFull(r, buf[1:2])
		if err != nil {
			if err == io.EOF {
				err = io.ErrUnexpectedEOF
			}
			return
		}
		if buf[1] >= 0x80 {
			warningCallback(newSMFDecodeError(pos, fmt.Errorf("invalid MIDI event % x", buf[:2])))
		}
		event = &EventChannelPressure{
			EventCommon: eventCommon,
			Velocity:    buf[1] & 0x7f,
		}
	case 0xe0:
		_, err = io.ReadFull(r, buf[1:3])
		if err != nil {
			if err == io.EOF {
				err = io.ErrUnexpectedEOF
			}
			return
		}
		if buf[1] >= 0x80 || buf[2] >= 0x80 {
			warningCallback(newSMFDecodeError(pos, fmt.Errorf("invalid MIDI event % x", buf[:3])))
		}
		event = &EventPitchWheelChange{
			EventCommon: eventCommon,
			Pitch:       (int16(buf[1]&0x7f)<<7 | int16(buf[2]&0x7f)) - 0x2000,
		}
	default:
		switch buf[0] {
		case 0xf0:
			var length VLQ
			length, err = DecodeVLQ(r, warningCallback)
			if err != nil {
				return
			}
			sysex := &EventSystemExclusive{
				EventCommon: eventCommon,
				Data:        make([]byte, length),
			}
			event = sysex
			var n int
			n, err = io.ReadFull(r, sysex.Data)
			if n != int(length) {
				sysex.Data = append(sysex.Data[:n], 0xf7)
			}
		case 0xf1:
			_, err = io.ReadFull(r, buf[1:2])
			if err != nil {
				if err == io.EOF {
					err = io.ErrUnexpectedEOF
				}
				return
			}
			if buf[1] >= 0x80 {
				warningCallback(newSMFDecodeError(pos, fmt.Errorf("invalid MIDI event % x", buf[:2])))
			}
			event = &EventTimeCodeQuarterFrame{
				EventCommon: eventCommon,
				MessageType: (buf[1] >> 4) & 0x7,
				Values:      buf[1] & 0xf,
			}
		case 0xf2:
			_, err = io.ReadFull(r, buf[1:3])
			if err != nil {
				if err == io.EOF {
					err = io.ErrUnexpectedEOF
				}
				return
			}
			if buf[1] >= 0x80 || buf[2] >= 0x80 {
				warningCallback(newSMFDecodeError(pos, fmt.Errorf("invalid MIDI event % x", buf[:3])))
			}
			event = &EventSongPositionPointer{
				EventCommon:  eventCommon,
				SongPosition: uint16(buf[1]&0x7f)<<7 | uint16(buf[2]&0x7f),
			}
		case 0xf3:
			_, err = io.ReadFull(r, buf[1:2])
			if err != nil {
				if err == io.EOF {
					err = io.ErrUnexpectedEOF
				}
				return
			}
			if buf[1] >= 0x80 {
				warningCallback(newSMFDecodeError(pos, fmt.Errorf("invalid MIDI event % x", buf[:2])))
			}
			event = &EventSongSelect{
				EventCommon: eventCommon,
				SongNumber:  buf[1] & 0x7f,
			}
		case 0xf6:
			event = &EventTuneRequest{
				EventCommon: eventCommon,
			}
		case 0xf7:
			var length VLQ
			length, err = DecodeVLQ(r, warningCallback)
			if err != nil {
				return
			}
			escape := &EventEscape{
				EventCommon: eventCommon,
				Data:        make([]byte, length),
			}
			event = escape
			var n int
			n, err = io.ReadFull(r, escape.Data)
			if n == 0 {
				escape.Data = nil
			} else if n != int(length) {
				escape.Data = append(escape.Data[:n], 0xf7)
			}
		case 0xf8:
			event = &EventTimingClock{
				EventCommon: eventCommon,
			}
		case 0xfa:
			event = &EventStart{
				EventCommon: eventCommon,
			}
		case 0xfb:
			event = &EventContinue{
				EventCommon: eventCommon,
			}
		case 0xfc:
			event = &EventStop{
				EventCommon: eventCommon,
			}
		case 0xfe:
			event = &EventActiveSensing{
				EventCommon: eventCommon,
			}
		case 0xff:
			_, err = io.ReadFull(r, buf[1:2])
			if err != nil {
				if err == io.EOF {
					err = io.ErrUnexpectedEOF
				}
				return
			}
			if buf[1] >= 0x80 {
				warningCallback(newSMFDecodeError(pos, fmt.Errorf("invalid MIDI event % x", buf[:2])))
				return
			}
			var length VLQ
			length, err = DecodeVLQ(r, warningCallback)
			if err != nil {
				return
			}
			meta := &MetaEventUnknown{
				EventCommon: eventCommon,
				RawType:     buf[1] & 0x7f,
				RawData:     make([]byte, length),
			}
			var n int
			n, err = io.ReadFull(r, meta.RawData)
			meta.RawData = meta.RawData[:n]
			event = decodeMetaEvent(meta, warningCallback)
			if midiChannelPrefix, ok := event.(*MetaEventMIDIChannelPrefix); ok {
				*channel = midiChannelPrefix.Channel
			}
			if err != nil {
				return
			}
		default:
			event = &EventUnknown{
				EventCommon: eventCommon,
				RawData:     []byte{buf[0]},
			}
		}
	}
	return
}
