// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type FundAccountAccount struct {
	DataVersion               uint16
	Bump                      uint8
	ReserveAccountBump        uint8
	TreasuryAccountBump       uint8
	WrapAccountBump           uint8
	Padding                   [8]uint8
	TransferEnabled           uint8
	AddressLookupTableEnabled uint8
	AddressLookupTableAccount ag_solanago.PublicKey
	ReserveAccount            ag_solanago.PublicKey
	TreasuryAccount           ag_solanago.PublicKey

	// receipt token information
	ReceiptTokenMint             ag_solanago.PublicKey
	ReceiptTokenProgram          ag_solanago.PublicKey
	ReceiptTokenDecimals         uint8
	Padding2                     [7]uint8
	ReceiptTokenSupplyAmount     uint64
	OneReceiptTokenAsSol         uint64
	ReceiptTokenValueUpdatedSlot uint64
	ReceiptTokenValue            TokenValuePod

	// global withdrawal configurations
	WithdrawalBatchThresholdIntervalSeconds int64
	WithdrawalFeeRateBps                    uint16
	WithdrawalEnabled                       uint8
	DepositEnabled                          uint8
	DonationEnabled                         uint8
	Padding4                                [3]uint8

	// SOL deposit & withdrawal
	Sol AssetState

	// underlying assets
	Padding6           [15]uint8
	NumSupportedTokens uint8
	SupportedTokens    [30]SupportedToken

	// optional basket of underlying assets
	NormalizedToken NormalizedToken

	// investments
	Padding7           [15]uint8
	NumRestakingVaults uint8
	RestakingVaults    [16]RestakingVault
	Padding8           [112]uint8

	// fund operation state
	Operation OperationState

	// optional wrapped token of fund receipt token
	WrapAccount  ag_solanago.PublicKey
	WrappedToken WrappedToken

	// which DEX to use for swap between two tokens
	NumTokenSwapStrategies uint8
	Padding9               [7]uint8
	TokenSwapStrategies    [30]TokenSwapStrategy
	Reserved               [3616]uint8
}

var FundAccountAccountDiscriminator = [8]byte{49, 104, 168, 214, 134, 180, 173, 154}

func (obj FundAccountAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(FundAccountAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `DataVersion` param:
	err = encoder.Encode(obj.DataVersion)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `ReserveAccountBump` param:
	err = encoder.Encode(obj.ReserveAccountBump)
	if err != nil {
		return err
	}
	// Serialize `TreasuryAccountBump` param:
	err = encoder.Encode(obj.TreasuryAccountBump)
	if err != nil {
		return err
	}
	// Serialize `WrapAccountBump` param:
	err = encoder.Encode(obj.WrapAccountBump)
	if err != nil {
		return err
	}
	// Serialize `Padding` param:
	err = encoder.Encode(obj.Padding)
	if err != nil {
		return err
	}
	// Serialize `TransferEnabled` param:
	err = encoder.Encode(obj.TransferEnabled)
	if err != nil {
		return err
	}
	// Serialize `AddressLookupTableEnabled` param:
	err = encoder.Encode(obj.AddressLookupTableEnabled)
	if err != nil {
		return err
	}
	// Serialize `AddressLookupTableAccount` param:
	err = encoder.Encode(obj.AddressLookupTableAccount)
	if err != nil {
		return err
	}
	// Serialize `ReserveAccount` param:
	err = encoder.Encode(obj.ReserveAccount)
	if err != nil {
		return err
	}
	// Serialize `TreasuryAccount` param:
	err = encoder.Encode(obj.TreasuryAccount)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenMint` param:
	err = encoder.Encode(obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenProgram` param:
	err = encoder.Encode(obj.ReceiptTokenProgram)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenDecimals` param:
	err = encoder.Encode(obj.ReceiptTokenDecimals)
	if err != nil {
		return err
	}
	// Serialize `Padding2` param:
	err = encoder.Encode(obj.Padding2)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenSupplyAmount` param:
	err = encoder.Encode(obj.ReceiptTokenSupplyAmount)
	if err != nil {
		return err
	}
	// Serialize `OneReceiptTokenAsSol` param:
	err = encoder.Encode(obj.OneReceiptTokenAsSol)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenValueUpdatedSlot` param:
	err = encoder.Encode(obj.ReceiptTokenValueUpdatedSlot)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenValue` param:
	err = encoder.Encode(obj.ReceiptTokenValue)
	if err != nil {
		return err
	}
	// Serialize `WithdrawalBatchThresholdIntervalSeconds` param:
	err = encoder.Encode(obj.WithdrawalBatchThresholdIntervalSeconds)
	if err != nil {
		return err
	}
	// Serialize `WithdrawalFeeRateBps` param:
	err = encoder.Encode(obj.WithdrawalFeeRateBps)
	if err != nil {
		return err
	}
	// Serialize `WithdrawalEnabled` param:
	err = encoder.Encode(obj.WithdrawalEnabled)
	if err != nil {
		return err
	}
	// Serialize `DepositEnabled` param:
	err = encoder.Encode(obj.DepositEnabled)
	if err != nil {
		return err
	}
	// Serialize `DonationEnabled` param:
	err = encoder.Encode(obj.DonationEnabled)
	if err != nil {
		return err
	}
	// Serialize `Padding4` param:
	err = encoder.Encode(obj.Padding4)
	if err != nil {
		return err
	}
	// Serialize `Sol` param:
	err = encoder.Encode(obj.Sol)
	if err != nil {
		return err
	}
	// Serialize `Padding6` param:
	err = encoder.Encode(obj.Padding6)
	if err != nil {
		return err
	}
	// Serialize `NumSupportedTokens` param:
	err = encoder.Encode(obj.NumSupportedTokens)
	if err != nil {
		return err
	}
	// Serialize `SupportedTokens` param:
	err = encoder.Encode(obj.SupportedTokens)
	if err != nil {
		return err
	}
	// Serialize `NormalizedToken` param:
	err = encoder.Encode(obj.NormalizedToken)
	if err != nil {
		return err
	}
	// Serialize `Padding7` param:
	err = encoder.Encode(obj.Padding7)
	if err != nil {
		return err
	}
	// Serialize `NumRestakingVaults` param:
	err = encoder.Encode(obj.NumRestakingVaults)
	if err != nil {
		return err
	}
	// Serialize `RestakingVaults` param:
	err = encoder.Encode(obj.RestakingVaults)
	if err != nil {
		return err
	}
	// Serialize `Padding8` param:
	err = encoder.Encode(obj.Padding8)
	if err != nil {
		return err
	}
	// Serialize `Operation` param:
	err = encoder.Encode(obj.Operation)
	if err != nil {
		return err
	}
	// Serialize `WrapAccount` param:
	err = encoder.Encode(obj.WrapAccount)
	if err != nil {
		return err
	}
	// Serialize `WrappedToken` param:
	err = encoder.Encode(obj.WrappedToken)
	if err != nil {
		return err
	}
	// Serialize `NumTokenSwapStrategies` param:
	err = encoder.Encode(obj.NumTokenSwapStrategies)
	if err != nil {
		return err
	}
	// Serialize `Padding9` param:
	err = encoder.Encode(obj.Padding9)
	if err != nil {
		return err
	}
	// Serialize `TokenSwapStrategies` param:
	err = encoder.Encode(obj.TokenSwapStrategies)
	if err != nil {
		return err
	}
	// Serialize `Reserved` param:
	err = encoder.Encode(obj.Reserved)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FundAccountAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(FundAccountAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[49 104 168 214 134 180 173 154]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `DataVersion`:
	err = decoder.Decode(&obj.DataVersion)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `ReserveAccountBump`:
	err = decoder.Decode(&obj.ReserveAccountBump)
	if err != nil {
		return err
	}
	// Deserialize `TreasuryAccountBump`:
	err = decoder.Decode(&obj.TreasuryAccountBump)
	if err != nil {
		return err
	}
	// Deserialize `WrapAccountBump`:
	err = decoder.Decode(&obj.WrapAccountBump)
	if err != nil {
		return err
	}
	// Deserialize `Padding`:
	err = decoder.Decode(&obj.Padding)
	if err != nil {
		return err
	}
	// Deserialize `TransferEnabled`:
	err = decoder.Decode(&obj.TransferEnabled)
	if err != nil {
		return err
	}
	// Deserialize `AddressLookupTableEnabled`:
	err = decoder.Decode(&obj.AddressLookupTableEnabled)
	if err != nil {
		return err
	}
	// Deserialize `AddressLookupTableAccount`:
	err = decoder.Decode(&obj.AddressLookupTableAccount)
	if err != nil {
		return err
	}
	// Deserialize `ReserveAccount`:
	err = decoder.Decode(&obj.ReserveAccount)
	if err != nil {
		return err
	}
	// Deserialize `TreasuryAccount`:
	err = decoder.Decode(&obj.TreasuryAccount)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenMint`:
	err = decoder.Decode(&obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenProgram`:
	err = decoder.Decode(&obj.ReceiptTokenProgram)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenDecimals`:
	err = decoder.Decode(&obj.ReceiptTokenDecimals)
	if err != nil {
		return err
	}
	// Deserialize `Padding2`:
	err = decoder.Decode(&obj.Padding2)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenSupplyAmount`:
	err = decoder.Decode(&obj.ReceiptTokenSupplyAmount)
	if err != nil {
		return err
	}
	// Deserialize `OneReceiptTokenAsSol`:
	err = decoder.Decode(&obj.OneReceiptTokenAsSol)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenValueUpdatedSlot`:
	err = decoder.Decode(&obj.ReceiptTokenValueUpdatedSlot)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenValue`:
	err = decoder.Decode(&obj.ReceiptTokenValue)
	if err != nil {
		return err
	}
	// Deserialize `WithdrawalBatchThresholdIntervalSeconds`:
	err = decoder.Decode(&obj.WithdrawalBatchThresholdIntervalSeconds)
	if err != nil {
		return err
	}
	// Deserialize `WithdrawalFeeRateBps`:
	err = decoder.Decode(&obj.WithdrawalFeeRateBps)
	if err != nil {
		return err
	}
	// Deserialize `WithdrawalEnabled`:
	err = decoder.Decode(&obj.WithdrawalEnabled)
	if err != nil {
		return err
	}
	// Deserialize `DepositEnabled`:
	err = decoder.Decode(&obj.DepositEnabled)
	if err != nil {
		return err
	}
	// Deserialize `DonationEnabled`:
	err = decoder.Decode(&obj.DonationEnabled)
	if err != nil {
		return err
	}
	// Deserialize `Padding4`:
	err = decoder.Decode(&obj.Padding4)
	if err != nil {
		return err
	}
	// Deserialize `Sol`:
	err = decoder.Decode(&obj.Sol)
	if err != nil {
		return err
	}
	// Deserialize `Padding6`:
	err = decoder.Decode(&obj.Padding6)
	if err != nil {
		return err
	}
	// Deserialize `NumSupportedTokens`:
	err = decoder.Decode(&obj.NumSupportedTokens)
	if err != nil {
		return err
	}
	// Deserialize `SupportedTokens`:
	err = decoder.Decode(&obj.SupportedTokens)
	if err != nil {
		return err
	}
	// Deserialize `NormalizedToken`:
	err = decoder.Decode(&obj.NormalizedToken)
	if err != nil {
		return err
	}
	// Deserialize `Padding7`:
	err = decoder.Decode(&obj.Padding7)
	if err != nil {
		return err
	}
	// Deserialize `NumRestakingVaults`:
	err = decoder.Decode(&obj.NumRestakingVaults)
	if err != nil {
		return err
	}
	// Deserialize `RestakingVaults`:
	err = decoder.Decode(&obj.RestakingVaults)
	if err != nil {
		return err
	}
	// Deserialize `Padding8`:
	err = decoder.Decode(&obj.Padding8)
	if err != nil {
		return err
	}
	// Deserialize `Operation`:
	err = decoder.Decode(&obj.Operation)
	if err != nil {
		return err
	}
	// Deserialize `WrapAccount`:
	err = decoder.Decode(&obj.WrapAccount)
	if err != nil {
		return err
	}
	// Deserialize `WrappedToken`:
	err = decoder.Decode(&obj.WrappedToken)
	if err != nil {
		return err
	}
	// Deserialize `NumTokenSwapStrategies`:
	err = decoder.Decode(&obj.NumTokenSwapStrategies)
	if err != nil {
		return err
	}
	// Deserialize `Padding9`:
	err = decoder.Decode(&obj.Padding9)
	if err != nil {
		return err
	}
	// Deserialize `TokenSwapStrategies`:
	err = decoder.Decode(&obj.TokenSwapStrategies)
	if err != nil {
		return err
	}
	// Deserialize `Reserved`:
	err = decoder.Decode(&obj.Reserved)
	if err != nil {
		return err
	}
	return nil
}

type FundWithdrawalBatchAccountAccount struct {
	DataVersion               uint16
	Bump                      uint8
	ReceiptTokenMint          ag_solanago.PublicKey
	SupportedTokenMint        *ag_solanago.PublicKey `bin:"optional"`
	SupportedTokenProgram     *ag_solanago.PublicKey `bin:"optional"`
	BatchId                   uint64
	NumRequests               uint64
	NumClaimedRequests        uint64
	ReceiptTokenAmount        uint64
	ClaimedReceiptTokenAmount uint64

	// asset to be withdrawn
	AssetUserAmount        uint64
	ClaimedAssetUserAmount uint64

	// informative: withdrawal fee is already paid to the treasury account, just informative.
	AssetFeeAmount uint64
	ProcessedAt    int64
	Reserved       [32]uint8
}

var FundWithdrawalBatchAccountAccountDiscriminator = [8]byte{85, 200, 99, 175, 79, 125, 149, 220}

func (obj FundWithdrawalBatchAccountAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(FundWithdrawalBatchAccountAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `DataVersion` param:
	err = encoder.Encode(obj.DataVersion)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenMint` param:
	err = encoder.Encode(obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	// Serialize `SupportedTokenMint` param (optional):
	{
		if obj.SupportedTokenMint == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.SupportedTokenMint)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `SupportedTokenProgram` param (optional):
	{
		if obj.SupportedTokenProgram == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.SupportedTokenProgram)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `BatchId` param:
	err = encoder.Encode(obj.BatchId)
	if err != nil {
		return err
	}
	// Serialize `NumRequests` param:
	err = encoder.Encode(obj.NumRequests)
	if err != nil {
		return err
	}
	// Serialize `NumClaimedRequests` param:
	err = encoder.Encode(obj.NumClaimedRequests)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenAmount` param:
	err = encoder.Encode(obj.ReceiptTokenAmount)
	if err != nil {
		return err
	}
	// Serialize `ClaimedReceiptTokenAmount` param:
	err = encoder.Encode(obj.ClaimedReceiptTokenAmount)
	if err != nil {
		return err
	}
	// Serialize `AssetUserAmount` param:
	err = encoder.Encode(obj.AssetUserAmount)
	if err != nil {
		return err
	}
	// Serialize `ClaimedAssetUserAmount` param:
	err = encoder.Encode(obj.ClaimedAssetUserAmount)
	if err != nil {
		return err
	}
	// Serialize `AssetFeeAmount` param:
	err = encoder.Encode(obj.AssetFeeAmount)
	if err != nil {
		return err
	}
	// Serialize `ProcessedAt` param:
	err = encoder.Encode(obj.ProcessedAt)
	if err != nil {
		return err
	}
	// Serialize `Reserved` param:
	err = encoder.Encode(obj.Reserved)
	if err != nil {
		return err
	}
	return nil
}

func (obj *FundWithdrawalBatchAccountAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(FundWithdrawalBatchAccountAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[85 200 99 175 79 125 149 220]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `DataVersion`:
	err = decoder.Decode(&obj.DataVersion)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenMint`:
	err = decoder.Decode(&obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	// Deserialize `SupportedTokenMint` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.SupportedTokenMint)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `SupportedTokenProgram` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.SupportedTokenProgram)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `BatchId`:
	err = decoder.Decode(&obj.BatchId)
	if err != nil {
		return err
	}
	// Deserialize `NumRequests`:
	err = decoder.Decode(&obj.NumRequests)
	if err != nil {
		return err
	}
	// Deserialize `NumClaimedRequests`:
	err = decoder.Decode(&obj.NumClaimedRequests)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenAmount`:
	err = decoder.Decode(&obj.ReceiptTokenAmount)
	if err != nil {
		return err
	}
	// Deserialize `ClaimedReceiptTokenAmount`:
	err = decoder.Decode(&obj.ClaimedReceiptTokenAmount)
	if err != nil {
		return err
	}
	// Deserialize `AssetUserAmount`:
	err = decoder.Decode(&obj.AssetUserAmount)
	if err != nil {
		return err
	}
	// Deserialize `ClaimedAssetUserAmount`:
	err = decoder.Decode(&obj.ClaimedAssetUserAmount)
	if err != nil {
		return err
	}
	// Deserialize `AssetFeeAmount`:
	err = decoder.Decode(&obj.AssetFeeAmount)
	if err != nil {
		return err
	}
	// Deserialize `ProcessedAt`:
	err = decoder.Decode(&obj.ProcessedAt)
	if err != nil {
		return err
	}
	// Deserialize `Reserved`:
	err = decoder.Decode(&obj.Reserved)
	if err != nil {
		return err
	}
	return nil
}

type NormalizedTokenPoolAccountAccount struct {
	DataVersion                     uint16
	Bump                            uint8
	NormalizedTokenMint             ag_solanago.PublicKey
	NormalizedTokenProgram          ag_solanago.PublicKey
	SupportedTokens                 []NormalizedSupportedToken
	NormalizedTokenDecimals         uint8
	NormalizedTokenSupplyAmount     uint64
	NormalizedTokenValue            TokenValue
	NormalizedTokenValueUpdatedSlot uint64
	OneNormalizedTokenAsSol         uint64
	Reserved                        [128]uint8
}

var NormalizedTokenPoolAccountAccountDiscriminator = [8]byte{7, 113, 233, 177, 153, 66, 175, 56}

func (obj NormalizedTokenPoolAccountAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(NormalizedTokenPoolAccountAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `DataVersion` param:
	err = encoder.Encode(obj.DataVersion)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `NormalizedTokenMint` param:
	err = encoder.Encode(obj.NormalizedTokenMint)
	if err != nil {
		return err
	}
	// Serialize `NormalizedTokenProgram` param:
	err = encoder.Encode(obj.NormalizedTokenProgram)
	if err != nil {
		return err
	}
	// Serialize `SupportedTokens` param:
	err = encoder.Encode(obj.SupportedTokens)
	if err != nil {
		return err
	}
	// Serialize `NormalizedTokenDecimals` param:
	err = encoder.Encode(obj.NormalizedTokenDecimals)
	if err != nil {
		return err
	}
	// Serialize `NormalizedTokenSupplyAmount` param:
	err = encoder.Encode(obj.NormalizedTokenSupplyAmount)
	if err != nil {
		return err
	}
	// Serialize `NormalizedTokenValue` param:
	err = encoder.Encode(obj.NormalizedTokenValue)
	if err != nil {
		return err
	}
	// Serialize `NormalizedTokenValueUpdatedSlot` param:
	err = encoder.Encode(obj.NormalizedTokenValueUpdatedSlot)
	if err != nil {
		return err
	}
	// Serialize `OneNormalizedTokenAsSol` param:
	err = encoder.Encode(obj.OneNormalizedTokenAsSol)
	if err != nil {
		return err
	}
	// Serialize `Reserved` param:
	err = encoder.Encode(obj.Reserved)
	if err != nil {
		return err
	}
	return nil
}

func (obj *NormalizedTokenPoolAccountAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(NormalizedTokenPoolAccountAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[7 113 233 177 153 66 175 56]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `DataVersion`:
	err = decoder.Decode(&obj.DataVersion)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `NormalizedTokenMint`:
	err = decoder.Decode(&obj.NormalizedTokenMint)
	if err != nil {
		return err
	}
	// Deserialize `NormalizedTokenProgram`:
	err = decoder.Decode(&obj.NormalizedTokenProgram)
	if err != nil {
		return err
	}
	// Deserialize `SupportedTokens`:
	err = decoder.Decode(&obj.SupportedTokens)
	if err != nil {
		return err
	}
	// Deserialize `NormalizedTokenDecimals`:
	err = decoder.Decode(&obj.NormalizedTokenDecimals)
	if err != nil {
		return err
	}
	// Deserialize `NormalizedTokenSupplyAmount`:
	err = decoder.Decode(&obj.NormalizedTokenSupplyAmount)
	if err != nil {
		return err
	}
	// Deserialize `NormalizedTokenValue`:
	err = decoder.Decode(&obj.NormalizedTokenValue)
	if err != nil {
		return err
	}
	// Deserialize `NormalizedTokenValueUpdatedSlot`:
	err = decoder.Decode(&obj.NormalizedTokenValueUpdatedSlot)
	if err != nil {
		return err
	}
	// Deserialize `OneNormalizedTokenAsSol`:
	err = decoder.Decode(&obj.OneNormalizedTokenAsSol)
	if err != nil {
		return err
	}
	// Deserialize `Reserved`:
	err = decoder.Decode(&obj.Reserved)
	if err != nil {
		return err
	}
	return nil
}

type NormalizedTokenWithdrawalAccountAccount struct {
	DataVersion           uint16
	Bump                  uint8
	WithdrawalAuthority   ag_solanago.PublicKey
	NormalizedTokenMint   ag_solanago.PublicKey
	NormalizedTokenPool   ag_solanago.PublicKey
	NormalizedTokenAmount uint64
	ClaimableTokens       []NormalizedClaimableToken
	CreatedAt             int64
	Reserved              [32]uint8
}

var NormalizedTokenWithdrawalAccountAccountDiscriminator = [8]byte{93, 156, 243, 244, 209, 190, 231, 249}

func (obj NormalizedTokenWithdrawalAccountAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(NormalizedTokenWithdrawalAccountAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `DataVersion` param:
	err = encoder.Encode(obj.DataVersion)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `WithdrawalAuthority` param:
	err = encoder.Encode(obj.WithdrawalAuthority)
	if err != nil {
		return err
	}
	// Serialize `NormalizedTokenMint` param:
	err = encoder.Encode(obj.NormalizedTokenMint)
	if err != nil {
		return err
	}
	// Serialize `NormalizedTokenPool` param:
	err = encoder.Encode(obj.NormalizedTokenPool)
	if err != nil {
		return err
	}
	// Serialize `NormalizedTokenAmount` param:
	err = encoder.Encode(obj.NormalizedTokenAmount)
	if err != nil {
		return err
	}
	// Serialize `ClaimableTokens` param:
	err = encoder.Encode(obj.ClaimableTokens)
	if err != nil {
		return err
	}
	// Serialize `CreatedAt` param:
	err = encoder.Encode(obj.CreatedAt)
	if err != nil {
		return err
	}
	// Serialize `Reserved` param:
	err = encoder.Encode(obj.Reserved)
	if err != nil {
		return err
	}
	return nil
}

func (obj *NormalizedTokenWithdrawalAccountAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(NormalizedTokenWithdrawalAccountAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[93 156 243 244 209 190 231 249]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `DataVersion`:
	err = decoder.Decode(&obj.DataVersion)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `WithdrawalAuthority`:
	err = decoder.Decode(&obj.WithdrawalAuthority)
	if err != nil {
		return err
	}
	// Deserialize `NormalizedTokenMint`:
	err = decoder.Decode(&obj.NormalizedTokenMint)
	if err != nil {
		return err
	}
	// Deserialize `NormalizedTokenPool`:
	err = decoder.Decode(&obj.NormalizedTokenPool)
	if err != nil {
		return err
	}
	// Deserialize `NormalizedTokenAmount`:
	err = decoder.Decode(&obj.NormalizedTokenAmount)
	if err != nil {
		return err
	}
	// Deserialize `ClaimableTokens`:
	err = decoder.Decode(&obj.ClaimableTokens)
	if err != nil {
		return err
	}
	// Deserialize `CreatedAt`:
	err = decoder.Decode(&obj.CreatedAt)
	if err != nil {
		return err
	}
	// Deserialize `Reserved`:
	err = decoder.Decode(&obj.Reserved)
	if err != nil {
		return err
	}
	return nil
}

type RewardAccountAccount struct {
	DataVersion      uint16
	Bump             uint8
	ReceiptTokenMint ag_solanago.PublicKey
	MaxHolders       uint8
	MaxRewards       uint16
	MaxRewardPools   uint8
	NumHolders       uint8
	NumRewards       uint16
	NumRewardPools   uint8
	Padding          [5]uint8
	Holders1         [4]RewardPoolHolder
	Rewards1         [16]Reward
	RewardPools1     [4]RewardPool
}

var RewardAccountAccountDiscriminator = [8]byte{225, 81, 31, 253, 84, 234, 171, 129}

func (obj RewardAccountAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(RewardAccountAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `DataVersion` param:
	err = encoder.Encode(obj.DataVersion)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenMint` param:
	err = encoder.Encode(obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	// Serialize `MaxHolders` param:
	err = encoder.Encode(obj.MaxHolders)
	if err != nil {
		return err
	}
	// Serialize `MaxRewards` param:
	err = encoder.Encode(obj.MaxRewards)
	if err != nil {
		return err
	}
	// Serialize `MaxRewardPools` param:
	err = encoder.Encode(obj.MaxRewardPools)
	if err != nil {
		return err
	}
	// Serialize `NumHolders` param:
	err = encoder.Encode(obj.NumHolders)
	if err != nil {
		return err
	}
	// Serialize `NumRewards` param:
	err = encoder.Encode(obj.NumRewards)
	if err != nil {
		return err
	}
	// Serialize `NumRewardPools` param:
	err = encoder.Encode(obj.NumRewardPools)
	if err != nil {
		return err
	}
	// Serialize `Padding` param:
	err = encoder.Encode(obj.Padding)
	if err != nil {
		return err
	}
	// Serialize `Holders1` param:
	err = encoder.Encode(obj.Holders1)
	if err != nil {
		return err
	}
	// Serialize `Rewards1` param:
	err = encoder.Encode(obj.Rewards1)
	if err != nil {
		return err
	}
	// Serialize `RewardPools1` param:
	err = encoder.Encode(obj.RewardPools1)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RewardAccountAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(RewardAccountAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[225 81 31 253 84 234 171 129]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `DataVersion`:
	err = decoder.Decode(&obj.DataVersion)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenMint`:
	err = decoder.Decode(&obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	// Deserialize `MaxHolders`:
	err = decoder.Decode(&obj.MaxHolders)
	if err != nil {
		return err
	}
	// Deserialize `MaxRewards`:
	err = decoder.Decode(&obj.MaxRewards)
	if err != nil {
		return err
	}
	// Deserialize `MaxRewardPools`:
	err = decoder.Decode(&obj.MaxRewardPools)
	if err != nil {
		return err
	}
	// Deserialize `NumHolders`:
	err = decoder.Decode(&obj.NumHolders)
	if err != nil {
		return err
	}
	// Deserialize `NumRewards`:
	err = decoder.Decode(&obj.NumRewards)
	if err != nil {
		return err
	}
	// Deserialize `NumRewardPools`:
	err = decoder.Decode(&obj.NumRewardPools)
	if err != nil {
		return err
	}
	// Deserialize `Padding`:
	err = decoder.Decode(&obj.Padding)
	if err != nil {
		return err
	}
	// Deserialize `Holders1`:
	err = decoder.Decode(&obj.Holders1)
	if err != nil {
		return err
	}
	// Deserialize `Rewards1`:
	err = decoder.Decode(&obj.Rewards1)
	if err != nil {
		return err
	}
	// Deserialize `RewardPools1`:
	err = decoder.Decode(&obj.RewardPools1)
	if err != nil {
		return err
	}
	return nil
}

type UserFundAccountAccount struct {
	DataVersion        uint16
	Bump               uint8
	ReceiptTokenMint   ag_solanago.PublicKey
	User               ag_solanago.PublicKey
	ReceiptTokenAmount uint64
	Reserved           [32]uint8
	WithdrawalRequests []WithdrawalRequest
}

var UserFundAccountAccountDiscriminator = [8]byte{208, 166, 47, 241, 179, 76, 157, 212}

func (obj UserFundAccountAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(UserFundAccountAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `DataVersion` param:
	err = encoder.Encode(obj.DataVersion)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenMint` param:
	err = encoder.Encode(obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenAmount` param:
	err = encoder.Encode(obj.ReceiptTokenAmount)
	if err != nil {
		return err
	}
	// Serialize `Reserved` param:
	err = encoder.Encode(obj.Reserved)
	if err != nil {
		return err
	}
	// Serialize `WithdrawalRequests` param:
	err = encoder.Encode(obj.WithdrawalRequests)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UserFundAccountAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(UserFundAccountAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[208 166 47 241 179 76 157 212]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `DataVersion`:
	err = decoder.Decode(&obj.DataVersion)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenMint`:
	err = decoder.Decode(&obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenAmount`:
	err = decoder.Decode(&obj.ReceiptTokenAmount)
	if err != nil {
		return err
	}
	// Deserialize `Reserved`:
	err = decoder.Decode(&obj.Reserved)
	if err != nil {
		return err
	}
	// Deserialize `WithdrawalRequests`:
	err = decoder.Decode(&obj.WithdrawalRequests)
	if err != nil {
		return err
	}
	return nil
}

type UserRewardAccountAccount struct {
	DataVersion        uint16
	Bump               uint8
	ReceiptTokenMint   ag_solanago.PublicKey
	User               ag_solanago.PublicKey
	NumUserRewardPools uint8
	MaxUserRewardPools uint8
	Padding            [11]uint8
	UserRewardPools1   [4]UserRewardPool
}

var UserRewardAccountAccountDiscriminator = [8]byte{55, 245, 122, 238, 147, 89, 164, 198}

func (obj UserRewardAccountAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(UserRewardAccountAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `DataVersion` param:
	err = encoder.Encode(obj.DataVersion)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `ReceiptTokenMint` param:
	err = encoder.Encode(obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `NumUserRewardPools` param:
	err = encoder.Encode(obj.NumUserRewardPools)
	if err != nil {
		return err
	}
	// Serialize `MaxUserRewardPools` param:
	err = encoder.Encode(obj.MaxUserRewardPools)
	if err != nil {
		return err
	}
	// Serialize `Padding` param:
	err = encoder.Encode(obj.Padding)
	if err != nil {
		return err
	}
	// Serialize `UserRewardPools1` param:
	err = encoder.Encode(obj.UserRewardPools1)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UserRewardAccountAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(UserRewardAccountAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[55 245 122 238 147 89 164 198]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `DataVersion`:
	err = decoder.Decode(&obj.DataVersion)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `ReceiptTokenMint`:
	err = decoder.Decode(&obj.ReceiptTokenMint)
	if err != nil {
		return err
	}
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `NumUserRewardPools`:
	err = decoder.Decode(&obj.NumUserRewardPools)
	if err != nil {
		return err
	}
	// Deserialize `MaxUserRewardPools`:
	err = decoder.Decode(&obj.MaxUserRewardPools)
	if err != nil {
		return err
	}
	// Deserialize `Padding`:
	err = decoder.Decode(&obj.Padding)
	if err != nil {
		return err
	}
	// Deserialize `UserRewardPools1`:
	err = decoder.Decode(&obj.UserRewardPools1)
	if err != nil {
		return err
	}
	return nil
}
