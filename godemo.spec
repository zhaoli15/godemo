Name:       godemo
Version:    1.0
Release:    325789
Summary:    load godemo config manager
License: GPL


URL:  github.com/zhaoli15/godemo

%description
load godemo

%define godemo_srcdir ${GOPATH}/src/godemo
%define godemo_confdir ${GOPATH}/src/godemo/conf
%define godemo_homedir /opt/app/godemo
%define godemo_binname godemo

%prep

%build
export MAKE=%{__make}
cd %{godemo_srcdir} && make all

%install
rm -rf %{buildroot}

install -d -m 755 %{buildroot}/%{godemo_homedir}/bin/
install -d -m 755 %{buildroot}/%{godemo_homedir}/conf/
install -d -m 755 %{buildroot}/%{godemo_homedir}/logs/
install -d -m 755 %{buildroot}/%{_initddir}/

install -p -D -m 755 ${GOPATH}/bin/%{godemo_binname} \
%{buildroot}/%{godemo_homedir}/bin/%{name}
install -p -D -m 755 %{godemo_srcdir}/conf/config.json \
%{buildroot}/%{godemo_homedir}/conf/config.json
install -p -D -m 755 %{godemo_srcdir}/%{name}.service \
%{buildroot}/%{_initddir}/%{name}

%post
/sbin/chkconfig --add %{name}

%preun
if [ $1 = 0 ];then
    /sbin/service %{name} stop
    /sbin/chkconfig --del %{name}
fi

%postun
if [ "$1" -ge "1" ];then
    /sbin/service %{name} restart > /dev/null 2>&1 
fi


%clean
[ "%{buildroot}" != "/" ] && %{__rm} -rf %{buildroot}

%files
%{godemo_homedir}/bin/%{name}
%config(noreplace) %{godemo_homedir}/conf/config.json
%dir %attr(0755,root,root) %{godemo_homedir}/logs/
%{_initddir}/%{name}
