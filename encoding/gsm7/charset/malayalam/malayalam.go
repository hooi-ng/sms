// Copyright © 2018 Kent Gibson <warthog618@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package malayalam

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
		0x01: '\u0d02',
		0x02: '\u0d03',
		0x03: '\u0d05',
		0x04: '\u0d06',
		0x05: '\u0d07',
		0x06: '\u0d08',
		0x07: '\u0d09',
		0x08: '\u0d0a',
		0x09: '\u0d0b',
		0x0a: '\n',
		0x0b: '\u0d0c',
		0x0d: '\r',
		0x0e: '\u0d0e',
		0x0f: '\u0d0f',
		0x10: '\u0d10',
		0x12: '\u0d12',
		0x13: '\u0d13',
		0x14: '\u0d14',
		0x15: '\u0d15',
		0x16: '\u0d16',
		0x17: '\u0d17',
		0x18: '\u0d18',
		0x19: '\u0d19',
		0x1a: '\u0d1a',
		0x1b: 0x1b,
		0x1c: '\u0d1b',
		0x1d: '\u0d1c',
		0x1e: '\u0d1d',
		0x1f: '\u0d1e',
		0x20: 0x20,
		0x21: '!',
		0x22: '\u0d1f',
		0x23: '\u0d20',
		0x24: '\u0d21',
		0x25: '\u0d22',
		0x26: '\u0d23',
		0x27: '\u0d24',
		0x28: ')',
		0x29: '(',
		0x2a: '\u0d25',
		0x2b: '\u0d26',
		0x2c: ',',
		0x2d: '\u0d27',
		0x2e: '.',
		0x2f: '\u0d28',
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
		0x3d: '\u0d2a',
		0x3e: '\u0d2b',
		0x3f: '?',
		0x40: '\u0d2c',
		0x41: '\u0d2d',
		0x42: '\u0d2e',
		0x43: '\u0d2f',
		0x44: '\u0d30',
		0x45: '\u0d31',
		0x46: '\u0d32',
		0x47: '\u0d33',
		0x48: '\u0d34',
		0x49: '\u0d35',
		0x4a: '\u0d36',
		0x4b: '\u0d37',
		0x4c: '\u0d38',
		0x4d: '\u0d39',
		0x4f: '\u0d3d',
		0x50: '\u0d3e',
		0x51: '\u0d3f',
		0x52: '\u0d40',
		0x53: '\u0d41',
		0x54: '\u0d42',
		0x55: '\u0d43',
		0x56: '\u0d44',
		0x58: '\u0d46',
		0x59: '\u0d47',
		0x5a: '\u0d48',
		0x5c: '\u0d4a',
		0x5d: '\u0d4b',
		0x5e: '\u0d4c',
		0x5f: '\u0d4d',
		0x60: '\u0d57',
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
		0x7b: '\u0d60',
		0x7c: '\u0d61',
		0x7d: '\u0d62',
		0x7e: '\u0d63',
		0x7f: '\u0d79',
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
		0x1c: '\u0d66',
		0x1d: '\u0d67',
		0x1e: '\u0d68',
		0x1f: '\u0d69',
		0x20: '\u0d6a',
		0x21: '\u0d6b',
		0x22: '\u0d6c',
		0x23: '\u0d6d',
		0x24: '\u0d6e',
		0x25: '\u0d6f',
		0x26: '\u0d70',
		0x27: '\u0d71',
		0x28: '{',
		0x29: '}',
		0x2a: '\u0d72',
		0x2b: '\u0d73',
		0x2c: '\u0d74',
		0x2d: '\u0d75',
		0x2e: '\u0d7a',
		0x2f: '\\',
		0x30: '\u0d7b',
		0x31: '\u0d7c',
		0x32: '\u0d7d',
		0x33: '\u0d7e',
		0x34: '\u0d7f',
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
