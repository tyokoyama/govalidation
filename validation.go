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
	"net/http"
)

type Validators struct {
	errorMap map[string] []string	// 検証エラーメッセージ
	rules []Validator				// 検証ルール
}

// 検証インターフェイス
type Validator interface {
	Create(string) Validator
	Validate(r *http.Request) bool
	DefaultMessage() string
	Message() string				// メッセージ取得
	ParameterName() string			// パラメータ取得
}

// 検証ルールの追加
func (v *Validators) Add(paramName string, rule ...Validator) {
	for _, r := range rule {
		v.rules = append(v.rules, r.Create(paramName))
	}
}

// 検証
func (v *Validators) Validate(r *http.Request) bool {
	var result bool = true
	for pos, validator := range v.rules {
		if validate := validator.Validate(r); validate == false {
			// 検証失敗
			message := v.rules[pos].Message()
			if message == "" {
				message = v.rules[pos].DefaultMessage()
			}

			if v.errorMap == nil {
				v.errorMap = make(map[string] []string)
			}
			v.errorMap[v.rules[pos].ParameterName()] = append(v.errorMap[v.rules[pos].ParameterName()], message)
			result = false
		}
	}

	return result
}

// 検証エラー取得
func (v *Validators) Errors() map[string] []string {
	return v.errorMap
}

// パラメータ毎の検証エラー取得
func (v *Validators) Error(paramName string) []string {
	return v.errorMap[paramName]
}