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
	"strconv"
)

type Key uint8

type KeySignature int16

const (
	KeyCMaj  KeySignature = 0x0000
	KeyCMin               = -0x2ff
	KeyCsMaj              = 0x0700
	KeyCsMin              = 0x0401
	KeyDbMaj              = -0x500
	KeyDMaj               = 0x0200
	KeyDMin               = -0x0ff
	KeyDsMin              = 0x0601
	KeyEbMaj              = -0x300
	KeyEbMin              = -0x5ff
	KeyEMaj               = 0x0400
	KeyEMin               = 0x0101
	KeyFMaj               = -0x100
	KeyFMin               = -0x3ff
	KeyFsMaj              = 0x0600
	KeyFsMin              = 0x0301
	KeyGbMaj              = -0x600
	KeyGMaj               = 0x0100
	KeyGMin               = -0x1ff
	KeyGsMin              = 0x0501
	KeyAbMaj              = -0x400
	KeyAbMin              = -0x6ff
	KeyAMaj               = 0x0300
	KeyAMin               = 0x0001
	KeyAsMin              = 0x0701
	KeyBbMaj              = -0x200
	KeyBbMin              = -0x4ff
	KeyBMaj               = 0x0500
	KeyBMin               = 0x0201
	KeyCbMaj              = -0x700
)

var (
	keyToString = [128]string{
		"C-1", "C#-1", "D-1", "Eb-1", "E-1", "F-1", "F#-1", "G-1", "Ab-1", "A-1", "Bb-1", "B-1", "C0", "C#0", "D0", "Eb0", "E0", "F0", "F#0", "G0", "Ab0", "A0", "Bb0", "B0", "C1", "C#1", "D1", "Eb1", "E1", "F1", "F#1", "G1", "Ab1", "A1", "Bb1", "B1", "C2", "C#2", "D2", "Eb2", "E2", "F2", "F#2", "G2", "Ab2", "A2", "Bb2", "B2", "C3", "C#3", "D3", "Eb3", "E3", "F3", "F#3", "G3", "Ab3", "A3", "Bb3", "B3", "C4", "C#4", "D4", "Eb4", "E4", "F4", "F#4", "G4", "Ab4", "A4", "Bb4", "B4", "C5", "C#5", "D5", "Eb5", "E5", "F5", "F#5", "G5", "Ab5", "A5", "Bb5", "B5", "C6", "C#6", "D6", "Eb6", "E6", "F6", "F#6", "G6", "Ab6", "A6", "Bb6", "B6", "C7", "C#7", "D7", "Eb7", "E7", "F7", "F#7", "G7", "Ab7", "A7", "Bb7", "B7", "C8", "C#8", "D8", "Eb8", "E8", "F8", "F#8", "G8", "Ab8", "A8", "Bb8", "B8", "C9", "C#9", "D9", "Eb9", "E9", "F9", "F#9", "G9",
	}
	stringToKey = map[string]Key{
		"C-1": 0x00, "C#-1": 0x01, "Db-1": 0x01, "D-1": 0x02, "D#-1": 0x03, "Eb-1": 0x03, "E-1": 0x04, "F-1": 0x05, "F#-1": 0x06, "Gb-1": 0x06, "G-1": 0x07, "G#-1": 0x08, "Ab-1": 0x08, "A-1": 0x09, "A#-1": 0x0a, "Bb-1": 0x0a, "B-1": 0x0b, "C0": 0x0c, "C#0": 0x0d, "Db0": 0x0d, "D0": 0x0e, "D#0": 0x0f, "Eb0": 0x0f, "E0": 0x10, "F0": 0x11, "F#0": 0x12, "Gb0": 0x12, "G0": 0x13, "G#0": 0x14, "Ab0": 0x14, "A0": 0x15, "A#0": 0x16, "Bb0": 0x16, "B0": 0x17, "C1": 0x18, "C#1": 0x19, "Db1": 0x19, "D1": 0x1a, "D#1": 0x1b, "Eb1": 0x1b, "E1": 0x1c, "F1": 0x1d, "F#1": 0x1e, "Gb1": 0x1e, "G1": 0x1f, "G#1": 0x20, "Ab1": 0x20, "A1": 0x21, "A#1": 0x22, "Bb1": 0x22, "B1": 0x23, "C2": 0x24, "C#2": 0x25, "Db2": 0x25, "D2": 0x26, "D#2": 0x27, "Eb2": 0x27, "E2": 0x28, "F2": 0x29, "F#2": 0x2a, "Gb2": 0x2a, "G2": 0x2b, "G#2": 0x2c, "Ab2": 0x2c, "A2": 0x2d, "A#2": 0x2e, "Bb2": 0x2e, "B2": 0x2f, "C3": 0x30, "C#3": 0x31, "Db3": 0x31, "D3": 0x32, "D#3": 0x33, "Eb3": 0x33, "E3": 0x34, "F3": 0x35, "F#3": 0x36, "Gb3": 0x36, "G3": 0x37, "G#3": 0x38, "Ab3": 0x38, "A3": 0x39, "A#3": 0x3a, "Bb3": 0x3a, "B3": 0x3b, "C4": 0x3c, "C#4": 0x3d, "Db4": 0x3d, "D4": 0x3e, "D#4": 0x3f, "Eb4": 0x3f, "E4": 0x40, "F4": 0x41, "F#4": 0x42, "Gb4": 0x42, "G4": 0x43, "G#4": 0x44, "Ab4": 0x44, "A4": 0x45, "A#4": 0x46, "Bb4": 0x46, "B4": 0x47, "C5": 0x48, "C#5": 0x49, "Db5": 0x49, "D5": 0x4a, "D#5": 0x4b, "Eb5": 0x4b, "E5": 0x4c, "F5": 0x4d, "F#5": 0x4e, "Gb5": 0x4e, "G5": 0x4f, "G#5": 0x50, "Ab5": 0x50, "A5": 0x51, "A#5": 0x52, "Bb5": 0x52, "B5": 0x53, "C6": 0x54, "C#6": 0x55, "Db6": 0x55, "D6": 0x56, "D#6": 0x57, "Eb6": 0x57, "E6": 0x58, "F6": 0x59, "F#6": 0x5a, "Gb6": 0x5a, "G6": 0x5b, "G#6": 0x5c, "Ab6": 0x5c, "A6": 0x5d, "A#6": 0x5e, "Bb6": 0x5e, "B6": 0x5f, "C7": 0x60, "C#7": 0x61, "Db7": 0x61, "D7": 0x62, "D#7": 0x63, "Eb7": 0x63, "E7": 0x64, "F7": 0x65, "F#7": 0x66, "Gb7": 0x66, "G7": 0x67, "G#7": 0x68, "Ab7": 0x68, "A7": 0x69, "A#7": 0x6a, "Bb7": 0x6a, "B7": 0x6b, "C8": 0x6c, "C#8": 0x6d, "Db8": 0x6d, "D8": 0x6e, "D#8": 0x6f, "Eb8": 0x6f, "E8": 0x70, "F8": 0x71, "F#8": 0x72, "Gb8": 0x72, "G8": 0x73, "G#8": 0x74, "Ab8": 0x74, "A8": 0x75, "A#8": 0x76, "Bb8": 0x76, "B8": 0x77, "C9": 0x78, "C#9": 0x79, "Db9": 0x79, "D9": 0x7a, "D#9": 0x7b, "Eb9": 0x7b, "E9": 0x7c, "F9": 0x7d, "F#9": 0x7e, "Gb9": 0x7e, "G9": 0x7f,
	}
	keySignatureToString = map[KeySignature]string{
		KeyCMaj: "C Maj", KeyCMin: "C Min", KeyCsMaj: "C# Maj", KeyCsMin: "C# Min", KeyDbMaj: "Db Maj", KeyDMaj: "D Maj", KeyDMin: "D Min", KeyDsMin: "D# Min", KeyEbMaj: "Eb Maj", KeyEbMin: "Eb Min", KeyEMaj: "E Maj", KeyEMin: "E Min", KeyFMaj: "F Maj", KeyFMin: "F Min", KeyFsMaj: "F# Maj", KeyFsMin: "F# Min", KeyGbMaj: "Gb Maj", KeyGMaj: "G Maj", KeyGMin: "G Min", KeyGsMin: "G# Min", KeyAbMaj: "Ab Maj", KeyAbMin: "Ab Min", KeyAMaj: "A Maj", KeyAMin: "A Min", KeyAsMin: "A# Min", KeyBbMaj: "Bb Maj", KeyBbMin: "Bb Min", KeyBMaj: "B Maj", KeyBMin: "B Min", KeyCbMaj: "Cb Maj",
	}
	stringToKeySignature = map[string]KeySignature{
		"C Maj": KeyCMaj, "C Min": KeyCMin, "C": KeyCMaj, "c": KeyCMin, "C# Maj": KeyCsMaj, "C# Min": KeyCsMin, "C#": KeyCsMaj, "c#": KeyCsMin, "C#m": KeyCsMin, "Cb Maj": KeyCbMaj, "Cb": KeyCbMaj, "Cm": KeyCMin, "D Maj": KeyDMaj, "D Min": KeyDMin, "D": KeyDMaj, "d": KeyDMin, "D# Min": KeyDsMin, "d#": KeyDsMin, "D#m": KeyDsMin, "Db Maj": KeyDbMaj, "Db": KeyDbMaj, "Dm": KeyDMin, "E Maj": KeyEMaj, "E Min": KeyEMin, "E": KeyEMaj, "e": KeyEMin, "Eb Maj": KeyEbMaj, "Eb Min": KeyEbMin, "Eb": KeyDbMaj, "eb": KeyEbMin, "Ebm": KeyEbMin, "Em": KeyEMin, "F Maj": KeyFMaj, "F Min": KeyFMin, "F": KeyFMaj, "f": KeyFMin, "F# Maj": KeyFsMaj, "F# Min": KeyFsMin, "F#": KeyFsMaj, "f#": KeyFsMin, "F#m": KeyFsMin, "Fm": KeyFMin, "G Maj": KeyGMaj, "G Min": KeyGMin, "G": KeyGMaj, "g": KeyGMin, "G# Min": KeyGsMin, "g#": KeyGsMin, "G#m": KeyGsMin, "Gb Maj": KeyGbMaj, "Gb": KeyGbMaj, "Gm": KeyGMin, "A Maj": KeyAMaj, "A Min": KeyAMin, "A": KeyAMaj, "a": KeyAMin, "A# Min": KeyAsMin, "a#": KeyAsMin, "A#m": KeyAsMin, "Ab Maj": KeyAbMaj, "Ab Min": KeyAbMin, "Ab": KeyAbMaj, "ab": KeyAbMin, "Abm": KeyAbMin, "Am": KeyAMin, "B Maj": KeyBMaj, "B Min": KeyBMin, "B": KeyBMaj, "b": KeyBMin, "Bb Maj": KeyBbMaj, "Bb Min": KeyBbMin, "Bb": KeyBbMaj, "bb": KeyBbMin, "Bbm": KeyBbMin, "Bm": KeyBMin,
	}
)

func (key Key) String() string {
	if key >= 0x80 {
		return fmt.Sprintf("%#02x", uint8(key))
	}
	return keyToString[key]
}

func ParseKey(str string) (Key, error) {
	if key, ok := stringToKey[str]; ok {
		return key, nil
	}
	if key, err := strconv.ParseUint(str, 0, 7); err == nil {
		return Key(key), nil
	}
	return 0xff, fmt.Errorf("unrecognized key name %q", str)
}

func (ks KeySignature) String() string {
	if str, ok := keySignatureToString[ks]; ok {
		return str
	}
	return fmt.Sprintf("%#04x", int16(ks))
}

func ParseKeySignature(str string) (KeySignature, error) {
	if ks, ok := stringToKeySignature[str]; ok {
		return ks, nil
	}
	if ks, err := strconv.ParseInt(str, 0, 17); err == nil {
		return KeySignature(ks), nil
	}
	return 0, fmt.Errorf("unrecognized key signature %q", str)
}
