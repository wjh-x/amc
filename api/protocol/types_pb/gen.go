// Copyright 2022 The AmazeChain Authors
// This file is part of the AmazeChain library.
//
// The AmazeChain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The AmazeChain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the AmazeChain library. If not, see <http://www.gnu.org/licenses/>.

package types_pb

////go:generate protoc --plugin=/Users/mac/go/bin/protoc-gen-go-cast -I=../ -I=. -I=../include --go-cast_out=plugins=protoc-gen-go-cast,paths=source_relative:. types.proto
//go:generate protoc  -I=../ -I=. -I=../include --go-cast_out=paths=source_relative:. types.proto
//go:generate sszgen -path=. -objs=H128,H160,H256,H384,H768,H512,H1024,H2048,Header,Body,Block,Transaction -output=generated.ssz.go
