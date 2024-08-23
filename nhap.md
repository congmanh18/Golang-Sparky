Cấu hình tệp Hosts

Tệp "/etc/hosts" phải chứa tên máy đầy đủ (FQDN) cho máy chủ.

<địa-chỉ-IP> <tên-máy-đầy-đủ> <tên-máy>

Ví dụ:
127.0.0.1       localhost localhost.localdomain localhost4 localhost4.localdomain4
192.168.56.107  ol9-19.localdomain ol9-19

Thiết lập tên máy chính xác trong tệp "/etc/hostname".

ol9-19.localdomain


Các yêu cầu trước khi cài đặt Oracle

Sử dụng gói "oracle-database-preinstall-19c" để thực hiện tất cả các thiết lập yêu cầu trước, chạy lệnh sau:

dnf install -y oracle-database-preinstall-19c

Nên chạy cập nhật toàn bộ hệ thống, tuy không bắt buộc:

dnf update -y

Tạo nhóm và người dùng

1. Tạo các nhóm mới:

groupadd -g 54321 oinstall groupadd -g 54322 dba groupadd -g 54323 oper

2. Tạo người dùng:

useradd -u 54321 -g oinstall -G dba,oper oracle

3. Đặt mật khẩu cho người dùng "oracle":

passwd oracle

4. Thiết lập SELinux ở chế độ permissive bằng cách chỉnh sửa tệp "/etc/selinux/config":

SELINUX=permissive

Sau khi chỉnh sửa, khởi động lại hệ thống hoặc chạy lệnh sau:

setenforce Permissive

5. Nếu tường lửa Linux được bật, bạn cần tắt hoặc cấu hình nó:

systemctl stop firewalld systemctl disable firewalld


6. Tạo các thư mục cài đặt Oracle:

mkdir -p /u01/app/oracle/product/19.0.0/dbhome_1 mkdir -p /u02/oradata

chown -R oracle:oinstall /u01 /u02 chmod -R 775 /u01 /u02

7. Tạo các tệp kịch bản 

Tạo thư mục "scripts":

mkdir /home/oracle/scripts

Tạo tệp môi trường "setEnv.sh":

cat > /home/oracle/scripts/setEnv.sh <<EOF # Oracle Settings

export TMP=/tmp
export TMPDIR=$TMP

export ORACLE_HOSTNAME=ol9-19.localdomain
export ORACLE_UNQNAME=cdb1
export ORACLE_BASE=/u01/app/oracle
export ORACLE_HOME=$ORACLE_BASE/product/19.0.0/dbhome_1
export ORA_INVENTORY=/u01/app/oraInventory
export ORACLE_SID=cdb1
export PDB_NAME=pdb1
export DATA_DIR=/u02/oradata

export PATH=/usr/sbin:/usr/local/bin:$PATH
export PATH=$ORACLE_HOME/bin:$PATH

export LD_LIBRARY_PATH=$ORACLE_HOME/lib:/lib:/usr/lib
export CLASSPATH=$ORACLE_HOME/jlib:$ORACLE_HOME/rdbms/jlib EOF

8. Thêm tham chiếu đến tệp "setEnv.sh" vào cuối tệp "/home/oracle/.bash_profile":

echo ". /home/oracle/scripts/setEnv.sh" >> /home/oracle/.bash_profile

9. Tạo các tệp kịch bản "start_all.sh" và "stop_all.sh":

cat > /home/oracle/scripts/start_all.sh <<EOF

#!/bin/bash
. /home/oracle/scripts/setEnv.sh

export ORAENV_ASK=NO
. oraenv

export ORAENV_ASK=YES
dbstart $ORACLE_HOME
EOF

cat > /home/oracle/scripts/stop_all.sh <<EOF

#!/bin/bash
. /home/oracle/scripts/setEnv.sh

export ORAENV_ASK=NO
. oraenv
export ORAENV_ASK=YES

dbshut $ORACLE_HOME
EOF

10. Đặt quyền sở hữu và quyền truy cập cho các kịch bản:

chown -R oracle:oinstall /home/oracle/scripts
chmod u+x /home/oracle/scripts/\*.sh


Cài đặt Oracle
11. Chuyển đến thư mục ORACLE_HOME, giải nén phần mềm và bắt đầu Oracle Universal Installer (OUI):

export SOFTWARE_DIR=/u01/software
export OPATCH_FILE="p6880880_190000_Linux-x86-64.zip"
export PATCH_FILE="p35742413_190000_Linux-x86-64.zip"
export PATCH_TOP=${SOFTWARE_DIR}/35742413
export PATCH_PATH1=${PATCH_TOP}/35643107
export PATCH_PATH2=${PATCH_TOP}/35648110

cd $ORACLE_HOME
unzip -oq ${SOFTWARE_DIR}/LINUX.X64_193000_db_home.zip
unzip -oq ${SOFTWARE_DIR}/${OPATCH_FILE}

cd ${SOFTWARE_DIR}
unzip -oq ${SOFTWARE_DIR}/${PATCH_FILE}

cd $ORACLE_HOME
export CV_ASSUME_DISTID=OL8

./runInstaller

12. Chạy các kịch bản root khi được yêu cầu:

/u01/app/oraInventory/orainstRoot.sh
/u01/app/oracle/product/19.0.0/dbhome_1/root.sh

Tạo cơ sở dữ liệu

13. Khởi động listener:

lsnrctl start

14. Tạo cơ sở dữ liệu bằng Database Configuration Assistant (DBCA):

dbca -silent -createDatabase\-template
Name General_Purpose.dbc\-gdbname ${ORACLE_SID} -sid ${ORACLE_SID} -responseFile NO_VALUE
\-characterSet AL32UTF8
\-sysPassword SysPassword1
\-systemPassword SysPassword1
\-createAsContainerDatabase true
\-numberOfPDBs 1
\-pdbName ${PDB_NAME}
\-pdbAdminPassword PdbPassword1
\-databaseType MULTIPURPOSE
\-memoryMgmtType auto_sga
\-totalMemory 2000
\-storageType FS
\-datafileDestination "${DATA_DIR}"
\-redoLogFileSize 50
\-emConfiguration NONE
\-ignorePreReqs

## Sau cài đặt

15. Chỉnh sửa tệp "/etc/oratab" để thiết lập cờ khởi động lại cho mỗi instance:

cdb1:/u01/app/oracle/product/19.0.0/dbhome_1:Y

16. Kích hoạt Oracle Managed Files (OMF) và đảm bảo PDB khởi động cùng với instance:

sqlplus / as sysdba <<EOF
alter system set db_create_file_dest='${DATA_DIR}'; alter pluggable database ${PDB_NAME} save state; exit;
EOF