package main

import (
	"testing"
)

func TestGreeting(t *testing.T)  {
	caseTestes := []struct {
		message string
		result string
	} {
		{"testando", "<b>testando</b>"},
		{"Code.education Rocks!", "<b>Code.education Rocks!</b>"},
		{"Rocks!", "<b>Rocks!</b>"},
		{"Code!", "<b>Code!</b>"},
	}
	for _, test := range caseTestes {
		t.Logf("Test message %v \n",test.message)
		caseTest := Greeting(test.message)

		if caseTest != test.result {
			t.Errorf("O valor esperadao é: %v, mas o retornado é %v", test.result, caseTest)
		}
	}
}