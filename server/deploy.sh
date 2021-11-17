cd ../client
yarn generate

cd ../server
CGO_ENABLED=0 go build -ldflags "-s -w" 
ssh honeyfox-pve 'systemctl stop drazil'
scp drazil honeyfox-pve:/opt/drazil/drazil
ssh honeyfox-pve 'chown drazil /opt/drazil/drazil && systemctl start drazil'