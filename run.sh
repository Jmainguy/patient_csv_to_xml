#!/bin/bash
docker run --name patient_csv_to_xml --restart always -p 8000:8000 -d patient_csv_to_xml 
