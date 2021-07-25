package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// 定义解析错误
var (
	ErrTypeError           = errors.New("type error")
	ErrConvertToIntError   = errors.New("conver to int error")
	ErrConvertToFloatError = errors.New("conver to float error")
	ErrConvertToBoolError  = errors.New("conver to bool error")
)

// ParseInt64 从interface{}解析出int64
func ParseInt64(in interface{}) (int64, error) {
	switch x := in.(type) {
	case int:
		return int64(x), nil
	case int8:
		return int64(x), nil
	case int16:
		return int64(x), nil
	case int32:
		return int64(x), nil
	case int64:
		return x, nil
	case uint:
		return int64(x), nil
	case uint8:
		return int64(x), nil
	case uint16:
		return int64(x), nil
	case uint32:
		return int64(x), nil
	case uint64:
		return int64(x), nil
	case float32:
		return int64(x), nil
	case float64:
		return int64(x), nil
	case string:
		i, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			return 0, ErrConvertToIntError
		}
		return i, nil
	default:
		return 0, ErrTypeError
	}
}

// MustInt64 按int64去解析
func MustInt64(in interface{}) int64 {
	v, _ := ParseInt64(in)
	return v
}

// ParseFloat64 从interface{}解析出float64
func ParseFloat64(in interface{}) (float64, error) {
	switch x := in.(type) {
	case int:
		return float64(x), nil
	case int8:
		return float64(x), nil
	case int16:
		return float64(x), nil
	case int32:
		return float64(x), nil
	case int64:
		return float64(x), nil
	case uint:
		return float64(x), nil
	case uint8:
		return float64(x), nil
	case uint16:
		return float64(x), nil
	case uint32:
		return float64(x), nil
	case uint64:
		return float64(x), nil
	case float32:
		return float64(x), nil
	case float64:
		return x, nil
	case string:
		f, err := strconv.ParseFloat(x, 64)
		if err != nil {
			return 0.0, ErrConvertToFloatError
		}
		return f, nil
	default:
		return 0.0, ErrTypeError
	}
}

// MustFloat64 按float64去解析
func MustFloat64(in interface{}) float64 {
	v, _ := ParseFloat64(in)
	return v
}

// ParseBool 从interface{}解析出bool
func ParseBool(in interface{}) (bool, error) {
	switch x := in.(type) {
	case bool:
		return x, nil
	case string:
		b, err := strconv.ParseBool(x)
		if err != nil {
			return false, ErrConvertToBoolError
		}
		return b, nil
	default:
		return false, ErrTypeError
	}
}

// MustBool 从interface{}解析出bool
func MustBool(in interface{}) bool {
	v, _ := ParseBool(in)
	return v
}

// ParseStringSlice 从interface{}解析出字符串列表
func ParseStringSlice(in interface{}) ([]string, error) {
	switch x := in.(type) {
	case []string:
		return x, nil
	case []interface{}:
		vs := make([]string, 0, len(x))
		for _, i := range x {
			vs = append(vs, fmt.Sprintf("%v", i))
		}
		return vs, nil
	case string:
		vs := make([]string, 0, 10)
		err := json.Unmarshal([]byte(x), &vs)
		if err != nil {
			return []string{x}, nil
		}
		return vs, nil
	default:
		return []string{}, ErrTypeError
	}
}

// MustStringSlice 从interface{}解析出字符串列表
func MustStringSlice(in interface{}) []string {
	v, _ := ParseStringSlice(in)
	return v
}

// ParseInt64Slice 从interface{}解析出字符串列表
func ParseInt64Slice(in interface{}) ([]int64, error) {
	switch x := in.(type) {
	case []int64:
		return x, nil
	// TODO: 其他类型强转
	default:
		return []int64{}, ErrTypeError
	}
}

// MustInt64Slice 从interface{}解析出字符串列表
func MustInt64Slice(in interface{}) []int64 {
	v, _ := ParseInt64Slice(in)
	return v
}

// PathExists 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDirIfNotExist 如不存在则创建目录
func CreateDirIfNotExist(path string) error {
	ok, err := PathExists(path)
	if ok {
		return err
	}
	return os.MkdirAll(path, 0666)
}

// CreateFileIfNotExist 如不存在则创建文件
func CreateFileIfNotExist(path string) error {
	ok, err := PathExists(path)
	if ok {
		return err
	}
	return ioutil.WriteFile(path, []byte(""), 0666)
}

//BuildMD5 计算MD5值
func BuildMD5(value string) string {
	h := md5.New()
	h.Write([]byte(value))

	return hex.EncodeToString(h.Sum(nil))
}
