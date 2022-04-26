module github.com/LassiHeikkila/bluetooth

go 1.15

replace tinygo.org/x/bluetooth => ./

require (
	github.com/JuulLabs-OSS/cbgo v0.0.2
	github.com/go-ole/go-ole v1.2.4
	github.com/godbus/dbus/v5 v5.0.3
	github.com/muka/go-bluetooth v0.0.0-20210812063148-b6c83362e27d
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	tinygo.org/x/bluetooth v0.0.0-00010101000000-000000000000
	tinygo.org/x/drivers v0.19.1-0.20220421091018-f0a260be6630
	tinygo.org/x/tinyterm v0.1.0
)
