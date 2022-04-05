package infraascode

import "fmt"

func GetIACSummary(q TypIACQuery) []TypSummary {
	var summary []TypSummary
	switch q.Filetype {
	case "arm":
		// Summarize ARM JSON file to resource and location
		// TODO: Need to retrieve Variables and Parameters values inside resource type "Microsoft.Resources/deployments"
		summary = armSummary(q.Filename)
	}
	// Print out summarized ARM data
	PrintSummary(&summary)
	return summary

}

func PrintSummary(summary *[]TypSummary) {

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
		fmt.Println(s.Resource)
		for _, sz := range s.Sizes {
			fmt.Printf("- Size/SKU: %v\n", sz.Size)
			for _, d := range sz.Details {
				fmt.Printf("  - %v in %v\n", d.Count, d.Location)
				count += d.Count
				total += d.Count
			}
		}
		fmt.Printf("TOTAL: %v\n", count)
		fmt.Println()
	}
	fmt.Printf("Total number of resources: %v\n\n", total)
}

type TypIACQuery struct {
	Filetype string
	Filename string
}

type TypSummary struct {
	Resource string
	Sizes    []TypSizes
}

type TypSizes struct {
	Size    string
	Details []TypSummaryDetails
}

type TypSummaryDetails struct {
	Location string
	Count    int
}
