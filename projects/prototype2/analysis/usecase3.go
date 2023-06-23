package analysis

import (
	"os"
	"prototype2/sexpr"
)

// Analysis of C, bash and (internally) preprocessor, see usecase3 files

func Usecase3_Analyzer() []module {
	modules := make([]module, 0, 128)
	modules = append(modules, analyzeC2()...)
	modules = append(modules, analyzeBash2()...)
	return modules
}

func analyzeC2() []module {
	modules := make([]module, 0, 128)
	S := sexpr.S

	home, _ := os.UserHomeDir()
	cwd := home + "/dev/mag/language-analysis/projects/prototype2/usecases/usecase3"

	libC := NewExport(
		S("File"),
		S(cwd+"/lib.c"),
		S("File:0:0:lib.c"),
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
		S("File"),
		S(cwd+"/main.c"),
		S("File:0:0:main.c"),
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

	modules = append(modules,
		module{
			priority: 1,
			lang:     "C",
			imports:  []import_{iVar, iF},
			exports:  []export{f, libC, mainC, eMain},
			intralinks: []struct {
				from statement
				to   statement
			}{
				{from: eMain.statement,
					to: iF.statement,
				},
			},
		},
		module{
			priority: 0,
			lang:     "C",
			imports:  []import_{iF},
			exports:  []export{g, libC, mainC, eMain},
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

	return modules
}

func analyzeBash2() []module {
	modules := make([]module, 0, 128)
	S := sexpr.S

	home, _ := os.UserHomeDir()
	cwd := home + "/dev/mag/language-analysis/projects/prototype2/usecases/usecase3"

	buildSh := NewExport(
		S("File"),
		S(cwd+"/build.sh"),
		S("File:0:0:build.sh"),
	)
	eVar := NewExport(
		S("String"),
		S("VAR"),
		S("CommandFlag:2:5:build.sh"),
	)
	iLibC := NewImport(
		S("File"),
		S(cwd+"/lib.c"),
		S("Command:2:10:build.sh"),
	)
	iMainC := NewImport(
		S("File"),
		S(cwd+"/main.c"),
		S("Command:2:16:build.sh"),
	)

	modules = append(modules,
		module{
			priority:   0,
			lang:       "Sh",
			imports:    []import_{iLibC, iMainC},
			exports:    []export{buildSh, eVar},
			intralinks: nil,
		},
	)

	return modules
}
