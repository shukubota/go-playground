/**
 * @fileoverview gRPC-Web generated client stub for chat
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.1
// 	protoc              v3.12.4
// source: protobuf/chat.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as protobuf_chat_pb from '../protobuf/chat_pb';


export class ChatClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorConnect = new grpcWeb.MethodDescriptor(
    '/chat.Chat/Connect',
    grpcWeb.MethodType.SERVER_STREAMING,
    protobuf_chat_pb.ChatConnectRequest,
    protobuf_chat_pb.ChatConnectResponse,
    (request: protobuf_chat_pb.ChatConnectRequest) => {
      return request.serializeBinary();
    },
    protobuf_chat_pb.ChatConnectResponse.deserializeBinary
  );

  connect(
    request: protobuf_chat_pb.ChatConnectRequest,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<protobuf_chat_pb.ChatConnectResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/chat.Chat/Connect',
      request,
      metadata || {},
      this.methodDescriptorConnect);
  }

  methodDescriptorSendData = new grpcWeb.MethodDescriptor(
    '/chat.Chat/SendData',
    grpcWeb.MethodType.UNARY,
    protobuf_chat_pb.ChatSendDataRequest,
    protobuf_chat_pb.ChatSendDataResponse,
    (request: protobuf_chat_pb.ChatSendDataRequest) => {
      return request.serializeBinary();
    },
    protobuf_chat_pb.ChatSendDataResponse.deserializeBinary
  );

  sendData(
    request: protobuf_chat_pb.ChatSendDataRequest,
    metadata: grpcWeb.Metadata | null): Promise<protobuf_chat_pb.ChatSendDataResponse>;

  sendData(
    request: protobuf_chat_pb.ChatSendDataRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: protobuf_chat_pb.ChatSendDataResponse) => void): grpcWeb.ClientReadableStream<protobuf_chat_pb.ChatSendDataResponse>;

  sendData(
    request: protobuf_chat_pb.ChatSendDataRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: protobuf_chat_pb.ChatSendDataResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/chat.Chat/SendData',
        request,
        metadata || {},
        this.methodDescriptorSendData,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/chat.Chat/SendData',
    request,
    metadata || {},
    this.methodDescriptorSendData);
  }

}

