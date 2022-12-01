/**
 * @fileoverview gRPC-Web generated client stub for drawing_share
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.1
// 	protoc              v3.12.4
// source: protobuf/drawing_share.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as protobuf_drawing_share_pb from '../protobuf/drawing_share_pb';


export class DrawingShareClient {
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
    '/drawing_share.DrawingShare/Connect',
    grpcWeb.MethodType.SERVER_STREAMING,
    protobuf_drawing_share_pb.ConnectRequest,
    protobuf_drawing_share_pb.ConnectResponse,
    (request: protobuf_drawing_share_pb.ConnectRequest) => {
      return request.serializeBinary();
    },
    protobuf_drawing_share_pb.ConnectResponse.deserializeBinary
  );

  connect(
    request: protobuf_drawing_share_pb.ConnectRequest,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<protobuf_drawing_share_pb.ConnectResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/drawing_share.DrawingShare/Connect',
      request,
      metadata || {},
      this.methodDescriptorConnect);
  }

  methodDescriptorSendDrawing = new grpcWeb.MethodDescriptor(
    '/drawing_share.DrawingShare/SendDrawing',
    grpcWeb.MethodType.UNARY,
    protobuf_drawing_share_pb.SendDrawingRequest,
    protobuf_drawing_share_pb.SendDrawingResponse,
    (request: protobuf_drawing_share_pb.SendDrawingRequest) => {
      return request.serializeBinary();
    },
    protobuf_drawing_share_pb.SendDrawingResponse.deserializeBinary
  );

  sendDrawing(
    request: protobuf_drawing_share_pb.SendDrawingRequest,
    metadata: grpcWeb.Metadata | null): Promise<protobuf_drawing_share_pb.SendDrawingResponse>;

  sendDrawing(
    request: protobuf_drawing_share_pb.SendDrawingRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: protobuf_drawing_share_pb.SendDrawingResponse) => void): grpcWeb.ClientReadableStream<protobuf_drawing_share_pb.SendDrawingResponse>;

  sendDrawing(
    request: protobuf_drawing_share_pb.SendDrawingRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: protobuf_drawing_share_pb.SendDrawingResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/drawing_share.DrawingShare/SendDrawing',
        request,
        metadata || {},
        this.methodDescriptorSendDrawing,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/drawing_share.DrawingShare/SendDrawing',
    request,
    metadata || {},
    this.methodDescriptorSendDrawing);
  }

  methodDescriptorDisConnect = new grpcWeb.MethodDescriptor(
    '/drawing_share.DrawingShare/DisConnect',
    grpcWeb.MethodType.UNARY,
    protobuf_drawing_share_pb.DisConnectRequest,
    protobuf_drawing_share_pb.DisConnectResponse,
    (request: protobuf_drawing_share_pb.DisConnectRequest) => {
      return request.serializeBinary();
    },
    protobuf_drawing_share_pb.DisConnectResponse.deserializeBinary
  );

  disConnect(
    request: protobuf_drawing_share_pb.DisConnectRequest,
    metadata: grpcWeb.Metadata | null): Promise<protobuf_drawing_share_pb.DisConnectResponse>;

  disConnect(
    request: protobuf_drawing_share_pb.DisConnectRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: protobuf_drawing_share_pb.DisConnectResponse) => void): grpcWeb.ClientReadableStream<protobuf_drawing_share_pb.DisConnectResponse>;

  disConnect(
    request: protobuf_drawing_share_pb.DisConnectRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: protobuf_drawing_share_pb.DisConnectResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/drawing_share.DrawingShare/DisConnect',
        request,
        metadata || {},
        this.methodDescriptorDisConnect,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/drawing_share.DrawingShare/DisConnect',
    request,
    metadata || {},
    this.methodDescriptorDisConnect);
  }

}

