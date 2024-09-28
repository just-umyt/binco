package start

import (
	"fmt"
	"os"
	"time"

	"github.com/just-umyt/just-game/internal/models"
)

type Det models.Details

// ⚪🟡🟢🔵🔴🟣🟤⚫

func Start(pro bool) {

	detail := Det{
		Colors: [16]string{"⬜", "🟨", "🟩", "🟦", "🟥", "🟪", "🟫", "⬛", "⬜", "🟨", "🟩", "🟦", "🟥", "🟪", "🟫", "⬛"},
		StartBalls: []string{
			"⚪", "⚪", "⚪", "⚪", "🟡", "🟡", "🟡", "🟡", "🟢", "🟢", "🟢", "🟢", "🔵", "🔵", "🔵", "🔵", "🔴", "🔴", "🔴", "🔴", "🟣", "🟣", "🟣", "🟣", "🟤", "🟤", "🟤", "🟤", "⚫", "⚫", "⚫", "⚫",
		},
		EndBalls: [2][16]string{
			{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		},
		CorrectBalls: [2][16]string{
			{"⚪", "🟡", "🟢", "🔵", "🔴", "🟣", "🟤", "⚫", "⚪", "🟡", "🟢", "🔵", "🔴", "🟣", "🟤", "⚫"},
			{"⚪", "🟡", "🟢", "🔵", "🔴", "🟣", "🟤", "⚫", "⚪", "🟡", "🟢", "🔵", "🔴", "🟣", "🟤", "⚫"},
		},
	}

	var gameOver bool = false
	timer := time.Now()

	for !gameOver {

		if len(detail.StartBalls) == 1 {
			gameOver = true
		}
		genBal := detail.GeneratorAndRemove()

		if pro {
			detail.ViewPro(genBal)
		} else {
			detail.View(genBal)
		}
		var input string
		fmt.Scan(&input)

		ind, err := ValidateInput(input)
		if err != nil {
			fmt.Println("GAME OVER! Only 0 and 1 numbers. Input length must be 4. Try again! ", time.Since(timer.Round(time.Second)))
			os.Exit(1)
		}

		if !detail.CheckProcess(ind, genBal) {
			fmt.Println("GAME OVER! Wrong color. Try again! ", time.Since(timer.Round(time.Second)))
			os.Exit(1)
		}
		ClearScreen()
	}

	detail.View("")
	fmt.Println("Congrats! ", time.Since(timer.Round(time.Second)))

}
