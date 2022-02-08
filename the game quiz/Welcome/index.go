package welcome

import (
	"fmt"
)

func IntroAboutQuiz() {

	fmt.Println("---------------------------------------------------------------------------------------------------------------")
	fmt.Println("|________________________________________   About The Quiz   __________________________________________________|")
	fmt.Println("|                                                                                                              |")
	fmt.Println("|       1.  In This Quiz Game You Will Given Some Simple Math Problems                                         |")
	fmt.Println("|       2.  When You Start The Quiz , You Will See The Problems On The Console With Their Numbers              |")
	fmt.Println("|       3.  You Have Solve The Problem And Type The Answer Just After The Quiz                                 |")
	fmt.Println("|       4.  There Will Be Total 10 Problems Which You Are Gonna Solve And Answer Them                          |")
	fmt.Println("|       5.  Each Right Answer Will Add 1 Number To Your Score !                                                |")
	fmt.Println("|       6.  Each Wrong Answer Will Deduct ! Number From Your Score !                                           |")
	fmt.Println("|       7.  To Pass This Quiz You Have To Score Min 06/10 Of Total OtherWise You Will Loose This Quiz !        |")
	fmt.Println("|______________________________________________________________________________________________________________|")

	fmt.Println("                                            Are You Ready To Score ?")
	var options int
	fmt.Println(" Press 1 To Enter To Quiz Or 0 To Exit The Program .")
	fmt.Scanf("%d", &options)
	if options == 1 {
		fmt.Println("                                                Lets Get Started                                           ")
	} else {
		fmt.Println("It Feels So Bad To Hear That You Choose To Exit , See You Soon .")
	}

}
