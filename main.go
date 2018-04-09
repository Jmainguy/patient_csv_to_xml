package main

import (
    "encoding/csv"
    "os"
    "encoding/xml"
    "fmt"
    "io"
    "strings"
)

func rm_lead_space(old_value string) (new_value string) {
    new_value = strings.TrimLeft(old_value, " ")
    return
}

func main() {

    filename := "/tmp/jon.csv"
    v := &Import{}

    // Open CSV file
    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Read File into a Variable
    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        panic(err)
    }

    // Loop through lines & turn into object
    for _, line := range lines {
        patFullName := strings.Split(line[9], ",")
        data := Patient{
            PatFirstName: rm_lead_space(patFullName[1]),
            PatLastName: patFullName[0],
            PatAddress: rm_lead_space(line[16]),
            PatBirthDate: rm_lead_space(line[10]),
            PatCity: rm_lead_space(line[22]),
            PatPhoneNo: rm_lead_space(line[27]),
            PatSigOnFile: "1",
            PatState: rm_lead_space(line[23]),
            PatZip: rm_lead_space(line[26]),
            Patient_Insured: Patient_Insured{
                InsAddress: rm_lead_space(line[21]),
                InsBirthDate: rm_lead_space(line[37]),
                InsCity: rm_lead_space(line[24]),
                InsFirstName: rm_lead_space(line[15]),
                InsIDNumber: rm_lead_space(line[8]),
                InsPhone: rm_lead_space(line[30]),
                InsState: rm_lead_space(line[25]),
                InsZip: rm_lead_space(line[29]),
                PatInsRelationToInsured: "1",
            },
            Claim: Claim{
                ClaBillDate: rm_lead_space(line[222]),
                Claim_Insured: Claim_Insured{
                    ClaInsAcceptAssignment: "1",
                    ClaInsPriorAuthorizationNumber: rm_lead_space(line[102]),
                },
                Service_Line: Service_Line{
                    SrvCharges: fmt.Sprintf("%s.%s", rm_lead_space(line[118]), rm_lead_space(line[119])),
                    SrvFromDate: rm_lead_space(line[106]),
                    SrvPlace: rm_lead_space(line[112]),
                    SrvProcedureCode: rm_lead_space(line[114]),
                    SrvToDate: rm_lead_space(line[109]),
                    SrvUnits: rm_lead_space(line[120]),
                },
            },
        }
        v.Patients = append(v.Patients, data)
    }

    writefilename := "/tmp/patients.xml"
    file, _ := os.Create(writefilename)

    xmlWriter := io.Writer(file)

    enc := xml.NewEncoder(xmlWriter)
    enc.Indent("", "  ")
    if err := enc.Encode(v); err != nil {
        fmt.Printf("error: %v\n", err)
    }
}

