# Setting up an IDE Setup on CodeAnywhere

Starting with a Blank Development Stack with Ubuntu 16.04 and SSH into the container

## Install Node.js and NPM

NodeSource maintains a repository containing the latest version of Node.js and npm. Using NodeSource gives a fast and easy way to install the pecific version of Node that you need. For this purpose, we are going to be installing Node.Js v12.x. 

Opening up an SSH Terminal, add the NodeSource repository by running the following command. This will add the NodeSource signing key to your system, create an apt sources repository file, install all necessary packages and finally refresh the apt cache.

```
$ curl -sL https://deb.nodesource.com/setup_12.x | sudo -E bash -
```

Once the NodeSource repository is setup, install Node.js and npm:
```
$ sudo apt install nodejs
```

Now you can verify that Node.js and npm were successfully installed by printing their version
```
$ node -v
v12.14.1
$ npm -v
6.13.4
```

## Configure npm package root

Now that NPM is installed, if we tried to install Angular CLI using npm. If you just run the command, you will most likely get an Missing write access error. Of course you can just run the command with sudo, and it will work. but it's not the correct way. Npm modules should not be install with sudo or you might run into other issues in the future. So one way to get around this issue is to tell NPM to use a different package root in your homedir to hold global packages.

First create this new global packages folder.
```
$ NPM_PACKAGES="$HOME/.npm-packages"
$ mkdir -p "$NPM_PACKAGES"
```

Then set NPM to use that directory for global package installs.
```
$ echo "prefix = $NPM_PACKAGES" >> ~/.npmrc
```

Then configure your PATH and MANPATH to see commands in the new global packages folder by adding the following to your .bashrc.
```
$ vi ~/.bashrc
```
And insert the following lines

```
# NPM packages in homedir
NPM_PACKAGES="$HOME/.npm-packages"
   
# Tell our environment about user-installed node tools
PATH="$NPM_PACKAGES/bin:$PATH"
# Unset manpath so we can inherit from /etc/manpath via the `manpath` command
unset MANPATH  # delete if you already modified MANPATH elsewhere in your configuration
MANPATH="$NPM_PACKAGES/share/man:$(manpath)"
   
# Tell Node about these packages
NODE_PATH="$NPM_PACKAGES/lib/node_modules:$NODE_PATH"
```

And then reload your .bashrc
```
$ . ~/.bashrc
```

Now when you do an npm install -g, NPM will install the libraries into ~/.npm-packages/lib/node_modules link executable tools into ~/.npm-packages/bin, which is in your PATH. 

## Install Angular CLI

Install Angular CLI is now pretty easy using npm.
```
$ npm install -g @angular/cli
```
And verify install by running
```
$ ng version
Angular CLI: 8.3.22
...
```

## Install GoLang

The version of Go that is being installed is 1.11. The reason for not using 1.12 or 1.13 is compatability issue with the Google App Engine. 
```
$ wget https://dl.google.com/go/go1.11.13.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go1.11.13.linux-amd64.tar.gz
$ export PATH=$PATH:/usr/local/go/bin
$ rm go1.11.13.linux-amd64.tar.gz

$ go version
go version go1.11.13 linux/amd64
```

Remember to add the PATH export to your .bashrc file as well
```
# Set PATH for Go
export PATH=$PATH:/usr/local/go/bin
```

## Install Google Could SDK

Finally, installing Google Cloud SDK along with Google App Engine for Go, and the Cloud Datastore Emulator.
```
$ echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
$ sudo apt-get install apt-transport-https ca-certificates gnupg
$ curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
$ sudo apt-get update && sudo apt-get install google-cloud-sdk
$ sudo apt-get install google-cloud-sdk-app-engine-go
$ sudo apt-get install google-cloud-sdk-datastore-emulator

$ gcloud -v
Google Cloud SDK 275.0.0
alpha 2020.01.03
app-engine-go
app-engine-python 1.9.88
beta 2020.01.03
bq 2.0.51
cloud-datastore-emulator 2.1.0
core 2020.01.03
gsutil 4.46
kubectl 2020.01.03
```

This IDE should now be able to install the Project Base and run it on a development Google App Engine.

*Back to [README.md](README.md)*