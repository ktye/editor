package sam

import (
	"testing"
)

// 01234567890	lineAddresses
// alpha|	0-5
// beta|	6-10
// gamma|	11-16
//
// 01234567890123456
// alpha|beta|gamma|

type testcase struct {
	cmd      string
	initDot  string
	expected string
}

func TestAddr(t *testing.T) {
	var testcases []testcase = []testcase{
		testcase{",", "", "1:1,3:6"},
		testcase{"1,$", "", "1:1,3:6"},
		testcase{"1", "", "1:1,1:6"},
		testcase{"2", "", "2:1,2:5"},
		testcase{"3", "", "3:1,3:6"},
		testcase{"2,3", "", "2:1,3:6"},
		testcase{"2-1", "", "1:1,1:6"},
		testcase{"/ph/", "", "1:3,1:5"},
		testcase{"/eta/", "", "2:2,2:5"},
		testcase{"/eta/+1", "", "3:1,3:6"},
		testcase{"/eta/+1-1", "", "2:1,2:5"},
		testcase{"/eta/+-", "", "2:1,2:5"},
		testcase{"/mm/+-", "", "3:1,3:6"},
		testcase{"/amma/+1", "", "3:6,3:6"},
		testcase{"/eta/+2", "", "3:6,3:6"},
		testcase{"/mm/", "", "3:3,3:5"},
		testcase{"/ph/,/mm/", "", "1:3,3:5"},
		testcase{"/a/", "1:5,1:6", "2:4,2:5"},
		testcase{"/a/", "1:5,1:6;3:1,3:1", "2:4,2:5;3:2,3:3"},
		testcase{"/l/", "2:1,2:1", "1:2,1:3"},
		testcase{"-/a/", "2:1,2:1", "1:5,1:6"},
		testcase{"$-/m/", "", "3:4,3:5"},
		testcase{"/p/+1,2", "", "2:1,2:5"},
	}
	file := "alpha\nbeta\ngamma\n"
	for i := 0; i < len(testcases); i++ {
		cmd := testcases[i].cmd
		initDot := testcases[i].initDot
		expected := testcases[i].expected
		_, addr, err := Edit([]byte(file), cmd, initDot)
		if err != nil {
			t.Error("cmd:", cmd, "initDot:", initDot, "dot:", addr, " [failed]", err)
		} else if addr != expected {
			t.Error("cmd:", cmd, "initDot:", initDot, "dot:", addr, " [failed]", err)
			t.Errorf("expected: " + expected + " but got: " + addr)
			//		} else {
			//			fmt.Println("cmd:", cmd, "initDot:", initDot, "dot:", addr, " [ok]")
		}
	}
}

func TestCmd(t *testing.T) {
	file := "alpha\n"
	var testcases []testcase = []testcase{
		testcase{"/ph/a/xx/", "1:1,1:6", "alphxxa\n"},
		testcase{"/ph/ a/xx/", "1:1,1:6", "alphxxa\n"},
		testcase{"/ph/i/xx/", "1:1,1:6", "alxxpha\n"},
		testcase{"/ph/c/xxx/", "1:1,1:6", "alxxxa\n"},
		testcase{"/ph/d/xx/", "1:1,1:6", "ala\n"},
		testcase{"s/a/b/", "1:1,1:6", "blpha\n"},
		testcase{"s/l(p)h/x${1}x/", "1:1,1:6", "axpxa\n"},
		testcase{"/a/,/a/s/a/b/", "1:1,1:6", "blpha\n"},
		testcase{"/l/,/a/s/a/b/", "1:1,1:6", "alphb\n"},
		testcase{"g/ph/ a/xx/", "1:1,1:6", "alphaxx\n"},
		testcase{"v/ph/ a/xx/", "1:1,1:6", "alpha\n"},
		testcase{"v/x/ i/m/", "1:1,1:6", "malpha\n"},
		testcase{"x/a/c/b/", "1:1,1:6", "blphb\n"},
		testcase{"x/a/c/123456/", "1:1,1:6", "123456lph123456\n"},
	}
	for i := 0; i < len(testcases); i++ {
		cmd := testcases[i].cmd
		initDot := testcases[i].initDot
		expected := testcases[i].expected
		b, _, err := Edit([]byte(file), cmd, initDot)
		if err != nil {
			t.Error("cmd:", cmd, " [failed]", err)
		} else if string(b) != expected {
			t.Error("cmd:", cmd, " [failed]")
			t.Errorf("expected: " + expected + " but got: " + string(b))
			//		} else {
			//			fmt.Println("cmd:", cmd, " [ok]")
		}
	}
}

func TestX(t *testing.T) {
	file := "alpha\n"
	cmd := "X/a/"
	initDot := "1:1,1:6"
	expectedDot := "1:1,1:2;1:5,1:6"
	_, dot, err := Edit([]byte(file), cmd, initDot)
	if err != nil {
		t.Error("cmd:", cmd, " [failed]")
		t.Error(err)
	} else if dot != expectedDot {
		t.Error("cmd:", cmd, " [failed]")
		t.Errorf("expected: " + expectedDot + " but got: " + dot)
		//	} else {
		//		fmt.Println("cmd:", cmd, " [ok]")
	}
}
