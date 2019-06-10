package sheets

//import (
//	"io/ioutil"
//	"log"
//	"golang.org/x/oauth2/google"
//	"github.com/ksw9599/alice/sheets/utils"
//	"google.golang.org/api/sheets/v4"
//)
//
//func Read_sheets(spreadsheetId string) (*ValueRange, error) {
//	b, err := ioutil.ReadFile("./auth_files/credentials.json")
//	if err != nil {
//		log.Fatalf("Unable to read client secret file: %v", err)
//	}
//
//	// If modifying these scopes, delete your previously saved token.json.
//	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
//	if err != nil {
//		log.Fatalf("Unable to parse client secret file to config: %v", err)
//	}
//	client := utils.GetClient(config)
//
//	srv, err := sheets.New(client)
//	if err != nil {
//		log.Fatalf("Unable to retrieve Sheets client: %v", err)
//	}
//
//	// Prints the names and majors of students in a sample spreadsheet:
//	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
//	//spreadsheetId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
//	readRange := "Class Data!A2:E"
//	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
//	if err != nil {
//		log.Fatalf("Unable to retrieve data from sheet: %v", err)
//	}
//
//	return resp, err
//}
