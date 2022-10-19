package tools

import "testing"

func TestSetPort(t *testing.T) {
	t.Run("Success ", func(t *testing.T) {
		expectedValue := ":5000"
		value := "5000"

		SetPort(&value)

		if value != expectedValue {
			t.Errorf("Port not formated")
		}
	})
}
