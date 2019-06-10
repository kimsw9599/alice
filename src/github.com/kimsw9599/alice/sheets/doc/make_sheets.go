package doc

import (
	"io/ioutil"
	"log"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"github.com/kimsw9599/alice/sheets/utils"
)

func Make_sheets(sheets_name string) (string, error) {
	b, err := ioutil.ReadFile("./auth_files/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return "", err
	}

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/drive")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
		return "", err
	}
	client := utils.GetClient(config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
		return "", err
	}

	sheetTitle := sheets.SpreadsheetProperties{Title:sheets_name}
	spreadsheet := sheets.Spreadsheet{}
	spreadsheet.Properties = &sheetTitle
	s, err := srv.Spreadsheets.Create(&spreadsheet).Do()

	if err!= nil{
		log.Fatalf("Create sheets error: %v", err)
		return "", err
	}
	print(s.SpreadsheetId)

	return s.SpreadsheetId, nil
}
