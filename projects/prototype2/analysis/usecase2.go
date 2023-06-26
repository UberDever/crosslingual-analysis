package analysis

import (
	"os"
	"prototype2/sexpr"
	"prototype2/util"
)

// Analysis of Python, bash and C FFI, see usecase2 files

func Usecase2_Analyzer() []fragment {
	fragments := make([]fragment, 0, 128)
	fragments = append(fragments, analyzePython()...)
	fragments = append(fragments, analyzeBash()...)
	fragments = append(fragments, analyzeC()...)
	return fragments
}

func analyzePython() []fragment {
	fragments := make([]fragment, 0, 128)
	S := sexpr.S

	home, _ := os.UserHomeDir()
	cwd := util.ShortenPath(home+"/dev/mag/language-analysis/projects/prototype2/usecases/usecase2", 2)

	scriptPy := NewExport(
		S("URI"),
		S(cwd+"/script.py"),
		S("URI:0:0:script.py"),
	)
	importLib := NewImport(
		S("URI"),
		S(cwd+"/liblib.so"),
		S("FunctionCall:17:3:script.py"),
	)
	importFunc := NewImport(
		Function("Unit", "Any"),
		S("doTwoPlusTwo"),
		S("PropertyAccess:7:9:script.py"),
	)

	fragments = append(fragments,
		fragment{
			path:         cwd + "/script.py",
			priority:     0,
			lang:         "Python",
			environments: []environment{importFunc, importLib},
			signatures:   []signature{scriptPy},
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

	return fragments
}

func analyzeBash() []fragment {
	fragments := make([]fragment, 0, 128)
	S := sexpr.S

	home, _ := os.UserHomeDir()
	cwd := util.ShortenPath(home+"/dev/mag/language-analysis/projects/prototype2/usecases/usecase2", 2)

	runSh := NewExport(
		S("URI"),
		S(cwd+"/run.sh"),
		S("URI:0:0:run.sh"),
	)
	libSo := NewExport(
		S("URI"),
		S(cwd+"/liblib.so"),
		S("Command:5:15:run.sh"),
	)
	importLibC := NewImport(
		S("URI"),
		S(cwd+"/lib.c"),
		S("Command:4:7:run.sh"),
	)
	importScriptPy := NewImport(
		S("URI"),
		S(cwd+"/script.py"),
		S("Command:6:9:run.sh"),
	)

	fragments = append(fragments,
		fragment{
			path:         cwd + "/run.sh",
			priority:     0,
			lang:         "Sh",
			environments: []environment{importLibC, importScriptPy},
			signatures:   []signature{libSo, runSh},
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

	return fragments
}

func analyzeC() []fragment {
	fragments := make([]fragment, 0, 128)
	S := sexpr.S

	home, _ := os.UserHomeDir()
	cwd := util.ShortenPath(home+"/dev/mag/language-analysis/projects/prototype2/usecases/usecase2", 2)

	libC := NewExport(
		S("URI"),
		S(cwd+"/lib.c"),
		S("URI:0:0:lib.c"),
	)
	fn := NewExport(
		Function("Unit", "Int"),
		S("doTwoPlusTwo"),
		S("FunctionDecl:2:5:lib.c"),
	)

	fragments = append(fragments,
		fragment{
			path:         cwd + "/lib.c",
			priority:     0,
			lang:         "C",
			environments: nil,
			signatures:   []signature{fn, libC},
			intralinks:   nil,
		},
	)

	return fragments
}
