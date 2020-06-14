/*generated by internal*/
package dgsdk

/*
#include "src/c/discord_game_sdk.h"
*/
import "C"

type Result C.enum_EDiscordResult

const (
	ResultOk                              Result = C.DiscordResult_Ok
	ResultServiceUnavailable                     = C.DiscordResult_ServiceUnavailable
	ResultInvalidVersion                         = C.DiscordResult_InvalidVersion
	ResultLockFailed                             = C.DiscordResult_LockFailed
	ResultInternalError                          = C.DiscordResult_InternalError
	ResultInvalidPayload                         = C.DiscordResult_InvalidPayload
	ResultInvalidCommand                         = C.DiscordResult_InvalidCommand
	ResultInvalidPermissions                     = C.DiscordResult_InvalidPermissions
	ResultNotFetched                             = C.DiscordResult_NotFetched
	ResultNotFound                               = C.DiscordResult_NotFound
	ResultConflict                               = C.DiscordResult_Conflict
	ResultInvalidSecret                          = C.DiscordResult_InvalidSecret
	ResultInvalidJoinSecret                      = C.DiscordResult_InvalidJoinSecret
	ResultNoEligibleActivity                     = C.DiscordResult_NoEligibleActivity
	ResultInvalidInvite                          = C.DiscordResult_InvalidInvite
	ResultNotAuthenticated                       = C.DiscordResult_NotAuthenticated
	ResultInvalidAccessToken                     = C.DiscordResult_InvalidAccessToken
	ResultApplicationMismatch                    = C.DiscordResult_ApplicationMismatch
	ResultInvalidDataUrl                         = C.DiscordResult_InvalidDataUrl
	ResultInvalidBase64                          = C.DiscordResult_InvalidBase64
	ResultNotFiltered                            = C.DiscordResult_NotFiltered
	ResultLobbyFull                              = C.DiscordResult_LobbyFull
	ResultInvalidLobbySecret                     = C.DiscordResult_InvalidLobbySecret
	ResultInvalidFilename                        = C.DiscordResult_InvalidFilename
	ResultInvalidFileSize                        = C.DiscordResult_InvalidFileSize
	ResultInvalidEntitlement                     = C.DiscordResult_InvalidEntitlement
	ResultNotInstalled                           = C.DiscordResult_NotInstalled
	ResultNotRunning                             = C.DiscordResult_NotRunning
	ResultInsufficientBuffer                     = C.DiscordResult_InsufficientBuffer
	ResultPurchaseCanceled                       = C.DiscordResult_PurchaseCanceled
	ResultInvalidGuild                           = C.DiscordResult_InvalidGuild
	ResultInvalidEvent                           = C.DiscordResult_InvalidEvent
	ResultInvalidChannel                         = C.DiscordResult_InvalidChannel
	ResultInvalidOrigin                          = C.DiscordResult_InvalidOrigin
	ResultRateLimited                            = C.DiscordResult_RateLimited
	ResultOAuth2Error                            = C.DiscordResult_OAuth2Error
	ResultSelectChannelTimeout                   = C.DiscordResult_SelectChannelTimeout
	ResultGetGuildTimeout                        = C.DiscordResult_GetGuildTimeout
	ResultSelectVoiceForceRequired               = C.DiscordResult_SelectVoiceForceRequired
	ResultCaptureShortcutAlreadyListening        = C.DiscordResult_CaptureShortcutAlreadyListening
	ResultUnauthorizedForAchievement             = C.DiscordResult_UnauthorizedForAchievement
	ResultInvalidGiftCode                        = C.DiscordResult_InvalidGiftCode
	ResultPurchaseError                          = C.DiscordResult_PurchaseError
	ResultTransactionAborted                     = C.DiscordResult_TransactionAborted
)
