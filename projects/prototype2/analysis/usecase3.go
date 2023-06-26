package analysis

import (
	"os"
	"prototype2/sexpr"
	"prototype2/util"
)

// Analysis of C, bash and (internally) preprocessor, see usecase3 files

func Usecase3_Analyzer() []fragment {
	fragments := make([]fragment, 0, 128)
	fragments = append(fragments, analyzeC2()...)
	fragments = append(fragments, analyzeBash2()...)
	return fragments
}

func analyzeC2() []fragment {
	fragments := make([]fragment, 0, 128)
	S := sexpr.S

	home, _ := os.UserHomeDir()
	cwd := util.ShortenPath(home+"/dev/mag/language-analysis/projects/prototype2/usecases/usecase3", 2)

	libC := NewExport(
		S("URI"),
		S(cwd+"/lib.c"),
		S("URI:0:0:lib.c"),
	)
	f := NewExport(
		Function("Unit", "Int"),
		S("f"),
		S("FunctionDecl:3:5:lib.c"),
	)
	g := NewExport(
		Function("Unit", "Int"),
		S("g"),
		S("FunctionDecl:9:5:lib.c"),
	)
	iVar := NewImport(
		S("String"),
		S("VAR"),
		S("PreprocessorIf:1:8:lib.c"),
	)

	mainC := NewExport(
		S("URI"),
		S(cwd+"/main.c"),
		S("URI:0:0:main.c"),
	)
	iF := NewImport(
		Function("Unit", "Int"),
		S("f"),
		S("FunctionDecl:1:5:main.c"),
	)
	eMain := NewExport(
		Function("Int", S("List", S("List", "Int")), "Int"),
		S("main"),
		S("FunctionDecl:3:5:main.c"),
	)

	fragments = append(fragments,
		fragment{
			path:         cwd + "/lib.c(2:6);main.c",
			priority:     1,
			lang:         "C",
			environments: []environment{iVar, iF},
			signatures:   []signature{f, libC, mainC, eMain},
			intralinks: []struct {
				from statement
				to   statement
			}{
				{from: eMain.statement,
					to: iF.statement,
				},
			},
		},
		fragment{
			path:         cwd + "/lib.c(8:12);main.c",
			priority:     0,
			lang:         "C",
			environments: []environment{iF},
			signatures:   []signature{g, libC, mainC, eMain},
			intralinks: []struct {
				from statement
				to   statement
			}{
				{from: eMain.statement,
					to: iF.statement,
				},
			},
		},
	)

	return fragments
}

func analyzeBash2() []fragment {
	fragments := make([]fragment, 0, 128)
	S := sexpr.S

	home, _ := os.UserHomeDir()
	cwd := util.ShortenPath(home+"/dev/mag/language-analysis/projects/prototype2/usecases/usecase3", 2)

	buildSh := NewExport(
		S("URI"),
		S(cwd+"/build.sh"),
		S("URI:0:0:build.sh"),
	)
	eVar := NewExport(
		S("String"),
		S("VAR"),
		S("CommandFlag:2:5:build.sh"),
	)
	iLibC := NewImport(
		S("URI"),
		S(cwd+"/lib.c"),
		S("Command:2:10:build.sh"),
	)
	iMainC := NewImport(
		S("URI"),
		S(cwd+"/main.c"),
		S("Command:2:16:build.sh"),
	)

	fragments = append(fragments,
		fragment{
			path:         cwd + "/build.sh",
			priority:     0,
			lang:         "Sh",
			environments: []environment{iLibC, iMainC},
			signatures:   []signature{buildSh, eVar},
			intralinks:   nil,
		},
	)

	return fragments
}
