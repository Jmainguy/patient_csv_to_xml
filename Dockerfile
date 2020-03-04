FROM busybox:latest
EXPOSE 8000
LABEL maintainer="Jonathan Seth Mainguy <jon@soh.re>"
ENV version=v0.0.2
RUN wget -O /usr/sbin/patient_csv_to_xml_Linux_x86_64.tar.gz https://github.com/Jmainguy/patient_csv_to_xml/releases/download/$version/patient_csv_to_xml_Linux_x86_64.tar.gz
RUN tar zxvf /usr/sbin/patient_csv_to_xml_Linux_x86_64.tar.gz -C /usr/sbin/
ADD upload.gtpl /usr/sbin/
WORKDIR /usr/sbin/
CMD ["/usr/sbin/patient_csv_to_xml"]
