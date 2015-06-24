#to be run as root

#download java
wget --no-cookies --no-check-certificate --header "Cookie: gpw_e24=http%3A%2F%2Fwww.oracle.com%2F; oraclelicense=accept-securebackup-cookie" "http://download.oracle.com/otn-pub/java/jdk/7u75-b13/jdk-7u75-linux-x64.tar.gz"

tar -zxf jdk-7u75-linux-x64.tar.gz -C /opt

rm /jdk-7u75-linux-x64.tar.gz

#GO IS NOT BEING INSTALLED **********
#install go
#wget -O go1.4.2.tar.gz https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz
#tar -zxf go1.4.2.tar.gz -C /usr/local

#rm /go1.4.2.tar.gz

#arduino
wget -O arduino.tgz http://arduino.cc/download.php?f=/arduino-1.0.6-linux64.tgz
tar -zxvf arduino.tgz
mv arduino-1.0.6/ /usr/local/share/arduino

rm /arduino.tgz

#ino
pip install ino
