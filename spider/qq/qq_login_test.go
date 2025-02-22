package qq

import (
	"testing"

	"go-demo/utils/env"
)

func TestQQLogin(t *testing.T) {
	if env.IsCI() {
		return
	}

	user, err := GetQQInfo(QQZone)
	if err != nil {
		t.Log(err)
	}
	t.Logf("%+v \n", user)

	user, err = GetQQInfo(QQFriend)
	if err != nil {
		t.Log(err)
	}
	t.Logf("%+v \n", user)
}

func TestGenderTtk(t *testing.T) {
	t.Log(genderGTK("@lasyOoUaj"))
}
