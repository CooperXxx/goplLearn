package display

import (
	"fmt"
	"reflect"
	"testing"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
	Directer        People
}
type xxx interface {
	sing()
}
type People struct {
	name      string
	age       int
	hobby     []string
	Abilities xxx
}

type Singer struct {
	kind string
}

func (s Singer) sing() {
	fmt.Println("sing a song")
}

func Test(t *testing.T) {

	//Display("slice",[]int{1,2,3})
	//Display("Struct", struct {
	//	name string
	//	nums []int
	//	maps map[string]int
	var singer xxx = Singer{"song"}
	//}{"xwd",[]int{1,2,3},map[string]int{"sss":1,"zzz":123413}})
	xwd := People{"xwd", 24, []string{"swimming", "working out"}, singer}

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
		Directer: xwd,
	}

	Display("strangelove", strangelove)

	rm := reflect.ValueOf(&strangelove)
	movie2 := rm.Elem().Addr().Interface().(*Movie)
	movie2.Year = 1999
	fmt.Println("---------------------------------------------------")
	Display("movie2", *movie2)
}
