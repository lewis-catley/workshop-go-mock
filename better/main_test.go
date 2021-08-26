package main

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lewis-catley/workshop-go-mock/better/mocks"
)

func Test_printUrlContent(t *testing.T) {
	type getReturn struct {
		r string
		e error
	}
	tests := []struct {
		name        string
		getRet      getReturn
		shouldPanic bool
	}{
		{
			name: "Should NOT panic when GetContent is successful",
			getRet: getReturn{
				r: "test response",
			},
		},
		{
			name: "Should panic when GetContent errors",
			getRet: getReturn{
				r: "",
				e: errors.New("foo"),
			},
			shouldPanic: true,
		},
	}

	for _, test := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockWeb := mocks.NewMockIWebClient(ctrl)
		mockWeb.EXPECT().GetContent("foo.bar").Times(1).Return(test.getRet.r, test.getRet.e)

		a := application{
			web: mockWeb,
		}

		defer func() {
			if r := recover(); r == nil && !test.shouldPanic {
				t.Errorf("The code did not panic")
			}
		}()

		t.Logf("Test Name: %s->%s", t.Name(), test.name)
		a.printUrlContent("foo.bar")
	}
}
