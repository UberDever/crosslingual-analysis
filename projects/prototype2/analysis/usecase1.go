package analysis

// Analysis of C# and javascript, see usecase1 files

import (
	"fmt"
	"prototype2/sexpr"
	"prototype2/util"
)

type Box = sexpr.Box

func Usecase1_CSharp() Box {
	S := sexpr.S
	return S(".TypeDeclaration",
		S(".Attributes",
			S(".Attribute", "Route", "api/[controller]"),
			S(".Attribute", "ApiController"),
		),
		S(".Class", "TodoItemsController", S(".ClassBase", "ControllerBase"),
			S(".ClassMemberDeclarations",
				// GetTodoItems
				S(".ClassMemberDeclaration",
					S(".Attributes",
						S(".Attribute", "HttpGet")),
					S(".AllMemberModifiers", "Public", "Async"),
					S(".TypedMemberDeclaration",
						S(".Type", "Task", S(".TypeArgumentList",
							S(".Type", "ActionResult", S(".TypeArgumentList",
								S(".Type", "IEnumerable", S(".TypeArgumentList",
									S(".Type", "TodoItem"))))))),
						S(".MethodDeclaration",
							S(".MethodMemberName", "GetTodoItems"),
							S(".FormalParameterList", nil),
							S(".MethodBody" /*TODO*/),
						))),
				// GetTodoItem
				S(".ClassMemberDeclaration",
					S(".Attributes",
						S(".Attribute", "HttpGet", "{id}")),
					S(".AllMemberModifiers", "Public", "Async"),
					S(".TypedMemberDeclaration",
						S(".Type", "Task", S(".TypeArgumentList",
							S(".Type", "ActionResult", S(".TypeArgumentList",
								S(".Type", "TodoItem"))))),
						S(".MethodDeclaration",
							S(".MethodMemberName", "GetTodoItem"),
							S(".FormalParameterList",
								S(".ArgDeclaration",
									S(".Type", "long"),
									"id",
								)),
							S(".MethodBody" /*TODO*/),
						))),
				// PutTodoItem
				S(".ClassMemberDeclaration",
					S(".Attributes",
						S(".Attribute", "HttpPut", "{id}")),
					S(".AllMemberModifiers", "Public", "Async"),
					S(".TypedMemberDeclaration",
						S(".Type", "Task", S(".TypeArgumentList",
							S(".Type", "IActionResult"))),
						S(".MethodDeclaration",
							S(".MethodMemberName", "PutTodoItem"),
							S(".FormalParameterList",
								S(".ArgDeclaration",
									S(".Type", "long"),
									"id"),
								S(".ArgDeclaration",
									S(".Type", "TodoItem"),
									"todoItem"),
							),
							S(".MethodBody" /*TODO*/),
						))),
				// PostTodoItem
				S(".ClassMemberDeclaration",
					S(".Attributes",
						S(".Attribute", "HttpPost")),
					S(".AllMemberModifiers", "Public", "Async"),
					S(".TypedMemberDeclaration",
						S(".Type", "Task", S(".TypeArgumentList",
							S(".Type", "ActionResult",
								S(".TypeArgumentList",
									S(".Type", "TodoItem"))))),
						S(".MethodDeclaration",
							S(".MethodMemberName", "PostTodoItem"),
							S(".FormalParameterList",
								S(".ArgDeclaration",
									S(".Type", "TodoItem"),
									"todoItem"),
							),
							S(".MethodBody" /*TODO*/),
						))),
				// DeleteTodoItem
				S(".ClassMemberDeclaration",
					S(".Attributes",
						S(".Attribute", "HttpDelete", "{id}")),
					S(".AllMemberModifiers", "Public", "Async"),
					S(".TypedMemberDeclaration",
						S(".Type", "Task", S(".TypeArgumentList",
							S(".Type", "IActionResult"))),
						S(".MethodDeclaration",
							S(".MethodMemberName", "DeleteTodoItem"),
							S(".FormalParameterList",
								S(".ArgDeclaration",
									S(".Type", "long"),
									"id"),
							),
							S(".MethodBody" /*TODO*/),
						))),
				// private member is not included
			)))
}

func Usecase1_JS() Box {
	S := sexpr.S
	return S(".Program",
		S(".VariableStatement", S(".VariableDeclarationList",
			// TODO: use partial evaluation in analysis to remember that constant and substitute it later
			S(".Const", S(".Identifier", "uri"),
				S(".StringLiteral", "api/todoitems")),
		)),
		S(".VariableStatement", S(".VariableDeclarationList",
			S(".Let", S(".Identifier", "todos"),
				S(".ArrayLiteral")),
		)),
		S(".FunctionDeclaration", "getItems",
			S(".FormalParameterList"),
			S(".FunctionBody",
				S(".SourceElements",
					// then chain is skipped
					S(".CallExpression", S(".Identifier", "fetch"),
						S(".Arguments", S(".Identifier", "uri"))),
				)),
		),
		S(".FunctionDeclaration", "addItem",
			S(".FormalParameterList"),
			S(".FunctionBody",
				S(".SourceElements",
					// variable declarations not interesting here, so they skipped
					// const addNameTextbox = ...
					// const item = {...}
					// then chain is skipped
					S(".CallExpression", S(".Identifier", "fetch"),
						S(".Arguments",
							S(".Identifier", "uri"),
							S(".ObjectLiteral",
								S(".PropertyAssignment",
									S(".Identifier", "method"),
									S(".StringLiteral", "POST"),
								),
								// headers: {...}
								// body: {...}
							))))),
		),
		S(".FunctionDeclaration", "deleteItem",
			S(".FormalParameterList", S(".Identifier", "id")),
			S(".FunctionBody",
				S(".SourceElements",
					// then chain is skipped
					S(".CallExpression", S(".Identifier", "fetch"),
						S(".Arguments",
							S(".TemplateStringLiteral",
								S(".Identifier", "uri"),
								S(".StringLiteral", "/"),
								S(".Identifier", "id"),
							),
							S(".ObjectLiteral",
								S(".PropertyAssignment",
									S(".Identifier", "method"),
									S(".StringLiteral", "DELETE"),
								),
							))))),
		),
		S(".FunctionDeclaration", "updateItem",
			S(".FormalParameterList"),
			S(".FunctionBody",
				S(".SourceElements",
					// const itemId = ...
					// const item = ...
					// then chain is skipped
					S(".CallExpression", S(".Identifier", "fetch"),
						S(".Arguments",
							S(".TemplateStringLiteral",
								S(".Identifier", "uri"),
								S(".StringLiteral", "/"),
								S(".Identifier", "itemId"),
							),
							S(".ObjectLiteral",
								S(".PropertyAssignment",
									S(".Identifier", "method"),
									S(".StringLiteral", "PUT"),
								),
								// headers: ...
								// body: ...
							))))),
		),
		// Then there are attribute sets on the html page, but I pretend
		// that functions are called directly
		// this will simplify analysis
		// TODO
	)
}

func Usecase1_Analyzer(csharpAST Box, jsAST Box) []node {
	nodes := make([]node, 0, 128)
	nodes = append(nodes, analyzeCsharp(csharpAST)...)
	nodes = append(nodes, analyzeJS(jsAST)...)
	return nodes
}

func analyzeCsharp(ast Box) []node {
	nodes := make([]node, 0, 128)
	S := sexpr.S
	Cons, Cdr, Car := sexpr.Cons, sexpr.Cdr, sexpr.Car

	Strings := util.NewStack[string]()
	Types := util.NewStack[Box]()
	// Arguments := util.NewStack[Box]()

	onEnter := func(n Box) {
		if !n.IsAtom() {
			return
		}

		s := n.Data.(string)
		if s[0] != '.' {
			Strings.Push(s)
			return
		}

		switch s {
		case ".Type":
			Types.Push(S(".Type", Strings.ForcePop()))
		case ".ArgDeclaration":
			args := S()
			for !Types.IsEmpty() {
				v, _ := Types.Pop()
				args = Cons(Car(Cdr(v)), args)
				fmt.Println(v.String())
			}
			fmt.Println(args)
		}
	}
	sexpr.TraversePostorder(ast, onEnter)

	return nodes
}

func analyzeJS(ast Box) []node {
	nodes := make([]node, 0, 128)

	onEnter := func(n Box) {
		if !n.IsAtom() {
			return
		}
	}
	sexpr.TraversePostorder(ast, onEnter)

	return nodes
}
