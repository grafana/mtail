// +build integration

package mtail_test

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"

	"github.com/google/mtail/internal/mtail"
)

func TestBadProgramFailsCompilation(t *testing.T) {
	progDir, rmProgDir := mtail.TestTempDir(t)
	defer rmProgDir()
	logDir, rmLogDir := mtail.TestTempDir(t)
	defer rmLogDir()

	err := ioutil.WriteFile(path.Join(progDir, "bad.mtail"), []byte("asdfasdf\n"), 0666)
	if err != nil {
		t.Fatal(err)
	}

	// Compile-only fails program compilation at server start, not after it's running.
	_, err = mtail.TestMakeServer(t, 0, false, mtail.ProgramPath(progDir), mtail.LogPathPatterns(logDir), mtail.CompileOnly)
	if err == nil {
		t.Error("expected error from mtail")
	}
	if !strings.Contains(err.Error(), "compile failed") {
		t.Error("compile failed not reported")
	}
}
