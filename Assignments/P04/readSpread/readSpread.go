package readSpread

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func read() {
	ctx := context.Background()

	b, err := ioutil.ReadFile("windy-art-405216-8dbb2cdfbd2d.json") // Update the path to your JSON key file
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// Configure the Google Sheets API client
	config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := config.Client(ctx)

	sheetService, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Specify the Spreadsheet ID and Range
	spreadsheetId := "14QSmG84jRhEgzXb96Ie9tw9zS7cY4PGXX9NGys-kc-o" // Replace with your Spreadsheet ID
	readRange := "Sheet1!A42:C46"                                   // Adjust the range accordingly

	// Read the data
	resp, err := sheetService.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			fmt.Printf("%s, %s\n", row[0], row[2])
		}
	}
}
