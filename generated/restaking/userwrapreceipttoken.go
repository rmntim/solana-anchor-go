// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UserWrapReceiptToken is the `user_wrap_receipt_token` instruction.
type UserWrapReceiptToken struct {
	Amount *uint64

	// [0] = [SIGNER] user
	//
	// [1] = [] fund_wrap_account
	//
	// [2] = [] receipt_token_program
	//
	// [3] = [] wrapped_token_program
	//
	// [4] = [WRITE] receipt_token_mint
	//
	// [5] = [WRITE] wrapped_token_mint
	//
	// [6] = [WRITE] user_receipt_token_account
	//
	// [7] = [WRITE] receipt_token_wrap_account
	//
	// [8] = [WRITE] user_wrapped_token_account
	//
	// [9] = [WRITE] fund_account
	//
	// [10] = [WRITE] user_fund_account
	//
	// [11] = [WRITE] reward_account
	//
	// [12] = [WRITE] user_reward_account
	//
	// [13] = [WRITE] fund_wrap_account_reward_account
	//
	// [14] = [] event_authority
	//
	// [15] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUserWrapReceiptTokenInstructionBuilder creates a new `UserWrapReceiptToken` instruction builder.
func NewUserWrapReceiptTokenInstructionBuilder() *UserWrapReceiptToken {
	nd := &UserWrapReceiptToken{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 16),
	}
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb"])
	nd.AccountMetaSlice[3] = ag_solanago.Meta(Addresses["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"])
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *UserWrapReceiptToken) SetAmount(amount uint64) *UserWrapReceiptToken {
	inst.Amount = &amount
	return inst
}

// SetUserAccount sets the "user" account.
func (inst *UserWrapReceiptToken) SetUserAccount(user ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(user).SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *UserWrapReceiptToken) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetFundWrapAccountAccount sets the "fund_wrap_account" account.
func (inst *UserWrapReceiptToken) SetFundWrapAccountAccount(fundWrapAccount ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(fundWrapAccount)
	return inst
}

func (inst *UserWrapReceiptToken) findFindFundWrapAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: fund_wrap
	seeds = append(seeds, []byte{byte(0x66), byte(0x75), byte(0x6e), byte(0x64), byte(0x5f), byte(0x77), byte(0x72), byte(0x61), byte(0x70)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindFundWrapAccountAddressWithBumpSeed calculates FundWrapAccount account address with given seeds and a known bump seed.
func (inst *UserWrapReceiptToken) FindFundWrapAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundWrapAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *UserWrapReceiptToken) MustFindFundWrapAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundWrapAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundWrapAccountAddress finds FundWrapAccount account address with given seeds.
func (inst *UserWrapReceiptToken) FindFundWrapAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundWrapAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *UserWrapReceiptToken) MustFindFundWrapAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundWrapAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundWrapAccountAccount gets the "fund_wrap_account" account.
func (inst *UserWrapReceiptToken) GetFundWrapAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetReceiptTokenProgramAccount sets the "receipt_token_program" account.
func (inst *UserWrapReceiptToken) SetReceiptTokenProgramAccount(receiptTokenProgram ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(receiptTokenProgram)
	return inst
}

// GetReceiptTokenProgramAccount gets the "receipt_token_program" account.
func (inst *UserWrapReceiptToken) GetReceiptTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetWrappedTokenProgramAccount sets the "wrapped_token_program" account.
func (inst *UserWrapReceiptToken) SetWrappedTokenProgramAccount(wrappedTokenProgram ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(wrappedTokenProgram)
	return inst
}

// GetWrappedTokenProgramAccount gets the "wrapped_token_program" account.
func (inst *UserWrapReceiptToken) GetWrappedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *UserWrapReceiptToken) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(receiptTokenMint).WRITE()
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *UserWrapReceiptToken) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetWrappedTokenMintAccount sets the "wrapped_token_mint" account.
func (inst *UserWrapReceiptToken) SetWrappedTokenMintAccount(wrappedTokenMint ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(wrappedTokenMint).WRITE()
	return inst
}

// GetWrappedTokenMintAccount gets the "wrapped_token_mint" account.
func (inst *UserWrapReceiptToken) GetWrappedTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetUserReceiptTokenAccountAccount sets the "user_receipt_token_account" account.
func (inst *UserWrapReceiptToken) SetUserReceiptTokenAccountAccount(userReceiptTokenAccount ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(userReceiptTokenAccount).WRITE()
	return inst
}

func (inst *UserWrapReceiptToken) findFindUserReceiptTokenAccountAddress(user ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// path: user
	seeds = append(seeds, user.Bytes())
	// path: receiptTokenProgram
	seeds = append(seeds, receiptTokenProgram.Bytes())
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	programID := Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"]

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, programID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, programID)
	}
	return
}

// FindUserReceiptTokenAccountAddressWithBumpSeed calculates UserReceiptTokenAccount account address with given seeds and a known bump seed.
func (inst *UserWrapReceiptToken) FindUserReceiptTokenAccountAddressWithBumpSeed(user ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserReceiptTokenAccountAddress(user, receiptTokenProgram, receiptTokenMint, bumpSeed)
	return
}

func (inst *UserWrapReceiptToken) MustFindUserReceiptTokenAccountAddressWithBumpSeed(user ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserReceiptTokenAccountAddress(user, receiptTokenProgram, receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserReceiptTokenAccountAddress finds UserReceiptTokenAccount account address with given seeds.
func (inst *UserWrapReceiptToken) FindUserReceiptTokenAccountAddress(user ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserReceiptTokenAccountAddress(user, receiptTokenProgram, receiptTokenMint, 0)
	return
}

func (inst *UserWrapReceiptToken) MustFindUserReceiptTokenAccountAddress(user ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserReceiptTokenAccountAddress(user, receiptTokenProgram, receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserReceiptTokenAccountAccount gets the "user_receipt_token_account" account.
func (inst *UserWrapReceiptToken) GetUserReceiptTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetReceiptTokenWrapAccountAccount sets the "receipt_token_wrap_account" account.
func (inst *UserWrapReceiptToken) SetReceiptTokenWrapAccountAccount(receiptTokenWrapAccount ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(receiptTokenWrapAccount).WRITE()
	return inst
}

func (inst *UserWrapReceiptToken) findFindReceiptTokenWrapAccountAddress(fundWrapAccount ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// path: fundWrapAccount
	seeds = append(seeds, fundWrapAccount.Bytes())
	// path: receiptTokenProgram
	seeds = append(seeds, receiptTokenProgram.Bytes())
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	programID := Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"]

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, programID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, programID)
	}
	return
}

// FindReceiptTokenWrapAccountAddressWithBumpSeed calculates ReceiptTokenWrapAccount account address with given seeds and a known bump seed.
func (inst *UserWrapReceiptToken) FindReceiptTokenWrapAccountAddressWithBumpSeed(fundWrapAccount ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindReceiptTokenWrapAccountAddress(fundWrapAccount, receiptTokenProgram, receiptTokenMint, bumpSeed)
	return
}

func (inst *UserWrapReceiptToken) MustFindReceiptTokenWrapAccountAddressWithBumpSeed(fundWrapAccount ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenWrapAccountAddress(fundWrapAccount, receiptTokenProgram, receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindReceiptTokenWrapAccountAddress finds ReceiptTokenWrapAccount account address with given seeds.
func (inst *UserWrapReceiptToken) FindReceiptTokenWrapAccountAddress(fundWrapAccount ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindReceiptTokenWrapAccountAddress(fundWrapAccount, receiptTokenProgram, receiptTokenMint, 0)
	return
}

func (inst *UserWrapReceiptToken) MustFindReceiptTokenWrapAccountAddress(fundWrapAccount ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindReceiptTokenWrapAccountAddress(fundWrapAccount, receiptTokenProgram, receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetReceiptTokenWrapAccountAccount gets the "receipt_token_wrap_account" account.
func (inst *UserWrapReceiptToken) GetReceiptTokenWrapAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetUserWrappedTokenAccountAccount sets the "user_wrapped_token_account" account.
func (inst *UserWrapReceiptToken) SetUserWrappedTokenAccountAccount(userWrappedTokenAccount ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(userWrappedTokenAccount).WRITE()
	return inst
}

func (inst *UserWrapReceiptToken) findFindUserWrappedTokenAccountAddress(user ag_solanago.PublicKey, wrappedTokenProgram ag_solanago.PublicKey, wrappedTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// path: user
	seeds = append(seeds, user.Bytes())
	// path: wrappedTokenProgram
	seeds = append(seeds, wrappedTokenProgram.Bytes())
	// path: wrappedTokenMint
	seeds = append(seeds, wrappedTokenMint.Bytes())

	programID := Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"]

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, programID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, programID)
	}
	return
}

// FindUserWrappedTokenAccountAddressWithBumpSeed calculates UserWrappedTokenAccount account address with given seeds and a known bump seed.
func (inst *UserWrapReceiptToken) FindUserWrappedTokenAccountAddressWithBumpSeed(user ag_solanago.PublicKey, wrappedTokenProgram ag_solanago.PublicKey, wrappedTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserWrappedTokenAccountAddress(user, wrappedTokenProgram, wrappedTokenMint, bumpSeed)
	return
}

func (inst *UserWrapReceiptToken) MustFindUserWrappedTokenAccountAddressWithBumpSeed(user ag_solanago.PublicKey, wrappedTokenProgram ag_solanago.PublicKey, wrappedTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserWrappedTokenAccountAddress(user, wrappedTokenProgram, wrappedTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserWrappedTokenAccountAddress finds UserWrappedTokenAccount account address with given seeds.
func (inst *UserWrapReceiptToken) FindUserWrappedTokenAccountAddress(user ag_solanago.PublicKey, wrappedTokenProgram ag_solanago.PublicKey, wrappedTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserWrappedTokenAccountAddress(user, wrappedTokenProgram, wrappedTokenMint, 0)
	return
}

func (inst *UserWrapReceiptToken) MustFindUserWrappedTokenAccountAddress(user ag_solanago.PublicKey, wrappedTokenProgram ag_solanago.PublicKey, wrappedTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserWrappedTokenAccountAddress(user, wrappedTokenProgram, wrappedTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserWrappedTokenAccountAccount gets the "user_wrapped_token_account" account.
func (inst *UserWrapReceiptToken) GetUserWrappedTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetFundAccountAccount sets the "fund_account" account.
func (inst *UserWrapReceiptToken) SetFundAccountAccount(fundAccount ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(fundAccount).WRITE()
	return inst
}

func (inst *UserWrapReceiptToken) findFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: fund
	seeds = append(seeds, []byte{byte(0x66), byte(0x75), byte(0x6e), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindFundAccountAddressWithBumpSeed calculates FundAccount account address with given seeds and a known bump seed.
func (inst *UserWrapReceiptToken) FindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *UserWrapReceiptToken) MustFindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAccountAddress finds FundAccount account address with given seeds.
func (inst *UserWrapReceiptToken) FindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *UserWrapReceiptToken) MustFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccountAccount gets the "fund_account" account.
func (inst *UserWrapReceiptToken) GetFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetUserFundAccountAccount sets the "user_fund_account" account.
func (inst *UserWrapReceiptToken) SetUserFundAccountAccount(userFundAccount ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(userFundAccount).WRITE()
	return inst
}

func (inst *UserWrapReceiptToken) findFindUserFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: user_fund
	seeds = append(seeds, []byte{byte(0x75), byte(0x73), byte(0x65), byte(0x72), byte(0x5f), byte(0x66), byte(0x75), byte(0x6e), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())
	// path: user
	seeds = append(seeds, user.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindUserFundAccountAddressWithBumpSeed calculates UserFundAccount account address with given seeds and a known bump seed.
func (inst *UserWrapReceiptToken) FindUserFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserFundAccountAddress(receiptTokenMint, user, bumpSeed)
	return
}

func (inst *UserWrapReceiptToken) MustFindUserFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserFundAccountAddress(receiptTokenMint, user, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserFundAccountAddress finds UserFundAccount account address with given seeds.
func (inst *UserWrapReceiptToken) FindUserFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserFundAccountAddress(receiptTokenMint, user, 0)
	return
}

func (inst *UserWrapReceiptToken) MustFindUserFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserFundAccountAddress(receiptTokenMint, user, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserFundAccountAccount gets the "user_fund_account" account.
func (inst *UserWrapReceiptToken) GetUserFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetRewardAccountAccount sets the "reward_account" account.
func (inst *UserWrapReceiptToken) SetRewardAccountAccount(rewardAccount ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(rewardAccount).WRITE()
	return inst
}

func (inst *UserWrapReceiptToken) findFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: reward
	seeds = append(seeds, []byte{byte(0x72), byte(0x65), byte(0x77), byte(0x61), byte(0x72), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindRewardAccountAddressWithBumpSeed calculates RewardAccount account address with given seeds and a known bump seed.
func (inst *UserWrapReceiptToken) FindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *UserWrapReceiptToken) MustFindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindRewardAccountAddress finds RewardAccount account address with given seeds.
func (inst *UserWrapReceiptToken) FindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *UserWrapReceiptToken) MustFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetRewardAccountAccount gets the "reward_account" account.
func (inst *UserWrapReceiptToken) GetRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetUserRewardAccountAccount sets the "user_reward_account" account.
func (inst *UserWrapReceiptToken) SetUserRewardAccountAccount(userRewardAccount ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(userRewardAccount).WRITE()
	return inst
}

func (inst *UserWrapReceiptToken) findFindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: user_reward
	seeds = append(seeds, []byte{byte(0x75), byte(0x73), byte(0x65), byte(0x72), byte(0x5f), byte(0x72), byte(0x65), byte(0x77), byte(0x61), byte(0x72), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())
	// path: user
	seeds = append(seeds, user.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindUserRewardAccountAddressWithBumpSeed calculates UserRewardAccount account address with given seeds and a known bump seed.
func (inst *UserWrapReceiptToken) FindUserRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserRewardAccountAddress(receiptTokenMint, user, bumpSeed)
	return
}

func (inst *UserWrapReceiptToken) MustFindUserRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserRewardAccountAddress(receiptTokenMint, user, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserRewardAccountAddress finds UserRewardAccount account address with given seeds.
func (inst *UserWrapReceiptToken) FindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserRewardAccountAddress(receiptTokenMint, user, 0)
	return
}

func (inst *UserWrapReceiptToken) MustFindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserRewardAccountAddress(receiptTokenMint, user, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserRewardAccountAccount gets the "user_reward_account" account.
func (inst *UserWrapReceiptToken) GetUserRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetFundWrapAccountRewardAccountAccount sets the "fund_wrap_account_reward_account" account.
func (inst *UserWrapReceiptToken) SetFundWrapAccountRewardAccountAccount(fundWrapAccountRewardAccount ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(fundWrapAccountRewardAccount).WRITE()
	return inst
}

func (inst *UserWrapReceiptToken) findFindFundWrapAccountRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, fundWrapAccount ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: user_reward
	seeds = append(seeds, []byte{byte(0x75), byte(0x73), byte(0x65), byte(0x72), byte(0x5f), byte(0x72), byte(0x65), byte(0x77), byte(0x61), byte(0x72), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())
	// path: fundWrapAccount
	seeds = append(seeds, fundWrapAccount.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindFundWrapAccountRewardAccountAddressWithBumpSeed calculates FundWrapAccountRewardAccount account address with given seeds and a known bump seed.
func (inst *UserWrapReceiptToken) FindFundWrapAccountRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, fundWrapAccount ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundWrapAccountRewardAccountAddress(receiptTokenMint, fundWrapAccount, bumpSeed)
	return
}

func (inst *UserWrapReceiptToken) MustFindFundWrapAccountRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, fundWrapAccount ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundWrapAccountRewardAccountAddress(receiptTokenMint, fundWrapAccount, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundWrapAccountRewardAccountAddress finds FundWrapAccountRewardAccount account address with given seeds.
func (inst *UserWrapReceiptToken) FindFundWrapAccountRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, fundWrapAccount ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundWrapAccountRewardAccountAddress(receiptTokenMint, fundWrapAccount, 0)
	return
}

func (inst *UserWrapReceiptToken) MustFindFundWrapAccountRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, fundWrapAccount ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundWrapAccountRewardAccountAddress(receiptTokenMint, fundWrapAccount, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundWrapAccountRewardAccountAccount gets the "fund_wrap_account_reward_account" account.
func (inst *UserWrapReceiptToken) GetFundWrapAccountRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *UserWrapReceiptToken) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(eventAuthority)
	return inst
}

func (inst *UserWrapReceiptToken) findFindEventAuthorityAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: __event_authority
	seeds = append(seeds, []byte{byte(0x5f), byte(0x5f), byte(0x65), byte(0x76), byte(0x65), byte(0x6e), byte(0x74), byte(0x5f), byte(0x61), byte(0x75), byte(0x74), byte(0x68), byte(0x6f), byte(0x72), byte(0x69), byte(0x74), byte(0x79)})

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindEventAuthorityAddressWithBumpSeed calculates EventAuthority account address with given seeds and a known bump seed.
func (inst *UserWrapReceiptToken) FindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindEventAuthorityAddress(bumpSeed)
	return
}

func (inst *UserWrapReceiptToken) MustFindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindEventAuthorityAddress finds EventAuthority account address with given seeds.
func (inst *UserWrapReceiptToken) FindEventAuthorityAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindEventAuthorityAddress(0)
	return
}

func (inst *UserWrapReceiptToken) MustFindEventAuthorityAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *UserWrapReceiptToken) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetProgramAccount sets the "program" account.
func (inst *UserWrapReceiptToken) SetProgramAccount(program ag_solanago.PublicKey) *UserWrapReceiptToken {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *UserWrapReceiptToken) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

func (inst UserWrapReceiptToken) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UserWrapReceiptToken,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UserWrapReceiptToken) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UserWrapReceiptToken) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.FundWrapAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ReceiptTokenProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.WrappedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.WrappedTokenMint is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.UserReceiptTokenAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.ReceiptTokenWrapAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.UserWrappedTokenAccount is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.FundAccount is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.UserFundAccount is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.RewardAccount is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.UserRewardAccount is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.FundWrapAccountRewardAccount is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *UserWrapReceiptToken) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UserWrapReceiptToken")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=16]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                     user", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("               fund_wrap_", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    receipt_token_program", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    wrapped_token_program", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       receipt_token_mint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("       wrapped_token_mint", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("      user_receipt_token_", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("      receipt_token_wrap_", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("      user_wrapped_token_", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                    fund_", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("               user_fund_", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                  reward_", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("             user_reward_", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("fund_wrap_account_reward_", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("          event_authority", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("                  program", inst.AccountMetaSlice.Get(15)))
					})
				})
		})
}

func (obj UserWrapReceiptToken) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UserWrapReceiptToken) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewUserWrapReceiptTokenInstruction declares a new UserWrapReceiptToken instruction with the provided parameters and accounts.
func NewUserWrapReceiptTokenInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	user ag_solanago.PublicKey,
	fundWrapAccount ag_solanago.PublicKey,
	receiptTokenProgram ag_solanago.PublicKey,
	wrappedTokenProgram ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	wrappedTokenMint ag_solanago.PublicKey,
	userReceiptTokenAccount ag_solanago.PublicKey,
	receiptTokenWrapAccount ag_solanago.PublicKey,
	userWrappedTokenAccount ag_solanago.PublicKey,
	fundAccount ag_solanago.PublicKey,
	userFundAccount ag_solanago.PublicKey,
	rewardAccount ag_solanago.PublicKey,
	userRewardAccount ag_solanago.PublicKey,
	fundWrapAccountRewardAccount ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *UserWrapReceiptToken {
	return NewUserWrapReceiptTokenInstructionBuilder().
		SetAmount(amount).
		SetUserAccount(user).
		SetFundWrapAccountAccount(fundWrapAccount).
		SetReceiptTokenProgramAccount(receiptTokenProgram).
		SetWrappedTokenProgramAccount(wrappedTokenProgram).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetWrappedTokenMintAccount(wrappedTokenMint).
		SetUserReceiptTokenAccountAccount(userReceiptTokenAccount).
		SetReceiptTokenWrapAccountAccount(receiptTokenWrapAccount).
		SetUserWrappedTokenAccountAccount(userWrappedTokenAccount).
		SetFundAccountAccount(fundAccount).
		SetUserFundAccountAccount(userFundAccount).
		SetRewardAccountAccount(rewardAccount).
		SetUserRewardAccountAccount(userRewardAccount).
		SetFundWrapAccountRewardAccountAccount(fundWrapAccountRewardAccount).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
