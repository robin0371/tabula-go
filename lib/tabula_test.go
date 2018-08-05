package tabula

import (
	"os"
	"path"
	"testing"
)

func TestGetCmdOptions(t *testing.T) {
	options := TabulaOptions{
		Area:    []string{"1.1", "2.2", "3.3", "4.4"},
		Batch:   "/home/user/test",
		Columns: []string{"0.1", "0.2", "0.3", "0.4"},
		Format:  "csv",
		Pages:   []string{"5"},
		Guess:   true,
		Lattice: true,
	}
	dir, _ := os.Getwd()
	expected_args := []string{
		"-jar", path.Join(dir, TabulaJar),
		"-a", "1.1,2.2,3.3,4.4",
		"-b", "/home/user/test",
		"-c", "0.1,0.2,0.3,0.4",
		"-f", "csv",
		"-p", "5",
		"-g",
		"-l",
	}

	args := GetCmdOptions(options)

	for i, v := range expected_args {
		if args[i] != v {
			t.Errorf("Expected %s, got %s", args[i], v)
		}
	}
}
