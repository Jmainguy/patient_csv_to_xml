FROM centos:latest
ADD Sohre.repo /etc/yum.repos.d/
RUN yum install -y patient_csv_to_xml-0.1-1.el7.centos && \
    yum clean all
RUN chgrp -R 0 /opt/patient_csv_to_xml \
    && chmod -R g+rwX /opt/patient_csv_to_xml
WORKDIR /opt/patient_csv_to_xml
USER patient_csv_to_xml
EXPOSE 8000
CMD ["/opt/patient_csv_to_xml/patient_csv_to_xml"]