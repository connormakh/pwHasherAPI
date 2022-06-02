package utils

import (
	"testing"
)

var hashTests = []struct {
	input    string
	expected string
}{
	{"", "z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg=="},
	{"test", "7iaw3Ur350mqGo7jwQrpkj9hiYB3Lkc/iBml1JQODbJ6wYX4oOHV+E+IvIh/1nsUNzLDBMxfqa2Ob1f1ACio/w=="},
	{"test123", "2u9JU7l4M2XK1mFSI3IFBsxGxRZ80Wq1APpZeqCP+WTrJPsZaH8012Zfd4/LbFNY/ApbgeFmLPkPc6JnHFP5kQ=="},
}

func TestHash(t *testing.T) {
	for _, test := range hashTests {
		result := Hash(test.input)
		if test.expected != result {
			t.Fatalf(`Hash(%q) = %q, Expected: %q`, test.input, result, test.expected)
		}
	}

}
