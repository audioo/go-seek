/*
Copyright © 2021 ax-i-om <addressaxiom@pm.me>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package http

import (
	"os"
	"testing"
)

func TestGetReq(t *testing.T) {
	_, err := GetReq("http://azenv.net")
	if err != nil {
		t.Error(err)
	}
}

func TestAuthGet(t *testing.T) {
	_, err := AuthGet("https://haveibeenpwned.com/api/v3/breachedaccount/email@example.com", "hibp-api-key", os.Getenv("BITCROOK_HIBP"))
	if err != nil {
		t.Error(err)
	}
}
