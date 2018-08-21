%define _unpackaged_files_terminate_build 0
%define  debug_package %{nil}
Name:	patient_csv_to_xml
Version: 0.4.2
Release: 1%{?dist}
Summary: A golang http server to convert generations csv to ez claim xml

License: GPLv2
URL: https://github.com/Jmainguy/patient_csv_to_xml
Source0: patient_csv_to_xml.tar.gz

%description
A golang http server to convert generations csv to ez claim xml

%prep
%setup -q -n patient_csv_to_xml
%build
export GOPATH=/usr/src/go
/usr/bin/go build
%install
mkdir -p $RPM_BUILD_ROOT/opt/patient_csv_to_xml
install -m 0755 $RPM_BUILD_DIR/patient_csv_to_xml/patient_csv_to_xml %{buildroot}/opt/patient_csv_to_xml/
install -m 0755 $RPM_BUILD_DIR/patient_csv_to_xml/upload.gtpl %{buildroot}/opt/patient_csv_to_xml/

%files
/opt/patient_csv_to_xml/patient_csv_to_xml
/opt/patient_csv_to_xml/upload.gtpl
%dir /opt/patient_csv_to_xml
%doc

%pre
getent group patient_csv_to_xml >/dev/null || groupadd -r patient_csv_to_xml
getent passwd patient_csv_to_xml >/dev/null || \
    useradd -r -g patient_csv_to_xml -d /opt/patient_csv_to_xml -s /sbin/nologin \
    -c "User to run patient_csv_to_xml" patient_csv_to_xml
exit 0
%post
chown -R patient_csv_to_xml:patient_csv_to_xml /opt/patient_csv_to_xml

%changelog
* Tue Aug 21 2018 Jonathan Mainguy <jon@soh.re> - 0.4.2-1
- Edit physType only to show if other Physican stuff is there

* Mon Aug 20 2018 Jonathan Mainguy <jon@soh.re> - 0.4-1
- Added physican and patsex

* Mon Aug 20 2018 Jonathan Mainguy <jon@soh.re> - 0.3-1
- Added ClaDiagnosis1 and ClaICDIndicator

* Mon Apr 23 2018 Jonathan Mainguy <jon@soh.re> - 0.2-1
- Added patClass

* Mon Apr 16 2018 Jonathan Mainguy <jon@soh.re> - 0.1-2
- Fixed url

* Mon Apr 16 2018 Jonathan Mainguy <jon@soh.re> - 0.1-1
- Initial release

