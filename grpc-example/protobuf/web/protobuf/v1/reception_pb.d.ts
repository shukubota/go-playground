import * as jspb from 'google-protobuf'



export class Patient extends jspb.Message {
  getId(): number;
  setId(value: number): Patient;

  getFullName(): string;
  setFullName(value: string): Patient;

  getFuriganaName(): string;
  setFuriganaName(value: string): Patient;

  getGender(): number;
  setGender(value: number): Patient;

  getBirthYear(): string;
  setBirthYear(value: string): Patient;

  getClinicalNumber(): string;
  setClinicalNumber(value: string): Patient;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Patient.AsObject;
  static toObject(includeInstance: boolean, msg: Patient): Patient.AsObject;
  static serializeBinaryToWriter(message: Patient, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Patient;
  static deserializeBinaryFromReader(message: Patient, reader: jspb.BinaryReader): Patient;
}

export namespace Patient {
  export type AsObject = {
    id: number,
    fullName: string,
    furiganaName: string,
    gender: number,
    birthYear: string,
    clinicalNumber: string,
  }
}

export class Examination extends jspb.Message {
  getId(): number;
  setId(value: number): Examination;

  getName(): string;
  setName(value: string): Examination;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Examination.AsObject;
  static toObject(includeInstance: boolean, msg: Examination): Examination.AsObject;
  static serializeBinaryToWriter(message: Examination, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Examination;
  static deserializeBinaryFromReader(message: Examination, reader: jspb.BinaryReader): Examination;
}

export namespace Examination {
  export type AsObject = {
    id: number,
    name: string,
  }
}

export class Memo extends jspb.Message {
  getId(): number;
  setId(value: number): Memo;

  getContent(): string;
  setContent(value: string): Memo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Memo.AsObject;
  static toObject(includeInstance: boolean, msg: Memo): Memo.AsObject;
  static serializeBinaryToWriter(message: Memo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Memo;
  static deserializeBinaryFromReader(message: Memo, reader: jspb.BinaryReader): Memo;
}

export namespace Memo {
  export type AsObject = {
    id: number,
    content: string,
  }
}

export class Label extends jspb.Message {
  getId(): number;
  setId(value: number): Label;

  getName(): string;
  setName(value: string): Label;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Label.AsObject;
  static toObject(includeInstance: boolean, msg: Label): Label.AsObject;
  static serializeBinaryToWriter(message: Label, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Label;
  static deserializeBinaryFromReader(message: Label, reader: jspb.BinaryReader): Label;
}

export namespace Label {
  export type AsObject = {
    id: number,
    name: string,
  }
}

export class Doctor extends jspb.Message {
  getId(): number;
  setId(value: number): Doctor;

  getName(): string;
  setName(value: string): Doctor;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Doctor.AsObject;
  static toObject(includeInstance: boolean, msg: Doctor): Doctor.AsObject;
  static serializeBinaryToWriter(message: Doctor, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Doctor;
  static deserializeBinaryFromReader(message: Doctor, reader: jspb.BinaryReader): Doctor;
}

export namespace Doctor {
  export type AsObject = {
    id: number,
    name: string,
  }
}

export class Reception extends jspb.Message {
  getId(): number;
  setId(value: number): Reception;

  getAppointmentTime(): string;
  setAppointmentTime(value: string): Reception;

  getReceptionTime(): string;
  setReceptionTime(value: string): Reception;

  getStatus(): number;
  setStatus(value: number): Reception;

  getPatient(): Patient | undefined;
  setPatient(value?: Patient): Reception;
  hasPatient(): boolean;
  clearPatient(): Reception;

  getExamination(): Examination | undefined;
  setExamination(value?: Examination): Reception;
  hasExamination(): boolean;
  clearExamination(): Reception;

  getDoctor(): Doctor | undefined;
  setDoctor(value?: Doctor): Reception;
  hasDoctor(): boolean;
  clearDoctor(): Reception;

  getMemo(): Memo | undefined;
  setMemo(value?: Memo): Reception;
  hasMemo(): boolean;
  clearMemo(): Reception;

  getLabelsList(): Array<Label>;
  setLabelsList(value: Array<Label>): Reception;
  clearLabelsList(): Reception;
  addLabels(value?: Label, index?: number): Label;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Reception.AsObject;
  static toObject(includeInstance: boolean, msg: Reception): Reception.AsObject;
  static serializeBinaryToWriter(message: Reception, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Reception;
  static deserializeBinaryFromReader(message: Reception, reader: jspb.BinaryReader): Reception;
}

export namespace Reception {
  export type AsObject = {
    id: number,
    appointmentTime: string,
    receptionTime: string,
    status: number,
    patient?: Patient.AsObject,
    examination?: Examination.AsObject,
    doctor?: Doctor.AsObject,
    memo?: Memo.AsObject,
    labelsList: Array<Label.AsObject>,
  }
}

