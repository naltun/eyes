package httpheader

import (
	"github.com/naltun/eyes/pkg/httpheader"
	"testing"
)

func testHttpHeader(t *testing.T) {
	testTarget1 := "https://github.com"
	testTarget2 := "https://www.cnet.com/"
	_, err := httpheader.Httpheader(testTarget1)
	if err != nil {
		t.Errorf("httpheader.Httpheader() failed:\n%s", err)
	}
	_, err := httpheader.Httpheader(testTarget2)
	if err != nil {
		t.Errorf("httpheader.Httpheader() failed:\n%s", err)
	}
}
