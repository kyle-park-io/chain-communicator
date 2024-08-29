// package: types
// file: trx.proto

import * as jspb from 'google-protobuf';

export class TrxProto extends jspb.Message {
  getVersion(): number;
  setVersion(value: number): void;

  getTime(): number;
  setTime(value: number): void;

  getNonce(): number;
  setNonce(value: number): void;

  getFrom(): Uint8Array | string;
  getFrom_asU8(): Uint8Array;
  getFrom_asB64(): string;
  setFrom(value: Uint8Array | string): void;

  getTo(): Uint8Array | string;
  getTo_asU8(): Uint8Array;
  getTo_asB64(): string;
  setTo(value: Uint8Array | string): void;

  getAmount(): Uint8Array | string;
  getAmount_asU8(): Uint8Array;
  getAmount_asB64(): string;
  setAmount(value: Uint8Array | string): void;

  getGas(): number;
  setGas(value: number): void;

  getGasprice(): Uint8Array | string;
  getGasprice_asU8(): Uint8Array;
  getGasprice_asB64(): string;
  setGasprice(value: Uint8Array | string): void;

  getType(): number;
  setType(value: number): void;

  getPayload(): Uint8Array | string;
  getPayload_asU8(): Uint8Array;
  getPayload_asB64(): string;
  setPayload(value: Uint8Array | string): void;

  getSig(): Uint8Array | string;
  getSig_asU8(): Uint8Array;
  getSig_asB64(): string;
  setSig(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TrxProto.AsObject;
  static toObject(includeInstance: boolean, msg: TrxProto): TrxProto.AsObject;
  static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
  static extensionsBinary: {
    [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>;
  };
  static serializeBinaryToWriter(
    message: TrxProto,
    writer: jspb.BinaryWriter,
  ): void;
  static deserializeBinary(bytes: Uint8Array): TrxProto;
  static deserializeBinaryFromReader(
    message: TrxProto,
    reader: jspb.BinaryReader,
  ): TrxProto;
}

export namespace TrxProto {
  export type AsObject = {
    version: number;
    time: number;
    nonce: number;
    from: Uint8Array | string;
    to: Uint8Array | string;
    amount: Uint8Array | string;
    gas: number;
    gasprice: Uint8Array | string;
    type: number;
    payload: Uint8Array | string;
    sig: Uint8Array | string;
  };
}

export class TrxPayloadContractProto extends jspb.Message {
  getData(): Uint8Array | string;
  getData_asU8(): Uint8Array;
  getData_asB64(): string;
  setData(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TrxPayloadContractProto.AsObject;
  static toObject(
    includeInstance: boolean,
    msg: TrxPayloadContractProto,
  ): TrxPayloadContractProto.AsObject;
  static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
  static extensionsBinary: {
    [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>;
  };
  static serializeBinaryToWriter(
    message: TrxPayloadContractProto,
    writer: jspb.BinaryWriter,
  ): void;
  static deserializeBinary(bytes: Uint8Array): TrxPayloadContractProto;
  static deserializeBinaryFromReader(
    message: TrxPayloadContractProto,
    reader: jspb.BinaryReader,
  ): TrxPayloadContractProto;
}

export namespace TrxPayloadContractProto {
  export type AsObject = {
    data: Uint8Array | string;
  };
}
