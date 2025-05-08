tinygo build -o .\sendFM\RadioPicoFM93_8MHz.uf2 -target=pico -size short .\sendFM\
tinygo flash -target=pico -size short .\sendFM\

tinygo build -o .\sendAM\RadioPicoAM999KHz.uf2 -target=pico -size short .\sendAM\
tinygo flash -target=pico -size short .\sendAM\
