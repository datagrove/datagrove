package main

import (
	"os"
	"testing"
)

func Test_build(t *testing.T) {
	os.Args = []string{"one", "build", "example"}
	main()
}
func Test_build2(t *testing.T) {
	os.Args = []string{"one", "build", "D:/dev/asi/asi1/packages/setup/dbdeli"}
	main()
}
func Test_start2(t *testing.T) {
	os.Args = []string{"one", "start", "D:/dev/asi/asi1/packages/setup/dbdeli"}
	main()
}
func Test_start(t *testing.T) {
	os.Args = []string{"one", "start", "example"}
	main()
}

func Test_restore(t *testing.T) {
	os.Args = []string{"", "restore", "example", "v10", "0"}
	main()
}

// normally we will create a golden copy with some program, or we might restore it.
// this will test our ability to load a backup as a utility
func Test_restore_golden(t *testing.T) {
	os.Args = []string{"", "load", "example", "/var/opt/mssql/backup/v10.bak", "v10"}
	main()
}

func Test_drop(t *testing.T) {
	os.Args = []string{"", "drop", "example"}
	main()
}
