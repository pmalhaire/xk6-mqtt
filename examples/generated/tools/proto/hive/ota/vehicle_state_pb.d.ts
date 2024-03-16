// package: hive.vehicle_state
// file: tools/proto/hive/ota/vehicle_state.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class CameraState extends jspb.Message {
  getActiveStatus(): CameraState.CameraActiveStatusMap[keyof CameraState.CameraActiveStatusMap];
  setActiveStatus(value: CameraState.CameraActiveStatusMap[keyof CameraState.CameraActiveStatusMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CameraState.AsObject;
  static toObject(includeInstance: boolean, msg: CameraState): CameraState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CameraState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CameraState;
  static deserializeBinaryFromReader(message: CameraState, reader: jspb.BinaryReader): CameraState;
}

export namespace CameraState {
  export type AsObject = {
    activeStatus: CameraState.CameraActiveStatusMap[keyof CameraState.CameraActiveStatusMap],
  }

  export interface CameraActiveStatusMap {
    CAMERA_ACTIVE_STATUS_UNSET: 0;
    ACTIVE: 1;
    INACTIVE: 2;
  }

  export const CameraActiveStatus: CameraActiveStatusMap;
}

export class ChargeState extends jspb.Message {
  getChargingStatus(): ChargeState.ChargingStatusMap[keyof ChargeState.ChargingStatusMap];
  setChargingStatus(value: ChargeState.ChargingStatusMap[keyof ChargeState.ChargingStatusMap]): void;

  getChargePercentage(): number;
  setChargePercentage(value: number): void;

  getChargingRateKw(): number;
  setChargingRateKw(value: number): void;

  getRangeKm(): number;
  setRangeKm(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChargeState.AsObject;
  static toObject(includeInstance: boolean, msg: ChargeState): ChargeState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ChargeState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChargeState;
  static deserializeBinaryFromReader(message: ChargeState, reader: jspb.BinaryReader): ChargeState;
}

export namespace ChargeState {
  export type AsObject = {
    chargingStatus: ChargeState.ChargingStatusMap[keyof ChargeState.ChargingStatusMap],
    chargePercentage: number,
    chargingRateKw: number,
    rangeKm: number,
  }

  export interface ChargingStatusMap {
    UNSET: 0;
    NOT_CHARGING: 1;
    CHARGING: 2;
  }

  export const ChargingStatus: ChargingStatusMap;
}

export class DoorState extends jspb.Message {
  getLockState(): DoorState.DoorLockStateMap[keyof DoorState.DoorLockStateMap];
  setLockState(value: DoorState.DoorLockStateMap[keyof DoorState.DoorLockStateMap]): void;

  getOpenCloseState(): DoorState.DoorOpenCloseStateMap[keyof DoorState.DoorOpenCloseStateMap];
  setOpenCloseState(value: DoorState.DoorOpenCloseStateMap[keyof DoorState.DoorOpenCloseStateMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DoorState.AsObject;
  static toObject(includeInstance: boolean, msg: DoorState): DoorState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DoorState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DoorState;
  static deserializeBinaryFromReader(message: DoorState, reader: jspb.BinaryReader): DoorState;
}

export namespace DoorState {
  export type AsObject = {
    lockState: DoorState.DoorLockStateMap[keyof DoorState.DoorLockStateMap],
    openCloseState: DoorState.DoorOpenCloseStateMap[keyof DoorState.DoorOpenCloseStateMap],
  }

  export interface DoorLockStateMap {
    LOCK_UNSET: 0;
    UNLOCKED: 1;
    LOCKED: 2;
  }

  export const DoorLockState: DoorLockStateMap;

  export interface DoorOpenCloseStateMap {
    OPENCLOSE_UNSET: 0;
    OPEN: 1;
    CLOSED: 2;
  }

  export const DoorOpenCloseState: DoorOpenCloseStateMap;
}

export class DriveState extends jspb.Message {
  getGearState(): DriveState.GearStateMap[keyof DriveState.GearStateMap];
  setGearState(value: DriveState.GearStateMap[keyof DriveState.GearStateMap]): void;

  getVelocity(): number;
  setVelocity(value: number): void;

  getOdometerKm(): number;
  setOdometerKm(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DriveState.AsObject;
  static toObject(includeInstance: boolean, msg: DriveState): DriveState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DriveState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DriveState;
  static deserializeBinaryFromReader(message: DriveState, reader: jspb.BinaryReader): DriveState;
}

export namespace DriveState {
  export type AsObject = {
    gearState: DriveState.GearStateMap[keyof DriveState.GearStateMap],
    velocity: number,
    odometerKm: number,
  }

  export interface GearStateMap {
    UNSET: 0;
    PARK: 1;
    REVERSE: 2;
    NEUTRAL: 3;
    DRIVE: 4;
  }

  export const GearState: GearStateMap;
}

export class GPSCoordinates extends jspb.Message {
  getLatitudeDegrees(): number;
  setLatitudeDegrees(value: number): void;

  getLongitudeDegrees(): number;
  setLongitudeDegrees(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GPSCoordinates.AsObject;
  static toObject(includeInstance: boolean, msg: GPSCoordinates): GPSCoordinates.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GPSCoordinates, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GPSCoordinates;
  static deserializeBinaryFromReader(message: GPSCoordinates, reader: jspb.BinaryReader): GPSCoordinates;
}

export namespace GPSCoordinates {
  export type AsObject = {
    latitudeDegrees: number,
    longitudeDegrees: number,
  }
}

export class HVACState extends jspb.Message {
  getAirCirculationButtonPressed(): boolean;
  setAirCirculationButtonPressed(value: boolean): void;

  hasFrontLeft(): boolean;
  clearFrontLeft(): void;
  getFrontLeft(): HVACState.SeatState | undefined;
  setFrontLeft(value?: HVACState.SeatState): void;

  hasFrontRight(): boolean;
  clearFrontRight(): void;
  getFrontRight(): HVACState.SeatState | undefined;
  setFrontRight(value?: HVACState.SeatState): void;

  hasBackLeft(): boolean;
  clearBackLeft(): void;
  getBackLeft(): HVACState.SeatState | undefined;
  setBackLeft(value?: HVACState.SeatState): void;

  hasBackRight(): boolean;
  clearBackRight(): void;
  getBackRight(): HVACState.SeatState | undefined;
  setBackRight(value?: HVACState.SeatState): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HVACState.AsObject;
  static toObject(includeInstance: boolean, msg: HVACState): HVACState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HVACState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HVACState;
  static deserializeBinaryFromReader(message: HVACState, reader: jspb.BinaryReader): HVACState;
}

export namespace HVACState {
  export type AsObject = {
    airCirculationButtonPressed: boolean,
    frontLeft?: HVACState.SeatState.AsObject,
    frontRight?: HVACState.SeatState.AsObject,
    backLeft?: HVACState.SeatState.AsObject,
    backRight?: HVACState.SeatState.AsObject,
  }

  export class SeatState extends jspb.Message {
    getAirTemperatureCelsius(): number;
    setAirTemperatureCelsius(value: number): void;

    getSeatHeating(): HVACState.SeatState.SeatHVACLevelMap[keyof HVACState.SeatState.SeatHVACLevelMap];
    setSeatHeating(value: HVACState.SeatState.SeatHVACLevelMap[keyof HVACState.SeatState.SeatHVACLevelMap]): void;

    getSeatVentilation(): HVACState.SeatState.SeatHVACLevelMap[keyof HVACState.SeatState.SeatHVACLevelMap];
    setSeatVentilation(value: HVACState.SeatState.SeatHVACLevelMap[keyof HVACState.SeatState.SeatHVACLevelMap]): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SeatState.AsObject;
    static toObject(includeInstance: boolean, msg: SeatState): SeatState.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SeatState, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SeatState;
    static deserializeBinaryFromReader(message: SeatState, reader: jspb.BinaryReader): SeatState;
  }

  export namespace SeatState {
    export type AsObject = {
      airTemperatureCelsius: number,
      seatHeating: HVACState.SeatState.SeatHVACLevelMap[keyof HVACState.SeatState.SeatHVACLevelMap],
      seatVentilation: HVACState.SeatState.SeatHVACLevelMap[keyof HVACState.SeatState.SeatHVACLevelMap],
    }

    export interface SeatHVACLevelMap {
      UNSET: 0;
      OFF: 1;
      LEVEL_1: 2;
      LEVEL_2: 3;
      LEVEL_3: 4;
    }

    export const SeatHVACLevel: SeatHVACLevelMap;
  }
}

export class VehicleLightsState extends jspb.Message {
  hasHeadlightsState(): boolean;
  clearHeadlightsState(): void;
  getHeadlightsState(): VehicleLightsState.HeadlightsState | undefined;
  setHeadlightsState(value?: VehicleLightsState.HeadlightsState): void;

  hasTaillightsState(): boolean;
  clearTaillightsState(): void;
  getTaillightsState(): VehicleLightsState.TaillightsState | undefined;
  setTaillightsState(value?: VehicleLightsState.TaillightsState): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VehicleLightsState.AsObject;
  static toObject(includeInstance: boolean, msg: VehicleLightsState): VehicleLightsState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: VehicleLightsState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VehicleLightsState;
  static deserializeBinaryFromReader(message: VehicleLightsState, reader: jspb.BinaryReader): VehicleLightsState;
}

export namespace VehicleLightsState {
  export type AsObject = {
    headlightsState?: VehicleLightsState.HeadlightsState.AsObject,
    taillightsState?: VehicleLightsState.TaillightsState.AsObject,
  }

  export class HeadlightsState extends jspb.Message {
    getFrontLeftDrl(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setFrontLeftDrl(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getFrontLeftBlinker(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setFrontLeftBlinker(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getFrontLeftHigh(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setFrontLeftHigh(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getFrontLeftLow(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setFrontLeftLow(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getFrontRightDrl(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setFrontRightDrl(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getFrontRightBlinker(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setFrontRightBlinker(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getFrontRightHigh(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setFrontRightHigh(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getFrontRightLow(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setFrontRightLow(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): HeadlightsState.AsObject;
    static toObject(includeInstance: boolean, msg: HeadlightsState): HeadlightsState.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: HeadlightsState, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): HeadlightsState;
    static deserializeBinaryFromReader(message: HeadlightsState, reader: jspb.BinaryReader): HeadlightsState;
  }

  export namespace HeadlightsState {
    export type AsObject = {
      frontLeftDrl: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      frontLeftBlinker: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      frontLeftHigh: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      frontLeftLow: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      frontRightDrl: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      frontRightBlinker: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      frontRightHigh: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      frontRightLow: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
    }
  }

  export class TaillightsState extends jspb.Message {
    getBackLeftDrl(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setBackLeftDrl(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getBackLeftBlinker(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setBackLeftBlinker(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getBackLeftBrake(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setBackLeftBrake(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getBackRightDrl(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setBackRightDrl(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getBackRightBlinker(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setBackRightBlinker(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getBackRightBrake(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setBackRightBrake(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getBackCenterFog(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setBackCenterFog(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getBackCenterReverse(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setBackCenterReverse(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getBackCenterLeftDrl(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setBackCenterLeftDrl(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getBackCenterRightDrl(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setBackCenterRightDrl(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getBackCenterLeftBrake(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setBackCenterLeftBrake(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    getBackCenterRightBrake(): VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap];
    setBackCenterRightBrake(value: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap]): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TaillightsState.AsObject;
    static toObject(includeInstance: boolean, msg: TaillightsState): TaillightsState.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TaillightsState, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TaillightsState;
    static deserializeBinaryFromReader(message: TaillightsState, reader: jspb.BinaryReader): TaillightsState;
  }

  export namespace TaillightsState {
    export type AsObject = {
      backLeftDrl: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      backLeftBlinker: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      backLeftBrake: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      backRightDrl: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      backRightBlinker: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      backRightBrake: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      backCenterFog: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      backCenterReverse: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      backCenterLeftDrl: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      backCenterRightDrl: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      backCenterLeftBrake: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
      backCenterRightBrake: VehicleLightsState.ExteriorLightStateMap[keyof VehicleLightsState.ExteriorLightStateMap],
    }
  }

  export interface ExteriorLightStateMap {
    UNSET: 0;
    ON: 1;
    OFF: 2;
  }

  export const ExteriorLightState: ExteriorLightStateMap;
}

export class VehicleCamerasState extends jspb.Message {
  hasLeftSideCameraState(): boolean;
  clearLeftSideCameraState(): void;
  getLeftSideCameraState(): CameraState | undefined;
  setLeftSideCameraState(value?: CameraState): void;

  hasRightSideCameraState(): boolean;
  clearRightSideCameraState(): void;
  getRightSideCameraState(): CameraState | undefined;
  setRightSideCameraState(value?: CameraState): void;

  hasFrontCameraState(): boolean;
  clearFrontCameraState(): void;
  getFrontCameraState(): CameraState | undefined;
  setFrontCameraState(value?: CameraState): void;

  hasRearCameraState(): boolean;
  clearRearCameraState(): void;
  getRearCameraState(): CameraState | undefined;
  setRearCameraState(value?: CameraState): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VehicleCamerasState.AsObject;
  static toObject(includeInstance: boolean, msg: VehicleCamerasState): VehicleCamerasState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: VehicleCamerasState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VehicleCamerasState;
  static deserializeBinaryFromReader(message: VehicleCamerasState, reader: jspb.BinaryReader): VehicleCamerasState;
}

export namespace VehicleCamerasState {
  export type AsObject = {
    leftSideCameraState?: CameraState.AsObject,
    rightSideCameraState?: CameraState.AsObject,
    frontCameraState?: CameraState.AsObject,
    rearCameraState?: CameraState.AsObject,
  }
}

export class TargetInfo extends jspb.Message {
  getHealthStatus(): TargetInfo.HealthStatusMap[keyof TargetInfo.HealthStatusMap];
  setHealthStatus(value: TargetInfo.HealthStatusMap[keyof TargetInfo.HealthStatusMap]): void;

  getVersion(): string;
  setVersion(value: string): void;

  getBoard(): string;
  setBoard(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TargetInfo.AsObject;
  static toObject(includeInstance: boolean, msg: TargetInfo): TargetInfo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TargetInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TargetInfo;
  static deserializeBinaryFromReader(message: TargetInfo, reader: jspb.BinaryReader): TargetInfo;
}

export namespace TargetInfo {
  export type AsObject = {
    healthStatus: TargetInfo.HealthStatusMap[keyof TargetInfo.HealthStatusMap],
    version: string,
    board: string,
  }

  export interface HealthStatusMap {
    UNSET: 0;
    HEALTHY: 1;
    UNHEALTHY: 2;
  }

  export const HealthStatus: HealthStatusMap;
}

export class TargetsState extends jspb.Message {
  getNameToTargetInfoMapMap(): jspb.Map<string, TargetInfo>;
  clearNameToTargetInfoMapMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TargetsState.AsObject;
  static toObject(includeInstance: boolean, msg: TargetsState): TargetsState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TargetsState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TargetsState;
  static deserializeBinaryFromReader(message: TargetsState, reader: jspb.BinaryReader): TargetsState;
}

export namespace TargetsState {
  export type AsObject = {
    nameToTargetInfoMapMap: Array<[string, TargetInfo.AsObject]>,
  }
}

export class VehicleDoorsState extends jspb.Message {
  hasFrontLeft(): boolean;
  clearFrontLeft(): void;
  getFrontLeft(): DoorState | undefined;
  setFrontLeft(value?: DoorState): void;

  hasFrontRight(): boolean;
  clearFrontRight(): void;
  getFrontRight(): DoorState | undefined;
  setFrontRight(value?: DoorState): void;

  hasBackLeft(): boolean;
  clearBackLeft(): void;
  getBackLeft(): DoorState | undefined;
  setBackLeft(value?: DoorState): void;

  hasBackRight(): boolean;
  clearBackRight(): void;
  getBackRight(): DoorState | undefined;
  setBackRight(value?: DoorState): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VehicleDoorsState.AsObject;
  static toObject(includeInstance: boolean, msg: VehicleDoorsState): VehicleDoorsState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: VehicleDoorsState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VehicleDoorsState;
  static deserializeBinaryFromReader(message: VehicleDoorsState, reader: jspb.BinaryReader): VehicleDoorsState;
}

export namespace VehicleDoorsState {
  export type AsObject = {
    frontLeft?: DoorState.AsObject,
    frontRight?: DoorState.AsObject,
    backLeft?: DoorState.AsObject,
    backRight?: DoorState.AsObject,
  }
}

export class VehicleTiresState extends jspb.Message {
  hasFrontLeft(): boolean;
  clearFrontLeft(): void;
  getFrontLeft(): TireState | undefined;
  setFrontLeft(value?: TireState): void;

  hasFrontRight(): boolean;
  clearFrontRight(): void;
  getFrontRight(): TireState | undefined;
  setFrontRight(value?: TireState): void;

  hasBackLeft(): boolean;
  clearBackLeft(): void;
  getBackLeft(): TireState | undefined;
  setBackLeft(value?: TireState): void;

  hasBackRight(): boolean;
  clearBackRight(): void;
  getBackRight(): TireState | undefined;
  setBackRight(value?: TireState): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VehicleTiresState.AsObject;
  static toObject(includeInstance: boolean, msg: VehicleTiresState): VehicleTiresState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: VehicleTiresState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VehicleTiresState;
  static deserializeBinaryFromReader(message: VehicleTiresState, reader: jspb.BinaryReader): VehicleTiresState;
}

export namespace VehicleTiresState {
  export type AsObject = {
    frontLeft?: TireState.AsObject,
    frontRight?: TireState.AsObject,
    backLeft?: TireState.AsObject,
    backRight?: TireState.AsObject,
  }
}

export class TireState extends jspb.Message {
  getTirePressureBar(): number;
  setTirePressureBar(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TireState.AsObject;
  static toObject(includeInstance: boolean, msg: TireState): TireState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TireState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TireState;
  static deserializeBinaryFromReader(message: TireState, reader: jspb.BinaryReader): TireState;
}

export namespace TireState {
  export type AsObject = {
    tirePressureBar: number,
  }
}

export class SentryModeState extends jspb.Message {
  getIsEnabled(): boolean;
  setIsEnabled(value: boolean): void;

  getSentryModeStatus(): SentryModeState.StatusMap[keyof SentryModeState.StatusMap];
  setSentryModeStatus(value: SentryModeState.StatusMap[keyof SentryModeState.StatusMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SentryModeState.AsObject;
  static toObject(includeInstance: boolean, msg: SentryModeState): SentryModeState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SentryModeState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SentryModeState;
  static deserializeBinaryFromReader(message: SentryModeState, reader: jspb.BinaryReader): SentryModeState;
}

export namespace SentryModeState {
  export type AsObject = {
    isEnabled: boolean,
    sentryModeStatus: SentryModeState.StatusMap[keyof SentryModeState.StatusMap],
  }

  export interface StatusMap {
    UNSET: 0;
    STANDBY: 1;
    ALERT: 2;
    ALARM: 3;
    OFF: 4;
  }

  export const Status: StatusMap;
}

export class ValetModeState extends jspb.Message {
  getValetModeStatus(): ValetModeState.StatusMap[keyof ValetModeState.StatusMap];
  setValetModeStatus(value: ValetModeState.StatusMap[keyof ValetModeState.StatusMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ValetModeState.AsObject;
  static toObject(includeInstance: boolean, msg: ValetModeState): ValetModeState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ValetModeState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ValetModeState;
  static deserializeBinaryFromReader(message: ValetModeState, reader: jspb.BinaryReader): ValetModeState;
}

export namespace ValetModeState {
  export type AsObject = {
    valetModeStatus: ValetModeState.StatusMap[keyof ValetModeState.StatusMap],
  }

  export interface StatusMap {
    UNSET: 0;
    OFF: 1;
    ON: 2;
    LOCKED: 3;
  }

  export const Status: StatusMap;
}

export class DriveModeState extends jspb.Message {
  getMode(): DriveModeState.ModeMap[keyof DriveModeState.ModeMap];
  setMode(value: DriveModeState.ModeMap[keyof DriveModeState.ModeMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DriveModeState.AsObject;
  static toObject(includeInstance: boolean, msg: DriveModeState): DriveModeState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DriveModeState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DriveModeState;
  static deserializeBinaryFromReader(message: DriveModeState, reader: jspb.BinaryReader): DriveModeState;
}

export namespace DriveModeState {
  export type AsObject = {
    mode: DriveModeState.ModeMap[keyof DriveModeState.ModeMap],
  }

  export interface ModeMap {
    UNSET: 0;
    NORMAL: 1;
    ECO: 2;
    SPORT: 3;
    RACETRACK: 4;
  }

  export const Mode: ModeMap;
}

export class VehicleOperatingState extends jspb.Message {
  getOperatingState(): VehicleOperatingState.OperatingStateMap[keyof VehicleOperatingState.OperatingStateMap];
  setOperatingState(value: VehicleOperatingState.OperatingStateMap[keyof VehicleOperatingState.OperatingStateMap]): void;

  getTransitionStatus(): VehicleOperatingState.TransitionStatusMap[keyof VehicleOperatingState.TransitionStatusMap];
  setTransitionStatus(value: VehicleOperatingState.TransitionStatusMap[keyof VehicleOperatingState.TransitionStatusMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VehicleOperatingState.AsObject;
  static toObject(includeInstance: boolean, msg: VehicleOperatingState): VehicleOperatingState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: VehicleOperatingState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VehicleOperatingState;
  static deserializeBinaryFromReader(message: VehicleOperatingState, reader: jspb.BinaryReader): VehicleOperatingState;
}

export namespace VehicleOperatingState {
  export type AsObject = {
    operatingState: VehicleOperatingState.OperatingStateMap[keyof VehicleOperatingState.OperatingStateMap],
    transitionStatus: VehicleOperatingState.TransitionStatusMap[keyof VehicleOperatingState.TransitionStatusMap],
  }

  export interface OperatingStateMap {
    OPERATING_UNSET: 0;
    DEEP_SLEEP: 1;
    SLEEP: 2;
    ACCESSORY_READY: 3;
    ACCESSORY: 4;
    DRIVE_READY: 5;
    DRIVE: 6;
    SLEEP_ACTIVE_OTA: 7;
    SLEEP_ACTIVE_SENTRY: 8;
    SLEEP_ACTIVE_REMOTE: 9;
  }

  export const OperatingState: OperatingStateMap;

  export interface TransitionStatusMap {
    TRANSITIONING_UNSET: 0;
    NOT_TRANSITIONING: 1;
    TRANSITIONING: 2;
    UNKNOWN: 3;
  }

  export const TransitionStatus: TransitionStatusMap;
}

export class TurnSignalState extends jspb.Message {
  getState(): TurnSignalState.StateMap[keyof TurnSignalState.StateMap];
  setState(value: TurnSignalState.StateMap[keyof TurnSignalState.StateMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TurnSignalState.AsObject;
  static toObject(includeInstance: boolean, msg: TurnSignalState): TurnSignalState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TurnSignalState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TurnSignalState;
  static deserializeBinaryFromReader(message: TurnSignalState, reader: jspb.BinaryReader): TurnSignalState;
}

export namespace TurnSignalState {
  export type AsObject = {
    state: TurnSignalState.StateMap[keyof TurnSignalState.StateMap],
  }

  export interface StateMap {
    UNSET: 0;
    INACTIVE: 1;
    SIGNALLING_LEFT: 2;
    SIGNALLING_RIGHT: 3;
  }

  export const State: StateMap;
}

export class ThermalManagementState extends jspb.Message {
  getFanRpm(): number;
  setFanRpm(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ThermalManagementState.AsObject;
  static toObject(includeInstance: boolean, msg: ThermalManagementState): ThermalManagementState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ThermalManagementState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ThermalManagementState;
  static deserializeBinaryFromReader(message: ThermalManagementState, reader: jspb.BinaryReader): ThermalManagementState;
}

export namespace ThermalManagementState {
  export type AsObject = {
    fanRpm: number,
  }
}

export class SpoilerState extends jspb.Message {
  getPosition(): number;
  setPosition(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SpoilerState.AsObject;
  static toObject(includeInstance: boolean, msg: SpoilerState): SpoilerState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SpoilerState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SpoilerState;
  static deserializeBinaryFromReader(message: SpoilerState, reader: jspb.BinaryReader): SpoilerState;
}

export namespace SpoilerState {
  export type AsObject = {
    position: number,
  }
}

export class VehicleState extends jspb.Message {
  getDoorsLocked(): boolean;
  setDoorsLocked(value: boolean): void;

  hasDriveState(): boolean;
  clearDriveState(): void;
  getDriveState(): DriveState | undefined;
  setDriveState(value?: DriveState): void;

  hasDoorsState(): boolean;
  clearDoorsState(): void;
  getDoorsState(): VehicleDoorsState | undefined;
  setDoorsState(value?: VehicleDoorsState): void;

  hasTimestampUtc(): boolean;
  clearTimestampUtc(): void;
  getTimestampUtc(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTimestampUtc(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasChargeState(): boolean;
  clearChargeState(): void;
  getChargeState(): ChargeState | undefined;
  setChargeState(value?: ChargeState): void;

  hasGpsCoordinates(): boolean;
  clearGpsCoordinates(): void;
  getGpsCoordinates(): GPSCoordinates | undefined;
  setGpsCoordinates(value?: GPSCoordinates): void;

  hasVehicleLightsState(): boolean;
  clearVehicleLightsState(): void;
  getVehicleLightsState(): VehicleLightsState | undefined;
  setVehicleLightsState(value?: VehicleLightsState): void;

  hasHvacState(): boolean;
  clearHvacState(): void;
  getHvacState(): HVACState | undefined;
  setHvacState(value?: HVACState): void;

  hasSentryModeState(): boolean;
  clearSentryModeState(): void;
  getSentryModeState(): SentryModeState | undefined;
  setSentryModeState(value?: SentryModeState): void;

  hasValetModeState(): boolean;
  clearValetModeState(): void;
  getValetModeState(): ValetModeState | undefined;
  setValetModeState(value?: ValetModeState): void;

  hasTiresState(): boolean;
  clearTiresState(): void;
  getTiresState(): VehicleTiresState | undefined;
  setTiresState(value?: VehicleTiresState): void;

  hasDriveModeState(): boolean;
  clearDriveModeState(): void;
  getDriveModeState(): DriveModeState | undefined;
  setDriveModeState(value?: DriveModeState): void;

  hasTurnSignalState(): boolean;
  clearTurnSignalState(): void;
  getTurnSignalState(): TurnSignalState | undefined;
  setTurnSignalState(value?: TurnSignalState): void;

  hasTmeState(): boolean;
  clearTmeState(): void;
  getTmeState(): ThermalManagementState | undefined;
  setTmeState(value?: ThermalManagementState): void;

  hasCamerasState(): boolean;
  clearCamerasState(): void;
  getCamerasState(): VehicleCamerasState | undefined;
  setCamerasState(value?: VehicleCamerasState): void;

  hasTargetsState(): boolean;
  clearTargetsState(): void;
  getTargetsState(): TargetsState | undefined;
  setTargetsState(value?: TargetsState): void;

  hasSpoilerState(): boolean;
  clearSpoilerState(): void;
  getSpoilerState(): SpoilerState | undefined;
  setSpoilerState(value?: SpoilerState): void;

  hasVehicleOperatingState(): boolean;
  clearVehicleOperatingState(): void;
  getVehicleOperatingState(): VehicleOperatingState | undefined;
  setVehicleOperatingState(value?: VehicleOperatingState): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VehicleState.AsObject;
  static toObject(includeInstance: boolean, msg: VehicleState): VehicleState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: VehicleState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VehicleState;
  static deserializeBinaryFromReader(message: VehicleState, reader: jspb.BinaryReader): VehicleState;
}

export namespace VehicleState {
  export type AsObject = {
    doorsLocked: boolean,
    driveState?: DriveState.AsObject,
    doorsState?: VehicleDoorsState.AsObject,
    timestampUtc?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    chargeState?: ChargeState.AsObject,
    gpsCoordinates?: GPSCoordinates.AsObject,
    vehicleLightsState?: VehicleLightsState.AsObject,
    hvacState?: HVACState.AsObject,
    sentryModeState?: SentryModeState.AsObject,
    valetModeState?: ValetModeState.AsObject,
    tiresState?: VehicleTiresState.AsObject,
    driveModeState?: DriveModeState.AsObject,
    turnSignalState?: TurnSignalState.AsObject,
    tmeState?: ThermalManagementState.AsObject,
    camerasState?: VehicleCamerasState.AsObject,
    targetsState?: TargetsState.AsObject,
    spoilerState?: SpoilerState.AsObject,
    vehicleOperatingState?: VehicleOperatingState.AsObject,
  }
}

