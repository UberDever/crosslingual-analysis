
.DEFAULT_GOAL := default

.clean:
	rm *.so

.build:
	gfortran -shared -fPIC -o lib.so lib.f90
	f2py3 -c -m compute_lib lib.f90

.run:
	python3 compute.py

default: .clean .build .run
