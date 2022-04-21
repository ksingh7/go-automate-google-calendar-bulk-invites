package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gocarina/gocsv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type TalkDetails struct {
	TalkTitle     string `csv:"talk_title"`
	TalkStartTime string `csv:"talk_start_time"`
	TalkEndTime   string `csv:"talk_end_time"`
	Email         string `csv:"guest_email"`
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main() {

	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	// Read CSV file
	in, err := os.Open("talk_details.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()
	talkDetails := []*TalkDetails{}
	if err := gocsv.UnmarshalFile(in, &talkDetails); err != nil {
		panic(err)
	}
	for _, td := range talkDetails {
		fmt.Println(td)
		fmt.Println("Email, ", td.Email)
		fmt.Println("Talk Title, ", td.TalkTitle)
		fmt.Println("Talk Start Time, ", td.TalkStartTime)
		fmt.Println("Talk End Time, ", td.TalkEndTime)
		event := &calendar.Event{
			Summary:     "Your KCD BLR Talk QA : " + td.TalkTitle,
			Location:    "https://streamyard.com/dqhtmvzbkm",
			Description: "Dear Speaker,\nThanks for showing interest in KCD BLR 23 April Virtual Event. We have already received the recording of your talk, and we would be playing the recording on the allocated time slot.\nWith this calendar invite, we request you to join the event / your talk for QA at the end.\nUse the following link to join the streaming studio https://streamyard.com/dqhtmvzbkm\nThanks",
			Start: &calendar.EventDateTime{
				DateTime: td.TalkStartTime,
				TimeZone: "Asia/Kolkata",
			},
			End: &calendar.EventDateTime{
				DateTime: td.TalkEndTime,
				TimeZone: "Asia/Kolkata",
			},
			//Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
			Attendees: []*calendar.EventAttendee{
				&calendar.EventAttendee{Email: "karasing@redhat.com"},
				&calendar.EventAttendee{Email: td.Email},
			},
		}

		calendarId := "primary"
		event, err = srv.Events.Insert(calendarId, event).Do()
		if err != nil {
			log.Fatalf("Unable to create event. %v\n", err)
		}
		fmt.Printf("Event created: %s\n", event.HtmlLink)
	}

}
