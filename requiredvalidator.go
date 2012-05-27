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
)

type RequiredValidator struct {
	ParamName string			// パラメータ名
	ErrorMessage string			// メッセージ
}

// 未入力検証インスタンス生成
func (v RequiredValidator) Create(paramName string) Validator {
	v.ParamName  = paramName

	return v
}

// 未入力検証
func (v RequiredValidator) Validate(r *http.Request) bool {
	var result bool
	param := r.FormValue(v.ParamName)
	if param == "" {
		result = false
	} else {
		result = true
	}

	return result
}

// 未入力時のデフォルトエラーメッセージ
func(v RequiredValidator) DefaultMessage() string {
	return fmt.Sprintf("%sは必須です", v.ParamName)
}

func(v RequiredValidator) ParameterName() string {
	return v.ParamName
}

func(v RequiredValidator) Message() string {
	return v.ErrorMessage
}