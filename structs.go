package main

import (
	"encoding/xml"
)

type ServiceLine struct {
	XMLName          xml.Name `xml:"Service_Line"`
	SrvCharges       string   `xml:"SrvCharges,omitempty"`
	SrvFromDate      string   `xml:"SrvFromDate,omitempty"`
	SrvPlace         string   `xml:"SrvPlace,omitempty"`
	SrvEMG           string   `xml:"SrvEMG,omitempty"`
	SrvProcedureCode string   `xml:"SrvProcedureCode,omitempty"`
	SrvModifier1     string   `xml:"SrvModifier1,omitempty"`
	SrvToDate        string   `xml:"SrvToDate,omitempty"`
	SrvUnits         string   `xml:"SrvUnits,omitempty"`
}

type ClaimInsured struct {
	XMLName                        xml.Name `xml:"Claim_Insured"`
	ClaInsAcceptAssignment         string   `xml:"ClaInsAcceptAssignment,omitempty"`
	ClaInsPriorAuthorizationNumber string   `xml:"ClaInsPriorAuthorizationNumber,omitempty"`
}

type Physician struct {
	XMLName xml.Name `xml:"Physician"`
	PhyName string   `xml:"PhyName,omitempty"`
	PhyNPI  string   `xml:"PhyNPI,omitempty"`
	PhyType string   `xml:"PhyType,omitempty"`
}

type Claim struct {
	XMLName         xml.Name     `xml:"Claim"`
	ClaBillDate     string       `xml:"ClaBillDate,omitempty"`
	ClaimInsured    ClaimInsured `xml:"Claim_Insured,omitempty"`
	ClaDiagnosis1   string       `xml:"ClaDiagnosis1,omitempty"`
	ClaICDIndicator string       `xml:"ClaICDIndicator,omitempty"`
	ServiceLine     ServiceLine  `xml:"Service_Line,omitempty"`
}

type PatientInsured struct {
	XMLName                 xml.Name `xml:"Patient_Insured"`
	InsAddress              string   `xml:"InsAddress,omitempty"`
	InsBirthDate            string   `xml:"InsBirthDate,omitempty"`
	InsCity                 string   `xml:"InsCity,omitempty"`
	InsFirstName            string   `xml:"InsFirstName,omitempty"`
	InsLastName             string   `xml:"InsLastName,omitempty"`
	InsIDNumber             string   `xml:"InsIDNumber,omitempty"`
	InsPhone                string   `xml:"InsPhone,omitempty"`
	InsState                string   `xml:"InsState,omitempty"`
	InsZip                  string   `xml:"InsZip,omitempty"`
	PatInsRelationToInsured string   `xml:"PatInsRelationToInsured,omitempty"`
}

type Patient struct {
	XMLName           xml.Name       `xml:"Patient"`
	PatFirstName      string         `xml:"PatFirstName"`
	PatLastName       string         `xml:"PatLastName"`
	PatAccountNo      string         `xml:"PatAccountNo,omitempty"`
	PatAddress        string         `xml:"PatAddress,omitempty"`
	PatBirthDate      string         `xml:"PatBirthDate,omitempty"`
	PatCity           string         `xml:"PatCity,omitempty"`
	PatPhoneNo        string         `xml:"PatPhoneNo,omitempty"`
	PatSigOnFile      string         `xml:"PatSigOnFile,omitempty"`
	PatState          string         `xml:"PatState,omitempty"`
	PatZip            string         `xml:"PatZip,omitempty"`
	PatSex            string         `xml:"PatSex,omitempty"`
	PatClassification string         `xml:"PatClassification,omitempty"`
	PatientInsured    PatientInsured `xml:"Patient_Insured,omitempty"`
	Physician         *Physician
	Claim             Claim `xml:"Claim"`
}

type Import struct {
	XMLName  xml.Name  `xml:"Import"`
	Patients []Patient `xml:"Patient"`
}
