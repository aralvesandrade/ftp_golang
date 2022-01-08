#!/bin/bash
echo "Install vsftpd"
sudo apt-get update
sudo apt-get install vsftpd
sudo cp /etc/vsftpd.conf /etc/vsftpd.conf.original
sudo chmod 777 /etc/vsftpd.conf
sudo adduser ftpuser
sudo mkdir /home/ftpuser/ftp
sudo chown nobody:nogroup /home/ftpuser/ftp
sudo chmod a-w /home/ftpuser/ftp
sudo mkdir /home/ftpuser/ftp/arquivos
sudo chown ftpuser:ftpuser /home/ftpuser/ftp/arquivos
echo "vsftpd sample file" | sudo tee /home/ftpuser/ftp/arquivos/sample.txt
echo '#
#Custom
anonymous_enable=NO
local_enable=YES
write_enable=YES
chroot_local_user=YES
user_sub_token=$USER
local_root=/home/$USER/ftp
pasv_min_port=40000
pasv_max_port=50000' >> /etc/vsftpd.conf
echo "Restart vsftpd"
sudo systemctl restart vsftpd
echo "Done!"