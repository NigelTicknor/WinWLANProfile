# WinWLANProfile
Golang library for interacting with the WLAN Profile files

# Installation
`go get github.com/NigelTicknor/WinWLANProfile`

`import "github.com/NigelTicknor/WinWLANProfile"`

# Usage
`WinWLANProfile.ReadFile("path-to-xml-profile") (WLANProfile, error)` - returns a WLANProfile struct with all data
`WinWLANProfile.WriteFile("path-to-xml-output",WLANProfile) (error)` - writes a WLANProfile to the system

# Coming Soon
* A better way to construct a profile
* Documentation

Here is some placeholder code to create a profile by hand (adjust as needed):
```golang
profile := WinWLANProfile.WLANProfile{
		Name:"ProfileNameOrSSID",
		Xmlns:"http://www.microsoft.com/networking/WLAN/profile/v1",
		SSIDConfig: WinWLANProfile.SSIDConfig{
			SSID: WinWLANProfile.SSID{
				Name:"SSID",
				Hex:hex.EncodeToString([]byte("SSID"))}},
		ConnectionType:"ESS",
		ConnectionMode:"auto",
		MSM:WinWLANProfile.MSM{
			Security:WinWLANProfile.Security{
				AuthEncryption:WinWLANProfile.AuthEncryption{
					Authentication:"WPA2PSK",
					Encryption:"AES",
					UseOneX:"false"},
				SharedKey:WinWLANProfile.SharedKey{
					KeyType:"passPhrase",
					Protected:"false",
					KeyMaterial:"PassKey"}}},
		MacRandomization:WinWLANProfile.MacRandomization{
			Xmlns:"http://www.microsoft.com/networking/WLAN/profile/v3",
			EnableRandomization:false}}
```
