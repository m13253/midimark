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
	"io/ioutil"

	"github.com/beevik/etree"
)

func (seq *Sequence) EncodeSMF(w io.Writer) error {
	if len(seq.Tracks) < 0xffff {
		seq.Header.NTrks = uint16(len(seq.Tracks))
	} else {
		seq.Header.NTrks = 0xffff
	}
	err := seq.Header.EncodeSMF(w)
	if err != nil {
		return err
	}
	for _, mtrk := range seq.Tracks {
		err = mtrk.EncodeSMF(w)
		if err != nil {
			return err
		}
	}
	_, err = w.Write(seq.Undecoded)
	return err
}

func (seq *Sequence) EncodeXML() *etree.Element {
	el := etree.NewElement("Sequence")
	el.AddChild(seq.Header.EncodeXML())
	for _, mtrk := range seq.Tracks {
		el.AddChild(mtrk.EncodeXML())
	}
	if len(seq.Undecoded) != 0 {
		undecoded := etree.NewElement("Undecoded")
		undecoded.SetText(fmt.Sprintf("% x", seq.Undecoded))
		el.AddChild(undecoded)
	}
	return el
}

func (seq *Sequence) EncodeXMLToDocument(w io.Writer) (n int64, err error) {
	doc := etree.NewDocument()
	el := seq.EncodeXML()
	doc.AddChild(el)
	doc.Indent(2)
	return doc.WriteTo(w)
}

func DecodeSequenceFromSMF(r io.ReadSeeker, warningCallback WarningCallback) (seq *Sequence, err error) {
	mthd, err := DecodeMThdFromSMF(r, warningCallback)
	if err != nil {
		return nil, err
	}

	seq = &Sequence{
		Header: mthd,
		Tracks: make([]*MTrk, 0),
	}
	defer func() {
		seq.CalculateNotePair()
		seq.CalculateTempoTable()
	}()

	pos := tell(r)
	if mthd.NTrks == 0 {
		warningCallback(newSMFDecodeError(pos+10, errors.New("MIDI file seems to contain no tracks")))
	}

	for i := uint16(0); mthd.NTrks == 0 || i < mthd.NTrks; i++ {
		var mtrk *MTrk
		mtrk, err = DecodeMTrkFromSMF(r, warningCallback)
		if mtrk != nil {
			seq.Tracks = append(seq.Tracks, mtrk)
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

	pos = tell(r)
	seq.Undecoded, err = ioutil.ReadAll(r)
	if err != nil {
		warningCallback(newSMFDecodeError(pos, err))
	}
	return
}

func DecodeSequenceFromXML(el *etree.Element) (*Sequence, error) {
	seq := &Sequence{}
	for _, child := range el.Child {
		if childEl, ok := child.(*etree.Element); ok {
			switch childEl.Tag {
			case "MThd":
				var err error
				seq.Header, err = DecodeMThdFromXML(childEl)
				if err != nil {
					return nil, err
				}
			case "MTrk":
				mtrk, err := DecodeMTrkFromXML(childEl)
				if err != nil {
					return nil, err
				}
				seq.Tracks = append(seq.Tracks, mtrk)
			case "Undecoded":
				var err error
				seq.Undecoded, err = parseHexDump(childEl.Text())
				if err != nil {
					return nil, newXMLDecodeError(childEl, fmt.Errorf("unable to decode tag <Undecoded>"))
				}
			default:
				return nil, newXMLDecodeError(childEl, fmt.Errorf("unexpected tag <%s>", childEl.Tag))
			}
		}
	}
	if seq.Header == nil {
		return nil, newXMLDecodeError(el, errors.New("can not find a MThd tag"))
	}
	seq.CalculateNotePair()
	seq.CalculateTempoTable()
	return seq, nil
}

func DecodeXMLFromDocument(r io.Reader) (seq *Sequence, n int64, err error) {
	doc := etree.NewDocument()
	doc.ReadSettings.Permissive = true
	n, err = doc.ReadFrom(r)
	if err != nil {
		return nil, n, newXMLDecodeError(doc, err)
	}
	root := doc.Root()
	if root == nil {
		return nil, n, newXMLDecodeError(doc, errors.New("XML file contains no root tag"))
	}
	seq, err = DecodeSequenceFromXML(root)
	return
}
