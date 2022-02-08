package main

import (
	welcome "github.com/webvillain/thegamquiz/Welcome"
	"github.com/webvillain/thegamquiz/problems"
	"github.com/webvillain/thegamquiz/user"
)

func main() {
	user.User_Info()
	welcome.IntroAboutQuiz()
	problems.StartQuiz()
	problems.TimeTaken()
	problems.QuizStatics()
	problems.FinalResult()
}
