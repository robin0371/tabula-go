package tabula

import (
	"log"
	"os/exec"
	"strings"
)

const tabula_version = "1.0.2"
const tabula_jar = "tabula-" + tabula_version + "-jar-with-dependencies.jar"

type TabulaOptions struct {
	Area    []string // Portion of the page to analyze. Accepts top, left, bottom, right. Example: ["269.875" "12.75" "790.5" "561"]
	Batch   string   // Convert all .pdfs in the provided directory
	Columns []string // Coordinates of column boundaries. Example: ["10.1" "20.2" "30.3"]
	Format  string   // Output format: (CSV,TSV,JSON). Default: CSV
	Pages   []string // Pages ["5", "6"]
	Guess   bool     // Guess the portion of the page to analyze per page. Default true
	Lattice bool     // Force PDF to be extracted using lattice-mode extraction. Default false
}

// Returns prepared array of command line options to run tabula
func GetCmdOptions(options TabulaOptions) []string {
	args := []string{"-jar", tabula_jar}

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

	return args
}

// Executes command for extract data using tabula
func ExtractTableData(options TabulaOptions) {
	args := GetCmdOptions(options)
	cmd := exec.Command("java", args...)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
