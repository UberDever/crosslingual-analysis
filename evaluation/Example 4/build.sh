# build
c++ -shared lib.cpp -o lib.so

# deploy
if [ -d lib_dir ]; then
    rm -rf lib_dir
fi
mkdir lib_dir
mv lib.so lib_dir