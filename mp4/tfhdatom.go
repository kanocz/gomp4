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
	"log"
	//"errors"
	"github.com/oikomi/gomp4/util"
)

type TfhdAtom struct {
	Offset                int64
	Size                  int64
	IsFullBox             bool
	Version               uint8
	Flag                  uint32
	TrackId               uint32
	BaseDataOffset        uint64
	SampleDescIndex       uint32
	DefaultSampleDuration uint32
	DefaultSampleSize     uint32
	DefaultSampleFlags    uint32
	AllBytes              []byte
}

/*
version = stream.ui8
    @flags = stream.ui24
    @base_data_offset_present = @flags & 0x01
    @sample_desc_index_present = @flags & 0x02
    @default_sample_duration_present = @flags & 0x08
    @default_sample_size_present = @flags & 0x10
    @default_sample_flags_present = @flags & 0x20
    @duration_is_empty = @flags & 0x010000

    @track_id = stream.ui32
    @base_data_offset = stream.ui64 if @base_data_offset_present > 0
    @sample_desc_index = stream.ui32 if @sample_desc_index_present > 0
    @default_sample_duration = stream.ui32 if @default_sample_duration_present > 0
    @default_sample_size = stream.ui32 if @default_sample_size_present > 0
    @default_sample_flags = stream.ui32 if @default_sample_flags_present > 0
*/

func TfhdRead(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error {

}
