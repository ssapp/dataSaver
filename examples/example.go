/*
 * Copyright (c) 2022 Vladislav Naydenov <v.naydenov@icloud.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"dataSaver"
	"fmt"
	"log"
	"time"
)

type exampleObject struct {
	Event string
	Time  time.Time
}

func main() {
	firstObject := &exampleObject{
		Event: "Just testing the dataSaver package",
		Time:  time.Now(),
	}
	if err := dataSaver.Save("./file.tmp", firstObject); err != nil {
		log.Fatalln(err)
	}

	// load it back
	var secondObject exampleObject
	if err := dataSaver.Load("./example.tmp", &secondObject); err != nil {
		log.Fatalln(err)
	}
	// o and o2 are now the same
	// and check out file.tmp - you'll see the JSON file

	fmt.Printf("%v", secondObject)
}
