package utils

const (

	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"

	LocalURL     = "43.139.145.135:7777"
	PlayURL      = "http://" + LocalURL + "/public/"
	CoverTestURL = "http://" + LocalURL + "/public/cover/covertest.jpg"

)
