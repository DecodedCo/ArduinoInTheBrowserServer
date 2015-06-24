docker stop $(docker ps -a | grep playduino | egrep "hours|minutes ago" | awk {'print $1'})
docker rm  $(docker ps -a | grep playduino | egrep "hours|minutes ago" | awk {'print $1'})
find /srv/codefiles -regex ".*\.\(ino\|hex\)" -mmin +5 -exec rm -f {} \;
