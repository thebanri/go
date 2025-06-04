#!/bin/bash


go build -o task 
if [ $? -ne 0 ]; then
    echo "Build failed. Please check the Go code for errors."
    exit 1
fi

sudo mkdir -p /etc/taskcli
sudo touch /etc/taskcli/tasks.json
sudo chmod 777 /etc/taskcli/tasks.json
if [ $? -ne 0 ]; then
    echo "Failed to copy tasks.json to ~/.config/taskcli/. Please check if the directory exists."
    exit 1
fi
sudo cp task /usr/local/bin/
if [ $? -ne 0 ]; then
    echo "Failed to copy the binary to /usr/local/bin/. Please check your permissions."
    exit 1
fi
echo "Installation successful! You can now run 'task' from anywhere."
