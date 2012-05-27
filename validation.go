package validation

import (
	"fmt"
	"net/http"
	"strconv"
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

type RequiredValidator struct {
	ParamName string			// パラメータ名
	ErrorMessage string			// メッセージ
}

type IntValidator struct {
	ParamName string			// パラメータ名
	BitSize int					// サイズ
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

// 整数型検証インスタンス生成
func (v IntValidator) Create(paramName string) Validator {
	v.ParamName  = paramName

	return v
}

// 整数型検証
func (v IntValidator) Validate(r *http.Request) bool {
	var result bool
	param := r.FormValue(v.ParamName)
	if param == "" {
		// 未入力はOKとする。
		result = true
	} else {
		// 取り敢えず型変換を試みる。
		if _, err := strconv.ParseInt(param, 10, v.BitSize); err != nil {
			result = false
		} else {
			result = true			
		}
	}

	return result	
}

// 整数型検証失敗時のデフォルトエラーメッセージ
func(v IntValidator) DefaultMessage() string {
	return fmt.Sprintf("%sは%dビット整数ではありません。", v.ParamName, v.BitSize)
}

func(v IntValidator) Message() string {
	return v.ErrorMessage
}

func(v IntValidator) ParameterName() string {
	return v.ParamName
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
			fmt.Printf("DEBUG: %s", v.rules[pos].ParameterName())
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