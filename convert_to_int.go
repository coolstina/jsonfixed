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

type IntegerType int

const (
	IntegerTypeOfUint8 = iota
	IntegerTypeOfUint16
	IntegerTypeOfUint32
	IntegerTypeOfUint64
	IntegerTypeOfUint
	IntegerTypeOfInt8
	IntegerTypeOfInt16
	IntegerTypeOfInt32
	IntegerTypeOfInt64
	IntegerTypeOfInt
)

type DestinationOfInteger struct {
	destination interface{}
}

func NewDestinationOfInteger(destination interface{}) *DestinationOfInteger {
	return &DestinationOfInteger{destination: destination}
}

func (c *DestinationOfInteger) GetOriginal() interface{} {
	return nil
}

func (c *DestinationOfInteger) GetDestination() interface{} {
	return c.destination
}
