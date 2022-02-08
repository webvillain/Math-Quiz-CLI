package problems

import (
	"fmt"
	"time"

	"github.com/webvillain/thegamquiz/user"
)

var QuizStartingTime time.Time
var QuizEndingTime time.Time
var score int
var AA = [10]int{}                                            // Answer By User
var RA = [10]int{20, 110, 50, 25, 10000, 25, 17, 121, 300, 1} // Real Anser
var ra, wa int

func StartQuiz() {
	QuizStartingTime = time.Now()
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(" 1. 10 + 10 = ?")
	fmt.Scanf("%d", &AA[0])
	if AA[0] == RA[0] {
		fmt.Println("Yeah , Right Answer . You Got 1 Point")
		score = score + 1
		// fmt.Println("Your Total Score Will Be Added By 1")
		// fmt.Println("Your Total Score After This Attempt Is : ", score)
		ra = ra + 1
	} else {
		fmt.Println("Ooopppssss . . . . Wrong Answer , The Right Answer Is :", RA[0])
		score = score - 1
		// fmt.Println("Your Total Score Will Be Deduct By 1")
		// fmt.Println("Your Total Score After This Attempt : ", score)
		wa = wa + 1
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(" 2. 100 + 10 = ?")
	fmt.Scanf("%d", &AA[1])
	if AA[1] == RA[1] {
		fmt.Println("Yeah , Right Answer . You Got 1 Point")
		score = score + 1
		// fmt.Println("Your Total Score Will Be Added By 1")
		// fmt.Println("Your Total Score After This Attempt Is : ", score)
		ra = ra + 1
	} else {
		fmt.Println("Ooopppssss . . . . Wrong Answer , The Right Answer Is :", RA[1])
		score = score - 1
		// fmt.Println("Your Total Score Will Be Deduct By 1")
		// fmt.Println("Your Total Score After This Attempt : ", score)
		wa = wa + 1
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(" 3. 25 + 25 = ?")
	fmt.Scanf("%d", &AA[2])
	if AA[2] == RA[2] {
		fmt.Println("Yeah , Right Answer . You Got 1 Point")
		score = score + 1
		// fmt.Println("Your Total Score Will Be Added By 1")
		// fmt.Println("Your Total Score After This Attempt Is : ", score)
		ra = ra + 1
	} else {
		fmt.Println("Ooopppssss . . . . Wrong Answer , The Right Answer Is :", RA[2])
		score = score - 1
		// fmt.Println("Your Total Score Will Be Deduct By 1")
		// fmt.Println("Your Total Score After This Attempt : ", score)
		wa = wa + 1
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(" 4. 5 * 5 = ?")
	fmt.Scanf("%d", &AA[3])
	if AA[3] == RA[3] {
		fmt.Println("Yeah , Right Answer . You Got 1 Point")
		score = score + 1
		// fmt.Println("Your Total Score Will Be Added By 1")
		// fmt.Println("Your Total Score After This Attempt Is : ", score)
		ra = ra + 1
	} else {
		fmt.Println("Ooopppssss . . . . Wrong Answer , The Right Answer Is :", RA[3])
		score = score - 1
		// fmt.Println("Your Total Score Will Be Deduct By 1")
		// fmt.Println("Your Total Score After This Attempt : ", score)
		wa = wa + 1
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(" 5. 100 * 100  = ?")
	fmt.Scanf("%d", &AA[4])
	if AA[4] == RA[4] {
		fmt.Println("Yeah , Right Answer . You Got 1 Point")
		score = score + 1
		// fmt.Println("Your Total Score Will Be Added By 1")
		// fmt.Println("Your Total Score After This Attempt Is : ", score)
		ra = ra + 1
	} else {
		fmt.Println("Ooopppssss . . . . Wrong Answer , The Right Answer Is :", RA[4])
		score = score - 1
		// fmt.Println("Your Total Score Will Be Deduct By 1")
		// fmt.Println("Your Total Score After This Attempt : ", score)
		wa = wa + 1
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(" 6. 28 - 3 = ? ")
	fmt.Scanf("%d", &AA[5])
	if AA[5] == RA[5] {
		fmt.Println("Yeah , Right Answer . You Got 1 Point")
		score = score + 1
		// fmt.Println("Your Total Score Will Be Added By 1")
		// fmt.Println("Your Total Score After This Attempt Is : ", score)
		ra = ra + 1
	} else {
		fmt.Println("Ooopppssss . . . . Wrong Answer , The Right Answer Is :", RA[5])
		score = score - 1
		// fmt.Println("Your Total Score Will Be Deduct By 1")
		// fmt.Println("Your Total Score After This Attempt : ", score)
		wa = wa + 1
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(" 7. 102 / 6 = ?")
	fmt.Scanf("%d", &AA[6])
	if AA[6] == RA[6] {
		fmt.Println("Yeah , Right Answer . You Got 1 Point")
		score = score + 1
		// fmt.Println("Your Total Score Will Be Added By 1")
		// fmt.Println("Your Total Score After This Attempt Is : ", score)
		ra = ra + 1
	} else {
		fmt.Println("Ooopppssss . . . . Wrong Answer , The Right Answer Is :", RA[6])
		score = score - 1
		// fmt.Println("Your Total Score Will Be Deduct By 1")
		// fmt.Println("Your Total Score After This Attempt : ", score)
		wa = wa + 1
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(" 8. 11 * 11 = ?")
	fmt.Scanf("%d", &AA[7])
	if AA[7] == RA[7] {
		fmt.Println("Yeah , Right Answer . You Got 1 Point")
		score = score + 1
		// fmt.Println("Your Total Score Will Be Added By 1")
		// fmt.Println("Your Total Score After This Attempt Is : ", score)
		ra = ra + 1
	} else {
		fmt.Println("Ooopppssss . . . . Wrong Answer , The Right Answer Is :", RA[7])
		score = score - 1
		// fmt.Println("Your Total Score Will Be Deduct By 1")
		// fmt.Println("Your Total Score After This Attempt : ", score)
		wa = wa + 1
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(" 9. 100 - 25 + 75 * 2 = ?")
	fmt.Scanf("%d", &AA[8])
	if AA[8] == RA[8] {
		fmt.Println("Yeah , Right Answer . You Got 1 Point")
		score = score + 1
		// fmt.Println("Your Total Score Will Be Added By 1")
		// fmt.Println("Your Total Score After This Attempt Is : ", score)
		ra = ra + 1
	} else {
		fmt.Println("Ooopppssss . . . . Wrong Answer , The Right Answer Is :", RA[8])
		score = score - 1
		// fmt.Println("Your Total Score Will Be Deduct By 1")
		// fmt.Println("Your Total Score After This Attempt : ", score)
		wa = wa + 1
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Println(" 10. 555 % 2 = ?")
	fmt.Scanf("%d", &AA[9])
	if AA[9] == RA[9] {
		fmt.Println("Yeah , Right Answer . You Got 1 Point")
		score = score + 1
		// fmt.Println("Your Total Score Will Be Added By 1")
		// fmt.Println("Your Total Score After This Attempt Is : ", score)
		ra = ra + 1
	} else {
		fmt.Println("Ooopppssss . . . . Wrong Answer , The Right Answer Is :", RA[9])
		score = score - 1
		// fmt.Println("Your Total Score Will Be Deduct By 1")
		// fmt.Println("Your Total Score After This Attempt : ", score)
		wa = wa + 1
	}
	fmt.Println("---------------------------------------------------------------")
	QuizEndingTime = time.Now()

}
func TimeTaken() {
	fmt.Printf("Quiz Starting Time Is %v\n:", QuizStartingTime)
	fmt.Printf("Quiz Ending Time Is : %v\n", QuizEndingTime)
}

func QuizStatics() {
	fmt.Println("\t\t\t\t\t\t_______________________________________________")
	fmt.Println("\t\t\t\t\t\t                                              ")
	fmt.Println("\t\t\t\t\t\t  ------------   Quiz Result  --------------- ")
	fmt.Println("\t\t\t\t\t\t        No. Of Questions In Quiz : 10         ")
	fmt.Printf("\t\t\t\t\t\t    Total No. Of Questions Attampted: %v      \n", len(AA))
	fmt.Printf("\t\t\t\t\t\t              Right Answer : %d               \n", ra)
	fmt.Printf("\t\t\t\t\t\t              Wrong Answer : %d               \n", wa)
	fmt.Println("\t\t\t\t\t\t----------------------------------------------")
	fmt.Printf("\t\t\t\t\t\t          Total Score  : %d                   \n", score)
	fmt.Println("\t\t\t\t\t\t----------------------------------------------|")
}

func FinalResult() {
	fmt.Println("---------------------------------------------------------------------------------------------------------------------")
	if ra >= 6 {
		fmt.Printf("       Congratulation , Mr. %s %s For Passing This Quiz .\n", user.Fname, user.Lname)

	} else {
		fmt.Printf("       Sorry , Mr. %s %s For Loosing The Quiz This Time , But You Can Try Again\n", user.Fname, user.Lname)
	}
	fmt.Println("----------------------------------------------------------------------------------------------------------------------")
}
