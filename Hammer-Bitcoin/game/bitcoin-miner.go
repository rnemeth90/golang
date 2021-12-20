package game

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

// OGH is how we address the player
const OGH = "O Great Gill Bates"

// default values
var year = 1
var employees = 100
var cash = 2800
var computers = 1000
var computerPrice int
var computersMaintained int
var cashPaidToEmployees int
var starved = 0
var marketCrashVictims = 0
var newEmployees = 5
var cashMined = 3000
var bitcoinGeneratedPerComputer = 3
var amountStolenByHackers = 200

// Play plays the game
func Play() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// write greeting
	printIntroductoryParagraph()

	stillInOffice := true

	// play for 10 years, or until kicked out
	for stillInOffice && year <= 10 {
		computerPrice = updateComputerPrice()
		printSummary()
		buyComputers()
		sellComputers()
		payEmployees()
		maintainComputers()
		marketCrashVictims = checkForCrash()
		employees = employees - marketCrashVictims

		if countStarvedEmployees() >= 45 {
			stillInOffice = false
		}

		newEmployees = countNewHires()
		employees = employees + newEmployees

		cash = cash + mineBitCoin(computersMaintained)

		checkForHackers()
		computerPrice = updateComputerPrice()
		year = year + 1
	}

	printFinalScore()

}

// printIntroductoryParagraph prints the intro paragraph
func printIntroductoryParagraph() {
	clearScreen()
	color.Yellow(`BITCOIN MINER
-------------
Congratulations! You are the newest CEO of Make Me Rich, Inc, elected for a ten year term. Your
duties are to dispense living expenses for employees, direct mining of bitcoin, and buy and sell
computers as needed to support the corporation.

Watch out for hackers and market crashes!

Cash is the general currency, measured in bitcoins.

The following will help you in your decisions:

	* Each employee needs at least 20 bitcoins converted to cash per year to survive
	* Each employee can maintain at most 10 computers
	* It takes 2 bitcoins to pay for electricity to mine bitcoin on a computer
	* The market price for computers fluctuates yearly

Lead the team wisely and you will be showered with appreciation at the end of your term.

Do it poorly and you will be terminated!

`)
}

// updateComputerPrice Randomly sets the new price of computers.
// returns the new price of a computer as an int.
// The price fluctuates from 17 to 26 bitcoin per computer.
func updateComputerPrice() int {
	return rand.Intn(10) + 17
}

// printSummary prints the year-end summary
func printSummary() {
	fmt.Printf("%s!", OGH)
	fmt.Println("")
	fmt.Println(fmt.Sprintf("You are in year %d of your rule.", year))

	if marketCrashVictims > 0 {
		color.Red("A terrible market crash wiped out %d of your team.", marketCrashVictims)
	}

	fmt.Println(fmt.Sprintf("In the previous year, %d of your team starved to death.", starved))
	fmt.Println(fmt.Sprintf("In the previous year, %d employee(s) got employed by the corporation.", newEmployees))
	fmt.Println(fmt.Sprintf("The employee head count is now %d.", employees))
	fmt.Println(fmt.Sprintf("We mined %d bitcoins at %d bitcoins per computer.", cashMined, bitcoinGeneratedPerComputer))

	if amountStolenByHackers > 0 {
		color.Red("*** Hackers stole %d bitcoins, leaving %d bitcoins in your online wallet.", amountStolenByHackers, cash)
	} else {
		fmt.Println(fmt.Sprintf("We have %d bitcoins of cash in storage.", cash))
	}

	fmt.Println(fmt.Sprintf("The corporation owns %d computers for mining.", computers))
	fmt.Println(fmt.Sprintf("Computers currently cost %d bitcoins each.", computerPrice))
	fmt.Println("")
}

// buyComputers Allows the player to buy computers. If a valid amount is entered, the available cash is reduced
// accordingly
func buyComputers() {
	question := "How many computers will you buy?"
	computersToBuy := getNumber(question)
	cost := computerPrice * computersToBuy
	for cost > cash {
		jest(fmt.Sprintf("We have but %d bitcoins of cash, not %d!", cash, cost))
		computersToBuy := getNumber(question)
		cost = computerPrice * computersToBuy
	}
	cash = cash - cost
	computers = computers + computersToBuy

	fmt.Println(fmt.Sprintf("%s, you now have %d computers", OGH, computers))
	fmt.Println(fmt.Sprintf("and %d bitcoins of cash.", cash))
}

// sellComputers allows the player to sell computers, if any are on hand. Available
// cash will be increased by the value of the computers sold
func sellComputers() {
	question := "How many computers will you sell?"
	computersToSell := getNumber(question)

	for computersToSell > computers {
		jest(fmt.Sprintf("The corporation only has %d computers!", computers))
		computersToSell = getNumber(question)
	}
	computers = computers - computersToSell
	cash = cash + computerPrice*computersToSell

	fmt.Println(fmt.Sprintf("%s, you now have %d computers", OGH, computers))
	fmt.Println(fmt.Sprintf("and %d bitcoins of cash.", cash))
}

// payEmployees allows the player to decide how much cash to use to feed people. If a valid
// amount is entered, the available cash is reduced accordingly
func payEmployees() {
	question := "How much bitcoin will you distribute to the employees?"
	cashPaidToEmployees = getNumber(question)

	for cashPaidToEmployees > cash {
		jest(fmt.Sprintf("We have but %d bitcoins!", cash))
		cashPaidToEmployees = getNumber(question)
	}
	cash = cash - cashPaidToEmployees

	fmt.Println(fmt.Sprintf("%s, %d bitcoins remain.", OGH, cash))
}

// maintainComputers allows the user to choose how much to spend on maintenance
func maintainComputers() {
	question := "How many bitcoins will you allocate for maintenance?"
	maintenanceAmount := 0
	haveGoodAnswer := false

	for !haveGoodAnswer {
		maintenanceAmount = getNumber(question)
		if maintenanceAmount > cash {
			jest(fmt.Sprintf("We have but %d bitcoins left!", cash))
		} else if maintenanceAmount > 2*computers {
			jest(fmt.Sprintf("We have but %d computers available for mining!", computers))
		} else if maintenanceAmount > 20*employees {
			jest(fmt.Sprintf("We have but %d people to maintain the computers!", employees))
		} else {
			haveGoodAnswer = true
		}
	}
	computersMaintained = maintenanceAmount / 2

	fmt.Println(fmt.Sprintf("%s, we now have %d bitcoins in storage.", OGH, cash))
}

// checkForCrash checks for market crash, and counts the victims
func checkForCrash() int {
	victims := 0

	diceRoll := rand.Intn(99) + 1
	if diceRoll <= 15 {
		color.Red("*** A terrible market crash wipes out half of the corporation's employees! ***")
		victims = employees / 2
	}
	return victims
}

// countNewHires counts how many new employees joined the company
func countNewHires() int {
	var newEmployees int
	if starved > 0 {
		newEmployees = 0
	} else {
		newEmployees = (20*computers+cash)/(100*employees) + 1
	}
	return newEmployees
}

// checkForHackers checks if hackers get into the system, and determines how much they stole.
func checkForHackers() {
	diceRoll := rand.Intn(99) + 1
	if diceRoll < 40 {
		percentHacked := 10 + rand.Intn(21)
		color.Red("*** Hackers steal %d percent of your bitcoins! ***", percentHacked)
		amountStolenByHackers = (percentHacked * cash) / 100
		cash = cash - amountStolenByHackers
	} else {
		amountStolenByHackers = 0
	}
}

// mineBitCoin collects the new cash mined
func mineBitCoin(computers int) int {
	bitcoinGeneratedPerComputer = rand.Intn(6) + 1
	cashMined = bitcoinGeneratedPerComputer * computers
	return cashMined
}

// countStarvedEmployees counts how many people starved, and removes them from the employees
func countStarvedEmployees() int {
	employeesPaid := cashPaidToEmployees / 20
	percentStarved := 0

	if employeesPaid >= employees {
		starved = 0
		fmt.Println("The corporation's employees are well fed and happy.")
	} else {
		starved = employees - employeesPaid
		color.Red("%d employees starved to death.", starved)
		percentStarved = (100 * starved) / employees
		employees = employees - starved
	}
	return percentStarved
}

// printFinalScore prints out the final score
func printFinalScore() {
	clearScreen()

	if starved >= (45*employees)/100 {
		color.Red(`O Once-Great %s,
%d of your team starved during the last year of your incompetent reign!
The few who remain hacked your bank account and changed your password, effectively evicting you!

Your final rating: TERRIBLE.`, OGH, starved)
		return
	}

	computerScore := computers

	if 20*employees < computerScore {
		computerScore = 20 * employees
	}

	if computerScore < 600 {
		color.Cyan(`Congratulations, %s.
You have ruled wisely,  but not well. You have led your people through ten difficult
years, but your corporation assets have shrunk to a mere %d computers.

Your final rating: ADEQUATE`, OGH, computers)
	} else if computerScore < 800 {
		color.Yellow(`Congratulations %s,
You  have ruled wisely, and shown the online world that it's possible to make money in cryptocurrency.

Your final rating: GOOD.`, OGH)
	} else {
		color.White(`Congratulations %s,
you  have ruled wisely and well, and expanded your holdings while keeping your team happy.
Altogether, a most impressive job!

Your final rating: SUPERB.`, OGH)
	}
}

// GetYesOrNo allows the player to try again, or quit
func GetYesOrNo(q string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(q)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if char == 'n' || char == 'N' {
			return false
		}
		return true
	}
}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// getNumber prints question q and asks for a number, then returns it
// as an int
func getNumber(q string) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(q)
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number!")
			continue
		} else {
			return num
		}
	}
}

// jest tells player that a request cannot be fulfilled
func jest(msg string) {
	fmt.Println("")
	color.Magenta("%s, you are dreaming!", OGH)
	fmt.Println(msg)
}
