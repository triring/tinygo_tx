# go mod init tinygo_tx
mkdir uf2
mkdir uf2/music
mkdir uf2/effects

tinygo build -o uf2/music/Famima.uf2             -target=pico -size short ./music/Famima
tinygo build -o uf2/music/Ievan_polkka.uf2       -target=pico -size short ./music/Ievan_polkka
tinygo build -o uf2/music/KimiGaYo.uf2           -target=pico -size short ./music/KimiGaYo
tinygo build -o uf2/music/Ode_to_Joy.uf2         -target=pico -size short ./music/Ode_to_Joy
tinygo build -o uf2/music/ShortShorts.uf2        -target=pico -size short ./music/ShortShorts
tinygo build -o uf2/music/WelcomeJapariPark.uf2  -target=pico -size short ./music/WelcomeJapariPark

tinygo build -o uf2/effects/CTUringtone.uf2       -target=pico -size short ./effects/CTUringtone
tinygo build -o uf2/effects/Jihou.uf2             -target=pico -size short ./effects/Jihou
tinygo build -o uf2/effects/PiPo.uf2              -target=pico -size short ./effects/PiPo
tinygo build -o uf2/effects/ThiroriSound.uf2      -target=pico -size short ./effects/ThiroriSound
tinygo build -o uf2/effects/Ultraman.uf2          -target=pico -size short ./effects/Ultraman
tinygo build -o uf2/effects/WestminsterChimes.uf2 -target=pico -size short ./effects/WestminsterChimes

# pandoc README.md -o README.html