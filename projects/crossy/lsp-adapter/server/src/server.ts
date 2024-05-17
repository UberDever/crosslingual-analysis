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
    DeclarationParams
} from 'vscode-languageserver/node';

import {
    TextDocument
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
            textDocumentSync: TextDocumentSyncKind.Full,
            declarationProvider: true,
            definitionProvider: true,
            // Tell the client that this server supports code completion.
            completionProvider: {
                resolveProvider: true
            },
            signatureHelpProvider: {
                triggerCharacters: ["(", ","],
            },
            diagnosticProvider: {
                interFileDependencies: false,
                workspaceDiagnostics: false
            },
            typeHierarchyProvider: true
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


connection.languages.diagnostics.on(async (params) => {
    const document = documents.get(params.textDocument.uri);
    if (document !== undefined) {
        return {
            kind: DocumentDiagnosticReportKind.Full,
            items: []
        } satisfies DocumentDiagnosticReport;
    } else {
        // We don't know the document. We can either try to read it from disk
        // or we don't report problems for it.
        return {
            kind: DocumentDiagnosticReportKind.Full,
            items: []
        } satisfies DocumentDiagnosticReport;
    }
});

// The content of a text document has changed. This event is emitted
// when the text document first opened or when its content has changed.
documents.onDidChangeContent(change => {
    console.log(change)
});

// This handler provides the initial list of the completion items.
connection.onCompletion(
    (textDocumentPosition: TextDocumentPositionParams): CompletionItem[] => {
        const doc = textDocumentPosition.textDocument

        return [
            {
                label: 'водил',
                kind: CompletionItemKind.Text,
                data: 1
            },
            {
                label: 'в кино',
                kind: CompletionItemKind.Text,
                data: 2
            },
            {
                label: 'маму',
                kind: CompletionItemKind.Text,
                data: 3
            },
            {
                label: 'твою',
                kind: CompletionItemKind.Text,
                data: 4
            },
            {
                label: 'Я',
                kind: CompletionItemKind.Text,
                data: 5
            }
        ];
    }
);

connection.onCompletionResolve(
    (item: CompletionItem): CompletionItem => {
        if (item.data === 1) {
            item.detail = 'TypeScript details';
            item.documentation = 'TypeScript documentation';
        } else if (item.data === 2) {
            item.detail = 'JavaScript details';
            item.documentation = 'JavaScript documentation';
        }
        return item;
    }
);

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
