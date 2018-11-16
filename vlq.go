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
	"io"
)

type VLQ uint32

const MaxVLQ = 0xfffffff

func (v VLQ) Encode(w io.Writer) error {
	switch {
	case v < 0x80:
		_, err := w.Write([]byte{uint8(v)})
		return err
	case v < 0x4000:
		_, err := w.Write([]byte{uint8(v>>7) | 0x80, uint8(v) & 0x7f})
		return err
	case v < 0x200000:
		_, err := w.Write([]byte{uint8(v>>14) | 0x80, uint8(v>>7) | 0x80, uint8(v) & 0x7f})
		return err
	case v < 0x10000000:
		_, err := w.Write([]byte{uint8(v>>21) | 0x80, uint8(v>>14) | 0x80, uint8(v>>7) | 0x80, uint8(v) & 0x7f})
		return err
	default:
		return newSMFEncodeError(&v, fmt.Errorf("invalid VLQ value %d", v))
	}
}

func (v VLQ) EncodeLen() (int64, error) {
	switch {
	case v < 0x80:
		return 1, nil
	case v < 0x4000:
		return 2, nil
	case v < 0x200000:
		return 3, nil
	case v < 0x10000000:
		return 4, nil
	default:
		return 0, newSMFEncodeError(&v, fmt.Errorf("invalid VLQ value %d", v))
	}
}

func DecodeVLQ(r io.ReadSeeker, warningCallback WarningCallback) (VLQ, error) {
	pos := tell(r)
	var buf [4]byte

	_, err := io.ReadFull(r, buf[:1])
	if err != nil {
		return 0, err
	}
	if buf[0] < 0x80 {
		return VLQ(buf[0]), nil
	}

	_, err = io.ReadFull(r, buf[1:2])
	if err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return VLQ(buf[0]), err
	}
	if buf[1] < 0x80 {
		return (VLQ(buf[0]&0x7f) << 7) | VLQ(buf[1]), nil
	}

	_, err = io.ReadFull(r, buf[2:3])
	if err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return (VLQ(buf[0]&0x7f) << 7) | VLQ(buf[1]), err
	}
	if buf[2] < 0x80 {
		return (VLQ(buf[0]&0x7f) << 14) | (VLQ(buf[1]&0x7f) << 7) | VLQ(buf[2]), nil
	}

	_, err = io.ReadFull(r, buf[3:4])
	if err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return (VLQ(buf[0]&0x7f) << 14) | (VLQ(buf[1]&0x7f) << 7) | VLQ(buf[2]), err
	}
	if buf[3] < 0x80 {
		return (VLQ(buf[0]&0x7f) << 21) | (VLQ(buf[1]&0x7f) << 14) | (VLQ(buf[2]&0x7f) << 7) | VLQ(buf[3]), nil
	}

	warningCallback(newSMFDecodeError(pos, fmt.Errorf("invalid VLQ encoding %#01x", buf)))
	// We might be decoding something other than VLQ, try to resync
	_, err = r.Seek(-4, io.SeekCurrent)
	if err != nil {
		return 0, err
	}
	return 0, nil
}
