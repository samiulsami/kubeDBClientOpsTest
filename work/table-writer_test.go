package work

import (
	"os"
	"testing"
	"time"

	"github.com/olekukonko/tablewriter"
)

func Test_table_writer(t *testing.T) {
	data := [][]string{
		{"1/1/2014", "Domain name", "1234", "$10.98"},
		{"1/1/2014", "January Hosting", "1234", "$54.95"},
		{"1/4/2014", "February Hosting", "3456", "$51.00"},
		{"1/4/2014", "February Extra Bandwidth", "4567", "$30.00"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
	table.SetFooter([]string{"", "", "Total", "$146.93"})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}

type DiagnosticResultTest struct {
	CheckType string
	Timestamp string
	Outputs   []DiagnosticOutputTest
}

type DiagnosticOutputTest struct {
	Message string
	Data    string
}

func Test_table_writer1(t *testing.T) {
	data := []DiagnosticResultTest{
		{
			CheckType: "InspectLogs",
			Timestamp: time.Now().String(),
			Outputs: []DiagnosticOutputTest{
				{Message: "Log 1", Data: "Lorem Ipsum is simply dummy text of the printing and typesetting industry."},
				{Message: "Log 2", Data: "Another diagnostic message here."},
			},
		},
		{
			CheckType: "HealthCheck",
			Timestamp: time.Now().String(),
			Outputs: []DiagnosticOutputTest{
				{Message: "Log 3", Data: "Some health check data here."},
				{Message: "Log 4", Data: "Additional health data with more details."},
			},
		},
	}

	// Set up table writer for the main table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"CheckType", "Timestamp", "Description", "Content"})
	table.SetAutoMergeCells(true) // This will allow merging cells in the "CheckType" and "Timestamp" columns
	table.SetRowLine(true)

	// Populate the table with DiagnosticResult data
	for _, result := range data {
		checkType := ""
		timestamp := ""
		for i, output := range result.Outputs {
			// Format each row, only showing CheckType and Timestamp for the first Output entry

			if i == 0 {
				checkType = result.CheckType
				timestamp = result.Timestamp
			}
			outputText := output.Message
			outputText1 := output.Data
			table.Append([]string{checkType, timestamp, outputText, outputText1})
		}
	}

	// Render the main table to the console
	table.Render()
}
