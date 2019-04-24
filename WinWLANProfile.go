/**
* WinWLANProfile - Package for handling the Windows WLAN Profile XML Files
* Author: Nigel Ticknor
**/


package WinWLANProfile

import (
	"os"
	"io/ioutil"
	"encoding/xml"
	)

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

func ReadFile(filename string) (WLANProfile, error) {
	var wlp WLANProfile
	xmlFile, err := os.Open(filename)
    if err != nil {
		return wlp, err
    }
    defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue,&wlp)
	return wlp, nil
}

func WriteFile(filename string,wlp WLANProfile) (error) {
	file, err  := xml.MarshalIndent(wlp,"","\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename,file,0644)
	if err != nil {
		return err
	}
	return nil
}