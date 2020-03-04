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
	data := &Patient{}
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
		insFullName := strings.Split(rmLeadSpace(line[16]), " ")
		if len(insFullName) > 1 {
			insFirstName = rmLeadSpace(insFullName[1])
			insLastName = insFullName[0]
		}
	}
	if line[121] != "" {
		formatedSrvCharges = fmt.Sprintf("%s.%s", rmLeadSpace(line[121]), rmLeadSpace(line[122]))
	}
	if line[28] != "" {
		fullPatPhoneNo = formatPhoneNumber(rmLeadSpace(line[28]), rmLeadSpace(line[29]))
	}
	if line[31] != "" {
		fullInsPhone = formatPhoneNumber(rmLeadSpace(line[31]), rmLeadSpace(line[32]))
	}
	if line[41] != "" {
		patSex = "M"
	} else if line[42] != "" {
		patSex = "F"
	}
	if line[75] != "" {
		phyType = "Rendering"
		*data = Patient{
			PatFirstName:      rmLeadSpace(patFullName[1]),
			PatLastName:       patFullName[0],
			PatAccountNo:      rmLeadSpace(line[226]),
			PatAddress:        rmLeadSpace(line[17]),
			PatBirthDate:      fullPatBirthDate,
			PatCity:           rmLeadSpace(line[23]),
			PatPhoneNo:        fullPatPhoneNo,
			PatSigOnFile:      "1",
			PatState:          rmLeadSpace(line[24]),
			PatZip:            rmLeadSpace(line[27]),
			PatSex:            patSex,
			PatClassification: patClass,
			PatientInsured: PatientInsured{
				InsAddress:              rmLeadSpace(line[22]),
				InsBirthDate:            fullInsBirthDate,
				InsCity:                 rmLeadSpace(line[25]),
				InsFirstName:            insFirstName,
				InsLastName:             insLastName,
				InsIDNumber:             rmLeadSpace(line[9]),
				InsPhone:                fullInsPhone,
				InsState:                rmLeadSpace(line[26]),
				InsZip:                  rmLeadSpace(line[30]),
				PatInsRelationToInsured: "1",
			},
			Physician: &Physician{
				PhyName: rmLeadSpace(line[75]),
				PhyNPI:  rmLeadSpace(line[76]),
				PhyType: phyType,
			},
			Claim: Claim{
				ClaBillDate: rmLeadSpace(line[223]),
				ClaimInsured: ClaimInsured{
					ClaInsAcceptAssignment:         "1",
					ClaInsPriorAuthorizationNumber: rmLeadSpace(line[103]),
				},
				ClaDiagnosis1:   rmLeadSpace(line[89]),
				ClaICDIndicator: "0",
				ServiceLine: ServiceLine{
					SrvCharges:       formatedSrvCharges,
					SrvFromDate:      fullSrvFromDate,
					SrvPlace:         rmLeadSpace(line[113]),
					SrvEMG:           rmLeadSpace(line[114]),
					SrvProcedureCode: rmLeadSpace(line[115]),
					SrvModifier1:     rmLeadSpace(line[116]),
					SrvToDate:        fullSrvToDate,
					SrvUnits:         rmLeadSpace(line[123]),
				},
			},
		}
	} else {
		*data = Patient{
			PatFirstName:      rmLeadSpace(patFullName[1]),
			PatLastName:       patFullName[0],
			PatAccountNo:      rmLeadSpace(line[226]),
			PatAddress:        rmLeadSpace(line[17]),
			PatBirthDate:      fullPatBirthDate,
			PatCity:           rmLeadSpace(line[23]),
			PatPhoneNo:        fullPatPhoneNo,
			PatSigOnFile:      "1",
			PatState:          rmLeadSpace(line[24]),
			PatZip:            rmLeadSpace(line[27]),
			PatSex:            patSex,
			PatClassification: patClass,
			PatientInsured: PatientInsured{
				InsAddress:              rmLeadSpace(line[22]),
				InsBirthDate:            fullInsBirthDate,
				InsCity:                 rmLeadSpace(line[25]),
				InsFirstName:            insFirstName,
				InsLastName:             insLastName,
				InsIDNumber:             rmLeadSpace(line[9]),
				InsPhone:                fullInsPhone,
				InsState:                rmLeadSpace(line[26]),
				InsZip:                  rmLeadSpace(line[30]),
				PatInsRelationToInsured: "1",
			},
			Claim: Claim{
				ClaBillDate: rmLeadSpace(line[223]),
				ClaimInsured: ClaimInsured{
					ClaInsAcceptAssignment:         "1",
					ClaInsPriorAuthorizationNumber: rmLeadSpace(line[103]),
				},
				ClaDiagnosis1:   rmLeadSpace(line[89]),
				ClaICDIndicator: "0",
				ServiceLine: ServiceLine{
					SrvCharges:       formatedSrvCharges,
					SrvFromDate:      fullSrvFromDate,
					SrvPlace:         rmLeadSpace(line[113]),
					SrvEMG:           rmLeadSpace(line[114]),
					SrvProcedureCode: rmLeadSpace(line[115]),
					SrvModifier1:     rmLeadSpace(line[116]),
					SrvToDate:        fullSrvToDate,
					SrvUnits:         rmLeadSpace(line[123]),
				},
			},
		}
	}
	v.Patients = append(v.Patients, *data)
}
