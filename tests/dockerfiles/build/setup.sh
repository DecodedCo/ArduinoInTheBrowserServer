#update bashrc:
echo "export JAVA_HOME=/opt/jdk1.7.0_75" >> ~/.bashrc
echo "export JRE_HOME=/opt/jdk1.7.0_75/jre" >> ~/.bashrc
echo "export PATH=$PATH:/opt/jdk1.7.0_75/bin:/opt/jdk1.7.0_75/jre/bin" >> ~/.bashrc

#checks
which python
python -V
which ino
cat ~/.bashrc
. ~/.bashrc
which java

#mkdir -p /tmp/inotest; cd /tmp/inotest; ino init -t blink; ino build;

echo "testing connection..."
ping -c 4 www.google.com
#mkdir ~/go/src/goapplication
