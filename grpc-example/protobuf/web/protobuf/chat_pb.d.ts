import * as jspb from 'google-protobuf'



export class ChatConnectRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): ChatConnectRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChatConnectRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ChatConnectRequest): ChatConnectRequest.AsObject;
  static serializeBinaryToWriter(message: ChatConnectRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChatConnectRequest;
  static deserializeBinaryFromReader(message: ChatConnectRequest, reader: jspb.BinaryReader): ChatConnectRequest;
}

export namespace ChatConnectRequest {
  export type AsObject = {
    token: string,
  }
}

export class ChatConnectResponse extends jspb.Message {
  getStatus(): string;
  setStatus(value: string): ChatConnectResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChatConnectResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ChatConnectResponse): ChatConnectResponse.AsObject;
  static serializeBinaryToWriter(message: ChatConnectResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChatConnectResponse;
  static deserializeBinaryFromReader(message: ChatConnectResponse, reader: jspb.BinaryReader): ChatConnectResponse;
}

export namespace ChatConnectResponse {
  export type AsObject = {
    status: string,
  }
}

export class ChatSendDataRequest extends jspb.Message {
  getData(): string;
  setData(value: string): ChatSendDataRequest;

  getFrom(): string;
  setFrom(value: string): ChatSendDataRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChatSendDataRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ChatSendDataRequest): ChatSendDataRequest.AsObject;
  static serializeBinaryToWriter(message: ChatSendDataRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChatSendDataRequest;
  static deserializeBinaryFromReader(message: ChatSendDataRequest, reader: jspb.BinaryReader): ChatSendDataRequest;
}

export namespace ChatSendDataRequest {
  export type AsObject = {
    data: string,
    from: string,
  }
}

export class ChatSendDataResponse extends jspb.Message {
  getStatus(): string;
  setStatus(value: string): ChatSendDataResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChatSendDataResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ChatSendDataResponse): ChatSendDataResponse.AsObject;
  static serializeBinaryToWriter(message: ChatSendDataResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChatSendDataResponse;
  static deserializeBinaryFromReader(message: ChatSendDataResponse, reader: jspb.BinaryReader): ChatSendDataResponse;
}

export namespace ChatSendDataResponse {
  export type AsObject = {
    status: string,
  }
}

