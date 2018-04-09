package main

import (
    "encoding/xml"
)

type Service_Line struct {
    XMLName xml.Name `xml:"Service_Line"`
    SrvCharges string `xml:"SrvCharges"`
    SrvFromDate string `xml:"SrvFromDate"`
    SrvPlace string `xml:"SrvPlace"`
    SrvProcedureCode string `xml:"SrvProcedureCode"`
    SrvToDate string `xml:"SrvToDate"`
    SrvUnits string `xml:"SrvUnits"`
}

type Claim_Insured struct {
    XMLName xml.Name `xml:"Claim_Insured"`
    ClaInsAcceptAssignment string `xml:"ClaInsAcceptAssignment"`
    ClaInsPriorAuthorizationNumber string `xml:"ClaInsPriorAuthorizationNumber"`
}

type Claim struct {
    XMLName xml.Name `xml:"Claim"`
    ClaBillDate string `xml:"ClaBillDate"`
    Claim_Insured Claim_Insured `xml:"Claim_Insured"`
    Service_Line Service_Line `xml:"Service_Line"`
}

type Patient_Insured struct {
    XMLName xml.Name `xml:"Patient_Insured"`
    InsAddress string `xml:"InsAddress"`
    InsBirthDate string `xml:"InsBirthDate"`
    InsCity string `xml:"InsCity"`
    InsFirstName string `xml:"InsFirstName"`
    InsIDNumber string `xml:"InsIDNumber"`
    InsPhone string `xml:"InsPhone"`
    InsState string `xml:"InsState"`
    InsZip string `xml:"InsZip"`
    PatInsRelationToInsured string `xml:"PatInsRelationToInsured"`
}

type Patient struct {
    XMLName xml.Name `xml:"Patient"`
    PatFirstName string `xml:"PatFirstName"`
    PatLastName string `xml:"PatLastName"`
    PatAddress string `xml:"PatAddress"`
    PatBirthDate string `xml:"PatBirthDate"`
    PatCity string  `xml:"PatCity"`
    PatPhoneNo string `xml:"PatPhoneNo"`
    PatSigOnFile string `xml:"PatSigOnFile"`
    PatState string `xml:"PatState"`
    PatZip  string `xml:"PatZip"`
    Patient_Insured Patient_Insured `xml:"Patient_Insured"`
    Claim Claim `xml:"Claim"`
}

type Import struct {
    XMLName xml.Name `xml:"Import"`
    Patients []Patient `xml:"Patient"`
}
