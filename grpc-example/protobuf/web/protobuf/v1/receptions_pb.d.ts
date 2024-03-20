import * as jspb from 'google-protobuf'

import * as protobuf_v1_reception_pb from '../../protobuf/v1/reception_pb';


export class ReceptionListRequest extends jspb.Message {
  getDate(): string;
  setDate(value: string): ReceptionListRequest;

  getAssignTo(): number;
  setAssignTo(value: number): ReceptionListRequest;

  getSortAppointmentTime(): string;
  setSortAppointmentTime(value: string): ReceptionListRequest;

  getSortReservationTime(): string;
  setSortReservationTime(value: string): ReceptionListRequest;

  getStatus(): number;
  setStatus(value: number): ReceptionListRequest;

  getPage(): number;
  setPage(value: number): ReceptionListRequest;

  getSize(): number;
  setSize(value: number): ReceptionListRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReceptionListRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReceptionListRequest): ReceptionListRequest.AsObject;
  static serializeBinaryToWriter(message: ReceptionListRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReceptionListRequest;
  static deserializeBinaryFromReader(message: ReceptionListRequest, reader: jspb.BinaryReader): ReceptionListRequest;
}

export namespace ReceptionListRequest {
  export type AsObject = {
    date: string,
    assignTo: number,
    sortAppointmentTime: string,
    sortReservationTime: string,
    status: number,
    page: number,
    size: number,
  }
}

export class ReceptionListResponse extends jspb.Message {
  getTotalPage(): number;
  setTotalPage(value: number): ReceptionListResponse;

  getCurrentPage(): number;
  setCurrentPage(value: number): ReceptionListResponse;

  getSize(): number;
  setSize(value: number): ReceptionListResponse;

  getReceptionsList(): Array<protobuf_v1_reception_pb.Reception>;
  setReceptionsList(value: Array<protobuf_v1_reception_pb.Reception>): ReceptionListResponse;
  clearReceptionsList(): ReceptionListResponse;
  addReceptions(value?: protobuf_v1_reception_pb.Reception, index?: number): protobuf_v1_reception_pb.Reception;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReceptionListResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ReceptionListResponse): ReceptionListResponse.AsObject;
  static serializeBinaryToWriter(message: ReceptionListResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReceptionListResponse;
  static deserializeBinaryFromReader(message: ReceptionListResponse, reader: jspb.BinaryReader): ReceptionListResponse;
}

export namespace ReceptionListResponse {
  export type AsObject = {
    totalPage: number,
    currentPage: number,
    size: number,
    receptionsList: Array<protobuf_v1_reception_pb.Reception.AsObject>,
  }
}

