package main

import (
	"testing"

	"github.com/naltun/eyes/pkg/httpheader"
)

func testHttpHeader(t *testing.T) {
	testTarget := "https://github.com"
	_, err := httpheader.Httpheader(testTarget)
	if err != nil {
		t.Errorf("httpheader.Httpheader() failed:\n %s", err)
	}
}
