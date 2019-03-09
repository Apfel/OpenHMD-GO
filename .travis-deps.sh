echo "Installing Dependencies"
sudo apt update
sudo apt install libudev-dev libusb-1.0-0-dev libtool -y

echo "Cloning repositories"
git clone https://github.com/OpenHMD/OpenHMD.git
git clone https://github.com/OpenHMD/hidapi.git

echo "Installing HIDAPI"
cd hidapi
autoreconf --install --verbose --force && ./configure
make && sudo make install
cd ..

echo "Installing OpenHMD"
cd OpenHMD
autoreconf --install --verbose --force && ./configure
make && sudo make install
cd ..

rm -rf hidapi OpenHMD
