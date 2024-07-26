package tap

import (
	"testing"
)

func create(t *testing.T) {
	got, err := CreateTap()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf(got.Name())
}

func TestCreateTap(t *testing.T) {
	create(t)
	create(t)
	create(t)
}
