//
// Copyright 2015 Zac Shenker. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mp4

import (
//"log"
//"errors"
//"github.com/oikomi/gomp4/util"
)

type TrunEntry struct {
	Duration              uint32
	Size                  uint32
	Flag                  uint32
	CompositionTimeOffset int32
}

type TrunAtom struct {
	Offset           int64
	Size             int64
	IsFullBox        bool
	Version          uint8
	Flag             uint32
	EntryCount       uint32
	Entries          []TrunEntry
	DataOffset       int32
	FirstSampleFlags uint32
	AllBytes         []byte
}

/*
  version = stream.ui8
    @flags = stream.ui24

    @data_offset_present = @flags & 0x01
    @first_sample_flags_present = @flags & 0x04
    @sample_duration_present = @flags & 0x0100
    @sample_size_present = @flags & 0x0200
    @sample_flags_present = @flags & 0x0400
    @sample_composition_time_offsets_present = @flags & 0x0800
    @duration_is_empty = @flags & 0x010000

    @entry_count = stream.ui32
    @data_offset = stream.si32 if @data_offset_present > 0
    @first_sample_flags = stream.ui32 if @first_sample_flags_present > 0

      @entries = []
    @entry_count.times do
      entry = {}
      entry[:duration] = stream.ui32 if @sample_duration_present > 0
      entry[:size] = stream.ui32 if @sample_size_present > 0
      entry[:flags] = stream.ui32 if @sample_flags_present > 0
      entry[:composition_time_offset] = stream.si32 if @sample_composition_time_offsets_present > 0
      @entries << entry
    end
*/

func trunRead(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error {
	return nil
}
