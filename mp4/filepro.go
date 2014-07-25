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
	"os"
	"log"
	"bufio"
	//"../util"
)

type Mp4FilePro struct {
	file *os.File
	r *bufio.Reader
	
}

func NewMp4FilePro () *Mp4FilePro {
	fp := &Mp4FilePro {
		
	}
	
	return fp
}

func (self * Mp4FilePro) Mp4Open(fs *Mp4FileSpec)  error {
	var err error
	//log.Println(fs.mp4Name)
	self.file, err = os.Open(fs.mp4Name)
	
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	self.r = bufio.NewReader(self.file)
	
	return nil
}

func (self * Mp4FilePro) Mp4Read() (size []byte, box []byte, err error) {
	buf := make([]byte, 8)
	_, err = self.r.Read(buf)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, nil, err
	}

	return buf[0:4], buf[4:8], nil
}


func (self * Mp4FilePro) Mp4Seek(offset int64)  (err error) {
	//ret, err = self.file.Seek(offset, whence)
	//if err != nil {
	//	log.Fatalln(err.Error())
	//	return
	//}
	buf := make([]byte, offset)
	_, err = self.r.Read(buf)
	if err != nil {
		log.Fatalln(err.Error())
		return  err
	}
	
	return nil
}


func (self * Mp4FilePro) Mp4FileStat(fs *Mp4FileSpec) error {
	fi ,err := self.file.Stat() 
	
	if err != nil {
		log.Fatalln(err.Error())
		return  err
	}
	
	fs.end = fi.Size()
	
	return nil
}

