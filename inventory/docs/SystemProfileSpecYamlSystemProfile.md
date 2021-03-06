# SystemProfileSpecYamlSystemProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | **string** |  | [optional] 
**BiosReleaseDate** | **string** |  | [optional] 
**BiosVendor** | **string** |  | [optional] 
**BiosVersion** | **string** |  | [optional] 
**CapturedDate** | **string** |  | [optional] 
**CloudProvider** | **string** |  | [optional] 
**CoresPerSocket** | **int32** |  | [optional] 
**CpuFlags** | **[]string** |  | [optional] 
**CpuModel** | **string** | The cpu model name | [optional] 
**DiskDevices** | [**[]SystemProfileSpecYamlDiskDevice**](system_profile.spec.yaml_DiskDevice.md) |  | [optional] 
**DnfModules** | [**[]SystemProfileSpecYamlDnfModule**](system_profile.spec.yaml_DnfModule.md) |  | [optional] 
**EnabledServices** | **[]string** |  | [optional] 
**InfrastructureType** | **string** |  | [optional] 
**InfrastructureVendor** | **string** |  | [optional] 
**InsightsClientVersion** | **string** |  | [optional] 
**InsightsEggVersion** | **string** |  | [optional] 
**InstalledPackages** | **[]string** |  | [optional] 
**InstalledPackagesDelta** | **[]string** |  | [optional] 
**InstalledProducts** | [**[]SystemProfileSpecYamlInstalledProduct**](system_profile.spec.yaml_InstalledProduct.md) |  | [optional] 
**InstalledServices** | **[]string** |  | [optional] 
**IsMarketplace** | **bool** | Indicates whether the host is part of a marketplace install from AWS, Azure, etc. | [optional] 
**KatelloAgentRunning** | **bool** |  | [optional] 
**KernelModules** | **[]string** |  | [optional] 
**LastBootTime** | **string** |  | [optional] 
**NetworkInterfaces** | [**[]SystemProfileSpecYamlNetworkInterface**](system_profile.spec.yaml_NetworkInterface.md) |  | [optional] 
**NumberOfCpus** | **int32** |  | [optional] 
**NumberOfSockets** | **int32** |  | [optional] 
**OperatingSystem** | [**SystemProfileSpecYamlSystemProfileOperatingSystem**](system_profile_spec_yaml_SystemProfile_operating_system.md) |  | [optional] 
**OsKernelVersion** | **string** | The kernel version represented with a three, optionally four, number scheme. | [optional] 
**OsRelease** | **string** |  | [optional] 
**OwnerId** | **string** | A UUID associated with the host&#39;s RHSM certificate | [optional] 
**RhcClientId** | **string** | A UUID associated with a cloud_connector | [optional] 
**Rhsm** | [**SystemProfileSpecYamlSystemProfileRhsm**](system_profile_spec_yaml_SystemProfile_rhsm.md) |  | [optional] 
**RunningProcesses** | **[]string** |  | [optional] 
**SapInstanceNumber** | **string** | The instance number of the SAP HANA system (a two-digit number between 00 and 99) | [optional] 
**SapSids** | **[]string** |  | [optional] 
**SapSystem** | **bool** | Indicates if SAP is installed on the system | [optional] 
**SapVersion** | **string** | The version of the SAP HANA lifecycle management program | [optional] 
**SatelliteManaged** | **bool** |  | [optional] 
**SelinuxConfigFile** | **string** | The SELinux mode provided in the config file | [optional] 
**SelinuxCurrentMode** | **string** | The current SELinux mode, either enforcing, permissive, or disabled | [optional] 
**SubscriptionAutoAttach** | **string** |  | [optional] 
**SubscriptionStatus** | **string** |  | [optional] 
**SystemMemoryBytes** | **int64** |  | [optional] 
**TunedProfile** | **string** | Current profile resulting from command tuned-adm active | [optional] 
**YumRepos** | [**[]SystemProfileSpecYamlYumRepo**](system_profile.spec.yaml_YumRepo.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


