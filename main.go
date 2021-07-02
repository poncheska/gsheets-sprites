package main

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

type SpritesService struct {
	SpreadsheetID string
	SheetID       int64
	X0            int64
	Y0            int64
	Services      []*sheets.Service
	Sprites       [][][]byte
	Height        int
	Width         int
}

func main() {
	x0, y0 := int64(0), int64(0)
	spreadsheetId := "1ubv_Q3bKTrUmFJ8o5zrdpBrEUNc8gix1ZC6_np7fQs8"
	sheetId := int64(0)
	configPaths := []string{
		"credential.json",
		"credentials2.json",
		"credentials3.json",
		"credentials4.json",
	}
	srvs := []*sheets.Service{}
	for _, v := range configPaths {
		client, err := ServiceAccount(v)
		if err != nil {
			panic(err)
		}

		srv, err := sheets.New(client)
		if err != nil {
			panic(err)
		}
		srvs = append(srvs, srv)
	}
	sprites := SpritesService{
		SpreadsheetID: spreadsheetId,
		SheetID:       sheetId,
		X0:            x0,
		Y0:            y0,
		Services:      srvs,
		Sprites:       SpritesSample,
		Height:        SampleHeight,
		Width:         SampleWidth,
	}
	sprites.StartAnimation(SampleColors)
}

func (s SpritesService) StartAnimation(clrs map[byte]Color) {
	spriteNum := 0
	cur := 0
	cd := time.Second / time.Duration(len(s.Services)*20)
	for {
		rb := &sheets.BatchUpdateSpreadsheetRequest{}
		for i := 0; i < s.Height; i++ {
			for j := 0; j < s.Width; j++ {
				rb.Requests = append(rb.Requests, &sheets.Request{
					RepeatCell: &sheets.RepeatCellRequest{
						Fields: "*",
						Range: &sheets.GridRange{
							SheetId:          s.SheetID,
							StartColumnIndex: s.X0 + int64(j),
							EndColumnIndex:   s.X0 + int64(j+1),
							StartRowIndex:    s.Y0 + int64(i),
							EndRowIndex:      s.Y0 + int64(i+1),
						},
						Cell: &sheets.CellData{
							UserEnteredFormat: &sheets.CellFormat{
								BackgroundColor: &sheets.Color{
									Alpha: 1.0,
									Red:   clrs[s.Sprites[spriteNum][i][j]].Red,
									Green: clrs[s.Sprites[spriteNum][i][j]].Green,
									Blue:  clrs[s.Sprites[spriteNum][i][j]].Blue,
								},
							},
						},
					},
				})
			}
		}
		srv := s.Services[cur]
		cur = (cur + 1) % len(s.Services)
		spriteNum = (spriteNum + 1) % len(s.Sprites)
			_, err := srv.Spreadsheets.BatchUpdate(s.SpreadsheetID, rb).Context(context.Background()).Do()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(cd)
	}
}

func ServiceAccount(credentialFile string) (*http.Client, error) {
	b, err := ioutil.ReadFile(credentialFile)
	if err != nil {
		return nil, err
	}
	var c = struct {
		Email      string `json:"client_email"`
		PrivateKey string `json:"private_key"`
	}{}
	json.Unmarshal(b, &c)
	config := &jwt.Config{
		Email:      c.Email,
		PrivateKey: []byte(c.PrivateKey),
		Scopes: []string{
			docs.DocumentsScope,
			sheets.SpreadsheetsScope,
		},
		TokenURL: google.JWTTokenURL,
	}
	client := config.Client(oauth2.NoContext)
	return client, nil
}
