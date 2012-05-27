package validation

import "testing"

func TestAdd(t *testing.T) {
	var validators Validators

	var required RequiredValidator
	validators.Add("param1", required)

	// 検証ルールを１個追加
	if len(validators.rules) != 1 {
		t.Errorf("validators.Add(\"param1\", required)")
	}

	
}