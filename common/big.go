// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package common

import (
	"github.com/holiman/uint256"
)

// Common big integers often used
var (
	Big1   = uint256.NewInt(1)
	Big2   = uint256.NewInt(2)
	Big3   = uint256.NewInt(3)
	Big0   = uint256.NewInt(0)
	Big32  = uint256.NewInt(32)
	Big256 = uint256.NewInt(256)
	Big257 = uint256.NewInt(257)

	U2560 = uint256.NewInt(0)
)
