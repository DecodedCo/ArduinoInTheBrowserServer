# The Playduino server

* Thanks to Mr Jones who inspired this https://github.com/mrjones/Chrome-Arduino
* The chrome plugin that goes with this project can be found here: https://github.com/DecodedCo/ArduinoInTheBrowser
* This project got closed down at Decoded and so I no longer maintain it at work, however I think its cool so Im more than happy to continue looking after it. With community help ;)
* Im alex at Decoded dot com if anyone wants me....

* The docker scripts handle this, but:
* It relies on the Arduino (v1.1 - 1.5 doesnt work with ino) being installed, 
* It also needs [ino](http://inotool.org/) to be installed and on the path.
* Arduino will need Java installed.
* This is written in Go. It has two major end points, `program` and `hex`.
* Each user, for each compilation is given a UUID. This is used to distinguish their code from someone elses.
* `program` is the end point hit when the playduino plugin "Compile" button is pressed. It takes the code, 
prefixes it with the necessary libraries, sets up a temporary directory for the user, then builds the code
and creates a hex file.
* `hex` endpoint is the endpoint that is hit when the "Upload" button is pressed in the Playduino plugin.
* It returns the user's hex file based on their UUID.

### Setup

* Im assuming docker is installed on the server

* I have included the dockerfiles I used to create the docker image in dockerfiles.tar.gz

* To get it right it can be a bit tricky.
* The problem I had to get over was concurrent builds. The only way I truly solved this was with docker containers to do the compiling.
* First in your home directory add a script called `command.sh` - this holds the command to start a docker container:

```
docker run -e USERID=$1 -t --volumes-from CODE --name $1 playduino:latest bash /build/buildapplication.sh
```
* Leave that there. The application will call it when a new compilation is requested.
* That will only work if there is already a docker volume container called `CODE`
* You need to run this on your server:

```
docker run -v /srv/codefiles:/srv/codefiles --name CODE centos:7 true
```

* Instead of the above, I tend to have a script in my home directory called setupPlayduino.sh which includes:

```
sudo setenforce 0
sudo service docker start
docker ps -a
docker rm CODE
docker run -v /srv/codefiles:/srv/codefiles --name CODE centos:7 true
```

* But this is just for ease. The first line is a security threat however, so if you don't understand it, please be careful

* That will only be successful if there is already a directory at `/srv/codefiles`. So create that directory. it needs to be owned by your user with read/write permissions.
* Inside there you need to put the `libraries.zip` file that comes with this repository, and unzip it. Leave the zipped version there also.

* On starting the server you will need to run:
* Run `./setupScript.sh`
* Start a screen
`screen`
* Start the playduino server
`revel run playduinoserver`
* Detach from screen: `CTRL-A D`
* Start nginx
`sudo service nginx start`

Go to your webbrowser and test going to your server url. You should see

~~~
{
	message: "nothing at this location",
	result: "error 404",
	identity: "NIL"
}
~~~

That's good. Now you can compile code.
* Leave the server safely:
* now type `exit` to leave the server


### Plugin

* Pull the repository of the plugin
* You will need to edit serialmonitor.js - see https://github.com/DecodedCo/ArduinoInTheBrowser

# License

The MIT License (MIT)

Copyright (c)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
