package controllers

type Controller struct {
	Auth      AuthController
	Resident  ResidentController
	Event     EventController
	Household HouseholdController
	Income    IncomeController
	Official  OfficialController
	Expense   ExpenseController
	Logbook   LogbookController
}
