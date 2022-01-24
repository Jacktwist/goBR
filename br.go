package main

import (
	"fmt"
	"github.com/mpatraw/gopherlibterminal"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"time"
)

var SCREENWIDTH int = 80
var SCREENHEIGHT int = 25

func main() {
	rand.Seed(time.Now().UnixNano())
	// Start a gopherlibterminal window
	// Documentation is slightly adapted from source docs in http://foo.wyrd.name/en:bearlibterminal
	// Ex. keystroke TkQ in blt is TK_Q in glb
	glt.Open()
	defer glt.Close()
	glt.Set("window: size=" + strconv.Itoa(SCREENWIDTH) + "x" + strconv.Itoa(SCREENHEIGHT) + ";font: /usr/share/fonts/truetype/Orbitron/Orbitron-Regular.ttf, size=12;input: mousecursor=true")

	var rl Roller
	rl.roll("4d6k3")
	glt.Print(0,0, rl.String())
	fmt.Printf("Roller: %v %v\n", rl.rolls, rl.result)

	 glt.Refresh()

	keyPressed := glt.Read()
	for keyPressed != glt.TkQ {
		keyPressed = glt.Read()
	}
} // end of main function

	// General rectangle structure
type Rect struct {
	// Variables for the top left and bottom right corner coordinates
	x1, x2, y1, y2 int
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

type Roller struct {
	rolls []int
	result int
}

func (r Roller)roll(s string){
	rx := regexp.MustCompile("^[1-9]+[0-9]*d[1-9]+[0-9]*(k[1-9]+[0-9]*)?")
	fmt.Printf("Roll command: %v\n", s)
	byteRoll := []byte(s)

	var num, sides, keep []byte
	mode := "num"
	rolls := make([]int, 0)
	sum := 0

	// Make sure the string matches the regexp then strip out all the numbers into separate string variables

	if rx.MatchString(string(byteRoll)) {
		for i := range byteRoll {

			if mode == "keep" {
				keep = append(keep, byteRoll[i])
			}
			if mode == "sides" {
				if byteRoll[i] == 'k' {
					mode = "keep"
				}else{
			 	sides = append(sides, byteRoll[i])
				}
			}
			if mode == "num" {
				if byteRoll[i] == 'd' {
					mode = "sides"
				}else{
				num = append(num, byteRoll[i])
				}
			}
		}
	} else {
		fmt.Println("Regex pattern not matched.")
	}

	// Make our string numbers variables
	numN, _	 := strconv.Atoi(string(num))
	sidesN, _ := strconv.Atoi(string(sides))
	keepN, _	 := strconv.Atoi(string(keep))


	for i := 0; i < numN ; i++ {
		rolls = append(rolls, rand.Intn(sidesN)+1)
	}
	fmt.Printf("Roll results: %v\n", rolls)

	// sort the results and trim the lowest results if needed
	sort.Ints(rolls)
	fmt.Printf("Sorted results: %v\n", rolls)
	if keepN > 0 && len(rolls)-1-keepN > 0 {
	rolls = rolls[len(rolls)-keepN:len(rolls)-1]
	}
	fmt.Printf("Trimmed results: %v\n", rolls)
	fmt.Printf("%v %v %v %v\n", r.rolls, r.result, rolls, sum)
	copy(r.rolls, rolls)
	r.result = sum
	fmt.Printf("%v %v %v %v\n", r.rolls, r.result, rolls, sum)
}

func (r Roller) String() string {
		var astring string
	for i := 0;i < len(r.rolls);i++ {
			astring += strconv.Itoa(r.rolls[i])
		}
		return fmt.Sprintf(" rolled %v which sums %v" , r.rolls, r.result)
}
// Game objects
