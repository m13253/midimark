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

func (mthd *MThd) EncodeSMF(w io.Writer) error {
	if len(mthd.Undecoded) > 0xffffffff-6 {
		return newSMFEncodeError(mthd, errors.New("MThd chunk is too large"))
	}
	if len(mthd.Tracks) < 0xffff {
		mthd.NTrks = uint16(len(mthd.Tracks))
	} else {
		mthd.NTrks = 0xffff
	}
	if mthd.Framerate == 0 {
		if mthd.Division >= 0x8000 {
			return newSMFEncodeError(mthd, fmt.Errorf("invalid division %d", mthd.Division))
		}
	} else {
		if mthd.Framerate >= 0x80 {
			return newSMFEncodeError(mthd, fmt.Errorf("invalid SMPTE framerate %d", mthd.Framerate))
		}
		if mthd.Division >= 0x100 {
			return newSMFEncodeError(mthd, fmt.Errorf("invalid SMPTE division %d", mthd.Division))
		}
	}
	var buf [14]byte
	copy(buf[:4], []byte{'M', 'T', 'h', 'd'})
	binary.BigEndian.PutUint32(buf[4:8], uint32(6+len(mthd.Undecoded)))
	binary.BigEndian.PutUint16(buf[8:10], mthd.Format)
	binary.BigEndian.PutUint16(buf[10:12], mthd.NTrks)
	binary.BigEndian.PutUint16(buf[12:14], uint16(-mthd.Framerate)<<8|mthd.Division)
	_, err := w.Write(buf[:])
	if err != nil {
		return err
	}
	_, err = w.Write(mthd.Undecoded)
	if err != nil {
		return err
	}
	for _, mtrk := range mthd.Tracks {
		err = mtrk.EncodeSMF(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (mthd *MThd) EncodeXML() *etree.Element {
	el := etree.NewElement("MThd")
	el.CreateAttr("pos", fmt.Sprintf("%#x", mthd.FilePosition))
	el.CreateAttr("format", fmt.Sprintf("%d", mthd.Format))
	el.CreateAttr("ntrks", fmt.Sprintf("%d", mthd.NTrks))
	el.CreateAttr("division", fmt.Sprintf("%d", mthd.Division))
	for _, mtrk := range mthd.Tracks {
		el.AddChild(mtrk.EncodeXML())
	}
	return el
}

func DecodeMThdFromSMF(r io.ReadSeeker, warningCallback WarningCallback) (mthd *MThd, err error) {
	pos := tell(r)
	var buf [14]byte

	_, err = io.ReadFull(r, buf[:4])
	if err != nil {
		return
	}
	if !bytes.Equal(buf[:4], []byte{'M', 'T', 'h', 'd'}) {
		warningCallback(newSMFDecodeError(pos, errors.New("invalid MThd chunk")))
		for {
			pos, err = r.Seek(-3, io.SeekCurrent)
			if err != nil {
				return
			}
			_, err = io.ReadFull(r, buf[:4])
			if err != nil {
				if err == io.EOF || err == io.ErrUnexpectedEOF {
					warningCallback(newSMFDecodeError(pos, errors.New("not a standard MIDI file")))
				}
				return
			}
			if bytes.Equal(buf[:4], []byte{'M', 'T', 'h', 'd'}) {
				break
			}
		}
	}

	_, err = io.ReadFull(r, buf[4:14])
	if err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return
	}
	length := binary.BigEndian.Uint32(buf[4:8])
	framerate := uint8(0)
	division := binary.BigEndian.Uint16(buf[12:14])
	if division >= 0x8000 {
		framerate, division = uint8((0xff-division)>>8), division&0xff
	}
	mthd = &MThd{
		FilePosition: pos,
		Format:       binary.BigEndian.Uint16(buf[8:10]),
		NTrks:        binary.BigEndian.Uint16(buf[10:12]),
		Framerate:    framerate,
		Division:     division,
		Tracks:       make([]*MTrk, 0),
	}

	// Strangely there are wild MIDI files with MThd length == 0
	if length < 6 {
		warningCallback(newSMFDecodeError(pos+4, fmt.Errorf("invalid MThd chunk length %d", length)))
		pos, err = r.Seek(int64(length)-6, io.SeekCurrent)
		if err != nil {
			return
		}
	} else if length > 6 {
		mthd.Undecoded = make([]byte, length-6)
		var n int
		n, err = io.ReadFull(r, mthd.Undecoded)
		mthd.Undecoded = mthd.Undecoded[:n]
		if err != nil {
			return
		}
	}

	if mthd.NTrks == 0 {
		warningCallback(newSMFDecodeError(pos+10, errors.New("MIDI file seems to contain no tracks")))
	}

	for i := uint16(0); mthd.NTrks == 0 || i < mthd.NTrks; i++ {
		var mtrk *MTrk
		mtrk, err = DecodeMTrkFromSMF(r, warningCallback)
		if mtrk != nil {
			mthd.Tracks = append(mthd.Tracks, mtrk)
		}
		if err != nil {
			if err == io.EOF {
				if mthd.NTrks != 0 {
					err = io.ErrUnexpectedEOF
				} else {
					err = nil
				}
			}
			return
		}
	}

	return
}

func DecodeMThdFromXML(el *etree.Element) (*MThd, error) {
	if el.Tag != "MThd" {
		return nil, newXMLDecodeError(el, errors.New("can not find an MThd tag"))
	}
	pos, err := strconv.ParseInt(el.SelectAttrValue("pos", "0"), 0, 64)
	if err != nil {
		return nil, newXMLDecodeError(el, fmt.Errorf("invalid attribute: pos=%q", el.SelectAttrValue("pos", "")))
	}
	format, err := strconv.ParseUint(el.SelectAttrValue("format", ""), 0, 16)
	if err != nil {
		return nil, newXMLDecodeError(el, fmt.Errorf("invalid attribute: format=%q", el.SelectAttrValue("format", "")))
	}
	ntrks, err := strconv.ParseUint(el.SelectAttrValue("ntrks", ""), 0, 16)
	if err != nil {
		return nil, newXMLDecodeError(el, fmt.Errorf("invalid attribute: ntrks=%q", el.SelectAttrValue("ntrks", "")))
	}
	framerate, err := strconv.ParseUint(el.SelectAttrValue("framerate", ""), 0, 8)
	if err != nil {
		return nil, newXMLDecodeError(el, fmt.Errorf("invalid attribute: framerate=%q", el.SelectAttrValue("framerate", "")))
	}
	division, err := strconv.ParseUint(el.SelectAttrValue("division", ""), 0, 16)
	if err != nil {
		return nil, newXMLDecodeError(el, fmt.Errorf("invalid attribute: division=%q", el.SelectAttrValue("division", "")))
	}
	undecoded, err := parseHexDump(el.SelectAttrValue("undecoded", ""))
	if err != nil {
		return nil, newXMLDecodeError(el, fmt.Errorf("invalid attribute: undecoded=%q", el.SelectAttrValue("undecoded", "")))
	}
	return &MThd{
		FilePosition: pos,
		Format:       uint16(format),
		NTrks:        uint16(ntrks),
		Framerate:    uint8(framerate),
		Division:     uint16(division),
		Undecoded:    undecoded,
	}, nil
}
