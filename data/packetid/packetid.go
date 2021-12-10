package packetid

// Login Clientbound
const (
	LoginDisconnect = iota
	LoginEncryptionRequest
	LoginSuccess
	SetCompression
	LoginPluginRequest
)

// Login Serverbound
const (
	LoginStart = iota
	LoginEncryptionResponse
	LoginPluginResponse
)

// Status Clientbound
const (
	StatusResponse = iota
	StatusPong
)

// Game Serverbound
const (
	StatusRequest = iota
	StatusPing
)

// Game Clientbound
const (
	ClientboundAddEntity = iota
	ClientboundAddExperienceOrb
	ClientboundAddMob
	ClientboundAddPainting
	ClientboundAddPlayer
	ClientboundAddVibrationSignal
	ClientboundAnimate
	ClientboundAwardStats
	ClientboundBlockBreakAck
	ClientboundBlockDestruction
	ClientboundBlockEntityData
	ClientboundBlockEvent
	ClientboundBlockUpdate
	ClientboundBossEvent
	ClientboundChangeDifficulty
	ClientboundChat
	ClientboundClearTitles
	ClientboundCommandSuggestions
	ClientboundCommands
	ClientboundContainerClose
	ClientboundContainerSetContent
	ClientboundContainerSetData
	ClientboundContainerSetSlot
	ClientboundCooldown
	ClientboundCustomPayload
	ClientboundCustomSound
	ClientboundDisconnect
	ClientboundEntityEvent
	ClientboundExplode
	ClientboundForgetLevelChunk
	ClientboundGameEvent
	ClientboundHorseScreenOpen
	ClientboundInitializeBorder
	ClientboundKeepAlive
	ClientboundLevelChunk
	ClientboundLevelEvent
	ClientboundLevelParticles
	ClientboundLightUpdate
	ClientboundLogin
	ClientboundMapItemData
	ClientboundMerchantOffers
	ClientboundEntityPosition
	ClientboundEntityPositionAndRotation
	ClientboundEntityRotation
	ClientboundMoveVehicle
	ClientboundOpenBook
	ClientboundOpenScreen
	ClientboundOpenSignEditor
	ClientboundPing
	ClientboundPlaceGhostRecipe
	ClientboundPlayerAbilities
	ClientboundPlayerCombatEnd
	ClientboundPlayerCombatEnter
	ClientboundPlayerCombatKill
	ClientboundPlayerInfo
	ClientboundPlayerLookAt
	ClientboundPlayerPosition
	ClientboundRecipe
	ClientboundRemoveEntities
	ClientboundRemoveMobEffect
	ClientboundResourcePack
	ClientboundRespawn
	ClientboundRotateHead
	ClientboundSectionBlocksUpdate
	ClientboundSelectAdvancementsTab
	ClientboundSetActionBarText
	ClientboundSetBorderCenter
	ClientboundSetBorderLerpSize
	ClientboundSetBorderSize
	ClientboundSetBorderWarningDelay
	ClientboundSetBorderWarningDistance
	ClientboundSetCamera
	ClientboundSetCarriedItem
	ClientboundSetChunkCacheCenter
	ClientboundSetChunkCacheRadius
	ClientboundSetDefaultSpawnPosition
	ClientboundSetDisplayObjective
	ClientboundSetEntityData
	ClientboundSetEntityLink
	ClientboundSetEntityMotion
	ClientboundSetEquipment
	ClientboundSetExperience
	ClientboundSetHealth
	ClientboundSetObjective
	ClientboundSetPassengers
	ClientboundSetPlayerTeam
	ClientboundSetScore
	ClientboundUpdateSimulationDistance
	ClientboundSetSubtitleText
	ClientboundSetTime
	ClientboundSetTitleText
	ClientboundSetTitlesAnimation
	ClientboundSoundEntity
	ClientboundSound
	ClientboundStopSound
	ClientboundTabList
	ClientboundTagQuery
	ClientboundTakeItemEntity
	ClientboundTeleportEntity
	ClientboundUpdateAdvancements
	ClientboundUpdateAttributes
	ClientboundUpdateMobEffect
	ClientboundUpdateRecipes
	ClientboundUpdateTags
)

// Game Serverbound
const (
	ServerboundAcceptTeleportation = iota
	ServerboundChangeDifficulty
	ServerboundChat
	ServerboundClientCommand
	ServerboundClientInformation
	ServerboundCommandSuggestion
	ServerboundContainerButtonClick
	ServerboundContainerClick
	ServerboundContainerClose
	ServerboundCustomPayload
	ServerboundEditBook
	ServerboundInteract
	ServerboundJigsawGenerate
	ServerboundKeepAlive
	ServerboundLockDifficulty
	ServerboundMoveVehicle
	ServerboundPaddleBoat
	ServerboundPickItem
	ServerboundPlaceRecipe
	ServerboundPlayerAbilities
	ServerboundPlayerAction
	ServerboundPlayerCommand
	ServerboundPlayerInput
	ServerboundPong
	ServerboundRecipeBookChangeSettings
	ServerboundRecipeBookSeenRecipe
	ServerboundRenameItem
	ServerboundResourcePack
	ServerboundSeenAdvancements
	ServerboundSelectTrade
	ServerboundSetBeacon
	ServerboundSetCarriedItem
	ServerboundSetCommandBlock
	ServerboundSetCommandMinecart
	ServerboundSetCreativeModeSlot
	ServerboundSetJigsawBlock
	ServerboundSetStructureBlock
	ServerboundSignUpdate
	ServerboundSwing
	ServerboundTeleportToEntity
	ServerboundUseItemOn
	ServerboundUseItem
)
