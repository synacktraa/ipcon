<h1 align="center">
<br>
<img src="https://i.imgur.com/bpxQTm8.png" alt="ipcon">
<br>
<h3 align="center">Converts IP addresses to decimal and octal formats which helps in bypassing the black-lists and regex checks</h3>

## Installation

```sh
git clone https://github.com/synacktraa/ipcon
cd ./ipcon && go build
sudo mv ./ipcon /usr/local/bin
cd ../ && rm -rf ./ipcon
```

## Usage

####  IPv4 address to Decimal
```sh
$ ipcon 192.168.2.69
3232236101
```

####  Decimal to IPv4 address
```sh
$ ipcon 168430350  
10.10.11.14
