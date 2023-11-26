package main

// import fyne

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// new app
	a := app.New()
	//New title and window
	w := a.NewWindow("Weather API & Fyne")
	// Resize
	w.Resize(fyne.NewSize(400, 300))
	//  Creat UI

	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=London,uk&APPID=20de98b797248a2bb700f4a3987fd684")
	if err != nil {
		fmt.Println("Error:", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}

	weather, err := UnmarshalForecast(body)
	if err != nil {
		fmt.Println("Error:", err)
	}

	img := canvas.NewImageFromFile("C:/Project/GO/logo.jpg")
	img.FillMode = canvas.ImageFillOriginal

	label1 := canvas.NewText("Weather Api & fyne", color.Black)
	label1.TextStyle = fyne.TextStyle{Bold: true}

	label_Country := widget.NewLabel(fmt.Sprintf("Country: %s", weather.Sys.Country))

	label_City := widget.NewLabel(fmt.Sprintf("City: %s", weather.Name))

	label_WindSpeed := widget.NewLabel(fmt.Sprintf("Wind speed: %.2f m/s", weather.Wind.Speed))

	TempCel := weather.Main.Temp - 273.15 // Convert Kelvin to Celsius
	label_MainTemp := widget.NewLabel(fmt.Sprintf("Temperature: %.2f °C", TempCel))

	w.SetContent(
		container.NewVBox(

			label1,
			label_Country,
			label_City,
			label_WindSpeed,
			label_MainTemp,
			img,
		),
	)

	w.ShowAndRun() // show and run app
}

// This part of file was generated from JSON using quicktype.io !

func UnmarshalForecast(data []byte) (Forecast, error) {
	var result Forecast
	err := json.Unmarshal(data, &result)
	return result, err
}

type Forecast struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `json:"timezone"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        int64     `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}
