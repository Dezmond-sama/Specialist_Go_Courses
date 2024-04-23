package main

type About struct {
	Name     string
	Position string
	Phone    string
	Telegram string
	Email    string
}

type Job struct {
	Title       string
	From        string
	To          string
	Description string
}

var about = About{
	Name:     "Evgeny",
	Position: "Software developer",
	Phone:    "+79999999999",
	Email:    "dezmond-sama@mail.ru",
	Telegram: "https://t.me/dezmond_sama",
}

var jobs = []Job{
	{
		Title: "VNIIAES",
		From:  "2010",
		To:    "2016",
		Description: `
		Lorem ipsum dolor sit amet consectetur adipisicing elit. Alias nihil voluptates id consequatur nemo. 
		Cupiditate doloremque amet omnis sit, architecto placeat natus pariatur? 
		Blanditiis reprehenderit rerum repellendus dolore cupiditate corporis.
		`,
	},
	{
		Title: "JSC RASU",
		From:  "2016",
		To:    "2024",
		Description: `
		Sed, mollitia rerum ex amet, fugiat corporis perferendis obcaecati architecto deserunt dolorum nemo aliquam, 
		itaque aut veniam! Quas suscipit expedita aliquam ad nesciunt totam repellendus veniam aperiam repellat 
		doloribus voluptas voluptate, architecto molestiae provident pariatur consequatur a quisquam vel illo, tempore et?
		`,
	},
}
var skills = []string{
	"C#",
	"Qt C++",
	"Golang",
	"Delphi",
	"Unity C#",
	"JavaScript",
	"HTML",
	"CSS",
	"React JS",
	"Next JS",
	"Node JS",
	"Python",
	"Java",
	"Flutter",
	"VBA",
	"PostgreSQL",
	"Interbase",
	"Oracle",
}
