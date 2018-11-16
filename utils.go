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
)

type limitReadSeeker struct {
	r   io.ReadSeeker
	pos int64
	end int64
}

func tell(r io.ReadSeeker) int64 {
	pos, err := r.Seek(0, io.SeekCurrent)
	if err != nil {
		return -1
	}
	return pos
}

func newLimitReadSeeker(r io.ReadSeeker, n int64) (*limitReadSeeker, error) {
	pos, err := r.Seek(0, io.SeekCurrent)
	if err != nil {
		return nil, err
	}
	end := pos + n
	if end < pos {
		end = -1
	}
	return &limitReadSeeker{
		r:   r,
		pos: pos,
		end: end,
	}, nil
}

func (r *limitReadSeeker) Read(p []byte) (n int, err error) {
	if r.end >= 0 {
		if r.pos >= r.end {
			return 0, io.EOF
		}
		if r.pos+int64(len(p)) > r.end {
			p = p[:r.end-r.pos]
		}
	}
	n, err = r.r.Read(p)
	r.pos += int64(n)
	return
}

func (r *limitReadSeeker) Seek(offset int64, whence int) (pos int64, err error) {
	pos, err = r.r.Seek(offset, whence)
	if err == nil {
		r.pos = pos
	}
	return
}
