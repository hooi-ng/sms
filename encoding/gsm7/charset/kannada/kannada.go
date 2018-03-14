// Copyright © 2018 Kent Gibson <warthog618@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package kannada

// Decoder provides a mapping from GSM7 byte to UTF8 rune.
type Decoder map[byte]rune

// Encoder provides a mapping from UTF8 rune to GSM7 byte.
type Encoder map[rune]byte

// NewDecoder returns the mapping table from GSM7 to UTF8.
func NewDecoder() Decoder {
	return dset
}

// NewExtDecoder returns the extension mapping table from GSM7 to UTF8.
func NewExtDecoder() Decoder {
	return dext
}

// NewEncoder returns the mapping table from UTF8 to GSM7.
func NewEncoder() Encoder {
	return eset
}

// NewExtEncoder returns the extention mapping table from UTF8 to GSM7.
func NewExtEncoder() Encoder {
	return eext
}

var (
	dset = Decoder{
		0x01: '\u0c82',
		0x02: '\u0c83',
		0x03: '\u0c85',
		0x04: '\u0c86',
		0x05: '\u0c87',
		0x06: '\u0c88',
		0x07: '\u0c89',
		0x08: '\u0c8a',
		0x09: '\u0c8b',
		0x0a: '\n',
		0x0b: '\u0c8c',
		0x0d: '\r',
		0x0e: '\u0c8e',
		0x0f: '\u0c8f',
		0x10: '\u0c90',
		0x12: '\u0c92',
		0x13: '\u0c93',
		0x14: '\u0c94',
		0x15: '\u0c95',
		0x16: '\u0c96',
		0x17: '\u0c97',
		0x18: '\u0c98',
		0x19: '\u0c99',
		0x1a: '\u0c9a',
		0x1b: 0x1b,
		0x1c: '\u0c9b',
		0x1d: '\u0c9c',
		0x1e: '\u0c9d',
		0x1f: '\u0c9e',
		0x20: 0x20,
		0x21: '!',
		0x22: '\u0c9f',
		0x23: '\u0ca0',
		0x24: '\u0ca1',
		0x25: '\u0ca2',
		0x26: '\u0ca3',
		0x27: '\u0ca4',
		0x28: ')',
		0x29: '(',
		0x2a: '\u0ca5',
		0x2b: '\u0ca6',
		0x2c: ',',
		0x2d: '\u0ca7',
		0x2e: '.',
		0x2f: '\u0ca8',
		0x30: '0',
		0x31: '1',
		0x32: '2',
		0x33: '3',
		0x34: '4',
		0x35: '5',
		0x36: '6',
		0x37: '7',
		0x38: '8',
		0x39: '9',
		0x3a: ':',
		0x3b: ';',
		0x3d: '\u0caa',
		0x3e: '\u0cab',
		0x3f: '?',
		0x40: '\u0cac',
		0x41: '\u0cad',
		0x42: '\u0cae',
		0x43: '\u0caf',
		0x44: '\u0cb0',
		0x45: '\u0cb1',
		0x46: '\u0cb2',
		0x47: '\u0cb3',
		0x49: '\u0cb5',
		0x4a: '\u0cb6',
		0x4b: '\u0cb7',
		0x4c: '\u0cb8',
		0x4d: '\u0cb9',
		0x4e: '\u0cbc',
		0x4f: '\u0cbd',
		0x50: '\u0cbe',
		0x51: '\u0cbf',
		0x52: '\u0cc0',
		0x53: '\u0cc1',
		0x54: '\u0cc2',
		0x55: '\u0cc3',
		0x56: '\u0cc4',
		0x58: '\u0cc6',
		0x59: '\u0cc7',
		0x5a: '\u0cc8',
		0x5c: '\u0cca',
		0x5d: '\u0ccb',
		0x5e: '\u0ccc',
		0x5f: '\u0ccd',
		0x60: '\u0cd5',
		0x61: 'a',
		0x62: 'b',
		0x63: 'c',
		0x64: 'd',
		0x65: 'e',
		0x66: 'f',
		0x67: 'g',
		0x68: 'h',
		0x69: 'i',
		0x6a: 'j',
		0x6b: 'k',
		0x6c: 'l',
		0x6d: 'm',
		0x6e: 'n',
		0x6f: 'o',
		0x70: 'p',
		0x71: 'q',
		0x72: 'r',
		0x73: 's',
		0x74: 't',
		0x75: 'u',
		0x76: 'v',
		0x77: 'w',
		0x78: 'x',
		0x79: 'y',
		0x7a: 'z',
		0x7b: '\u0cd6',
		0x7c: '\u0ce0',
		0x7d: '\u0ce1',
		0x7e: '\u0ce2',
		0x7f: '\u0ce3',
	}
	dext = Decoder{
		0x00: '@',
		0x01: '£',
		0x02: '$',
		0x03: '¥',
		0x04: '¿',
		0x05: '"',
		0x06: '¤',
		0x07: '%',
		0x08: '&',
		0x09: '\'',
		0x0a: '\f',
		0x0b: '*',
		0x0c: '+',
		0x0d: '\r',
		0x0e: '-',
		0x0f: '/',
		0x10: '<',
		0x11: '=',
		0x12: '>',
		0x13: '¡',
		0x14: '^',
		0x15: '¡',
		0x16: '_',
		0x17: '#',
		0x18: '*',
		0x19: '\u0964',
		0x1a: '\u0965',
		0x1b: 0x1b,
		0x1c: '\u0ce6',
		0x1d: '\u0ce7',
		0x1e: '\u0ce8',
		0x1f: '\u0ce9',
		0x20: '\u0cea',
		0x21: '\u0ceb',
		0x22: '\u0cec',
		0x23: '\u0ced',
		0x24: '\u0cee',
		0x25: '\u0cef',
		0x26: '\u0cde',
		0x27: '\u0cf1',
		0x28: '{',
		0x29: '}',
		0x2a: '\u0cf2',
		0x2f: '\\',
		0x3c: '[',
		0x3d: '~',
		0x3e: ']',
		0x40: '|',
		0x41: 'A',
		0x42: 'B',
		0x43: 'C',
		0x44: 'D',
		0x45: 'E',
		0x46: 'F',
		0x47: 'G',
		0x48: 'H',
		0x49: 'I',
		0x4a: 'J',
		0x4b: 'K',
		0x4c: 'L',
		0x4d: 'M',
		0x4e: 'N',
		0x4f: 'O',
		0x50: 'P',
		0x51: 'Q',
		0x52: 'R',
		0x53: 'S',
		0x54: 'T',
		0x55: 'U',
		0x56: 'V',
		0x57: 'W',
		0x58: 'X',
		0x59: 'Y',
		0x5a: 'Z',
		0x65: '€',
	}
	eset Encoder
	eext Encoder
)

func init() {
	eset = make(Encoder, len(dset))
	for k, v := range dset {
		eset[v] = k
	}
	eext = make(Encoder, len(dext))
	for k, v := range dext {
		if ko, ok := eext[v]; !ok || ko > k {
			eext[v] = k
		}
	}
}
