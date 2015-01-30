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

type TfraAtom struct {
	Offset    int64
	Size      int64
	IsFullBox bool
	Version   uint8
	Flag      uint32

	AllBytes []byte
}

/*
 @version,@flags = stream.ui8, stream.ui24
    @track_id = stream.ui32
    length_sizes = stream.ui32
    size_of_traf_num = (length_sizes & 0b110000) >> 4
    size_of_trun_num = (length_sizes & 0b1100) >> 2
    size_of_sample_num = (length_sizes & 0b11)
    @entry_count = stream.ui32
    @entry_count.times do
      entry = {}
      if @version == 1
        entry[:time] = stream.ui64
        entry[:moof_offset] = stream.ui64
      else
        entry[:time] = stream.ui32
        entry[:moof_offset] = stream.ui32
      end
      entry[:traf_number] = stream.send(SIZE_READ_METHODS[size_of_traf_num])
      entry[:trun_number] = stream.send(SIZE_READ_METHODS[size_of_trun_num])
      entry[:sample_number] = stream.send(SIZE_READ_METHODS[size_of_sample_num])
      @entries << entry
    end
*/

func TfraRead(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error {

}
