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

import "time"

type TimeFormat string

func (t TimeFormat) String() string {
	return string(t)
}

const (
	TimeFormatOfTimestamp   TimeFormat = "timestamp"
	TimeFormatOfChina24     TimeFormat = "2006-01-02 15:04:05"
	TimeFormatOfChinaDate   TimeFormat = "2006-01-02"
	TimeFormatOfRFC3339     TimeFormat = time.RFC3339
	TimeFormatOfRFC3339Nano TimeFormat = time.RFC3339Nano
	TimeFormatOfRFC1123     TimeFormat = time.RFC1123
	TimeFormatOfRFC1123Z    TimeFormat = time.RFC1123Z
)

type DestinationOfTime struct {
	Original    TimeFormat
	Destination TimeFormat
}

func NewDestinationOfTime(original TimeFormat, destination TimeFormat) *DestinationOfTime {
	return &DestinationOfTime{Original: original, Destination: destination}
}

func (c *DestinationOfTime) GetOriginal() interface{} {
	return c.Original
}

func (c *DestinationOfTime) GetDestination() interface{} {
	return c.Destination
}
