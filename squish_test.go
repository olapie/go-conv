package conv_test

import (
	"testing"
	"time"

	"code.olapie.com/conv"
)

func TestSquishString(t *testing.T) {
	s := "\n\n\t\t1 \t2\n3  4\n\t \n5   "
	equal(t, "1 2 3 4 5", conv.SquishString(s))
}

func TestSquishStringFields(t *testing.T) {
	type FullName struct {
		FirstName string
		LastName  string
	}

	type User struct {
		ID        int64
		Name      FullName
		Address   string
		CreatedAt time.Time
	}

	u := &User{
		ID: 1,
		Name: FullName{
			FirstName: "Tom  ",
			LastName:  "  Jim ",
		},
		Address: "\n\n \t Toronto   Canada   ",
	}
	conv.SquishStringFields(u)
	equal(t, "Tom", u.Name.FirstName)
	equal(t, "Jim", u.Name.LastName)
	equal(t, "Toronto Canada", u.Address)
}

func TestRemoveAllSpaces(t *testing.T) {
	s := "\n\ra b c  d "
	res := conv.RemoveAllSpaces(s)
	equal(t, "abcd", res)
}

func TestRemoveBullet(t *testing.T) {
	s := "1. hello"
	equal(t, "hello", conv.RemoveBullet(s))

	s = "12. hello"
	equal(t, "hello", conv.RemoveBullet(s))

	s = ". hello"
	equal(t, "hello", conv.RemoveBullet(s))

	s = "* hello"
	equal(t, "hello", conv.RemoveBullet(s))

	s = "*hello"
	equal(t, "hello", conv.RemoveBullet(s))

	s = "12.hello"
	equal(t, "hello", conv.RemoveBullet(s))
}
