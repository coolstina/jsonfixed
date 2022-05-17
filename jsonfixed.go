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
	"encoding/json"
	"strconv"
	"time"
)

// Convert JSON strings to handle specified field values to a specific type.
func Convert(data []byte, template Template) ([]byte, error) {
	r, err := mapper(data)
	if err != nil {
		return nil, err
	}

	for k, v := range r {
		dst, ok := template[k]

		if ok && v != nil {
			original := dst.GetOriginal()
			destination := dst.GetDestination()

			switch dst.(type) {
			case *DestinationOfTime:
				parsed, err := timer(r, original, destination, v)
				if err != nil {
					return nil, err
				}
				r[k] = parsed
			case *DestinationOfNull:
				switch v.(type) {
				case string:
					if original.(string) == v.(string) {
						r[k] = nil
					}
				}
			case *DestinationOfBool:
				switch v.(type) {
				case float64:
					r[k] = int(v.(float64)) == 1
				case string:
					i, err := strconv.Atoi(v.(string))
					if err != nil {
						return nil, err
					}
					r[k] = i == 1
				}
			case *DestinationOfInteger:
				val, err := inter(v, IntegerType(destination.(int)))
				if err != nil {
					return nil, err
				}
				r[k] = val
			case *DestinationOfFloat:
				val, err := floater(v, FloatType(destination.(int)))
				if err != nil {
					return nil, err
				}
				r[k] = val
			}
		}
	}

	marshal, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return marshal, nil
}

func inter(original interface{}, destination IntegerType) (interface{}, error) {

	i := func(dst int) interface{} {

		switch destination {
		case IntegerTypeOfUint8:
			original = uint8(dst)
		case IntegerTypeOfUint16:
			original = uint16(dst)
		case IntegerTypeOfUint32:
			original = uint32(dst)
		case IntegerTypeOfUint64:
			original = uint64(dst)
		case IntegerTypeOfUint:
			original = uint(dst)
		case IntegerTypeOfInt8:
			original = int8(dst)
		case IntegerTypeOfInt16:
			original = int16(dst)
		case IntegerTypeOfInt32:
			original = int32(dst)
		case IntegerTypeOfInt64:
			original = int64(dst)
		case IntegerTypeOfInt:
			original = int(dst)
		}

		return original
	}

	switch val := original.(type) {
	case string:
		dst, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		original = i(dst)
	}

	return original, nil
}

func floater(original interface{}, destination FloatType) (interface{}, error) {
	switch val := original.(type) {
	case string:
		var bit int
		switch destination {
		case FloatTypeOfFloat32:
			bit = 32
		case FloatTypeOfFloat64:
			bit = 64
		}

		float, err := strconv.ParseFloat(val, bit)
		if err != nil {
			return nil, err
		}

		return float, nil
	}

	return original, nil
}

func timer(r map[string]interface{}, original interface{}, destination interface{}, val interface{}) (interface{}, error) {
	switch original {
	case TimeFormatOfChinaDate:
		fallthrough
	case TimeFormatOfRFC3339:
		fallthrough
	case TimeFormatOfRFC3339Nano:
		fallthrough
	case TimeFormatOfRFC1123:
		fallthrough
	case TimeFormatOfRFC1123Z:
		fallthrough
	case TimeFormatOfChina24:
		tm, err := time.Parse(original.(TimeFormat).String(), val.(string))
		if err != nil {
			return nil, err
		}
		return tm.Format(destination.(TimeFormat).String()), nil
	case TimeFormatOfTimestamp:
		var vi int64

		switch val.(type) {
		case string:
			i, err := strconv.Atoi(val.(string))
			if err != nil {
				return nil, err
			}
			vi = int64(i)
		case float64:
			vi = int64(val.(float64))
		}

		return time.Unix(vi, 0).Format(destination.(TimeFormat).String()), nil
	}
	return original, nil
}

func mapper(data []byte) (map[string]interface{}, error) {
	var r map[string]interface{}
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}
	return r, nil
}
