/*generated by internal*/
package dgsdk

/*
#include "src/c/discord_game_sdk.h"
*/
import "C"

type SkuType C.enum_EDiscordSkuType

const (
	SkuTypeApplication SkuType = C.DiscordSkuType_Application
	SkuTypeDLC                 = C.DiscordSkuType_DLC
	SkuTypeConsumable          = C.DiscordSkuType_Consumable
	SkuTypeBundle              = C.DiscordSkuType_Bundle
)
