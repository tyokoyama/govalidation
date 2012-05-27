package validation

import (
	"net/http"
)

type Validators struct {
	errorMap map[string] string		// 検証エラーメッセージ
	rules []Validator				// 検証ルール
}

// 検証インターフェイス
type Validator interface {
	Validate(r *http.Request) bool
}

type RequiredValidator struct {
	ParamName string			// パラメータ名
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

// 検証ルールの追加
func (v *Validators) Add(paramName string, rule ...Validator) {
	for _, r := range rule {
		v.rules = append(v.rules, r)
	}
}

// 検証
func (v *Validators) Validate(r *http.Request) bool {
	for _, validator := range v.rules {
		if result := validator.Validate(r); result == false {
			// 検証失敗

		}
	}

	return true
}
