package tabula

import (
	"log"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

const TabulaVersion = "1.0.2"
const TabulaJar = "tabula-" + TabulaVersion + "-jar-with-dependencies.jar"

type TabulaOptions struct {
	Area    []string // Portion of the page to analyze. Accepts top, left, bottom, right. Example: ["269.875" "12.75" "790.5" "561"]
	Batch   string   // Convert all .pdfs in the provided directory
	Columns []string // Coordinates of column boundaries. Example: ["10.1" "20.2" "30.3"]
	Format  string   // Output format: (CSV,TSV,JSON). Default: CSV
	Pages   []string // Pages ["5", "6"]
	Guess   bool     // Guess the portion of the page to analyze per page. Default true
	Lattice bool     // Force PDF to be extracted using lattice-mode extraction. Default false
	Path    string   // Path to target file
}

// Returns prepared array of command line options to run tabula
func GetCmdOptions(options TabulaOptions) []string {
	_, filename, _, _ := runtime.Caller(1)
	args := []string{"-jar", path.Join(path.Dir(filename), TabulaJar)}

	if len(options.Area) > 0 {
		area := strings.Join(options.Area, ",")
		args = append(args, "-a", area)
	}
	if len(options.Batch) > 0 {
		args = append(args, "-b", options.Batch)
	}
	if len(options.Columns) > 0 {
		columns := strings.Join(options.Columns, ",")
		args = append(args, "-c", columns)
	}
	if len(options.Format) > 0 {
		args = append(args, "-f", options.Format)
	}
	if len(options.Pages) > 0 {
		pages := strings.Join(options.Pages, ",")
		args = append(args, "-p", pages)
	}
	if options.Guess {
		args = append(args, "-g")
	}
	if options.Lattice {
		args = append(args, "-l")
	}

	if len(options.Path) > 0 {
		args = append(args, options.Path)
	}

	return args
}

// Executes command for extract data using tabula
// by tabula options as struct
func ExtractTableData(args TabulaOptions) string {
	options := GetCmdOptions(args)
	cmd := exec.Command("java", options...)
	log.Printf("Command: %s\n", strings.Join(cmd.Args, " "))

	output, err := cmd.Output()

	if err != nil {
		log.Printf("Error: %s\nResult: %s", err.Error(), output)
	} else {
		log.Printf("Success\nResult: %s", output)
	}

	return string(output)
}
