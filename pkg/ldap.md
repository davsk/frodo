# frodo install ldap

## console commands
``` bash
sudo eopkg it -c system.devel
sudo eopkg it openldap openldap-devel
systemctl status vboxdrv.service
journalctl -xe
cd ~/Downloads/
mkdir ~/tmp
sudo tar zxvf nss-pam-ldapd-0.9.11.tar.gz -C ~/tmp/ && cd ~/tmp/nss-pam-ldapd-0.9.11
./configure --prefix=/usr && make
sudo make install
sudo nano /etc/nslcd.conf 
sudo useradd --system nslcd
sudo nano /etc/nsswitch.conf 
sudo /usr/sbin/nslcd -d
sudo nano /etc/systemd/system/nslcd.service
sudo systemctl daemon-reload
sudo systemctl enable nslcd.service
sudo systemctl start nslcd.service
sudo journalctl -u nslcd.service
sudo su
```
