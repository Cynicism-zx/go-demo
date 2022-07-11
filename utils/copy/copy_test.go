package copy

import "testing"

type User struct {
	Name string
}

//func TestDeepCopy(t *testing.T) {
//	user1 := &User{Name: "pibigstar"}
//
//	var user2 User
//	err := DeepCopy(user1, &user2)
//	if err != nil {
//		t.Error(err)
//	}
//	t.Log(user2)
//
//	var user3 User
//	err = Copy(user1, &user3)
//	if err != nil {
//		t.Error(err)
//	}
//	t.Log(user3)
//}

func BenchmarkDeepCopy(b *testing.B) {
	user1 := &User{Name: "pibigstar"}
	for i := 0; i < b.N; i++ {
		var user2 User
		err := DeepCopy(user1, &user2)
		if err != nil {
			b.Error(err)
		}
		b.Log(user2)
	}
}

func BenchmarkCopy(b *testing.B) {
	user1 := &User{Name: "pibigstar"}
	for i := 0; i < b.N; i++ {
		var user2 User
		err := Copy(user1, &user2)
		if err != nil {
			b.Error(err)
		}
		b.Log(user2)
	}
}
