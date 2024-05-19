/* --------------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 * ------------------------------------------------------------------------------------------ */
import {
    createConnection,
    TextDocuments,
    Diagnostic,
    DiagnosticSeverity,
    ProposedFeatures,
    InitializeParams,
    DidChangeConfigurationNotification,
    CompletionItem,
    CompletionItemKind,
    TextDocumentPositionParams,
    TextDocumentSyncKind,
    InitializeResult,
    DocumentDiagnosticReportKind,
    type DocumentDiagnosticReport,
    TypeHierarchyItem,
    Range,
    Position,
    LocationLink,
    DeclarationParams,
    CompletionParams,
    TextDocumentItem,
    TextDocumentContentChangeEvent,
    SignatureHelpParams,
    SignatureHelp,
    RenameParams,
    WorkspaceEdit,
    PrepareRenameParams,
    TextEdit
} from 'vscode-languageserver/node';

import { URI } from 'vscode-uri'

import {
    TextDocument,
} from 'vscode-languageserver-textdocument';
import path = require('path');
import os = require('os')
import url = require('url')

// Create a connection for the server, using Node's IPC as a transport.
// Also include all preview / proposed LSP features.
const connection = createConnection(ProposedFeatures.all);

// Create a simple text document manager.
const documents: TextDocuments<TextDocument> = new TextDocuments(TextDocument);

let hasConfigurationCapability = false;
let hasWorkspaceFolderCapability = false;
let hasDiagnosticRelatedInformationCapability = false;

connection.onInitialize((params: InitializeParams) => {
    const capabilities = params.capabilities;

    // Does the client support the `workspace/configuration` request?
    // If not, we fall back using global settings.
    hasConfigurationCapability = !!(
        capabilities.workspace && !!capabilities.workspace.configuration
    );
    hasWorkspaceFolderCapability = !!(
        capabilities.workspace && !!capabilities.workspace.workspaceFolders
    );
    hasDiagnosticRelatedInformationCapability = !!(
        capabilities.textDocument &&
        capabilities.textDocument.publishDiagnostics &&
        capabilities.textDocument.publishDiagnostics.relatedInformation
    );

    const result: InitializeResult = {
        capabilities: {
            textDocumentSync: TextDocumentSyncKind.Incremental,
            declarationProvider: true,
            definitionProvider: true,
            // Tell the client that this server supports code completion.
            completionProvider: {
                resolveProvider: true,
                triggerCharacters: ["."]
            },
            signatureHelpProvider: {
                triggerCharacters: ["(", ","],
            },
            diagnosticProvider: {
                interFileDependencies: true,
                workspaceDiagnostics: false
            },
            typeHierarchyProvider: true,
            renameProvider: true,
            workspace: {
                fileOperations: {
                    willRename: {
                        filters: [
                            {
                                scheme: 'file',
                                pattern: {
                                    glob: '**/*.{go,js}',
                                    matches: 'file',
                                },
                            }
                        ],
                    },
                },
            },
        },
        serverInfo: {
            name: "Multilingual LSP"
        }
    };
    if (hasWorkspaceFolderCapability) {
        result.capabilities.workspace = {
            workspaceFolders: {
                supported: true
            }
        };
    }
    return result;
});

connection.onInitialized(() => {
    if (hasConfigurationCapability) {
        // Register for all configuration changes.
        connection.client.register(DidChangeConfigurationNotification.type, undefined);
    }
    if (hasWorkspaceFolderCapability) {
        connection.workspace.onDidChangeWorkspaceFolders(_event => {
            connection.console.log('Workspace folder change event received.');
        });
    }
});

// The example settings
interface ExampleSettings {
    maxNumberOfProblems: number;
}

// The global settings, used when the `workspace/configuration` request is not supported by the client.
// Please note that this is not the case when using this server with the client provided in this example
// but could happen with other clients.
const defaultSettings: ExampleSettings = { maxNumberOfProblems: 1000 };

// Cache the settings of all open documents
const documentSettings: Map<string, Thenable<ExampleSettings>> = new Map();

connection.onDidChangeConfiguration(change => {
    if (hasConfigurationCapability) {
        // Reset all cached document settings
        documentSettings.clear();
    } else {
    }
    // Refresh the diagnostics since the `maxNumberOfProblems` could have changed.
    // We could optimize things here and re-fetch the setting first can compare it
    // to the existing setting, but this is out of scope for this example.
    connection.languages.diagnostics.refresh();
});

// Only keep settings for open documents
documents.onDidClose(e => {
    documentSettings.delete(e.document.uri);
});

const golang_path = withHome("dev/mag/crosslingual-analysis/projects/crossy/lsp-adapter/examples/Example 3/backend/server.go")
const js_path = withHome("dev/mag/crosslingual-analysis/projects/crossy/lsp-adapter/examples/Example 3/frontend/script.js")
const countTagRange: Range = {
    start: {
        line: 9, character: 21,
    },
    end: {
        line: 9, character: 26,
    }
}
const countRange: Range = {
    start: {
        line: 13, character: 99,
    },
    end: {
        line: 13, character: 104,
    }
}

let countName = "count"

function findCount(uri: string): [string, Range] | undefined {
    console.log(uri)
    if (uri === js_path) {
        const d = documents.get(js_path)
        if (!d) return undefined
        const jsDoc = new LspDocument(d, js_path)
        const countRegex = /\bdata\s*\.\s*(\w+)\b/g
        const matches = jsDoc.getText().matchAll(countRegex)
        for (const match of matches) {
            if (!match.index) continue
            const pos = jsDoc.positionAt(match.index)
            const [data, count] = [match[0], match[1]]
            return [
                count,
                {
                    start: {
                        line: pos.line,
                        character: pos.character + data.length - count.length
                    },
                    end: {
                        line: pos.line,
                        character: pos.character + data.length
                    }
                }
            ]
        }
    } else if (uri === golang_path) {
        const d = documents.get(golang_path)
        if (!d) return undefined
        const doc = new LspDocument(d, golang_path)
        const countRegex = /json:"(\w+)"/g
        const matches = doc.getText().matchAll(countRegex)
        for (const match of matches) {
            if (!match.index) continue
            const pos = doc.positionAt(match.index) // this reports incorrect position idk why
            const shift = -1
            const [tag, count] = [match[0], match[1]]
            return [
                count,
                {
                    start: { line: pos.line, character: pos.character + tag.length - count.length + shift },
                    end: { line: pos.line, character: pos.character + tag.length + shift }
                }
            ]
        }
    }
    return undefined
}

connection.onRenameRequest(async (params: RenameParams): Promise<WorkspaceEdit> => {
    if (params.textDocument.uri !== golang_path) {
        return {}
    }

    let edits: { [uri: string]: TextEdit[]; } = {}
    const createAndPush = (uri: string, edit: TextEdit) => {
        if (!edits[uri]) {
            edits[uri] = []
        }
        edits[uri].push(edit)
    }
    const jsDoc = new LspDocument(documents.get(js_path)!, js_path)
    const countLine = jsDoc.getLine(countRange.start.line)
    const countRegex = /\bdata\.\w+\b/g
    const matches = countRegex.exec(countLine)
    if (matches?.length === 1) {
        const count = matches[0].split(".")[1]
        if (count === countName) {
            countName = params.newName
            createAndPush(js_path, TextEdit.replace(countRange, params.newName))
        }
    }

    return { changes: edits }
})

connection.languages.diagnostics.on(async (params): Promise<DocumentDiagnosticReport> => {
    const empty: DocumentDiagnosticReport = {
        kind: DocumentDiagnosticReportKind.Full,
        items: []
    };
    if (params.textDocument.uri === js_path) {
        const countFromGo = findCount(golang_path)
        if (!countFromGo) return empty
        const countFromJS = findCount(js_path)
        if (!countFromJS) return empty
        const [countGo, countGoRange] = countFromGo
        const [countJS, countJsRange] = countFromJS
        if (countGo === countJS) {
            return empty
        }
        return {
            kind: DocumentDiagnosticReportKind.Full,
            items: [
                {
                    severity: DiagnosticSeverity.Warning,
                    message: `Expected field "${countGo}" but got "${countJS}"`,
                    relatedInformation: [{ location: { uri: golang_path, range: countGoRange }, message: "struct tag declared here" }],
                    range: { ...countJsRange, end: { ...countJsRange.end, character: countJsRange.start.character + 1 } }
                }
            ]
        };
    }
    return empty
});

// The content of a text document has changed. This event is emitted
// when the text document first opened or when its content has changed.
documents.onDidChangeContent(change => {
    //console.log(change)
});

const python_path = withHome("dev/mag/crosslingual-analysis/projects/crossy/lsp-adapter/examples/Example 2/test.py")
const cpp_path = withHome("dev/mag/crosslingual-analysis/projects/crossy/lsp-adapter/examples/Example 2/lib.cpp")

// from https://github.com/typescript-language-server/typescript-language-server/
export class LspDocument {
    private _document: TextDocument;
    private _uri: URI;
    private _filepath: string;

    constructor(doc: TextDocument, filepath: string) {
        this._document = doc;
        this._uri = URI.parse(doc.uri);
        this._filepath = filepath;
    }

    get uri(): URI {
        return this._uri;
    }

    get filepath(): string {
        return this._filepath;
    }

    get languageId(): string {
        return this._document.languageId;
    }

    get version(): number {
        return this._document.version;
    }

    getText(range?: Range): string {
        return this._document.getText(range);
    }

    positionAt(offset: number): Position {
        return this._document.positionAt(offset);
    }

    offsetAt(position: Position): number {
        return this._document.offsetAt(position);
    }

    get lineCount(): number {
        return this._document.lineCount;
    }

    getLine(line: number): string {
        const lineRange = this.getLineRange(line);
        return this.getText(lineRange);
    }

    getLineRange(line: number): Range {
        const lineStart = this.getLineStart(line);
        const lineEnd = this.getLineEnd(line);
        return Range.create(lineStart, lineEnd);
    }

    getLineEnd(line: number): Position {
        const nextLine = line + 1;
        const nextLineOffset = this.getLineOffset(nextLine);
        // If next line doesn't exist then the offset is at the line end already.
        return this.positionAt(nextLine < this._document.lineCount ? nextLineOffset - 1 : nextLineOffset);
    }

    getLineOffset(line: number): number {
        const lineStart = this.getLineStart(line);
        return this.offsetAt(lineStart);
    }

    getLineStart(line: number): Position {
        return Position.create(line, 0);
    }

    getFullRange(): Range {
        return Range.create(
            Position.create(0, 0),
            this.getLineEnd(Math.max(this.lineCount - 1, 0)),
        );
    }

    applyEdit(version: number, change: TextDocumentContentChangeEvent): void {
        const content = this.getText();
        let newContent = change.text;
        if (TextDocumentContentChangeEvent.isIncremental(change)) {
            const start = this.offsetAt(change.range.start);
            const end = this.offsetAt(change.range.end);
            newContent = content.substr(0, start) + change.text + content.substr(end);
        }
        this._document = TextDocument.create(this._uri.toString(), this.languageId, version, newContent);
    }
}

const accessRegex = /\b\w+\./g
const counterFunctionTypes = new Map(Object.entries({
    "counter_new": "Unit -> (Ptr Counter)",
    "counter_free": "(Ptr Counter) -> Unit",
    "counter_get": "(Ptr Counter) -> Integer",
    "counter_reset": "(Ptr Counter) -> Unit",
    "counter_inc": "(Ptr Counter) -> Unit",
}))

// This handler provides the initial list of the completion items.
connection.onCompletion((params: CompletionParams): CompletionItem[] => {
    const d = documents.get(python_path)
    if (!d) {
        return []
    }
    const doc = new LspDocument(d, params.textDocument.uri)
    const pos = params.position
    const detail = (T: string, from: string) => {
        return T + " -- " + "From " + from
    }
    if (params.textDocument.uri === python_path) {
        const line = doc.getLine(pos.line)
        const matches = accessRegex.exec(line)
        if (matches) {
            const id = matches[0]
            if (id === "lib.") {
                return [
                    {
                        label: "counter_new",
                        kind: CompletionItemKind.Function,
                        detail: detail(counterFunctionTypes.get("counter_new")!, cpp_path)
                    },
                    {
                        label: "counter_free",
                        kind: CompletionItemKind.Function,
                        detail: detail(counterFunctionTypes.get("counter_free")!, cpp_path)
                    },
                    {
                        label: "counter_get",
                        kind: CompletionItemKind.Function,
                        detail: detail(counterFunctionTypes.get("counter_get")!, cpp_path)
                    },
                    {
                        label: "counter_reset",
                        kind: CompletionItemKind.Function,
                        detail: detail(counterFunctionTypes.get("counter_reset")!, cpp_path)
                    },
                    {
                        label: "counter_inc",
                        kind: CompletionItemKind.Function,
                        detail: detail(counterFunctionTypes.get("counter_inc")!, cpp_path)
                    },
                ]
            }
        }
    }
    return []
});

connection.onCompletionResolve(
    (item: CompletionItem): CompletionItem => {
        return item
    }
);

connection.onSignatureHelp((params: SignatureHelpParams): SignatureHelp => {
    const doc = new LspDocument(documents.get(python_path)!, "")
    const pos = params.position
    if (params.textDocument.uri === python_path) {
        const line = doc.getLine(pos.line)
        const libPrefix = "lib."
        for (let [func, T] of counterFunctionTypes) {
            if (line.includes(libPrefix + func)) {
                return {
                    signatures: [
                        { label: T }
                    ]
                }
            }
        }
    }
    return { signatures: [] }
})

const csharp_path = withHome("dev/mag/crosslingual-analysis/projects/crossy/lsp-adapter/examples/Example 1/CSharp/Program.cs")
const vb_path = withHome("dev/mag/crosslingual-analysis/projects/crossy/lsp-adapter/examples/Example 1/VB/Class1.vb")

function onDeclOrDef(params: DeclarationParams): LocationLink[] {
    const field_ranges: Range[] = [
        {
            start: { line: 8, character: 12 },
            end: { line: 8, character: 17 }
        },
        {
            start: { line: 12, character: 42 },
            end: { line: 12, character: 47 }
        },
    ]
    const vb_field_range = Range.create(1, 11, 1, 16)
    const field_location = LocationLink.create(vb_path, vb_field_range, vb_field_range)
    const included = (p: Position, r: Range): boolean => {
        return r.start.line == p.line &&
            r.start.character <= p.character &&
            p.line == r.end.line &&
            p.character <= r.end.character;
    }
    if (params.textDocument.uri === csharp_path) {
        for (const r of field_ranges) {
            if (included(params.position, r)) {
                return [field_location]
            }
        }
    }
    return []
}

connection.onDeclaration(async (params) => {
    return onDeclOrDef(params)
})

connection.onDefinition(async (params) => {
    return onDeclOrDef(params)
})


function withHome(s: string): string {
    return url.pathToFileURL(path.join(os.homedir(), s)).href
}
const A_range: Range = {
    start: {
        line: 4,
        character: 10
    },
    end: {
        line: 4,
        character: 10
    }
}
const BaseVB_range: Range = {
    start: {
        line: 0,
        character: 13
    },
    end: {
        line: 0,
        character: 13
    }
}

const csharp_A: TypeHierarchyItem = {
    name: "A",
    kind: 5,
    uri: csharp_path,
    range: A_range,
    selectionRange: A_range,
};

const vb_baseVB: TypeHierarchyItem =
{
    name: "BaseVB",
    kind: 5,
    uri: vb_path,
    range: BaseVB_range,
    selectionRange: BaseVB_range,
};

connection.languages.typeHierarchy.onPrepare(async (params) => {
    if (params.textDocument.uri.endsWith("cs")) {
        return [
            csharp_A
        ]
    } else {
        return [
            vb_baseVB
        ]
    }
})

connection.languages.typeHierarchy.onSupertypes(async (params) => {
    if (params.item.uri.endsWith("cs")) {
        return [
            vb_baseVB
        ]
    }
    return []
})

connection.languages.typeHierarchy.onSubtypes(async (params) => {
    if (params.item.uri.endsWith("vb")) {
        return [
            csharp_A
        ]
    }
    return []
})

// Make the text document manager listen on the connection
// for open, change and close text document events
documents.listen(connection);

// Listen on the connection
connection.listen();
