tinygo build -o sendAM.uf2 -target=pico -size short .\sendFM\
tinygo flash -target=pico -size short .\sendFM\

tinygo build -o sendAM.uf2 -target=pico -size short .\sendAM\
tinygo flash -target=pico -size short .\sendAM\

tinygo build -o RadioPicoAM999KHz.uf2  -target=pico -size short .\RadioPicoAM999KHz\
tinygo flash -target=pico -size short .\RadioPicoAM999KHz\

tinygo build -o RadioPicoFM83_4MHz.uf2 -target=pico -size short .\RadioPicoFM83_4MHz\
tinygo flash -target=pico -size short .\RadioPicoFM83_4MHz\
