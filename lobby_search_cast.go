/*generated by internal*/
package dgsdk

/*
#include "src/c/discord_game_sdk.h"
*/
import "C"

type LobbySearchCast C.enum_EDiscordLobbySearchCast

const (
	LobbySearchCastString LobbySearchCast = C.DiscordLobbySearchCast_String
	LobbySearchCastNumber                 = C.DiscordLobbySearchCast_Number
)
