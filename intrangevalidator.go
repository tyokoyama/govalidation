/*
Copyright 2012 Takashi Yokoyama

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
package validation

import (
	"fmt"
	"net/http"
	"strconv"
)

type IntRangeValidator struct {
	ParamName string			// パラメータ名
	BitSize int					// サイズ
	Min int						// 最小値
	Max int						// 最大値
	ErrorMessage string			// メッセージ
}

// 整数型検証インスタンス生成
func (v IntRangeValidator) Create(paramName string) Validator {
	v.ParamName  = paramName

	return v
}

// 整数型検証
func (v IntRangeValidator) Validate(r *http.Request) bool {
	var result bool
	param := r.FormValue(v.ParamName)
	if param == "" {
		// 未入力はOKとする。
		result = true
	} else {
		// 取り敢えず型変換を試みる。
		if conv, err := strconv.ParseInt(param, 10, v.BitSize); err != nil {
			result = false
		} else {
			if (int64(v.Min) <= conv && conv <= int64(v.Max)) {
				result = true
			} else {
				result = false
			}
		}
	}

	return result	
}

// 整数型検証失敗時のデフォルトエラーメッセージ
func(v IntRangeValidator) DefaultMessage() string {
	return fmt.Sprintf("%sは%dから%dの範囲内である必要があります。", v.ParamName, v.Min, v.Max)
}

func(v IntRangeValidator) Message() string {
	return v.ErrorMessage
}

func(v IntRangeValidator) ParameterName() string {
	return v.ParamName
}
