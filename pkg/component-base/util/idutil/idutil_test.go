package idutil

import "testing"

func TestGetUUID36(t *testing.T) {
	t.Log(GetUUID36(""))
}

func TestGetManyUuid(t *testing.T) {
	for i := 0; i < 10000; i++ {
		testID := GetUUID36("")
		if len(testID) != 14 {
			t.Errorf("GetUUID failed, expected uuid length 14, got: %d", len(testID))
		}
	}
}

func TestNewSecretID(t *testing.T) {
	t.Log(NewSecretID())
}

func TestNewSecretKey(t *testing.T) {
	t.Log(NewSecretKey())
}
