package utils

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"
)

func NewArrayUtils() *ArrayUtils {
	return &ArrayUtils{}
}

type ArrayUtils struct {
}

func (t *ArrayUtils) InArray(need interface{}, needArr []interface{}) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

func (t *ArrayUtils) InArrayUint8(need uint8, needArr []uint8) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

func (t *ArrayUtils) InArrayUint64(need uint64, needArr []uint64) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

func (t *ArrayUtils) InArrayInt64(need int64, needArr []int64) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

func (t *ArrayUtils) ContainAnyArrayUint64(arr []uint64, needArr []uint64) bool {
	for _, v := range needArr {
		for _, u := range arr {
			if u == v {
				return true
			}
		}
	}
	return false
}

func (t *ArrayUtils) ContainAnyArrayString(arr []string, needArr []string) bool {
	for _, v := range needArr {
		for _, u := range arr {
			if u == v {
				return true
			}
		}
	}
	return false
}

func (t *ArrayUtils) InArrayString(need string, needArr []string) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

func (t *ArrayUtils) IndexArrayString(need string, needArr []string) bool {
	for _, v := range needArr {
		if strings.Index(need, v) >= 0 {
			return true
		}
	}
	return false
}

func (t *ArrayUtils) ArrayUniqueUint64(arr []uint64) (newArr []uint64) {
	newArr = make([]uint64, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// Contain 判断obj是否在target中，target支持的类型arrary,slice,map
func (t *ArrayUtils) Contain(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}

func (t *ArrayUtils) CloneDataMap(tags map[uint64]map[string]int64) (map[uint64]map[string]int64, error) {
	str, err := json.Marshal(tags)
	if err != nil {
		return nil, err
	}
	p := make(map[uint64]map[string]int64)
	err = json.Unmarshal(str, &p)
	return p, err
}

func (t *ArrayUtils) RemoveDuplicate(in []string) (out []string) {
	// 先排序
	sort.Strings(in)
	// 后对比
	for i, s := range in {
		if i > 0 && in[i-1] == in[i] {
			continue
		}
		out = append(out, s)
	}
	return out
}

// Merge arr2覆盖 arr1
func (t *ArrayUtils) Merge(arr1, arr2 map[string]string) map[string]string {
	for k, v := range arr2 {
		arr1[k] = v
	}
	return arr1
}
