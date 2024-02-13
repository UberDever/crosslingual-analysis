# build
cc -shared lib.c -o lib.so
# deploy
mkdir lib_dir
mv lib.so lib_dir