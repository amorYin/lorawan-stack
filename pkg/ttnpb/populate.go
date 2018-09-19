// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ttnpb

var PopulatorConfig struct {
	LoRaWAN struct {
		AppendMHDR       func(b []byte, mhdr MHDR) ([]byte, error)
		AppendFHDR       func(b []byte, fhdr FHDR, isUplink bool) ([]byte, error)
		AppendMessage    func(dst []byte, msg Message) ([]byte, error)
		MarshalMessage   func(msg Message) ([]byte, error)
		UnmarshalMessage func(b []byte, msg *Message) error
	}
}