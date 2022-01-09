#!/bin/bash
echo "Install vsftpd"
sudo apt-get update
sudo apt-get install vsftpd

echo "Create dir ftp/arquivos"
sudo mkdir -p /var/ftp/arquivos
sudo chown nobody:nogroup /var/ftp/arquivos
sudo chmod 777 -R /var/ftp/arquivos
 
echo "Create file sample.txt"
echo "vsftpd test file" | sudo tee /var/ftp/arquivos/sample.txt

echo "Alter file vsftpd.conf"
sudo chmod 777 /etc/vsftpd.conf
echo '#
anonymous_enable=YES
local_enable=NO
anon_root=/var/ftp/
no_anon_password=YES
hide_ids=YES
pasv_min_port=40000
pasv_max_port=50000
write_enable=YES
anon_upload_enable=YES
anon_mkdir_write_enable=YES' >> /etc/vsftpd.conf

echo "Restart vsftpd"
sudo systemctl restart vsftpd

echo "Done!"

#sudo apt-get remove --auto-remove vsftpd
#sudo apt-get purge --auto-remove vsftpd