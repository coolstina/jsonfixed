// Copyright 2021 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jsonfixed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	const data = `{
		"id":1,
		"username":"helloshaohua",
		"signature":"",
		"birth_date":"2019-08-27",
		"logged_int":1,
		"logged_str":"1",
		"logout_int":0,
		"logout_str":"0",
		"level":1,
		"created_at":"2021-09-13 16:56:59",
		"updated_at":"2021-09-13 17:20:57",
		"graduated_int_at":1433841657,
		"graduated_str_at":"1433841657",
		"deleted_at":null,
		"fail_data":"<nil>",
		"qq":"100000"
	}`

	template := Template{
		"created_at":       NewDestinationOfTime(TimeFormatOfChina24, TimeFormatOfRFC3339),
		"updated_at":       NewDestinationOfTime(TimeFormatOfChina24, TimeFormatOfRFC3339),
		"birth_date":       NewDestinationOfTime(TimeFormatOfChinaDate, TimeFormatOfRFC3339),
		"graduated_int_at": NewDestinationOfTime(TimeFormatOfTimestamp, TimeFormatOfRFC3339),
		"graduated_str_at": NewDestinationOfTime(TimeFormatOfTimestamp, TimeFormatOfRFC3339),
		"signature":        NewDestinationOfNull(""),
		"fail_data":        NewDestinationOfNull("<nil>"),
		"logged_int":       NewDestinationOfBool(1),
		"logged_str":       NewDestinationOfBool("1"),
		"logout_int":       NewDestinationOfBool(0),
		"logout_str":       NewDestinationOfBool("0"),
		"qq":               NewDestinationOfInteger(IntegerTypeOfInt),
	}

	actual, err := Convert([]byte(data), template)
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}
