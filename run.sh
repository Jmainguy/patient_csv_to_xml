#!/bin/bash
docker kill patient_csv_to_xml
docker rm patient_csv_to_xml
docker run --name patient_csv_to_xml --restart always -p 8000:8000 -d hub.soh.re/soh.re/patient_csv_to_xml 
