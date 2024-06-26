// Code generated by "stringer -type ClientboundPacketID"; DO NOT EDIT.

package packetid

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ClientboundLoginLoginDisconnect-0]
	_ = x[ClientboundLoginHello-1]
	_ = x[ClientboundLoginGameProfile-2]
	_ = x[ClientboundLoginLoginCompression-3]
	_ = x[ClientboundLoginCustomQuery-4]
	_ = x[ClientboundLoginCookieRequest-5]
	_ = x[ClientboundStatusStatusResponse-0]
	_ = x[ClientboundStatusPongResponse-1]
	_ = x[ClientboundConfigCookieRequest-0]
	_ = x[ClientboundConfigCustomPayload-1]
	_ = x[ClientboundConfigDisconnect-2]
	_ = x[ClientboundConfigFinishConfiguration-3]
	_ = x[ClientboundConfigKeepAlive-4]
	_ = x[ClientboundConfigPing-5]
	_ = x[ClientboundConfigResetChat-6]
	_ = x[ClientboundConfigRegistryData-7]
	_ = x[ClientboundConfigResourcePackPop-8]
	_ = x[ClientboundConfigResourcePackPush-9]
	_ = x[ClientboundConfigStoreCookie-10]
	_ = x[ClientboundConfigTransfer-11]
	_ = x[ClientboundConfigUpdateEnabledFeatures-12]
	_ = x[ClientboundConfigUpdateTags-13]
	_ = x[ClientboundConfigSelectKnownPacks-14]
	_ = x[ClientboundConfigCustomReportDetails-15]
	_ = x[ClientboundConfigServerLinks-16]
	_ = x[BundleDelimiter-0]
	_ = x[ClientboundAddEntity-1]
	_ = x[ClientboundAddExperienceOrb-2]
	_ = x[ClientboundAnimate-3]
	_ = x[ClientboundAwardStats-4]
	_ = x[ClientboundBlockChangedAck-5]
	_ = x[ClientboundBlockDestruction-6]
	_ = x[ClientboundBlockEntityData-7]
	_ = x[ClientboundBlockEvent-8]
	_ = x[ClientboundBlockUpdate-9]
	_ = x[ClientboundBossEvent-10]
	_ = x[ClientboundChangeDifficulty-11]
	_ = x[ClientboundChunkBatchFinished-12]
	_ = x[ClientboundChunkBatchStart-13]
	_ = x[ClientboundChunksBiomes-14]
	_ = x[ClientboundClearTitles-15]
	_ = x[ClientboundCommandSuggestions-16]
	_ = x[ClientboundCommands-17]
	_ = x[ClientboundContainerClose-18]
	_ = x[ClientboundContainerSetContent-19]
	_ = x[ClientboundContainerSetData-20]
	_ = x[ClientboundContainerSetSlot-21]
	_ = x[ClientboundCookieRequest-22]
	_ = x[ClientboundCooldown-23]
	_ = x[ClientboundCustomChatCompletions-24]
	_ = x[ClientboundCustomPayload-25]
	_ = x[ClientboundDamageEvent-26]
	_ = x[ClientboundDebugSample-27]
	_ = x[ClientboundDeleteChat-28]
	_ = x[ClientboundDisconnect-29]
	_ = x[ClientboundDisguisedChat-30]
	_ = x[ClientboundEntityEvent-31]
	_ = x[ClientboundExplode-32]
	_ = x[ClientboundForgetLevelChunk-33]
	_ = x[ClientboundGameEvent-34]
	_ = x[ClientboundHorseScreenOpen-35]
	_ = x[ClientboundHurtAnimation-36]
	_ = x[ClientboundInitializeBorder-37]
	_ = x[ClientboundKeepAlive-38]
	_ = x[ClientboundLevelChunkWithLight-39]
	_ = x[ClientboundLevelEvent-40]
	_ = x[ClientboundLevelParticles-41]
	_ = x[ClientboundLightUpdate-42]
	_ = x[ClientboundLogin-43]
	_ = x[ClientboundMapItemData-44]
	_ = x[ClientboundMerchantOffers-45]
	_ = x[ClientboundMoveEntityPos-46]
	_ = x[ClientboundMoveEntityPosRot-47]
	_ = x[ClientboundMoveEntityRot-48]
	_ = x[ClientboundMoveVehicle-49]
	_ = x[ClientboundOpenBook-50]
	_ = x[ClientboundOpenScreen-51]
	_ = x[ClientboundOpenSignEditor-52]
	_ = x[ClientboundPing-53]
	_ = x[ClientboundPongResponse-54]
	_ = x[ClientboundPlaceGhostRecipe-55]
	_ = x[ClientboundPlayerAbilities-56]
	_ = x[ClientboundPlayerChat-57]
	_ = x[ClientboundPlayerCombatEnd-58]
	_ = x[ClientboundPlayerCombatEnter-59]
	_ = x[ClientboundPlayerCombatKill-60]
	_ = x[ClientboundPlayerInfoRemove-61]
	_ = x[ClientboundPlayerInfoUpdate-62]
	_ = x[ClientboundPlayerLookAt-63]
	_ = x[ClientboundPlayerPosition-64]
	_ = x[ClientboundRecipe-65]
	_ = x[ClientboundRemoveEntities-66]
	_ = x[ClientboundRemoveMobEffect-67]
	_ = x[ClientboundResetScore-68]
	_ = x[ClientboundResourcePackPop-69]
	_ = x[ClientboundResourcePackPush-70]
	_ = x[ClientboundRespawn-71]
	_ = x[ClientboundRotateHead-72]
	_ = x[ClientboundSectionBlocksUpdate-73]
	_ = x[ClientboundSelectAdvancementsTab-74]
	_ = x[ClientboundServerData-75]
	_ = x[ClientboundSetActionBarText-76]
	_ = x[ClientboundSetBorderCenter-77]
	_ = x[ClientboundSetBorderLerpSize-78]
	_ = x[ClientboundSetBorderSize-79]
	_ = x[ClientboundSetBorderWarningDelay-80]
	_ = x[ClientboundSetBorderWarningDistance-81]
	_ = x[ClientboundSetCamera-82]
	_ = x[ClientboundSetCarriedItem-83]
	_ = x[ClientboundSetChunkCacheCenter-84]
	_ = x[ClientboundSetChunkCacheRadius-85]
	_ = x[ClientboundSetDefaultSpawnPosition-86]
	_ = x[ClientboundSetDisplayObjective-87]
	_ = x[ClientboundSetEntityData-88]
	_ = x[ClientboundSetEntityLink-89]
	_ = x[ClientboundSetEntityMotion-90]
	_ = x[ClientboundSetEquipment-91]
	_ = x[ClientboundSetExperience-92]
	_ = x[ClientboundSetHealth-93]
	_ = x[ClientboundSetObjective-94]
	_ = x[ClientboundSetPassengers-95]
	_ = x[ClientboundSetPlayerTeam-96]
	_ = x[ClientboundSetScore-97]
	_ = x[ClientboundSetSimulationDistance-98]
	_ = x[ClientboundSetSubtitleText-99]
	_ = x[ClientboundSetTime-100]
	_ = x[ClientboundSetTitleText-101]
	_ = x[ClientboundSetTitlesAnimation-102]
	_ = x[ClientboundSoundEntity-103]
	_ = x[ClientboundSound-104]
	_ = x[ClientboundStartConfiguration-105]
	_ = x[ClientboundStopSound-106]
	_ = x[ClientboundStoreCookie-107]
	_ = x[ClientboundSystemChat-108]
	_ = x[ClientboundTabList-109]
	_ = x[ClientboundTagQuery-110]
	_ = x[ClientboundTakeItemEntity-111]
	_ = x[ClientboundTeleportEntity-112]
	_ = x[ClientboundTickingState-113]
	_ = x[ClientboundTickingStep-114]
	_ = x[ClientboundTransfer-115]
	_ = x[ClientboundUpdateAdvancements-116]
	_ = x[ClientboundUpdateAttributes-117]
	_ = x[ClientboundUpdateMobEffect-118]
	_ = x[ClientboundUpdateRecipes-119]
	_ = x[ClientboundUpdateTags-120]
	_ = x[ClientboundProjectilePower-121]
	_ = x[ClientboundCustomReportDetails-122]
	_ = x[ClientboundServerLinks-123]
	_ = x[ClientboundPacketIDGuard-124]
}

const _ClientboundPacketID_name = "ClientboundLoginLoginDisconnectClientboundLoginHelloClientboundLoginGameProfileClientboundLoginLoginCompressionClientboundLoginCustomQueryClientboundLoginCookieRequestClientboundConfigResetChatClientboundConfigRegistryDataClientboundConfigResourcePackPopClientboundConfigResourcePackPushClientboundConfigStoreCookieClientboundConfigTransferClientboundConfigUpdateEnabledFeaturesClientboundConfigUpdateTagsClientboundConfigSelectKnownPacksClientboundConfigCustomReportDetailsClientboundConfigServerLinksClientboundCommandsClientboundContainerCloseClientboundContainerSetContentClientboundContainerSetDataClientboundContainerSetSlotClientboundCookieRequestClientboundCooldownClientboundCustomChatCompletionsClientboundCustomPayloadClientboundDamageEventClientboundDebugSampleClientboundDeleteChatClientboundDisconnectClientboundDisguisedChatClientboundEntityEventClientboundExplodeClientboundForgetLevelChunkClientboundGameEventClientboundHorseScreenOpenClientboundHurtAnimationClientboundInitializeBorderClientboundKeepAliveClientboundLevelChunkWithLightClientboundLevelEventClientboundLevelParticlesClientboundLightUpdateClientboundLoginClientboundMapItemDataClientboundMerchantOffersClientboundMoveEntityPosClientboundMoveEntityPosRotClientboundMoveEntityRotClientboundMoveVehicleClientboundOpenBookClientboundOpenScreenClientboundOpenSignEditorClientboundPingClientboundPongResponseClientboundPlaceGhostRecipeClientboundPlayerAbilitiesClientboundPlayerChatClientboundPlayerCombatEndClientboundPlayerCombatEnterClientboundPlayerCombatKillClientboundPlayerInfoRemoveClientboundPlayerInfoUpdateClientboundPlayerLookAtClientboundPlayerPositionClientboundRecipeClientboundRemoveEntitiesClientboundRemoveMobEffectClientboundResetScoreClientboundResourcePackPopClientboundResourcePackPushClientboundRespawnClientboundRotateHeadClientboundSectionBlocksUpdateClientboundSelectAdvancementsTabClientboundServerDataClientboundSetActionBarTextClientboundSetBorderCenterClientboundSetBorderLerpSizeClientboundSetBorderSizeClientboundSetBorderWarningDelayClientboundSetBorderWarningDistanceClientboundSetCameraClientboundSetCarriedItemClientboundSetChunkCacheCenterClientboundSetChunkCacheRadiusClientboundSetDefaultSpawnPositionClientboundSetDisplayObjectiveClientboundSetEntityDataClientboundSetEntityLinkClientboundSetEntityMotionClientboundSetEquipmentClientboundSetExperienceClientboundSetHealthClientboundSetObjectiveClientboundSetPassengersClientboundSetPlayerTeamClientboundSetScoreClientboundSetSimulationDistanceClientboundSetSubtitleTextClientboundSetTimeClientboundSetTitleTextClientboundSetTitlesAnimationClientboundSoundEntityClientboundSoundClientboundStartConfigurationClientboundStopSoundClientboundStoreCookieClientboundSystemChatClientboundTabListClientboundTagQueryClientboundTakeItemEntityClientboundTeleportEntityClientboundTickingStateClientboundTickingStepClientboundTransferClientboundUpdateAdvancementsClientboundUpdateAttributesClientboundUpdateMobEffectClientboundUpdateRecipesClientboundUpdateTagsClientboundProjectilePowerClientboundCustomReportDetailsClientboundServerLinksClientboundPacketIDGuard"

var _ClientboundPacketID_index = [...]uint16{0, 31, 52, 79, 111, 138, 167, 193, 222, 254, 287, 315, 340, 378, 405, 438, 474, 502, 521, 546, 576, 603, 630, 654, 673, 705, 729, 751, 773, 794, 815, 839, 861, 879, 906, 926, 952, 976, 1003, 1023, 1053, 1074, 1099, 1121, 1137, 1159, 1184, 1208, 1235, 1259, 1281, 1300, 1321, 1346, 1361, 1384, 1411, 1437, 1458, 1484, 1512, 1539, 1566, 1593, 1616, 1641, 1658, 1683, 1709, 1730, 1756, 1783, 1801, 1822, 1852, 1884, 1905, 1932, 1958, 1986, 2010, 2042, 2077, 2097, 2122, 2152, 2182, 2216, 2246, 2270, 2294, 2320, 2343, 2367, 2387, 2410, 2434, 2458, 2477, 2509, 2535, 2553, 2576, 2605, 2627, 2643, 2672, 2692, 2714, 2735, 2753, 2772, 2797, 2822, 2845, 2867, 2886, 2915, 2942, 2968, 2992, 3013, 3039, 3069, 3091, 3115}

func (i ClientboundPacketID) String() string {
	if i < 0 || i >= ClientboundPacketID(len(_ClientboundPacketID_index)-1) {
		return "ClientboundPacketID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ClientboundPacketID_name[_ClientboundPacketID_index[i]:_ClientboundPacketID_index[i+1]]
}
