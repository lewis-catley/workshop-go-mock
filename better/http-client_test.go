package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lewis-catley/workshop-go-mock/better/mocks"
	"github.com/stretchr/testify/assert"
)

type DummyReader struct{ Error error }

func (dr *DummyReader) Read([]byte) (int, error) {
	return 0, dr.Error
}

func TestGetContent(t *testing.T) {
	type getOut struct {
		r *http.Response
		e error
	}

	type expected struct {
		s string
		e error
	}

	tests := []struct {
		name   string
		input  string
		getOut getOut
		exp    expected
	}{
		{
			name:  "Test Success",
			input: "https://test.test",
			getOut: getOut{
				r: &http.Response{
					Body: ioutil.NopCloser(strings.NewReader("test output")),
				},
			},
			exp: expected{
				s: "test output",
			},
		},
		{
			name:  "Get Fails",
			input: "https://test.test",
			getOut: getOut{
				e: errors.New("Get failure is handled"),
			},
			exp: expected{
				e: errors.New("Get failure is handled"),
			},
		},
		{
			name:  "Read all fails (invalid response)",
			input: "https://test.test",
			getOut: getOut{
				r: &http.Response{
					Body: ioutil.NopCloser(&DummyReader{
						errors.New("Reader failed"),
					}),
				},
			},
			exp: expected{
				e: errors.New("Reader failed"),
			},
		},
	}

	for _, test := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		webMock := mocks.NewMockIHTTPClient(ctrl)

		webMock.EXPECT().Get(test.input).Times(1).Return(test.getOut.r, test.getOut.e)

		web := &WebClient{
			c: webMock,
		}

		outS, outE := web.GetContent(test.input)

		assert.Equal(t, test.exp.s, outS)

		assert.Equal(t, test.exp.e, outE)
	}
}
