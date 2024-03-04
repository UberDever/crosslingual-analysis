
.DEFAULT_GOAL := default

.clean:
	rm -f *.so #Makefile:5

.build:
	gfortran -shared -fPIC -o lib.so lib.f90 #Makefile:8
	f2py3 -c -m compute_lib lib.f90 #Makefile:9

.run:
	python3 compute.py #Makefile:12

default: .clean .build .run
