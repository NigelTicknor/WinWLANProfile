package WinWLANProfile

import ("encoding/xml")

type WLANProfile struct {
	XMLName xml.Name `xml:"WLANProfile"`
	Xmlns string `xml:"xmlns,attr"`
	Name string `xml:"name"`
	SSIDConfig SSIDConfig `xml:"SSIDConfig"`
	ConnectionType string `xml:"connectionType"`
	ConnectionMode string `xml:"connectionMode"`
	MSM MSM `xml:"MSM"`
	MacRandomization MacRandomization `xml:"macRandomization"`
}

type SSIDConfig struct {
	XMLName xml.Name `xml:"SSIDConfig"`
	SSID SSID `xml:"SSID"`
}

type SSID struct {
	XMLName xml.Name `xml:"SSID"`
	Hex string `xml:"hex"`
	Name string `xml:"name"`
}

type MSM struct {
	XMLName xml.Name `xml:"MSM"`
	Security Security `xml:"security"`
}

type Security struct {
	XMLName xml.Name `xml:"security"`
	AuthEncryption AuthEncryption `xml:"authEncryption"`
	SharedKey SharedKey `xml:"sharedKey"`
}

type AuthEncryption struct {
	XMLName xml.Name `xml:"authEncryption"`
	Authentication string `xml:"authentication"`
	Encryption string `xml:"encryption"`
	UseOneX string `xml:"useOneX"`
}

type SharedKey struct {
	XMLName xml.Name `xml:"sharedKey"`
	KeyType string `xml:"keyType"`
	Protected string `xml:"protected"`
	KeyMaterial string `xml:"keyMaterial"`
}

type MacRandomization struct {
	XMLName xml.Name `xml:"macRandomization"`
	Xmlns string `xml:"xmlns,attr"`
	EnableRandomization bool `xml:"enableRandomization"`
}
