package analysis

import (
	"os"
	"prototype2/sexpr"
	"prototype2/util"
)

// Analysis of Python, bash and C FFI, see usecase2 files

func Usecase2_Analyzer() []module {
	modules := make([]module, 0, 128)
	modules = append(modules, analyzePython()...)
	modules = append(modules, analyzeBash()...)
	modules = append(modules, analyzeC()...)
	return modules
}

func analyzePython() []module {
	modules := make([]module, 0, 128)
	S := sexpr.S

	home, _ := os.UserHomeDir()
	cwd := util.ShortenPath(home+"/dev/mag/language-analysis/projects/prototype2/usecases/usecase2", 2)

	scriptPy := NewExport(
		S("File"),
		S(cwd+"/script.py"),
		S("File:0:0:script.py"),
	)
	importLib := NewImport(
		S("File"),
		S(cwd+"/liblib.so"),
		S("FunctionCall:17:3:script.py"),
	)
	importFunc := NewImport(
		Function("Unit", "Any"),
		S("doTwoPlusTwo"),
		S("PropertyAccess:7:9:script.py"),
	)

	modules = append(modules,
		module{
			path:     cwd + "/script.py",
			priority: 0,
			lang:     "Python",
			imports:  []import_{importFunc, importLib},
			exports:  []export{scriptPy},
			intralinks: []struct {
				from statement
				to   statement
			}{
				{from: importFunc.statement,
					to: importLib.statement,
				},
			},
		},
	)

	return modules
}

func analyzeBash() []module {
	modules := make([]module, 0, 128)
	S := sexpr.S

	home, _ := os.UserHomeDir()
	cwd := util.ShortenPath(home+"/dev/mag/language-analysis/projects/prototype2/usecases/usecase2", 2)

	runSh := NewExport(
		S("File"),
		S(cwd+"/run.sh"),
		S("File:0:0:run.sh"),
	)
	libSo := NewExport(
		S("File"),
		S(cwd+"/liblib.so"),
		S("Command:5:15:run.sh"),
	)
	importLibC := NewImport(
		S("File"),
		S(cwd+"/lib.c"),
		S("Command:4:7:run.sh"),
	)
	importScriptPy := NewImport(
		S("File"),
		S(cwd+"/script.py"),
		S("Command:6:9:run.sh"),
	)

	modules = append(modules,
		module{
			path:     cwd + "/run.sh",
			priority: 0,
			lang:     "Sh",
			imports:  []import_{importLibC, importScriptPy},
			exports:  []export{libSo, runSh},
			intralinks: []struct {
				from statement
				to   statement
			}{
				{from: libSo.statement,
					to: importLibC.statement,
				},
			},
		},
	)

	return modules
}

func analyzeC() []module {
	modules := make([]module, 0, 128)
	S := sexpr.S

	home, _ := os.UserHomeDir()
	cwd := util.ShortenPath(home+"/dev/mag/language-analysis/projects/prototype2/usecases/usecase2", 2)

	libC := NewExport(
		S("File"),
		S(cwd+"/lib.c"),
		S("File:0:0:lib.c"),
	)
	fn := NewExport(
		Function("Unit", "Int"),
		S("doTwoPlusTwo"),
		S("FunctionDecl:2:5:lib.c"),
	)

	modules = append(modules,
		module{
			path:       cwd + "/lib.c",
			priority:   0,
			lang:       "C",
			imports:    nil,
			exports:    []export{fn, libC},
			intralinks: nil,
		},
	)

	return modules
}
