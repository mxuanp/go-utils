package fmt_test

import (
	"github.com/mxuanp/go-utils/xfmt"
	"testing"
)

func TestPrintMsg(t *testing.T) {
	xfmt.PrintMsg("hello", xfmt.Red)
}
