package start

import "fmt"

func (det *Det) View(ball string) {
	//scheme
	fmt.Println("                       ", ball)
	fmt.Println("             __________0|1__________")
	fmt.Println("       ____0|1____             ____0|1____")
	fmt.Println("    _0|1_       _0|1_       _0|1_       _0|1_")
	fmt.Println("  0|1   0|1   0|1   0|1   0|1   0|1   0|1   0|1")

	//bottom balls place
	for _, balls := range det.EndBalls {
		fmt.Print("|")

		for _, ball := range balls {
			if ball == " " {
				fmt.Print(ball, " |")
			} else {
				fmt.Print(ball, "|")

			}

		}
		fmt.Println()
	}

	//colors
	fmt.Println(det.Colors)
	for i := 0; i < 16; i++ {
		if i > 9 {
			fmt.Printf(" %d", i)

		} else {
			fmt.Printf("  %d", i)
		}
	}
	fmt.Println()
}

func (det *Det) ViewPro(ball string) {
	//scheme
	fmt.Println("	 ____                 _____")
	fmt.Println(" 	|  _ \\  ", ball, "         / ____| ")
	fmt.Println(" 	| |_) |  _   _ __   | |        ___")
	fmt.Println(" 	|  _ <  | | | '_ \\  | |       / _ \\ ")
	fmt.Println(" 	| |_) | | | | | | | | |____  | (_) |")
	fmt.Println(" 	|____/  |_| |_| |_|  \\_____|  \\___/ ")

	//bottom balls place
	for _, balls := range det.EndBalls {
		fmt.Print("|")

		for _, ball := range balls {
			if ball == " " {
				fmt.Print(ball, " |")
			} else {
				fmt.Print(ball, "|")

			}

		}
		fmt.Println()
	}

	//colors
	fmt.Println(det.Colors)
	for i := 0; i < 16; i++ {
		if i > 9 {
			fmt.Printf(" %d", i)

		} else {
			fmt.Printf("  %d", i)
		}
	}
	fmt.Println()
}
