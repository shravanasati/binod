package main

import _ "embed"

//go:embed web/templates/index.html
var indexTemplate string

//go:embed web/templates/leaderboard.html
var leaderboardTemplate string

//go:embed web/templates/footer.html
var footerTemplate string

//go:embed web/templates/navbase.html
var navbaseTemplate string

//go:embed web/templates/404.html
var notFoundTemplate string

//go:embed web/templates/signup.html
var signUpTemplate string
