/***************************************************************************
   copyright            : (C) 2014 by Nick Sellen
   email                : code@nicksellen.co.uk
***************************************************************************/

/***************************************************************************
 *   This library is free software; you can redistribute it and/or modify  *
 *   it  under the terms of the GNU Lesser General Public License version  *
 *   2.1 as published by the Free Software Foundation.                     *
 *                                                                         *
 *   This library is distributed in the hope that it will be useful, but   *
 *   WITHOUT ANY WARRANTY; without even the implied warranty of            *
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU     *
 *   Lesser General Public License for more details.                       *
 *                                                                         *
 *   You should have received a copy of the GNU Lesser General Public      *
 *   License along with this library; if not, write to the Free Software   *
 *   Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA  02111-1307  *
 *   USA                                                                   *
 ***************************************************************************/

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nicksellen/audiotags"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("pass path to file")
		return
	}

	f, err := audiotags.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	props := f.ReadTags()
	for k, v := range props {
		fmt.Printf("%s %s\n", k, strings.Replace(strings.Replace(v, "\r\n", "\\n", -1), "\n", "\\n", -1))
	}

	audioProps := f.ReadAudioProperties()
	fmt.Printf("length %d\nbitrate %d\nsamplerate %d\nchannels %d\n",
		audioProps.Length, audioProps.Bitrate, audioProps.Samplerate, audioProps.Channels)

	images := f.ReadImages()
	for k, v := range images {
		fmt.Printf("Image %s size: %d\n", k, len(v))
	}
}
