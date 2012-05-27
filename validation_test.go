package validation

import (
	"net/http"
	"net/url"
	"testing"
)

func createRequest(t *testing.T) *http.Request {
	r, err := http.NewRequest("Get", "http://localhost", nil)
	if err != nil {
		t.Errorf("Request Create Error")
	}

	return r
}

func TestAdd(t *testing.T) {
	var validators Validators

	var required RequiredValidator
	var integer IntValidator

	validators.Add("param1", required)

	// 検証ルールを１個追加
	if len(validators.rules) != 1 {
		t.Errorf("validators.Add(\"param1\", required)")
	}

	// 別のパラメータで検証ルールを２個追加
	validators.Add("param2", required, integer)

	if len(validators.rules) != 3 {
		t.Errorf("validators.Add(\"param2\", required, int32type)")		
	}
}

// リクエストパラメータなし（必須チェック）
func TestValidateNoParam(t *testing.T) {
	r := createRequest(t)

	var validators Validators

	var param1Required RequiredValidator

	validators.Add("param1", param1Required)

	if result := validators.Validate(r); result == true {
		t.Errorf("validators.Validate(r) is failed")
	}

	if len(validators.Error("param1")) != 1 {
		t.Errorf("validators errorMap is wrong. [%d != %d]", len(validators.Error("param1")), 1)
	}

	if validators.Error("param1")[0] != "param1は必須です" {
		t.Errorf("validators errorMap is wrong. [%s != %s]", validators.Error("param1")[0], "param1は必須です")
	}
}

// リクエストパラメータなし（整数型、必須ではない）
func TestValidateNoParamInteger(t *testing.T) {
	r := createRequest(t)

	var validators Validators

	var param1Integer IntValidator
	param1Integer.BitSize = 32

	validators.Add("param1", param1Integer)

	if result := validators.Validate(r); result == false {
		t.Errorf("validators.Validate(r) is failed")
	}

	if len(validators.Error("param1")) != 0 {
		t.Errorf("validators errorMap is wrong.")
	}
}

// 整数型パラメータチェック（整数型のチェックに文字列を入れる）
func TestIntegerValidator1(t *testing.T) {
	r := createRequest(t)

	r.Form = make(url.Values)

	r.Form.Add("param1", "aaa")

	var validators Validators

	var param1Integer IntValidator
	param1Integer.BitSize = 32

	validators.Add("param1", param1Integer)

	if result := validators.Validate(r); result == true {
		t.Errorf("validators.Validate(r) is failed")
	}

	if len(validators.Error("param1")) != 1 {
		t.Errorf("validators errorMap is wrong.")
	}

	if validators.Error("param1")[0] != "param1は32ビット整数ではありません。" {
		t.Errorf("validators errorMap is wrong. [%s != %s]", validators.Error("param1")[0], "param1は32ビット整数ではありません。")
	}
}

// 整数型パラメータチェック（整数型のチェックに整数を入れる）
func TestIntegerValidator2(t *testing.T) {
	r := createRequest(t)

	r.Form = make(url.Values)

	r.Form.Add("param1", "123")

	var validators Validators

	var param1Integer IntValidator
	param1Integer.BitSize = 32

	validators.Add("param1", param1Integer)

	if result := validators.Validate(r); result == false {
		t.Errorf("validators.Validate(r) is failed")
	}

	if len(validators.Error("param1")) != 0 {
		t.Errorf("validators errorMap is wrong. [%s]", validators.Error("param1")[0])
	}
}

// 整数型範囲チェック
func TestIntRangeValidator1(t *testing.T) {
	r := createRequest(t)

	r.Form = make(url.Values)

	r.Form.Add("param1", "6")

	var validatorsMax Validators
	var validatorsMin Validators
	var validators Validators

	var param1IntRange IntRangeValidator
	param1IntRange.BitSize = 32
	param1IntRange.Min = 0
	param1IntRange.Max = 5

	validatorsMax.Add("param1", param1IntRange)

	if result := validatorsMax.Validate(r); result == true {
		t.Errorf("validators.Validate(r) is failed")
	}

	if len(validatorsMax.Error("param1")) != 1 {
		t.Errorf("validators errorMap is wrong.[%d]", len(validatorsMax.Error("param1")))
	}

	if validatorsMax.Error("param1")[0] != "param1は0から5の範囲内である必要があります。" {
		t.Errorf("validators errorMap is wrong. [%s != %s]", validatorsMax.Error("param1")[0], "param1は0から5の範囲内である必要があります。")
	}

	r.Form.Set("param1", "-1")

	validatorsMin.Add("param1", param1IntRange)

	if result := validatorsMin.Validate(r); result == true {
		t.Errorf("validators.Validate(r) is failed")
	}

	if len(validatorsMin.Error("param1")) != 1 {
		t.Errorf("validators errorMap is wrong.[%d]", len(validatorsMin.Error("param1")))
	}

	if validatorsMin.Error("param1")[0] != "param1は0から5の範囲内である必要があります。" {
		t.Errorf("validators errorMap is wrong. [%s != %s]", validatorsMin.Error("param1")[0], "param1は0から5の範囲内である必要があります。")
	}

	r.Form.Set("param1", "0")

	validators.Add("param1", param1IntRange)

	if result := validators.Validate(r); result == false {
		t.Errorf("validators.Validate(r) is failed")
	}

	if len(validators.Error("param1")) != 0 {
		t.Errorf("validators errorMap is wrong. [%s]", validators.Error("param1")[0])
	}

	r.Form.Set("param1", "5")
	if result := validators.Validate(r); result == false {
		t.Errorf("validators.Validate(r) is failed")
	}

	if len(validators.Error("param1")) != 0 {
		t.Errorf("validators errorMap is wrong. [%s]", validators.Error("param1")[0])
	}

}