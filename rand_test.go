/*
 * Copyright (C) 2015-2018 Virgil Security Inc.
 *
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     (1) Redistributions of source code must retain the above copyright
 *     notice, this list of conditions and the following disclaimer.
 *
 *     (2) Redistributions in binary form must reproduce the above copyright
 *     notice, this list of conditions and the following disclaimer in
 *     the documentation and/or other materials provided with the
 *     distribution.
 *
 *     (3) Neither the name of the copyright holder nor the names of its
 *     contributors may be used to endorse or promote products derived from
 *     this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE AUTHOR ''AS IS'' AND ANY EXPRESS OR
 * IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT,
 * INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
 * STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING
 * IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 *
 * Lead Maintainer: Virgil Security Inc. <support@virgilsecurity.com>
 */

package phe

import (
	"bytes"
	"crypto/rand"
)

var randBytes = []byte{
	0x56, 0x83, 0x3a, 0x7a, 0x1c, 0xdb, 0x2e, 0xc9,
	0x91, 0xca, 0x33, 0x8c, 0xd4, 0xa2, 0xed, 0x66,
	0x37, 0x70, 0x40, 0x0f, 0xab, 0x2c, 0xd7, 0x98,
	0xbe, 0xe9, 0xf3, 0xed, 0x37, 0x87, 0x96, 0xbd,
	0x40, 0x31, 0x8b, 0x72, 0x73, 0x91, 0xe8, 0x59,
	0x25, 0x55, 0xd3, 0x41, 0xc3, 0x25, 0xef, 0x2b,
	0xd2, 0xcc, 0x20, 0xaa, 0x40, 0x40, 0xec, 0x9a,
	0xd2, 0xb8, 0x51, 0xbc, 0x15, 0xb6, 0xfe, 0xf5,
	0x56, 0x83, 0x3a, 0x7a, 0x1c, 0xdb, 0x2e, 0xc9,
	0x91, 0xca, 0x33, 0x8c, 0xd4, 0xa2, 0xed, 0x66,
	0x37, 0x70, 0x40, 0x0f, 0xab, 0x2c, 0xd7, 0x98,
	0xbe, 0xe9, 0xf3, 0xed, 0x37, 0x87, 0x96, 0xbd,
}

func MockRandom() {
	randReader = bytes.NewReader(randBytes)
}

func EndMock() {
	randReader = rand.Reader
}
