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
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/beevik/etree"
)

func (mtrk *MTrk) EncodeSMF(w io.Writer) error {
	length := int64(0)
	status := uint8(0)
	channel := uint8(0)
	for _, event := range mtrk.Events {
		appendLength, err := event.EncodeSMFLen(&status, &channel)
		if err != nil {
			return err
		}
		length += appendLength
		if length > 0xffffffff {
			return newSMFEncodeError(mtrk, errors.New("track too long"))
		}
	}

	var buf [8]byte
	copy(buf[:4], []byte{'M', 'T', 'r', 'k'})
	binary.BigEndian.PutUint32(buf[4:8], uint32(length))
	_, err := w.Write(buf[:])
	if err != nil {
		return err
	}
	status = uint8(0)
	channel = uint8(0)
	for _, event := range mtrk.Events {
		err = event.EncodeSMF(w, &status, &channel)
		if err != nil {
			return err
		}
	}
	return nil
}

func (mtrk *MTrk) EncodeXML() *etree.Element {
	el := etree.NewElement("MTrk")
	el.CreateAttr("pos", fmt.Sprintf("%#x", mtrk.FilePosition))
	for _, event := range mtrk.Events {
		el.AddChild(event.EncodeXML())
	}
	return el
}

func DecodeMTrkFromSMF(r io.ReadSeeker, warningCallback WarningCallback) (mtrk *MTrk, err error) {
	pos := tell(r)
	var buf [8]byte

	_, err = io.ReadFull(r, buf[:4])
	if err != nil {
		return
	}
	if !bytes.Equal(buf[:4], []byte{'M', 'T', 'r', 'k'}) {
		warningCallback(newSMFDecodeError(pos, errors.New("invalid MTrk chunk")))
		for {
			pos, err = r.Seek(-3, io.SeekCurrent)
			if err != nil {
				return
			}
			_, err = io.ReadFull(r, buf[:4])
			if err != nil {
				return
			}
			if bytes.Equal(buf[:4], []byte{'M', 'T', 'r', 'k'}) {
				break
			}
		}
	}

	_, err = io.ReadFull(r, buf[4:8])
	if err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return
	}
	length := binary.BigEndian.Uint32(buf[4:8])

	mtrk = &MTrk{
		FilePosition: pos,
		Events:       make([]Event, 0),
	}
	status := uint8(0x80)
	channel := uint8(0)

	// Strangely there are wild MIDI files with MTrk length == 0
	if length == 0 {
		warningCallback(newSMFDecodeError(pos+4, errors.New("MIDI track seems to contain no events")))
	} else {
		r, err = newLimitReadSeeker(r, int64(length))
		if err != nil {
			return
		}
	}

	for {
		pos = tell(r)

		// If length == 0, guess the actual length by searching for the next track
		if length == 0 {
			_, err = io.ReadFull(r, buf[:4])
			if err != nil {
				if err == io.EOF {
					err = nil
					return
				} else if err == io.ErrUnexpectedEOF {
					_, err = r.Seek(pos, io.SeekStart)
					if err != nil {
						return
					}
				} else {
					return
				}
			} else {
				_, err = r.Seek(pos, io.SeekStart)
				if err != nil {
					return
				}
				if bytes.Equal(buf[:], []byte{'M', 'T', 'r', 'k'}) {
					return
				}
			}
		}

		var event Event
		event, err = DecodeEventFromSMF(r, &status, &channel, warningCallback)
		if event != nil {
			mtrk.Events = append(mtrk.Events, event)
		}
		if err != nil {
			if err == io.EOF {
				if length != 0 && pos < mtrk.FilePosition+int64(length) {
					err = io.ErrUnexpectedEOF
				} else {
					err = nil
				}
			} else if err == io.ErrUnexpectedEOF {
				if length == 0 || pos < mtrk.FilePosition+int64(length) {
					err = io.ErrUnexpectedEOF
				} else {
					warningCallback(newSMFDecodeError(pos, errors.New("MIDI track is incomplete")))
					err = nil
				}
			}
			return
		}
	}
}

func DecodeMTrkFromXML(el *etree.Element) (*MTrk, error) {
	if el.Tag != "MTrk" {
		return nil, newXMLDecodeError(el, fmt.Errorf("expect an <MTrk> tag, but got <%s>", el.Tag))
	}
	pos, err := strconv.ParseInt(el.SelectAttrValue("pos", "0"), 0, 64)
	if err != nil {
		return nil, newXMLDecodeError(el, fmt.Errorf("invalid attribute for MTrk tag: pos=%q", el.SelectAttrValue("pos", "")))
	}
	events := make([]Event, 0)
	for _, child := range el.Child {
		if childEl, ok := child.(*etree.Element); ok {
			event, err := DecodeEventFromXML(childEl)
			if err != nil {
				return nil, err
			}
			events = append(events, event)
		}
	}
	return &MTrk{
		FilePosition: pos,
		Events:       events,
	}, nil
}
