package validation

import (
	"net/http"
	"testing"
)

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
	r, err := http.NewRequest("Get", "http://localhost", nil)
	if err != nil {
		t.Errorf("Request Create Error")
	}

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
	r, err := http.NewRequest("Get", "http://localhost", nil)
	if err != nil {
		t.Errorf("Request Create Error")
	}

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