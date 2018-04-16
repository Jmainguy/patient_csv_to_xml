# This expects a few requirements
# one, that https://github.com/Jmainguy/docker_rpmbuild is cloned into ~/Github/docker_rpmbuild
# two, that docker is installed and running
# three, that ~/Github/docker_rpmbuild/dockerbuild/build.sh centos7 has been run
rpm:
	@tar -czvf ~/Github/docker_rpmbuild/rpmbuild/SOURCES/patient_csv_to_xml.tar.gz ../patient_csv_to_xml
	@cp patient_csv_to_xml.spec ~/Github/docker_rpmbuild/rpmbuild/SPECS/patient_csv_to_xml.spec
	@cd ~/Github/docker_rpmbuild/; ./run.sh centos7 patient_csv_to_xml
	@ls -ltrh ~/Github/docker_rpmbuild/rpmbuild/RPMS/x86_64/patient_csv_to_xml*
