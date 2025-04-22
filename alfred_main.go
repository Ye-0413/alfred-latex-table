package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/crispgm/alfred-latex-table/internal/namesgenerator"

	aw "github.com/deanishe/awgo"
)

var (
	wf *aw.Workflow
)

func run() {
	var (
		colNum   int = 2
		rowNum   int = 3
		testData bool
	)
	argc := len(os.Args)
	if argc == 1 {
		return
	}
	if argc > 1 {
		params := os.Args[1]
		realArgs := strings.Split(params, " ")
		count := 0
		for _, arg := range realArgs {
			if arg == "" || arg == " " {
				continue
			}
			val, err := strconv.Atoi(arg)
			if err == nil {
				if count == 0 {
					colNum = val
				} else if count == 1 {
					rowNum = val
				}
			} else if count == 2 && (arg == "data" || arg == "test") {
				testData = true
			} else {
				break
			}
			count++
		}
	}
	wf.NewItem(getSubtitle(colNum, rowNum, testData)).
		Arg(buildLatexTable(colNum, rowNum, testData)).
		Valid(true)
	wf.SendFeedback()
	return
}

func getSubtitle(colNum, rowNum int, testData bool) string {
	subtitle := fmt.Sprintf("Generate a %dx%d LaTeX table", colNum, rowNum)
	if testData {
		subtitle += " with test data"
	}
	return subtitle
}

func buildLatexTable(colNum, rowNum int, testData bool) string {
	var sb strings.Builder

	// Begin table environment
	sb.WriteString("\\begin{table}[htbp]\n")
	sb.WriteString("\\centering\n")

	// Begin tabular environment with column specifications
	sb.WriteString("\\begin{tabular}{")
	for i := 0; i < colNum; i++ {
		sb.WriteString("|c")
	}
	sb.WriteString("|}\n")
	sb.WriteString("\\hline\n") // Top line

	// Generate header row
	headerRow := generateRow(colNum, testData)
	for i, cell := range headerRow {
		sb.WriteString(cell)
		if i < colNum-1 {
			sb.WriteString(" & ")
		} else {
			sb.WriteString(" \\\\\n")
			sb.WriteString("\\hline\n") // Line after header
		}
	}

	// Generate data rows
	for i := 0; i < rowNum; i++ {
		row := generateRow(colNum, testData)
		for j, cell := range row {
			sb.WriteString(cell)
			if j < colNum-1 {
				sb.WriteString(" & ")
			} else {
				sb.WriteString(" \\\\")
				// Only add \hline after the last row
				if i == rowNum-1 {
					sb.WriteString("\n\\hline")
				}
				sb.WriteString("\n")
			}
		}
	}

	// End environments
	sb.WriteString("\\end{tabular}\n")
	sb.WriteString("\\caption{}\n")
	sb.WriteString("\\label{tab:}\n")
	sb.WriteString("\\end{table}")

	return sb.String()
}

func generateRow(colNum int, testData bool) []string {
	row := make([]string, colNum)
	for index := range row {
		if testData {
			row[index] = namesgenerator.GetRandomName(0)
		} else {
			row[index] = " "
		}
	}
	return row
}

func main() {
	wf = aw.New()
	wf.Run(run)
}
