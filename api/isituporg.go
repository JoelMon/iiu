/*
Copyright © 2019 Joel Montes de Oca

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getInfo(url string) int {

	type Message struct {
		domain        string
		port          int
		status_code   int
		response_ip   string
		response_code int
		response_time float64
	}

	var m Message

	resp, err := http.Get("https://isitup.org/" + url + ".json")
	if err != nil {
		log.Fatal(err)
	}

	jmsg, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jmsg, &m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Printing of jmsg: %s\n", jmsg)
	fmt.Println("----------------------")

	fmt.Printf("%v\n", m)

	return 0
}
