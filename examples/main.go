package main

import (
	"github.com/nilpntr/gogeojson"
	"log"
)

func main() {
	ls := gogeojson.NewLineStringField([][]float64{
		{
			40,
			5,
		},
		{
			41,
			6,
		},
	})
	if err := ls.Validate(); err != nil {
		log.Fatal(err)
	}

	pg := gogeojson.NewPolygonField([][][]float64{
		{
			{
				0,
				0,
			},
			{
				3,
				6,
			},
			{
				6,
				1,
			},
			{
				0,
				0,
			},
		},
	})

	if err := pg.Validate(); err != nil {
		log.Fatal(err)
	}
}