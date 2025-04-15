package function

var (
	Version = "1.24.2"
	Build   = "Release"
)

func GetAppVersion() string {
	return Version + " (" + Build + ")"
}
