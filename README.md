# configs-saver CLI tool

## Overview

I work on two systems: WSL2 and macOS and I have many different configuration files.
For security, I decided to save them somewhere in the cloud.
So, for simplicity, I decided to make a CLI tool in golang (actually, I just wanted to make some pet-project in golang since I'm starting to learn it)
## Table of Contents

- [Installation](#installation)
- [Usage](#usage)

## Installation

* Be sure you have golang, try `go version`, if there are no command, install latest golang [here](https://go.dev/doc/install)

```bash
   1. git clone https://github.com/timur-khakhalev/configs-saver.git
   2. cd configs-saver
   3. go get .
   4. go build -o bin/configs-saver main.go
   # configs-saver built for your platform and saved to ./bin/configs-saver
   # hint: if you need to build this tool for another platform, 
   # now you need only add this binary to /usr/local/bin (if mac)
   # for me its:
   
   sudo cp bin/configs-saver /usr/local/configs-saver
   ln -sf /usr/local/configs-saver /usr/local/bin/configs-saver
   
   # now just open new terminal window and you are welcome
   ```

## Usage
* Create *.ini configs file according to configs-saver.ini.example (or just rename this one, file type should be *.ini)
* Save this file where you want (for example, ~/.configs-saver.ini)
* Do not forget to specify [AWS Credentials](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) in this config file.
  * There are not only AWS S3 supported, this tool can interact with any S3-like services, you only need to specify credentials, endpoint url and region.

```bash
  configs-saver -c ~/.configs-saver.ini
  
  # Archive successfully uploaded to s3. Object key: unix-configs/mac/mba13_19-58-47_14-1-2024.tar.gz
  ```
