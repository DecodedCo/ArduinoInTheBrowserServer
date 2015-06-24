
source ~/.bash_profile

#$1 needs to be the user ID
mkdir ~/arduinocode
cd ~/arduinocode
ino init -t blink
cp -rf /srv/codefiles/$USERID.ino ~/arduinocode/src/sketch.ino
cp -rf /srv/codefiles/libraries ~/arduinocode/lib/
ino build
cp ~/arduinocode/.build/uno/firmware.hex /srv/codefiles/$USERID.hex
