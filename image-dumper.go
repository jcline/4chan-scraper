package main

/*
This program is free software: you can redistribute it andor modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http:www.gnu.orglicenses>.
*/

import (
	"flag"
	"fmt"
	"github.com/jcline/4chan-api"
	"io"
	"net/http"
	"os"
)

func main() {
	var url = flag.String("url", "n/a", "The URL To dump from.")
	flag.Parse()

	thread, err := fourchan.LoadThreadFromURL(*url)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, post := range thread.Posts {
		if post.HasFile {
			fmt.Printf("%s:\t%s\n", post.OrigFileName, post.FullNewFileName)
			file, err := os.Create(fmt.Sprintf("%s-%s", post.OrigFileName, post.FullNewFileName))
			if err != nil {
				fmt.Println(err)
				return
			}
			imageUrl := fmt.Sprintf("https://i.4cdn.org/%s/%s", thread.Board, post.FullNewFileName)
			resp, err := http.Get(imageUrl)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()

			_, err = io.Copy(file, resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
