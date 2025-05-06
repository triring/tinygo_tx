<<<<<<< HEAD
# go mod init tinygo_tx
# go get tinygo.org/x/drivers

mkdir uf2
mkdir uf2/demo
mkdir uf2/music
mkdir uf2/effects

tinygo build -o ./Templates.uf2                   -target=pico -size short ./Templates/
tinygo build -o uf2/demo/AM999KHz.uf2             -target=pico -size short ./demo/sendAM
tinygo build -o uf2/demo/FM083_4MHZ.uf2           -target=pico -size short ./demo/sendFM
tinygo build -o uf2/demo/FM093_8MHZ.uf2           -target=pico -size short ./demo/sendFM
tinygo build -o uf2/demo/FM104_1MHZ.uf2           -target=pico -size short ./demo/sendFM

tinygo build -o uf2/music/CheCheKoolay.uf2        -target=pico -size short ./music/CheCheKoolay
tinygo build -o uf2/music/Famima.uf2              -target=pico -size short ./music/Famima
tinygo build -o uf2/music/Ievan_polkka.uf2        -target=pico -size short ./music/Ievan_polkka
tinygo build -o uf2/music/KimiGaYo.uf2            -target=pico -size short ./music/KimiGaYo
tinygo build -o uf2/music/Ode_to_Joy.uf2          -target=pico -size short ./music/Ode_to_Joy
tinygo build -o uf2/music/ShortShorts.uf2         -target=pico -size short ./music/ShortShorts
tinygo build -o uf2/music/WelcomeJapariPark.uf2   -target=pico -size short ./music/WelcomeJapariPark

tinygo build -o uf2/effects/Beacon.uf2            -target=pico -size short ./effects/Beacon
tinygo build -o uf2/effects/CTUringtone.uf2       -target=pico -size short ./effects/CTUringtone
tinygo build -o uf2/effects/Jihou.uf2             -target=pico -size short ./effects/Jihou
tinygo build -o uf2/effects/Kingyo.uf2            -target=pico -size short ./effects/Kingyo
tinygo build -o uf2/effects/PiPo.uf2              -target=pico -size short ./effects/PiPo
tinygo build -o uf2/effects/Ramen.uf2             -target=pico -size short ./effects/Ramen
tinygo build -o uf2/effects/Saodake.uf2           -target=pico -size short ./effects/Saodake
tinygo build -o uf2/effects/ThiroriSound.uf2      -target=pico -size short ./effects/ThiroriSound
tinygo build -o uf2/effects/Ultraman.uf2          -target=pico -size short ./effects/Ultraman
tinygo build -o uf2/effects/WestminsterChimes.uf2 -target=pico -size short ./effects/WestminsterChimes

# pandoc -f markdown -t html -o README.html README.md
# pandoc README.md -o README.html
# copy ./demo/sendFM/main.go ./Templates/
# copy ./demo/sendFM/main.go ./music/CheCheKoolay/
# copy ./demo/sendFM/main.go ./music/Famima/
# copy ./demo/sendFM/main.go ./music/Ievan_polkka/
# copy ./demo/sendFM/main.go ./music/KimiGaYo/
# copy ./demo/sendFM/main.go ./music/Ode_to_Joy/
# copy ./demo/sendFM/main.go ./music/ShortShorts/
# copy ./demo/sendFM/main.go ./music/WelcomeJapariPark/
# copy ./demo/sendFM/main.go ./effects/CTUringtone/
# copy ./demo/sendFM/main.go ./effects/Jihou/
# copy ./demo/sendFM/main.go ./effects/Kingyo/
# copy ./demo/sendFM/main.go ./effects/PiPo/
# copy ./demo/sendFM/main.go ./effects/Ramen/
# copy ./demo/sendFM/main.go ./effects/Saodake/
# copy ./demo/sendFM/main.go ./effects/ThiroriSound/
# copy ./demo/sendFM/main.go ./effects/Ultraman/
# copy ./demo/sendFM/main.go ./effects/WestminsterChimes/


=======
# go mod init tinygo_tx
mkdir uf2
mkdir uf2/demo
mkdir uf2/music
mkdir uf2/effects

tinygo build -o ./Templates.uf2                   -target=pico -size short ./Templates/
tinygo build -o uf2/demo/AM999KHz.uf2             -target=pico -size short ./demo/sendAM
tinygo build -o uf2/demo/FM083_4MHZ.uf2           -target=pico -size short ./demo/sendFM
tinygo build -o uf2/demo/FM093_8MHZ.uf2           -target=pico -size short ./demo/sendFM
tinygo build -o uf2/demo/FM104_1MHZ.uf2           -target=pico -size short ./demo/sendFM

tinygo build -o uf2/music/CheCheKoolay.uf2        -target=pico -size short ./music/CheCheKoolay
tinygo build -o uf2/music/Famima.uf2              -target=pico -size short ./music/Famima
tinygo build -o uf2/music/Ievan_polkka.uf2        -target=pico -size short ./music/Ievan_polkka
tinygo build -o uf2/music/KimiGaYo.uf2            -target=pico -size short ./music/KimiGaYo
tinygo build -o uf2/music/Ode_to_Joy.uf2          -target=pico -size short ./music/Ode_to_Joy
tinygo build -o uf2/music/ShortShorts.uf2         -target=pico -size short ./music/ShortShorts
tinygo build -o uf2/music/WelcomeJapariPark.uf2   -target=pico -size short ./music/WelcomeJapariPark

tinygo build -o uf2/effects/CTUringtone.uf2       -target=pico -size short ./effects/CTUringtone
tinygo build -o uf2/effects/Jihou.uf2             -target=pico -size short ./effects/Jihou
tinygo build -o uf2/effects/Kingyo.uf2            -target=pico -size short ./effects/Kingyo
tinygo build -o uf2/effects/PiPo.uf2              -target=pico -size short ./effects/PiPo
tinygo build -o uf2/effects/Ramen.uf2             -target=pico -size short ./effects/Ramen
tinygo build -o uf2/effects/Saodake.uf2           -target=pico -size short ./effects/Saodake
tinygo build -o uf2/effects/ThiroriSound.uf2      -target=pico -size short ./effects/ThiroriSound
tinygo build -o uf2/effects/Ultraman.uf2          -target=pico -size short ./effects/Ultraman
tinygo build -o uf2/effects/WestminsterChimes.uf2 -target=pico -size short ./effects/WestminsterChimes

# pandoc -f markdown -t html -o README.html README.md
# pandoc README.md -o README.html
# copy ./demo/sendFM/main.go ./Templates/
# copy ./demo/sendFM/main.go ./music/CheCheKoolay/
# copy ./demo/sendFM/main.go ./music/Famima/
# copy ./demo/sendFM/main.go ./music/Ievan_polkka/
# copy ./demo/sendFM/main.go ./music/KimiGaYo/
# copy ./demo/sendFM/main.go ./music/Ode_to_Joy/
# copy ./demo/sendFM/main.go ./music/ShortShorts/
# copy ./demo/sendFM/main.go ./music/WelcomeJapariPark/
# copy ./demo/sendFM/main.go ./effects/CTUringtone/
# copy ./demo/sendFM/main.go ./effects/Jihou/
# copy ./demo/sendFM/main.go ./effects/Kingyo/
# copy ./demo/sendFM/main.go ./effects/PiPo/
# copy ./demo/sendFM/main.go ./effects/Ramen/
# copy ./demo/sendFM/main.go ./effects/Saodake/
# copy ./demo/sendFM/main.go ./effects/ThiroriSound/
# copy ./demo/sendFM/main.go ./effects/Ultraman/
# copy ./demo/sendFM/main.go ./effects/WestminsterChimes/


>>>>>>> origin/main
