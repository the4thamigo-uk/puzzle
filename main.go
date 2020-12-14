package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"unicode"
)

const (
	intro = `


                                                     *
  *                                                          *
                               *                  *        .--.
   \/ \/  \/  \/                                        ./   /=*
     \/     \/      *            *                ...  (_____)
      \ ^ ^/                                       \ \_((^o^))-.     *
      (o)(O)--)--------\.                           \   (   ) \  \._.
      |    |  ||================((~~~~~~~~~~~~~~~~~))|   ( )   |     \
       \__/             ,|        \. * * * * * * ./  (~~~~~~~~~~~)    \
*        ||^||\.____./|| |          \___________/     ~||~~~~|~'\____/ *
         || ||     || || A            ||    ||          ||    |   
  *      <> <>     <> <>          (___||____||_____)   ((~~~~~|   *
	`

	outro = `
                                  _
                               .-(_)
                              / _/
                           .-'   \
                          /       '.
                        ,-~--~-~-~-~-,
                       {__.._...__..._}             ,888,
       ,888,          /\##"  6  6  "##/\          ,88' ` + "`" + `88,
     ,88' '88,__     |(\` + "`" + `    (__)    ` + "`" + `/)|     __,88'     ` + "`" + `88
    ,88'   .8(_ \_____\_    '----'    _/_____/ _)8.       8'
    88    (___)\ \      '-.__    __.-'      / /(___)
    88    (___)88 |          '--'          | 88(___)
    8'      (__)88,___/                \___,88(__)
              __` + "`" + `88,_/__________________\_,88` + "`" + `__
             /    ` + "`" + `88,       |88|       ,88'    \
            /        ` + "`" + `88,    |88|    ,88'        \
           /____________` + "`" + `88,_\88/_,88` + "`" + `____________\
          /88888888888888888;8888;88888888888888888\
         /^^^^^^^^^^^^^^^^^^` + "`" + `/88\\^^^^^^^^^^^^^^^^^^\
   jgs  /                    |88| \============,     \
       /_  __  __  __   _ __ |88|_|^  MERRY    | _ ___\
       |;:.                  |88| | CHRISTMAS! |      |
       |;;:.                 |88| '============'      |
       |;;:.                 |88|                     |
       |::.                  |88|                     |
       |;;:'                 |88|                     |
       |:;,                  |88|                     |
       '---------------------""""---------------------'

------------------------------------------------
`
)

func main() {
	err := run()
	if err != nil {
		os.Exit(1)
	}
}

func run() error {
	if err := keyboard.Open(); err != nil {
		return err
	}
	defer func() {
		_ = keyboard.Close()
	}()

	m := map[rune]rune{
		'A': 'U',
		'B': 'J',
		'C': 'S',
		'D': 'G',
		'E': 'H',
		'F': 'R',
		'G': 'F',
		'H': 'B',
		'I': 'N',
		'J': 'M',
		'K': 'L',
		'L': 'A',
		'M': 'I',
		'N': 'W',
		'O': 'D',
		'P': 'Q',
		'Q': 'P',
		'R': 'Z',
		'S': 'E',
		'T': 'Y',
		'U': 'K',
		'V': 'X',
		'W': 'T',
		'X': 'V',
		'Y': 'O',
		'Z': ' ',
		' ': 'Z',
	}

	f := func(r rune) rune {
		r2, ok := m[r]
		if ok {
			return r2
		}
		return r
	}

	fmt.Println(intro)

	fmt.Printf("\n\nType the name of this tune played on the piano:\n\nEEE, EEE, EGCDE\n\n")

	for {

		fmt.Printf(">")

		var word string
		for {
			r, k, err := keyboard.GetKey()
			if err != nil {
				return err
			}

			r = unicode.ToUpper(r)
			if r >= 'A' && r <= 'Z' {
				r = f(r)
			} else if k == keyboard.KeySpace {
				r = f(' ')
			} else if k == keyboard.KeyEnter {
				fmt.Println("")
				fmt.Println("HAHA! WRONG!.... TRY AGAIN!")
				break
			} else if k == keyboard.KeyEsc {
				fmt.Println("HAHA! YOU GAVE UP!")
				return nil
			}

			word += string(r)

			if word == "JINGLE BELLS" {
				fmt.Println("")
				fmt.Println(outro)
				fmt.Println("WELL DONE")
				fmt.Println("now the next clue is....")
				return nil
			}

			fmt.Printf("%c", r)
		}
	}
}
