package utils

import "strconv"

// string类型的纯数字数据转换成int类型数据
func StringTOInt(s string) (int, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}
