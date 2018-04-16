package main

import (
    "strings"
    "fmt"
)


func addPatient(v *Import, line []string){
        patFullName := strings.Split(line[9], ",")
        fullPatBirthDate := ""
        fullSrvFromDate := ""
        fullSrvToDate := ""
        fullInsBirthDate := ""
        insFirstName := ""
        insLastName := ""
        formatedSrvCharges := ""
        fullPatPhoneNo := ""
        fullInsPhone := ""
        if line[10] != "" {
            fullPatBirthDate = formatBirthDate(line[10],line[11],line[12])
        }
        if line[106] != "" {
            fullSrvFromDate = formatBirthDate(line[106],line[107],line[108])
        }
        if line[109] != "" {
            fullSrvToDate = formatBirthDate(line[109],line[110],line[111])
        }
        if line[37] != "" {
            fullInsBirthDate = formatBirthDate(line[37],line[38],line[39])
        }
        if line[15] != "" {
            insFullName := strings.Split(rm_lead_space(line[15]), " ")
            if len(insFullName) > 1 {
                insFirstName = rm_lead_space(insFullName[1])
                insLastName = insFullName[0]
            }
        }
        if line[118] != "" {
            formatedSrvCharges = fmt.Sprintf("%s.%s", rm_lead_space(line[118]), rm_lead_space(line[119]))
        }
        if line[27] != "" {
            fullPatPhoneNo = formatPhoneNumber(rm_lead_space(line[27]), rm_lead_space(line[28]))
        }
        if line[30] != "" {
            fullInsPhone = formatPhoneNumber(rm_lead_space(line[30]), rm_lead_space(line[31]))
        }
        data := Patient{
            PatFirstName: rm_lead_space(patFullName[1]),
            PatLastName: patFullName[0],
            PatAccountNo: rm_lead_space(line[225]),
            PatAddress: rm_lead_space(line[16]),
            PatBirthDate: fullPatBirthDate,
            PatCity: rm_lead_space(line[22]),
            PatPhoneNo: fullPatPhoneNo,
            PatSigOnFile: "1",
            PatState: rm_lead_space(line[23]),
            PatZip: rm_lead_space(line[26]),
            PatClassification: "Affinity Health Services- Smithfield",
            Patient_Insured: Patient_Insured{
                InsAddress: rm_lead_space(line[21]),
                InsBirthDate: fullInsBirthDate,
                InsCity: rm_lead_space(line[24]),
                InsFirstName: insFirstName,
                InsLastName: insLastName,
                InsIDNumber: rm_lead_space(line[8]),
                InsPhone: fullInsPhone,
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
                    SrvCharges: formatedSrvCharges,
                    SrvFromDate: fullSrvFromDate,
                    SrvPlace: rm_lead_space(line[112]),
                    SrvEMG: rm_lead_space(line[113]),
                    SrvProcedureCode: rm_lead_space(line[114]),
                    SrvModifier1: rm_lead_space(line[115]),
                    SrvToDate: fullSrvToDate,
                    SrvUnits: rm_lead_space(line[120]),
                },
            },
        }
        v.Patients = append(v.Patients, data)
}
