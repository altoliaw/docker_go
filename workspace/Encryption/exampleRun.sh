# ##
# Compiling the static c libraries and copying the static header files and libraries to
# the proper location; after the previous compilation, the golang module will be executed
#
# @author Nick
# @date 2024/06/20
# ##
cd /home/SecuEyesDecoder
make clean
git pull -f
make archiveDLL
cd /programs/Encryption
rm -rf Libs
rm -rf Headers
mkdir -p Libs
cp /home/SecuEyesDecoder/Outputs/DOT_A/libjsonDecoder.a /programs/Encryption/Libs/libjsonDecoder.a && \
cp -pR /home/SecuEyesDecoder/Headers  /programs/Encryption


cd /programs/Encryption
go run Main.go
