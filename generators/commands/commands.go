package commands

type ZWaveCommand byte

const (
    Alarm = 0x71
    AlarmV2 = 0x71
    NotificationV3 = 0x71
    NotificationV4 = 0x71
    NotificationV5 = 0x71
    NotificationV6 = 0x71
    NotificationV7 = 0x71
    NotificationV8 = 0x71
    ApplicationStatus = 0x22
    AssociationCommandConfiguration = 0x9B
    Association = 0x85
    AssociationV2 = 0x85
    AvContentDirectoryMd = 0x95
    AvContentSearchMd = 0x97
    AvRendererStatus = 0x96
    AvTaggingMd = 0x99
    BasicTariffInfo = 0x36
    BasicWindowCovering = 0x50
    Basic = 0x20
    BasicV2 = 0x20
    Battery = 0x80
    ChimneyFan = 0x2A
    ClimateControlSchedule = 0x46
    Clock = 0x81
    Configuration = 0x70
    ConfigurationV2 = 0x70
    ConfigurationV3 = 0x70
    ConfigurationV4 = 0x70
    ControllerReplication = 0x21
    Crc16Encap = 0x56
    DcpConfig = 0x3A
    DcpMonitor = 0x3B
    DoorLockLogging = 0x4C
    DoorLock = 0x62
    DoorLockV2 = 0x62
    DoorLockV3 = 0x62
    EnergyProduction = 0x90
    FirmwareUpdateMd = 0x7A
    FirmwareUpdateMdV2 = 0x7A
    FirmwareUpdateMdV3 = 0x7A
    FirmwareUpdateMdV4 = 0x7A
    FirmwareUpdateMdV5 = 0x7A
    GeographicLocation = 0x8C
    GroupingName = 0x7B
    Hail = 0x82
    HrvControl = 0x39
    HrvStatus = 0x37
    Indicator = 0x87
    IndicatorV2 = 0x87
    IpConfiguration = 0x9A
    Language = 0x89
    Lock = 0x76
    ManufacturerProprietary = 0x91
    ManufacturerSpecific = 0x72
    ManufacturerSpecificV2 = 0x72
    Mark = 0xEF
    MeterPulse = 0x35
    MeterTblConfig = 0x3C
    MeterTblMonitor = 0x3D
    MeterTblMonitorV2 = 0x3D
    MeterTblPush = 0x3E
    Meter = 0x32
    MeterV2 = 0x32
    MeterV3 = 0x32
    MeterV4 = 0x32
    MtpWindowCovering = 0x51
    MultiChannelAssociationV2 = 0x8E
    MultiChannelAssociationV3 = 0x8E
    MultiChannelV2 = 0x60
    MultiChannelV3 = 0x60
    MultiChannelV4 = 0x60
    MultiCmd = 0x8F
    MultiInstanceAssociation = 0x8E
    MultiInstance = 0x60
    NetworkManagementProxy = 0x52
    NetworkManagementProxyV2 = 0x52
    NetworkManagementBasic = 0x4D
    NetworkManagementBasicV2 = 0x4D
    NetworkManagementInclusion = 0x34
    NetworkManagementInclusionV2 = 0x34
    NodeNaming = 0x77
    NonInteroperable = 0xF0
    Powerlevel = 0x73
    PrepaymentEncapsulation = 0x41
    Prepayment = 0x3F
    Proprietary = 0x88
    Protection = 0x75
    ProtectionV2 = 0x75
    RateTblConfig = 0x48
    RateTblMonitor = 0x49
    RemoteAssociationActivate = 0x7C
    RemoteAssociation = 0x7D
    SceneActivation = 0x2B
    SceneActuatorConf = 0x2C
    SceneControllerConf = 0x2D
    ScheduleEntryLock = 0x4E
    ScheduleEntryLockV2 = 0x4E
    ScheduleEntryLockV3 = 0x4E
    ScreenAttributes = 0x93
    ScreenAttributesV2 = 0x93
    ScreenMd = 0x92
    ScreenMdV2 = 0x92
    SecurityPanelMode = 0x24
    SecurityPanelZoneSensor = 0x2F
    SecurityPanelZone = 0x2E
    Security = 0x98
    SensorAlarm = 0x9C
    SensorBinary = 0x30
    SensorBinaryV2 = 0x30
    SensorConfiguration = 0x9E
    SensorMultilevel = 0x31
    SensorMultilevelV2 = 0x31
    SensorMultilevelV3 = 0x31
    SensorMultilevelV4 = 0x31
    SensorMultilevelV5 = 0x31
    SensorMultilevelV6 = 0x31
    SensorMultilevelV7 = 0x31
    SensorMultilevelV8 = 0x31
    SensorMultilevelV9 = 0x31
    SensorMultilevelV10 = 0x31
    SilenceAlarm = 0x9D
    SimpleAvControl = 0x94
    SwitchAll = 0x27
    SwitchBinary = 0x25
    SwitchBinaryV2 = 0x25
    SwitchMultilevel = 0x26
    SwitchMultilevelV2 = 0x26
    SwitchMultilevelV3 = 0x26
    SwitchMultilevelV4 = 0x26
    SwitchToggleBinary = 0x28
    SwitchToggleMultilevel = 0x29
    TariffConfig = 0x4A
    TariffTblMonitor = 0x4B
    ThermostatFanMode = 0x44
    ThermostatFanModeV2 = 0x44
    ThermostatFanModeV3 = 0x44
    ThermostatFanModeV4 = 0x44
    ThermostatFanState = 0x45
    ThermostatFanStateV2 = 0x45
    ThermostatHeating = 0x38
    ThermostatMode = 0x40
    ThermostatModeV2 = 0x40
    ThermostatModeV3 = 0x40
    ThermostatOperatingState = 0x42
    ThermostatOperatingStateV2 = 0x42
    ThermostatSetback = 0x47
    ThermostatSetpoint = 0x43
    ThermostatSetpointV2 = 0x43
    ThermostatSetpointV3 = 0x43
    TimeParameters = 0x8B
    Time = 0x8A
    TimeV2 = 0x8A
    TransportServiceV2 = 0x55
    TransportService = 0x55
    UserCode = 0x63
    Version = 0x86
    VersionV2 = 0x86
    WakeUp = 0x84
    WakeUpV2 = 0x84
    ZensorNet = 0x02
    Zip6lowpan = 0x4F
    Zip = 0x23
    ZipV2 = 0x23
    ZipV3 = 0x23
    Class = 0x01
    ApplicationCapability = 0x57
    SwitchColor = 0x33
    SwitchColorV2 = 0x33
    SwitchColorV3 = 0x33
    Schedule = 0x53
    ScheduleV2 = 0x53
    ScheduleV3 = 0x53
    NetworkManagementPrimary = 0x54
    ZipNd = 0x58
    AssociationGrpInfo = 0x59
    AssociationGrpInfoV2 = 0x59
    AssociationGrpInfoV3 = 0x59
    DeviceResetLocally = 0x5A
    CentralScene = 0x5B
    CentralSceneV2 = 0x5B
    CentralSceneV3 = 0x5B
    IpAssociation = 0x5C
    Antitheft = 0x5D
    AntitheftV2 = 0x5D
    ZwaveplusInfo = 0x5E
    ZwaveplusInfoV2 = 0x5E
    ZipGateway = 0x5F
    ZipPortal = 0x61
    Dmx = 0x65
    BarrierOperator = 0x66
    NetworkManagementInstallationMaintenance = 0x67
    ZipNaming = 0x68
    Mailbox = 0x69
    WindowCovering = 0x6A
    Security2 = 0x9F
    Irrigation = 0x6B
    Supervision = 0x6C
    HumidityControlSetpoint = 0x64
    HumidityControlMode = 0x6D
    HumidityControlOperatingState = 0x6E
    EntryControl = 0x6F
    InclusionController = 0x74
)
