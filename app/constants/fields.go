package constants

/*
	This file is to manage the constant values related to file processing
 */

import "regexp"

var EmailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

var Headers = []string{
	"Name",
	"Salary",
	"Email",
	"ID",
}

var EmailVariations = []string{
	"Email",
	"E-mail",
	"e-mail",
	"email",
}

var IdVariations = []string{
	"Number",
	"ID",
	"Employee Number",
	"emp id",
}

var FullNameVariations = []string{
	"Name",
}

var FirstNameVariations = []string{
	"First",
	"first name",
	"f. name",
}

var LastNameVariations = []string{
	"Last",
	"last name",
	"l. name",
}

var SalaryVariations = []string{
	"Wage",
	"Salary",
	"Rate",
	"wage",
}
