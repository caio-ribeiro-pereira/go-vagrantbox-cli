# GoVagrantBox

A CLI tool written in Go which list and install vagrant boxes from http://www.vagrantbox.es.

## How to use

#### Instalation

``` bash
go get github.com/caio-ribeiro-pereira/govagrantbox
```

#### Running

``` bash
govagrantbox
```

*Obs.:* This command will list all existing vagrant boxes and you just need to choose one typing the box id and setting a title for it.

#### Example

``` bash
Listing all vagrant boxes...
ID: 0
Description: Debian 7.3.0 64-bit Puppet 3.4.1 (Vagrant 1.4.0)
Size: 682 MB | Provider: VirtualBox 4.3.6
Link: https://dl.dropboxusercontent.com/u/29173892/vagrant-boxes/debian7.3.0-vbox4.3.6-puppet3.4.1.box
==================================================
ID: 1
Description: OpenBSD 5.5 64-bit + Chef 11.16.0 + Puppet 3.4.2
Size: 315 MB | Provider: VirtualBox
Link: https://github.com/jose-lpa/veewee-openbsd/releases/download/v0.5.5/openbsd55.box
==================================================
ID: 2
Description: OpenBSD 5.4 64-bit + Chef 11.8.2 (150GB HDD)
Size: 1800 MB | Provider: VirtualBox
Link: http://vagrant.inagile.org/vagrant-obsd54-amd64.box
==================================================
...
# Type here the vagrant box ID
Choose vagrant box id to install: 2 
Your vagrant box is:
OpenBSD 5.4 64-bit + Chef 11.8.2 (150GB HDD)
Size: 1800 MB | Provider: VirtualBox
Link: http://vagrant.inagile.org/vagrant-obsd54-amd64.box
# Set a vagrant box title
Vagrant box title: myvagrant 
==================================================
# After choose a box the instalation will run automaticaly...
Installing choosed vagrant box...
vagrant box add myvagrant http://vagrant.inagile.org/vagrant-obsd54-amd64.box
vagrant init myvagrant
vagrant up
Your box is up and running...see ya!
```

### Author

Caio Ribeiro Pereira - <caio.ribeiro.pereira@gmail.com>  
MIT License <http://caio-ribeiro-pereira.mit-license.org>
