type integer = number
type array = [any]


interface Message {
    jsonrpc: string;
}

interface RequestMessage extends Message {
    id: integer | string;
    method: string;
    params?: array | object;
}

interface ResponseMessage extends Message {
    id: integer | string | null;
    result?: string | integer | boolean | array | object | null;
    error?: ResponseError;
}

export namespace ErrorCodes {
    // Defined by JSON-RPC
    export const ParseError: integer// = -32700;
    export const InvalidRequest: integer// = -32600;
    export const MethodNotFound: integer// = -32601;
    export const InvalidParams: integer// = -32602;
    export const InternalError: integer// = -32603;

    /**
     * This is the start range of JSON-RPC reserved error codes.
     * It doesn't denote a real error code. No LSP error codes should
     * be defined between the start and end range. For backwards
     * compatibility the `ServerNotInitialized` and the `UnknownErrorCode`
     * are left in the range.
     *
     * @since 3.16.0
     */
    export const jsonrpcReservedErrorRangeStart: integer// = -32099;
    /** @deprecated use jsonrpcReservedErrorRangeStart */
    export const serverErrorStart: integer// = jsonrpcReservedErrorRangeStart;

    /**
     * Error code indicating that a server received a notification or
     * request before the server has received the `initialize` request.
     */
    export const ServerNotInitialized: integer// = -32002;
    export const UnknownErrorCode: integer// = -32001;

    /**
     * This is the end range of JSON-RPC reserved error codes.
     * It doesn't denote a real error code.
     *
     * @since 3.16.0
     */
    export const jsonrpcReservedErrorRangeEnd// = -32000;
    /** @deprecated use jsonrpcReservedErrorRangeEnd */
    export const serverErrorEnd: integer// = jsonrpcReservedErrorRangeEnd;

    /**
     * This is the start range of LSP reserved error codes.
     * It doesn't denote a real error code.
     *
     * @since 3.16.0
     */
    export const lspReservedErrorRangeStart: integer// = -32899;

    /**
     * A request failed but it was syntactically correct, e.g the
     * method name was known and the parameters were valid. The error
     * message should contain human readable information about why
     * the request failed.
     *
     * @since 3.17.0
     */
    export const RequestFailed: integer// = -32803;

    /**
     * The server cancelled the request. This error code should
     * only be used for requests that explicitly support being
     * server cancellable.
     *
     * @since 3.17.0
     */
    export const ServerCancelled: integer// = -32802;

    /**
     * The server detected that the content of a document got
     * modified outside normal conditions. A server should
     * NOT send this error code if it detects a content change
     * in it unprocessed messages. The result even computed
     * on an older state might still be useful for the client.
     *
     * If a client decides that a result is not of any use anymore
     * the client should cancel the request.
     */
    export const ContentModified: integer// = -32801;

    /**
     * The client has canceled a request and a server as detected
     * the cancel.
     */
    export const RequestCancelled: integer// = -32800;

    /**
     * This is the end range of LSP reserved error codes.
     * It doesn't denote a real error code.
     *
     * @since 3.16.0
     */
    export const lspReservedErrorRangeEnd: integer// = -32800;
}

interface ResponseError {
    code: integer;
    message: string;
    data?: string | integer | boolean | array | object | null;
}

// an event
interface NotificationMessage extends Message {
    method: string;
    params?: array | object;
}

declare class CancelMessage {
    method: string// = "$/cancelRequest"
    params?: CancelParams
}

interface CancelParams {
    id: integer | string;
}

type Language = string
type URI = string
type Code = string

interface Ontology {
    // will be filled as necessary
}

