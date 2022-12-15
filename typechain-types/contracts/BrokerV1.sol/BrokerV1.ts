/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BigNumberish,
  BytesLike,
  CallOverrides,
  ContractTransaction,
  Overrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import type {
  FunctionFragment,
  Result,
  EventFragment,
} from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "../../common";

export declare namespace BrokerV1 {
  export type BookingStruct = {
    index: PromiseOrValue<BigNumberish>;
    deprecated__vmTypeId: PromiseOrValue<BigNumberish>;
    miner: PromiseOrValue<string>;
    user: PromiseOrValue<string>;
    pricePerSecond: PromiseOrValue<BigNumberish>;
    bookedAt: PromiseOrValue<BigNumberish>;
    lastPayment: PromiseOrValue<BigNumberish>;
    offerIndex: PromiseOrValue<BigNumberish>;
  };

  export type BookingStructOutput = [
    BigNumber,
    BigNumber,
    string,
    string,
    BigNumber,
    BigNumber,
    BigNumber,
    BigNumber
  ] & {
    index: BigNumber;
    deprecated__vmTypeId: BigNumber;
    miner: string;
    user: string;
    pricePerSecond: BigNumber;
    bookedAt: BigNumber;
    lastPayment: BigNumber;
    offerIndex: BigNumber;
  };

  export type OfferStruct = {
    index: PromiseOrValue<BigNumberish>;
    miner: PromiseOrValue<string>;
    pricePerSecond: PromiseOrValue<BigNumberish>;
    machinesAvailable: PromiseOrValue<BigNumberish>;
    deprecated__vmTypeId: PromiseOrValue<BigNumberish>;
    specsIpfsHash: PromiseOrValue<BytesLike>;
  };

  export type OfferStructOutput = [
    BigNumber,
    string,
    BigNumber,
    BigNumber,
    BigNumber,
    string
  ] & {
    index: BigNumber;
    miner: string;
    pricePerSecond: BigNumber;
    machinesAvailable: BigNumber;
    deprecated__vmTypeId: BigNumber;
    specsIpfsHash: string;
  };
}

export interface BrokerV1Interface extends utils.Interface {
  functions: {
    "AddOffer(uint64,uint64,bytes32)": FunctionFragment;
    "Book(uint64)": FunctionFragment;
    "ClaimPayment(uint64)": FunctionFragment;
    "DepositCoin(uint256)": FunctionFragment;
    "FindBookingsByMiner(address)": FunctionFragment;
    "FindBookingsByUser(address)": FunctionFragment;
    "GetAvailableOffers()": FunctionFragment;
    "GetBooking(uint64)": FunctionFragment;
    "GetCoinBalance(address)": FunctionFragment;
    "GetMinerUrl(address)": FunctionFragment;
    "GetMinersOffers(address)": FunctionFragment;
    "GetTime()": FunctionFragment;
    "RemoveOffer(uint64)": FunctionFragment;
    "SECONDS_IN_WEEK()": FunctionFragment;
    "SetCoinAddress(address)": FunctionFragment;
    "SetCommunityContract(address)": FunctionFragment;
    "SetCommunityFee(uint64)": FunctionFragment;
    "SetMinerUrl(bytes32)": FunctionFragment;
    "Terminate(uint64,uint8)": FunctionFragment;
    "UpdateOffer(uint64,uint64,uint64)": FunctionFragment;
    "WithdrawCoin(uint256)": FunctionFragment;
    "coin()": FunctionFragment;
    "communityContract()": FunctionFragment;
    "communityFee()": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "AddOffer"
      | "Book"
      | "ClaimPayment"
      | "DepositCoin"
      | "FindBookingsByMiner"
      | "FindBookingsByUser"
      | "GetAvailableOffers"
      | "GetBooking"
      | "GetCoinBalance"
      | "GetMinerUrl"
      | "GetMinersOffers"
      | "GetTime"
      | "RemoveOffer"
      | "SECONDS_IN_WEEK"
      | "SetCoinAddress"
      | "SetCommunityContract"
      | "SetCommunityFee"
      | "SetMinerUrl"
      | "Terminate"
      | "UpdateOffer"
      | "WithdrawCoin"
      | "coin"
      | "communityContract"
      | "communityFee"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "AddOffer",
    values: [
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "Book",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "ClaimPayment",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "DepositCoin",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "FindBookingsByMiner",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "FindBookingsByUser",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "GetAvailableOffers",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "GetBooking",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "GetCoinBalance",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "GetMinerUrl",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "GetMinersOffers",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(functionFragment: "GetTime", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "RemoveOffer",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "SECONDS_IN_WEEK",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "SetCoinAddress",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "SetCommunityContract",
    values: [PromiseOrValue<string>]
  ): string;
  encodeFunctionData(
    functionFragment: "SetCommunityFee",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "SetMinerUrl",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "Terminate",
    values: [PromiseOrValue<BigNumberish>, PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "UpdateOffer",
    values: [
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "WithdrawCoin",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(functionFragment: "coin", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "communityContract",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "communityFee",
    values?: undefined
  ): string;

  decodeFunctionResult(functionFragment: "AddOffer", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "Book", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "ClaimPayment",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "DepositCoin",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "FindBookingsByMiner",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "FindBookingsByUser",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "GetAvailableOffers",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "GetBooking", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "GetCoinBalance",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "GetMinerUrl",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "GetMinersOffers",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "GetTime", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "RemoveOffer",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "SECONDS_IN_WEEK",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "SetCoinAddress",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "SetCommunityContract",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "SetCommunityFee",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "SetMinerUrl",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "Terminate", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "UpdateOffer",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "WithdrawCoin",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "coin", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "communityContract",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "communityFee",
    data: BytesLike
  ): Result;

  events: {
    "Complaint(address,address,uint8)": EventFragment;
    "Payment(address,address,uint256)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "Complaint"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Payment"): EventFragment;
}

export interface ComplaintEventObject {
  user: string;
  miner: string;
  reason: number;
}
export type ComplaintEvent = TypedEvent<
  [string, string, number],
  ComplaintEventObject
>;

export type ComplaintEventFilter = TypedEventFilter<ComplaintEvent>;

export interface PaymentEventObject {
  user: string;
  miner: string;
  amount: BigNumber;
}
export type PaymentEvent = TypedEvent<
  [string, string, BigNumber],
  PaymentEventObject
>;

export type PaymentEventFilter = TypedEventFilter<PaymentEvent>;

export interface BrokerV1 extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: BrokerV1Interface;

  queryFilter<TEvent extends TypedEvent>(
    event: TypedEventFilter<TEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TEvent>>;

  listeners<TEvent extends TypedEvent>(
    eventFilter?: TypedEventFilter<TEvent>
  ): Array<TypedListener<TEvent>>;
  listeners(eventName?: string): Array<Listener>;
  removeAllListeners<TEvent extends TypedEvent>(
    eventFilter: TypedEventFilter<TEvent>
  ): this;
  removeAllListeners(eventName?: string): this;
  off: OnEvent<this>;
  on: OnEvent<this>;
  once: OnEvent<this>;
  removeListener: OnEvent<this>;

  functions: {
    AddOffer(
      pricePerSecond: PromiseOrValue<BigNumberish>,
      machinesAvailable: PromiseOrValue<BigNumberish>,
      specsIpfsHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    Book(
      offerIndex: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    ClaimPayment(
      bookingId: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    DepositCoin(
      numTokens: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    FindBookingsByMiner(
      _miner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<
      [BrokerV1.BookingStructOutput[]] & {
        filteredBookings: BrokerV1.BookingStructOutput[];
      }
    >;

    FindBookingsByUser(
      _owner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<
      [BrokerV1.BookingStructOutput[]] & {
        filteredBookings: BrokerV1.BookingStructOutput[];
      }
    >;

    GetAvailableOffers(
      overrides?: CallOverrides
    ): Promise<
      [BrokerV1.OfferStructOutput[]] & {
        filteredOffers: BrokerV1.OfferStructOutput[];
      }
    >;

    GetBooking(
      index: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<
      [BrokerV1.BookingStructOutput] & { booking: BrokerV1.BookingStructOutput }
    >;

    GetCoinBalance(
      user: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[BigNumber, BigNumber]>;

    GetMinerUrl(
      _user: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    GetMinersOffers(
      miner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<
      [BrokerV1.OfferStructOutput[]] & {
        filteredOffers: BrokerV1.OfferStructOutput[];
      }
    >;

    GetTime(overrides?: CallOverrides): Promise<[BigNumber]>;

    RemoveOffer(
      offerIndex: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    SECONDS_IN_WEEK(overrides?: CallOverrides): Promise<[BigNumber]>;

    SetCoinAddress(
      newCoinAddress: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    SetCommunityContract(
      newCommunityAddress: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    SetCommunityFee(
      fee: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    SetMinerUrl(
      url: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    Terminate(
      bookingId: PromiseOrValue<BigNumberish>,
      reason: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    UpdateOffer(
      offerIndex: PromiseOrValue<BigNumberish>,
      machinesAvailable: PromiseOrValue<BigNumberish>,
      pps: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    WithdrawCoin(
      amt: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    coin(overrides?: CallOverrides): Promise<[string]>;

    communityContract(overrides?: CallOverrides): Promise<[string]>;

    communityFee(overrides?: CallOverrides): Promise<[BigNumber]>;
  };

  AddOffer(
    pricePerSecond: PromiseOrValue<BigNumberish>,
    machinesAvailable: PromiseOrValue<BigNumberish>,
    specsIpfsHash: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  Book(
    offerIndex: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  ClaimPayment(
    bookingId: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  DepositCoin(
    numTokens: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  FindBookingsByMiner(
    _miner: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<BrokerV1.BookingStructOutput[]>;

  FindBookingsByUser(
    _owner: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<BrokerV1.BookingStructOutput[]>;

  GetAvailableOffers(
    overrides?: CallOverrides
  ): Promise<BrokerV1.OfferStructOutput[]>;

  GetBooking(
    index: PromiseOrValue<BigNumberish>,
    overrides?: CallOverrides
  ): Promise<BrokerV1.BookingStructOutput>;

  GetCoinBalance(
    user: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<[BigNumber, BigNumber]>;

  GetMinerUrl(
    _user: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<string>;

  GetMinersOffers(
    miner: PromiseOrValue<string>,
    overrides?: CallOverrides
  ): Promise<BrokerV1.OfferStructOutput[]>;

  GetTime(overrides?: CallOverrides): Promise<BigNumber>;

  RemoveOffer(
    offerIndex: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  SECONDS_IN_WEEK(overrides?: CallOverrides): Promise<BigNumber>;

  SetCoinAddress(
    newCoinAddress: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  SetCommunityContract(
    newCommunityAddress: PromiseOrValue<string>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  SetCommunityFee(
    fee: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  SetMinerUrl(
    url: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  Terminate(
    bookingId: PromiseOrValue<BigNumberish>,
    reason: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  UpdateOffer(
    offerIndex: PromiseOrValue<BigNumberish>,
    machinesAvailable: PromiseOrValue<BigNumberish>,
    pps: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  WithdrawCoin(
    amt: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  coin(overrides?: CallOverrides): Promise<string>;

  communityContract(overrides?: CallOverrides): Promise<string>;

  communityFee(overrides?: CallOverrides): Promise<BigNumber>;

  callStatic: {
    AddOffer(
      pricePerSecond: PromiseOrValue<BigNumberish>,
      machinesAvailable: PromiseOrValue<BigNumberish>,
      specsIpfsHash: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    Book(
      offerIndex: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    ClaimPayment(
      bookingId: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    DepositCoin(
      numTokens: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    FindBookingsByMiner(
      _miner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BrokerV1.BookingStructOutput[]>;

    FindBookingsByUser(
      _owner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BrokerV1.BookingStructOutput[]>;

    GetAvailableOffers(
      overrides?: CallOverrides
    ): Promise<BrokerV1.OfferStructOutput[]>;

    GetBooking(
      index: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BrokerV1.BookingStructOutput>;

    GetCoinBalance(
      user: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<[BigNumber, BigNumber]>;

    GetMinerUrl(
      _user: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<string>;

    GetMinersOffers(
      miner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BrokerV1.OfferStructOutput[]>;

    GetTime(overrides?: CallOverrides): Promise<BigNumber>;

    RemoveOffer(
      offerIndex: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    SECONDS_IN_WEEK(overrides?: CallOverrides): Promise<BigNumber>;

    SetCoinAddress(
      newCoinAddress: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    SetCommunityContract(
      newCommunityAddress: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<void>;

    SetCommunityFee(
      fee: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    SetMinerUrl(
      url: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<void>;

    Terminate(
      bookingId: PromiseOrValue<BigNumberish>,
      reason: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    UpdateOffer(
      offerIndex: PromiseOrValue<BigNumberish>,
      machinesAvailable: PromiseOrValue<BigNumberish>,
      pps: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    WithdrawCoin(
      amt: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<boolean>;

    coin(overrides?: CallOverrides): Promise<string>;

    communityContract(overrides?: CallOverrides): Promise<string>;

    communityFee(overrides?: CallOverrides): Promise<BigNumber>;
  };

  filters: {
    "Complaint(address,address,uint8)"(
      user?: PromiseOrValue<string> | null,
      miner?: PromiseOrValue<string> | null,
      reason?: PromiseOrValue<BigNumberish> | null
    ): ComplaintEventFilter;
    Complaint(
      user?: PromiseOrValue<string> | null,
      miner?: PromiseOrValue<string> | null,
      reason?: PromiseOrValue<BigNumberish> | null
    ): ComplaintEventFilter;

    "Payment(address,address,uint256)"(
      user?: PromiseOrValue<string> | null,
      miner?: PromiseOrValue<string> | null,
      amount?: null
    ): PaymentEventFilter;
    Payment(
      user?: PromiseOrValue<string> | null,
      miner?: PromiseOrValue<string> | null,
      amount?: null
    ): PaymentEventFilter;
  };

  estimateGas: {
    AddOffer(
      pricePerSecond: PromiseOrValue<BigNumberish>,
      machinesAvailable: PromiseOrValue<BigNumberish>,
      specsIpfsHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    Book(
      offerIndex: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    ClaimPayment(
      bookingId: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    DepositCoin(
      numTokens: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    FindBookingsByMiner(
      _miner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    FindBookingsByUser(
      _owner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    GetAvailableOffers(overrides?: CallOverrides): Promise<BigNumber>;

    GetBooking(
      index: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    GetCoinBalance(
      user: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    GetMinerUrl(
      _user: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    GetMinersOffers(
      miner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    GetTime(overrides?: CallOverrides): Promise<BigNumber>;

    RemoveOffer(
      offerIndex: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    SECONDS_IN_WEEK(overrides?: CallOverrides): Promise<BigNumber>;

    SetCoinAddress(
      newCoinAddress: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    SetCommunityContract(
      newCommunityAddress: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    SetCommunityFee(
      fee: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    SetMinerUrl(
      url: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    Terminate(
      bookingId: PromiseOrValue<BigNumberish>,
      reason: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    UpdateOffer(
      offerIndex: PromiseOrValue<BigNumberish>,
      machinesAvailable: PromiseOrValue<BigNumberish>,
      pps: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    WithdrawCoin(
      amt: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    coin(overrides?: CallOverrides): Promise<BigNumber>;

    communityContract(overrides?: CallOverrides): Promise<BigNumber>;

    communityFee(overrides?: CallOverrides): Promise<BigNumber>;
  };

  populateTransaction: {
    AddOffer(
      pricePerSecond: PromiseOrValue<BigNumberish>,
      machinesAvailable: PromiseOrValue<BigNumberish>,
      specsIpfsHash: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    Book(
      offerIndex: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    ClaimPayment(
      bookingId: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    DepositCoin(
      numTokens: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    FindBookingsByMiner(
      _miner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    FindBookingsByUser(
      _owner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    GetAvailableOffers(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    GetBooking(
      index: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    GetCoinBalance(
      user: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    GetMinerUrl(
      _user: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    GetMinersOffers(
      miner: PromiseOrValue<string>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    GetTime(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    RemoveOffer(
      offerIndex: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    SECONDS_IN_WEEK(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    SetCoinAddress(
      newCoinAddress: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    SetCommunityContract(
      newCommunityAddress: PromiseOrValue<string>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    SetCommunityFee(
      fee: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    SetMinerUrl(
      url: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    Terminate(
      bookingId: PromiseOrValue<BigNumberish>,
      reason: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    UpdateOffer(
      offerIndex: PromiseOrValue<BigNumberish>,
      machinesAvailable: PromiseOrValue<BigNumberish>,
      pps: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    WithdrawCoin(
      amt: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    coin(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    communityContract(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    communityFee(overrides?: CallOverrides): Promise<PopulatedTransaction>;
  };
}
