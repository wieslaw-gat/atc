package migrations

import "github.com/concourse/atc/dbng/migration"

var Migrations = []migration.Migrator{
	InitialSchema,
	MoveSourceAndMetadataToVersionedResources,
	AddTypeToVersionedResources,
	RemoveTransitionalCurrentVersions,
	NonNullableVersionInfo,
	AddOneOffNameSequence,
	AddHijackURLToBuilds,
	AddTimestampsToBuilds,
	CreateLocks,
	AddBuildEvents,
	ReplaceBuildsAbortHijackURLsWithGuidAndEndpoint,
	ReplaceBuildEventsIDWithEventID,
	AddLocks,
	DropOldLocks,
	AddConfig,
	AddNameToBuildInputs,
	AddEngineAndEngineMetadataToBuilds,
	AddVersionToBuildEvents,
	AddCompletedToBuilds,
	AddWorkers,
	AddEnabledToBuilds,
	CreateEventIDSequencesForInFlightBuilds,
	AddResourceTypesToWorkers,
	AddPlatformAndTagsToWorkers,
	AddIdToConfig,
	ConvertJobBuildConfigToJobPlans,
	AddCheckErrorToResources,
	AddPausedToResources,
	AddPausedToJobs,
	CreateJobsSerialGroups,
	CreatePipes,
	RenameConfigToPipelines,
	RenamePipelineIDToVersionAddPrimaryKey,
	AddNameToPipelines,
	AddPipelineIDToResources,
	AddPipelineIDToJobs,
	AddPausedToPipelines,
	AddOrderingToPipelines,
	AddInputsDeterminedToBuilds,
	AddExplicitToBuildOutputs,
	AddLastCheckedToResources,
	AddLastTrackedToBuilds,
	AddLastScheduledToPipelines,
	AddCheckingToResources,
	AddUniqueConstraintToResources,
	RemoveSourceFromVersionedResources,
	AddIndexesToABunchOfStuff,
	DropLocks,
	AddBaggageclaimURLToWorkers,
	AddContainers,
	AddNameToWorkers,
	AddLastScheduledToBuilds,
	AddCheckTypeAndCheckSourceToContainers,
	AddStepLocationToContainers,
	AddVolumesAndCacheInvalidator,
	AddCompositeUniqueConstraintToVolumes,
	AddWorkingDirectoryToContainers,
	MakeContainerWorkingDirectoryNotNull,
	AddEnvVariablesToContainers,
	AddModifiedTimeToVersionedResourcesAndBuildOutputs,
	ReplaceStepLocationWithPlanID,
	AddTeamsColumnToPipelinesAndTeamsTable,
	CascadePipelineDeletes,
	AddTeamIDToPipelineNameUniqueness,
	MakeVolumesExpiresAtNullable,
	AddAuthFieldsToTeams,
	AddAdminToTeams,
	MakeContainersLinkToPipelineIds,
	MakeContainersLinkToResourceIds,
	MakeContainersBuildIdsNullable,
	MakeContainersLinkToWorkerIds,
	RemoveVolumesWithExpiredWorkers,
	AddWorkerIDToVolumes,
	RemoveWorkerIds,
	AddAttemptsToContainers,
	AddStageToContainers,
	AddImageResourceVersions,
	MakeContainerIdentifiersUnique,
	CleanUpMassiveUniqueConstraint,
	AddPipelineBuildEventsTables,
	AddBuildPreparation,
	DropCompletedFromBuildPreparation,
	AddInputsSatisfiedToBuildPreparation,
	AddOrderToVersionedResources,
	AddImageResourceTypeAndSourceToContainers,
	AddUserToContainer,
	ResetPendingBuilds,
	ResetCheckOrder,
	AddTTLToContainers,
	AddOriginalVolumeHandleToVolumes,
	DropNotNullResourceConstraintsOnVolumes,
	AddOutputNameToVolumes,
	CreateResourceTypes,
	AddLastCheckedAndCheckingToResourceTypes,
	AddHttpProxyHttpsProxyNoProxyToWorkers,
	AddModifiedTimeToBuildInputs,
	AddPathToVolumes,
	AddHostPathVersionToVolumes,
	AddBestIfUsedByToContainers,
	AddStartTimeToWorkers,
	AddReplicatedFromToVolumes,
	AddSizeToVolumes,
	AddFirstLoggedBuildIDToJobsAndReapTimeToBuildsAndLeases,
	AddMissingInputReasonsToBuildPreparation,
	MakeVolumeSizeBigint,
	MakeContainersExpiresAtNullable,
	AddContainerIDToVolumes,
	AddOnDeleteSetNullToFKeyContainerId,
	AddUAAAuthToTeams,
	AddTeamIDToBuilds,
	AddPublicToPipelines,
	AddTeamIDToWorkers,
	AddTeamIDToContainers,
	AddTeamIDToVolumes,
	AddNextBuildInputs,
	AddCaseInsenstiveUniqueIndexToTeamsName,
	AddNonEmptyConstraintToTeamName,
	AddGenericOAuthToTeams,
	MigrateFromLeasesToLocks,
	AddTeamNameToPipe,
	AddConfigToJobsResources,
	CascadeTeamDeletes,
	CascadeTeamDeletesOnPipes,
	RemoveResourceCheckingFromJobsAndAddManualyTriggeredToBuilds,
	AddStateToWorkers,
	AddWorkerForeignKeyToVolumesAndContainers,
	UpdateWorkerForeignKeyConstraint,
	AddRetiringWorkerState,
	AddRunningWorkerMustHaveAddrConstraint,
	AddInterruptibleToJob,
	AddLandedWorkerCannotHaveAddrConstraint,
	FixWorkerAddrConstraint,
	AddCertificatesPathToWorkers,
	CreateCaches,
	RemoveTTLFromVolumes,
	AddVolumeParentIdForeignKey,
	DeleteExtraParentConstrainOnVolume,
	AddNotNullConstraintToContainerHandle,
	RemoveTTLFromContainers,
	AddDiscontinuedToContainers,
	AddSourceHashToResources,
	AddWorkerBaseResourceTypeIdToContainers,
	AddInterceptibleToBuilds,
	AlterExpiresToIncludeTimezoneInWorkers,
	ChangeVolumeBaseResourceTypeToWorkerBaseResourceType,
	RenameWorkerBaseResourceTypesId,
	AddWorkerResourceCacheToContainers,
	AddWorkerResourceCacheToVolumes,
	RemoveLastTrackedFromBuilds,
	AddIndexesToABunchMoreStuff,
	RemoveDuplicateIndices,
	CleanUpContainerColumns,
	AddAuthToTeams,
	RemoveCertificatesPathToWorkers,
	AddVersionToWorkers,
	AddNonceToTeams,
	AddNonceToResourcesAndResourceTypes,
	AddMetadataToResourceCache,
	AddNonceToJobs,
	UseMd5ForVersions,
}
