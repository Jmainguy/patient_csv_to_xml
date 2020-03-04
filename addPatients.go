package main

import (
	"fmt"
	"strings"
)

func addPatient(v *Import, line []string, patClass string) {
	patFullName := strings.Split(line[10], ",")
	fullPatBirthDate := ""
	fullSrvFromDate := ""
	fullSrvToDate := ""
	fullInsBirthDate := ""
	insFirstName := ""
	insLastName := ""
	formatedSrvCharges := ""
	fullPatPhoneNo := ""
	fullInsPhone := ""
	patSex := ""
	phyType := ""
	data := Patient{}
	if line[11] != "" {
		fullPatBirthDate = formatBirthDate(line[11], line[12], line[13])
	}
	if line[107] != "" {
		fullSrvFromDate = formatDate(line[107], line[108], line[109])
	}
	if line[110] != "" {
		fullSrvToDate = formatDate(line[110], line[111], line[112])
	}
	if line[38] != "" {
		fullInsBirthDate = formatBirthDate(line[38], line[39], line[40])
	}
	if line[16] != "" {
		insFullName := strings.Split(rm_lead_space(line[16]), " ")
		if len(insFullName) > 1 {
			insFirstName = rm_lead_space(insFullName[1])
			insLastName = insFullName[0]
		}
	}
	if line[121] != "" {
		formatedSrvCharges = fmt.Sprintf("%s.%s", rm_lead_space(line[121]), rm_lead_space(line[122]))
	}
	if line[28] != "" {
		fullPatPhoneNo = formatPhoneNumber(rm_lead_space(line[28]), rm_lead_space(line[29]))
	}
	if line[31] != "" {
		fullInsPhone = formatPhoneNumber(rm_lead_space(line[31]), rm_lead_space(line[32]))
	}
	if line[41] != "" {
		patSex = "M"
	} else if line[42] != "" {
		patSex = "F"
	}
	if line[75] != "" {
		phyType = "Rendering"
		data = Patient{
			PatFirstName:      rm_lead_space(patFullName[1]),
			PatLastName:       patFullName[0],
			PatAccountNo:      rm_lead_space(line[226]),
			PatAddress:        rm_lead_space(line[17]),
			PatBirthDate:      fullPatBirthDate,
			PatCity:           rm_lead_space(line[23]),
			PatPhoneNo:        fullPatPhoneNo,
			PatSigOnFile:      "1",
			PatState:          rm_lead_space(line[24]),
			PatZip:            rm_lead_space(line[27]),
			PatSex:            patSex,
			PatClassification: patClass,
			Patient_Insured: Patient_Insured{
				InsAddress:              rm_lead_space(line[22]),
				InsBirthDate:            fullInsBirthDate,
				InsCity:                 rm_lead_space(line[25]),
				InsFirstName:            insFirstName,
				InsLastName:             insLastName,
				InsIDNumber:             rm_lead_space(line[9]),
				InsPhone:                fullInsPhone,
				InsState:                rm_lead_space(line[26]),
				InsZip:                  rm_lead_space(line[30]),
				PatInsRelationToInsured: "1",
			},
			Physician: &Physician{
				PhyName: rm_lead_space(line[75]),
				PhyNPI:  rm_lead_space(line[76]),
				PhyType: phyType,
			},
			Claim: Claim{
				ClaBillDate: rm_lead_space(line[223]),
				Claim_Insured: Claim_Insured{
					ClaInsAcceptAssignment:         "1",
					ClaInsPriorAuthorizationNumber: rm_lead_space(line[103]),
				},
				ClaDiagnosis1:   rm_lead_space(line[89]),
				ClaICDIndicator: "0",
				Service_Line: Service_Line{
					SrvCharges:       formatedSrvCharges,
					SrvFromDate:      fullSrvFromDate,
					SrvPlace:         rm_lead_space(line[113]),
					SrvEMG:           rm_lead_space(line[114]),
					SrvProcedureCode: rm_lead_space(line[115]),
					SrvModifier1:     rm_lead_space(line[116]),
					SrvToDate:        fullSrvToDate,
					SrvUnits:         rm_lead_space(line[123]),
				},
			},
		}
	} else {
		data = Patient{
			PatFirstName:      rm_lead_space(patFullName[1]),
			PatLastName:       patFullName[0],
			PatAccountNo:      rm_lead_space(line[226]),
			PatAddress:        rm_lead_space(line[17]),
			PatBirthDate:      fullPatBirthDate,
			PatCity:           rm_lead_space(line[23]),
			PatPhoneNo:        fullPatPhoneNo,
			PatSigOnFile:      "1",
			PatState:          rm_lead_space(line[24]),
			PatZip:            rm_lead_space(line[27]),
			PatSex:            patSex,
			PatClassification: patClass,
			Patient_Insured: Patient_Insured{
				InsAddress:              rm_lead_space(line[22]),
				InsBirthDate:            fullInsBirthDate,
				InsCity:                 rm_lead_space(line[25]),
				InsFirstName:            insFirstName,
				InsLastName:             insLastName,
				InsIDNumber:             rm_lead_space(line[9]),
				InsPhone:                fullInsPhone,
				InsState:                rm_lead_space(line[26]),
				InsZip:                  rm_lead_space(line[30]),
				PatInsRelationToInsured: "1",
			},
			Claim: Claim{
				ClaBillDate: rm_lead_space(line[223]),
				Claim_Insured: Claim_Insured{
					ClaInsAcceptAssignment:         "1",
					ClaInsPriorAuthorizationNumber: rm_lead_space(line[103]),
				},
				ClaDiagnosis1:   rm_lead_space(line[89]),
				ClaICDIndicator: "0",
				Service_Line: Service_Line{
					SrvCharges:       formatedSrvCharges,
					SrvFromDate:      fullSrvFromDate,
					SrvPlace:         rm_lead_space(line[113]),
					SrvEMG:           rm_lead_space(line[114]),
					SrvProcedureCode: rm_lead_space(line[115]),
					SrvModifier1:     rm_lead_space(line[116]),
					SrvToDate:        fullSrvToDate,
					SrvUnits:         rm_lead_space(line[123]),
				},
			},
		}
	}
	v.Patients = append(v.Patients, data)
}
