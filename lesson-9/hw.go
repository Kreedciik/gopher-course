package main

import (
	b "lesson-9/bank"
	c "lesson-9/course"
	fs "lesson-9/filesystem"
	p "lesson-9/payment"
	v "lesson-9/vehicle"
)

func main() {

	creditCard := p.CreditCardProcessor{CardNumber: "9989 8765 7890 7766", Transactions: []string{}}
	paypalAccount := p.PaypalProcessor{Email: "john@gmail.com", Transactions: []string{}}
	cryptoWallet := p.CryptoProcessor{WalletAddress: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", Transactions: []string{}}
	paymentSystems := []p.PaymentProcessor{&creditCard, &paypalAccount, &cryptoWallet}
	p.ShowTransactions(paymentSystems)

	txtFile := fs.TextFile{}
	csvFile := fs.CSVFile{}
	logFile := fs.LogFile{}

	txtFile.Open("./folder/a.txt")
	csvFile.Open("./folder/a.csv")
	logFile.Open("./folder/a.log")

	files := []fs.FileHandler{&txtFile, &csvFile, &logFile}

	fs.ReadAllTypeOfFiles(files)
	fs.CloseAllTypeOfFiles(files)

	mazda := v.Car{v.VehicleType{"Mazda", "petrol"}}
	byd := v.ElectricCar{v.VehicleType{"BYD", "electric"}, "70%"}
	volvo := v.Truck{v.VehicleType{"Volvo", "diesel"}}

	vehicles := []v.VehicleController{&mazda, &byd, &volvo}
	v.StartEngineOfAllVehicles(vehicles)
	v.DriveToDistance(vehicles, 12.0)
	v.StopEngineOfAllVehicles(vehicles)

	programmingCourse := c.ProgrammingCourse{Category: "Programming", Courses: []c.Course{}}
	designCourse := c.DesignCourse{Category: "Design", Courses: []c.Course{}}
	languageCourse := c.LanguageCourse{Category: "Language", Courses: []c.Course{}}

	c.AddAnyCourse(&programmingCourse, c.Course{1, "JavaScript"})
	c.AddAnyCourse(&designCourse, c.Course{1, "Figma"})
	c.AddAnyCourse(&languageCourse, c.Course{1, "English"})

	courses := []c.CourseManager{&programmingCourse, &designCourse, &languageCourse}
	c.ShowAnyCourses(courses)

	savingsAccount := b.SavingsAccount{}
	checkingAccount := b.CheckingAccount{}
	loanAccount := b.LoanAccount{}
	accounts := []b.BankAccount{&savingsAccount, &checkingAccount, &loanAccount}

	b.ReplenishAnyAccount(&savingsAccount, 100)
	b.ReplenishAnyAccount(&checkingAccount, 30)
	b.ReplenishAnyAccount(&loanAccount, 20)

	b.WithdrawFromAccount(&checkingAccount, 20)
	b.ShowAllBalances(accounts)
}
