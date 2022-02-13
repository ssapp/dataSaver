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

package dataSaver

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"sync"
)

var fileLock sync.Mutex

// Marshals object into an io.Reader.
// By default, it uses the JSON marshaller.
var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

// Unmarshals reader's data into the specified value.
// By default, it uses the JSON unmarshaller.
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// Saves a representation of v to the file at p.
func Save(p string, v interface{}) error {
	fileLock.Lock()
	defer fileLock.Unlock()
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

// Loads the file at p into v. Use os.IsNotExist() to see
// if the returned error is due to the file being missing.
func Load(p string, v interface{}) error {
	fileLock.Lock()
	defer fileLock.Unlock()
	f, err := os.Open(p)
	if err != nil {
		return err
	}
	defer f.Close()
	return Unmarshal(f, v)
}
