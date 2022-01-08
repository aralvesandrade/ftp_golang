#!/bin/bash
echo "Install vsftpd"
sudo apt-get update
sudo apt-get install vsftpd
echo "Restart vsftpd"
sudo systemctl restart vsftpd
echo "Done!"