module FirstGin

go 1.13

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/Unknwon/com v0.0.0-00010101000000-000000000000
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.51.1
	github.com/go-openapi/jsonreference v0.19.3 // indirect
	github.com/go-openapi/spec v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.6 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.11
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/robfig/cron v1.2.0
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.3
	github.com/tealeg/xlsx v1.0.5
	github.com/unknwon/com v1.0.1 // indirect
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	golang.org/x/sys v0.0.0-20191228213918-04cbcbbfeed8 // indirect
	golang.org/x/tools v0.0.0-20191230220329-2aa90c603ae3 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/ini.v1 v1.51.1 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace (
	FirstGin/conf => ~/go-application/FirstGin/pkg/conf v0.0.1
	FirstGin/docs => ~/go-application/FirstGin/docs v0.0.1
	FirstGin/middleware => ~/go-application/FirstGin/middleware v0.0.1
	FirstGin/models => ~/go-application/gFirstGin/models v0.0.1
	FirstGin/pkg/ => ~/go-application/FirstGin/pkg/setting v0.0.1
	FirstGin/pkg/app => ~/go-application/FirstGin/pkg/app v0.0.1
	FirstGin/pkg/setting => ~/go-application/FirstGin/pkg/setting v0.0.1
	FirstGin/routers => ~/go-application/FirstGin/routers v0.0.1
	github.com/Unknwon/com => github.com/unknwon/com v0.0.0-20190804042917-757f69c95f3e
)
