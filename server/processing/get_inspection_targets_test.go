package processing

import (
	"testing"
	"reflect"

	"github.com/ughvj/takaaki/types"
)

func Test_GeSubtractNEAERfromIHS_AllStudentsAreTarget(t *testing.T) {
	// Given.
	expected := &types.Students{
		types.Student{
			Id: 1,
			Name: "伊藤　博文",
			NyutaikunUserId: 100,
		},
		types.Student{
			Id: 2,
			Name: "黒田　清隆",
			NyutaikunUserId: 200,
		},
		types.Student{
			Id: 3,
			Name: "山縣　有朋",
			NyutaikunUserId: 300,
		},
	}

	// When.
	actual := SubtractNEAERfromIHS(
		types.GetStaticNyutaikunEntranceAndExitResult(),
		types.GetStaticInspectedHistoryCase0(),
		types.GetStaticStudents(),
	)

	// Then.
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Strings do not match. Expected: %v, Got: %v", expected, actual)
	}
}

func Test_GeSubtractNEAERfromIHS_2StudentsAreTarget(t *testing.T) {
	// Given.
	expected := &types.Students{
		types.Student{
			Id: 2,
			Name: "黒田　清隆",
			NyutaikunUserId: 200,
		},
		types.Student{
			Id: 3,
			Name: "山縣　有朋",
			NyutaikunUserId: 300,
		},
	}

	// When.
	actual := SubtractNEAERfromIHS(
		types.GetStaticNyutaikunEntranceAndExitResult(),
		types.GetStaticInspectedHistoryCase1(),
		types.GetStaticStudents(),
	)

	// Then.
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Strings do not match. Expected: %v, Got: %v", expected, actual)
	}
}

func Test_GeSubtractNEAERfromIHS_OneStudentIsTarget(t *testing.T) {
	// Given.
	expected := &types.Students{
		types.Student{
			Id: 1,
			Name: "伊藤　博文",
			NyutaikunUserId: 100,
		},
	}

	// When.
	actual := SubtractNEAERfromIHS(
		types.GetStaticNyutaikunEntranceAndExitResult(),
		types.GetStaticInspectedHistoryCase2(),
		types.GetStaticStudents(),
	)

	// Then.
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Strings do not match. Expected: %v, Got: %v", expected, actual)
	}
}
