package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestStorage_FindUser(t *testing.T) {
	testCases := map[string]struct {
		users        []User
		id           int
		expectedUser User
		shouldFail   bool
	}{
		"UserExists": {
			users: []User{
				{ID: 1, Name: "Jone"},
				{ID: 2, Name: "McVitae"},
				{ID: 3, Name: "Charles"},
			},
			id:           3,
			expectedUser: User{ID: 3, Name: "Charles"},
			shouldFail:   false,
		},
		"UserNotExists": {
			users: []User{
				{ID: 2, Name: "McViate"},
			},
			id:           1,
			expectedUser: User{},
			shouldFail:   true,
		},
	}

	for name, test := range testCases {
		testHeader := fmt.Sprintf("### TestCase : %s ###", name)
		t.Run(testHeader, func(t *testing.T) {
			storage := NewStorage(test.users)

			user, err := storage.FindUser(test.id)

			if test.shouldFail != (err != nil) {
				t.Fatalf("error expectancy doesn't match: shoudlFail is %v, but error != nil is %v", test.shouldFail, (err != nil))
			}

			if diff := cmp.Diff(test.expectedUser, user); diff != "" {
				t.Errorf("user expectancy doesn't match %v", diff)
			}
		})
	}
}
