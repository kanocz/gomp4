//
// Copyright 2014 Hong Miao. All Rights Reserved.
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
	"github.com/kanocz/gomp4/util"
)

type SampleCountOffset []uint32

type CttsAtom struct {
	Offset int64
	Size int64
	IsFullBox bool
	
	EntriesNum uint32
	
	CttsDataTable []SampleCountOffset

	AllBytes []byte
}

func cttsRead(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error {
	var err error
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.CttsAtomInstance.Offset = offset
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.CttsAtomInstance.IsFullBox = false
	err = fp.Mp4Seek(offset, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	size, _, err := fp.Mp4ReadHeader()
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	sizeInt := util.Bytes2Int(size)	
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.CttsAtomInstance.Size = sizeInt
		
	err = fp.Mp4Seek(offset, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	buf, err := fp.Mp4Read(fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.CttsAtomInstance.Size)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.CttsAtomInstance.AllBytes = buf

	err = fp.Mp4Seek(offset + 12, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}	

	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.CttsAtomInstance.EntriesNum = util.Byte42Uint32(size, 0)
		
	//log.Println(fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		//StblAtomInstance.CttsAtomAtomInstance.EntriesNum)
		
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.CttsAtomInstance.CttsDataTable = 
		make([]SampleCountOffset, fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.CttsAtomInstance.EntriesNum)
	
	var i uint32
	for i = 0; i < fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.CttsAtomInstance.EntriesNum; i++ {	
		
		buf, err := fp.Mp4Read(8)
		if err != nil {
			log.Fatalln(err.Error())
			return err
		}
			
		fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
			StblAtomInstance.CttsAtomInstance.CttsDataTable[i] = append(fs.MoovAtomInstance.
			TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.StblAtomInstance.
			CttsAtomInstance.CttsDataTable[i], util.Byte42Uint32(buf[0:4], 0))
		
		fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
			StblAtomInstance.CttsAtomInstance.CttsDataTable[i] = append(fs.MoovAtomInstance.
			TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.StblAtomInstance.
			CttsAtomInstance.CttsDataTable[i], util.Byte42Uint32(buf[4:8], 0))
	}
		
	return nil
}