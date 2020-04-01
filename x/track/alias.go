package track

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/bitsongofficial/go-bitsong/x/track/keeper"
	"github.com/bitsongofficial/go-bitsong/x/track/types"
)

const (
	ModuleName                  = types.ModuleName
	StoreKey                    = types.StoreKey
	RouterKey                   = types.RouterKey
	TypeMsgCreateTrack          = types.TypeMsgCreateTrack
	TypeMsgPlayTrack            = types.TypeMsgPlayTrack
	TypeMsgDeposit              = types.TypeMsgDeposit
	DefaultParamspace           = types.DefaultParamspace
	QueryParams                 = types.QueryParams
	QueryTracks                 = types.QueryTracks
	QueryTrack                  = types.QueryTrack
	QueryPlays                  = types.QueryPlays
	QueryShares                 = types.QueryShares
	QueryDeposits               = types.QueryDeposits
	DefaultCodespace            = types.DefaultCodespace
	CodeInvalidGenesis          = types.CodeInvalidGenesis
	CodeInvalidTrackTitle       = types.CodeInvalidTrackTitle
	CodeUnknownTrack            = types.CodeUnknownTrack
	CodeInvalidTrackStatus      = types.CodeInvalidTrackStatus
	CodeInvalidTrackMetadataURI = types.CodeInvalidTrackMetadataURI
	EventTypeCreateTrack        = types.EventTypeCreateTrack
	EventTypePlayTrack          = types.EventTypePlayTrack
	EventTypeSetTrackStatus     = types.EventTypeSetTrackStatus
	EventTypeDepositTrack       = types.EventTypeDepositTrack
	AttributeValueCategory      = types.AttributeValueCategory
	AttributeKeyTrackID         = types.AttributeKeyTrackID
	AttributeKeyTrackTitle      = types.AttributeKeyTrackTitle
	AttributeKeyTrackOwner      = types.AttributeKeyTrackOwner
	AttributeKeyTrackStatus     = types.AttributeKeyTrackStatus
	AttributeKeyPlayAccAddr     = types.AttributeKeyPlayAccAddr
	MaxTitleLength              = types.MaxTitleLength
	MaxDescriptionLength        = types.MaxDescriptionLength
	MaxCopyrightLength          = types.MaxCopyrightLength
	StatusNil                   = types.StatusNil
	StatusDepositPeriod         = types.StatusDepositPeriod
	StatusVerified              = types.StatusVerified
	StatusRejected              = types.StatusRejected
	StatusFailed                = types.StatusFailed
)

var (
	// functions aliases
	NewQuerier                 = keeper.NewQuerier
	NewHandler                 = keeper.NewHandler
	NewKeeper                  = keeper.NewKeeper
	PlaysKey                   = types.PlaysKey
	PlayKey                    = types.PlayKey
	ShareKey                   = types.ShareKey
	DepositsKey                = types.DepositsKey
	DepositKey                 = types.DepositKey
	NewMsgCreateTrack          = types.NewMsgCreateTrack
	NewMsgPlay                 = types.NewMsgPlay
	NewMsgDeposit              = types.NewMsgDeposit
	ParamKeyTable              = types.ParamKeyTable
	DefaultParams              = types.DefaultParams
	NewDepositParams           = types.NewDepositParams
	NewQueryTrackParams        = types.NewQueryTrackParams
	NewQueryTracksParams       = types.NewQueryTracksParams
	RegisterCodec              = types.RegisterCodec
	NewDeposit                 = types.NewDeposit
	ErrInvalidGenesis          = types.ErrInvalidGenesis
	ErrInvalidTrackTitle       = types.ErrInvalidTrackTitle
	ErrUnknownTrack            = types.ErrUnknownTrack
	ErrInvalidTrackStatus      = types.ErrInvalidTrackStatus
	ErrInvalidTrackMetadataURI = types.ErrInvalidTrackMetadataURI
	NewPlay                    = types.NewPlay
	NewShare                   = types.NewShare
	TrackKey                   = types.TrackKey
	NewTrack                   = types.NewTrack
	TrackStatusFromString      = types.TrackStatusFromString

	// variable aliases
	TracksKeyPrefix            = types.TracksKeyPrefix
	TrackIDKey                 = types.TrackIDKey
	PlaysKeyPrefix             = types.PlaysKeyPrefix
	SharesKeyPrefix            = types.SharesKeyPrefix
	DepositsKeyPrefix          = types.DepositsKeyPrefix
	ParamStoreKeyDepositParams = types.ParamStoreKeyDepositParams
	DefaultDepositParams       = types.DefaultDepositParams
	ModuleCdc                  = types.ModuleCdc
)

type (
	Keeper            = keeper.Keeper
	MsgCreateTrack    = types.MsgCreateTrack
	MsgPlay           = types.MsgPlay
	MsgDeposit        = types.MsgDeposit
	Params            = types.Params
	DepositParams     = types.DepositParams
	QueryTrackParams  = types.QueryTrackParams
	QueryTracksParams = types.QueryTracksParams
	Deposit           = types.Deposit
	Deposits          = types.Deposits
	Play              = types.Play
	Plays             = types.Plays
	Share             = types.Share
	Shares            = types.Shares
	Track             = types.Track
	Tracks            = types.Tracks
	TrackStatus       = types.TrackStatus
)
