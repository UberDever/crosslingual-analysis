all: build run

build-main:
	latexmk -xelatex -synctex=1 -jobname=main main.tex

clean:
	rm *.aux \
	*.fdb_latexmk \
	*.fls \
	*.lof \
	*.lot \
	*.log \
	*.out \
	*.synctex.gz \
	*.xdv \
	*.toc

build:
	docker build -t docker-latex .
	
run:
	docker run --rm -ti -v ${PWD}:/master-thesis:Z docker-latex bash -c "git config --global --add safe.directory /master-thesis && make build-main && make clean"
	docker run --rm -ti -v ${PWD}:/master-thesis:Z docker-latex bash -c "git config --global --add safe.directory /master-thesis && make -C presentation && make -C presentation clean"