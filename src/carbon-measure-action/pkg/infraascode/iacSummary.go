package infraascode

import "fmt"

func GetIACSummary(q TypIACQuery) []typSummary {
	var summary []typSummary
	switch q.Filetype {
	case "arm":
		// Summarize ARM JSON file to resource and location
		// TODO: Need to retrieve Variables and Parameters values inside resource type "Microsoft.Resources/deployments"
		summary = armSummary(q.Filename)
		break
	case "pulumi":
		summary = pulumiSummary(q.Filename, q.PreviewFileName)
	}

	// Print out summarized ARM data
	PrintSummary(&summary)
	return summary

}

func PrintSummary(summary *[]typSummary) {

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("*****************************************")
	fmt.Println("*** L I S T   O F   R E S O U R C E S ***")
	fmt.Println("*****************************************")
	fmt.Println()
	fmt.Println()
	fmt.Println()

	total := 0
	for _, s := range *summary {
		count := 0
		fmt.Println(s.resource)
		for _, sz := range s.sizes {
			fmt.Printf("- Size/SKU: %v\n", sz.size)
			for _, d := range sz.details {
				fmt.Printf("  - %v in %v\n", d.count, d.location)
				count += d.count
				total += d.count
			}
		}
		fmt.Printf("TOTAL: %v\n", count)
		fmt.Println()
	}
	fmt.Printf("Total number of resources: %v\n\n", total)
}

type TypIACQuery struct {
	Filetype        string
	Filename        string
	PreviewFileName string
}

type typSummary struct {
	resource string
	sizes    []typSizes
}

type typSizes struct {
	size    string
	details []typSummaryDetails
}

type typSummaryDetails struct {
	location string
	count    int
}
