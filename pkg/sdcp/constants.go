package sdcp

type NetworkStatus string

const (
	NetworkStatusWlan NetworkStatus = "wlan"
	NetworkStatusEth  NetworkStatus = "eth"
)

type UsbDiskStatus uint8

const (
	UsbDiskStatusDisconnected UsbDiskStatus = 0
	UbsDiskStatusConnected    UsbDiskStatus = 1
)

type Capabilities string

const (
	CapabilitiesFileTransfer Capabilities = "FILE_TRANSFER"
	CapabilitiesPrintControl Capabilities = "PRINT_CONTROL"
	CapabilitiesVideoStream  Capabilities = "VIDEO_STREAM"
)

type SupportedFileType string

const (
	SupportedFileTypeCTB SupportedFileType = "CTB"
)

type TempSensorStatusOfUVLED uint8

const (
	TempSensorStatusOfUVLEDDisconnected TempSensorStatusOfUVLED = 0
	TempSensorStatusOfUVLEDNormal       TempSensorStatusOfUVLED = 1
	TempSensorStatusOfUVLEDAbnormal     TempSensorStatusOfUVLED = 2
)

type LCDStatus uint8

const (
	LCDStatusDisconnected LCDStatus = 0
	LCDStatusConnected    LCDStatus = 1
)

type SgStatus uint8

const (
	SgStatusDisconnected      SgStatus = 0
	SgStatusNormal            SgStatus = 1
	SgStatusCalibrationFailed SgStatus = 2
)

type ZMotorStatus uint8

const (
	ZMotorStatusDisconnected ZMotorStatus = 0
	ZMotorStatusConnected    ZMotorStatus = 1
)

type RotateMotorStatus uint8

const (
	RotateMotorStatusDisconnected RotateMotorStatus = 0
	RotateMotorStatusConnected    RotateMotorStatus = 1
)

type ReleaseFilmState uint8

const (
	ReleaseFilmStateAbnormal ReleaseFilmState = 0
	ReleaseFilmStateNormal   ReleaseFilmState = 1
)

type XMotorStatus uint8

const (
	XMotorStatusDisconnected XMotorStatus = 0
	XMotorStatusConnected    XMotorStatus = 1
)

type TimeLapseStatus uint8

const (
	TimeLapseStatusOff TimeLapseStatus = 0
	TimeLapseStatusOn  TimeLapseStatus = 1
)

type StorageType uint8

const (
	StorageTypeInternal StorageType = 0
	StorageTypeExternal StorageType = 1
)

type FileType uint8

const (
	FileTypeFolder FileType = 0
	FileTypeFile   FileType = 1
)

type EnableDisable uint8

const (
	EnableDisableDisable EnableDisable = 0
	EnableDisableEnable  EnableDisable = 1
)

type From uint8

const (
	FromPC     From = 0 // Local PC Software Local Area Network
	FromWebPC  From = 1 // PC Software via WEB
	FromWeb    From = 2 // Web Client
	FromApp    From = 3 // APP
	FromServer From = 4 // Server
)

type FileTransferAck uint8

const (
	FileTransferAckSuccess     FileTransferAck = 0 // Success
	FileTransferAckNotTransfer FileTransferAck = 1 // The printer is not currently transferring files.
	FileTransferAckChecking    FileTransferAck = 2 // The printer is already in the file verification phase.
	FileTransferAckNotFound    FileTransferAck = 3 // File not found.
)

type PrintInfoError uint8

const (
	PrintInfoErrorNone              PrintInfoError = 0 // Normal
	PrintInfoErrorCheck             PrintInfoError = 1 // File MD5 Check Failed
	PrintInfoErrorFileIO            PrintInfoError = 2 // File Read Failed
	PrintInfoErrorInvalidResolution PrintInfoError = 3 // Resolution Mismatch
	PrintInfoErrorUnknownFormat     PrintInfoError = 4 // Format Mismatch
	PrintInfoErrorUnknownModel      PrintInfoError = 5 // Machine Model Mismatch
)

type PrintInfoStatus uint8

const (
	PrintInfoStatusIdle         PrintInfoStatus = 0  // Idle
	PrintInfoStatusHoming       PrintInfoStatus = 1  // Homing
	PrintInfoStatusDropping     PrintInfoStatus = 2  // Dropping
	PrintInfoStatusExposing     PrintInfoStatus = 3  // Exposing
	PrintInfoStatusLifting      PrintInfoStatus = 4  // Lifting
	PrintInfoStatusPausing      PrintInfoStatus = 5  // Pausing
	PrintInfoStatusPaused       PrintInfoStatus = 6  // Paused
	PrintInfoStatusStopping     PrintInfoStatus = 7  // Stopping
	PrintInfoStatusStopped      PrintInfoStatus = 8  // Stopped
	PrintInfoStatusComplete     PrintInfoStatus = 9  // Complete
	PrintInfoStatusFileChecking PrintInfoStatus = 10 // File Checking
)

type MachineStatus uint8

const (
	MachineStatusIdle             MachineStatus = 0 // Idle
	MachineStatusPrinting         MachineStatus = 1 // Printing
	MachineStatusFileTransferring MachineStatus = 2 // File Transferring
	MachineStatusExposureTesting  MachineStatus = 3 // Exposure Testing
	MachineStatusDevicesTesting   MachineStatus = 4 // Devices Testing
)

type TaskStatus uint8

const (
	TaskStatusOther       TaskStatus = 0 // Other
	TaskStatusCompleted   TaskStatus = 1 // Completed
	TaskStatusExceptional TaskStatus = 2 // Exceptional
	TaskStatusStopped     TaskStatus = 3 // Stopped
)

type TimeLapseVideoStatus uint8

const (
	TimeLapseVideoStatusNotShot        TimeLapseVideoStatus = 0 // Not shot
	TimeLapseVideoStatusTimeLapseExist TimeLapseVideoStatus = 1 // Time-lapse photography file exists
	TimeLapseVideoStatusDeleted        TimeLapseVideoStatus = 2 // Deleted
	TimeLapseVideoStatusGenerating     TimeLapseVideoStatus = 3 // Generating
	TimeLapseVideoStatusGenerationFail TimeLapseVideoStatus = 4 // Generation failed
)

type ControlAck uint8

const (
	ControlAckOk                ControlAck = 0 // OK
	ControlAckBusy              ControlAck = 1 // Busy
	ControlAckNotFound          ControlAck = 2 // File Not Found
	ControlAckMd5FailFailed     ControlAck = 3 // MD5 Verification Failed
	ControlAckFileIOFailed      ControlAck = 4 // File Read Failed
	ControlAckInvalidResolution ControlAck = 5 // Resolution Mismatch
	ControlAckUnknownFormat     ControlAck = 6 // Unrecognized File Format
	ControlAckUnknownModel      ControlAck = 7 // Machine Model Mismatch
)

type TaskError uint8

const (
	TaskErrorOk                    TaskError = 0  // Normal
	TaskErrorTempError             TaskError = 1  // Over-temperature
	TaskErrorCalibrateFailed       TaskError = 2  // Strain Gauge Calibration Failed
	TaskErrorResinLack             TaskError = 3  // Resin Level Low Detected
	TaskErrorResinOver             TaskError = 4  // The volume of resin required by the model exceeds the maximum capacity of the resin vat
	TaskErrorProbeFail             TaskError = 5  // No Resin Detected
	TaskErrorForeignBody           TaskError = 6  // Foreign Object Detected
	TaskErrorLevelFailed           TaskError = 7  // Auto-leveling Failed
	TaskErrorReleaseFailed         TaskError = 8  // Model Detachment Detected
	TaskErrorSgOffline             TaskError = 9  // Strain Gauge Not Connected
	TaskErrorLcdDetFailed          TaskError = 10 // LCD Screen Connection Abnormal
	TaskErrorReleaseOvercount      TaskError = 11 // The cumulative release film usage has reached the maximum value
	TaskErrorUdiskRemove           TaskError = 12 // USB drive detected as removed, printing has been stopped
	TaskErrorHomeFailedX           TaskError = 13 // Detection of X-axis motor anomaly, printing has been stopped
	TaskErrorHomeFailedZ           TaskError = 14 // Detection of Z-axis motor anomaly, printing has been stopped
	TaskErrorResinAbnormalHigh     TaskError = 15 // The resin level has been detected to exceed the maximum value, and printing has been stopped
	TaskErrorResinAbnormalLow      TaskError = 16 // Resin level detected as too low, printing has been stopped
	TaskErrorHomeFailed            TaskError = 17 // Home position calibration failed, please check if the motor or limit switch is functioning properly
	TaskErrorPlatFailed            TaskError = 18 // A model is detected on the platform; please clean it and then restart printing
	TaskErrorError                 TaskError = 19 // Printing Exception
	TaskErrorMoveAbnormal          TaskError = 20 // Motor Movement Abnormality
	TaskErrorAicModelNone          TaskError = 21 // No model detected, please troubleshoot
	TaskErrorAicModelWarp          TaskError = 22 // Warping of the model detected, please investigate
	TaskErrorHomeFailedY           TaskError = 23 // Deprecated
	TaskErrorFileError             TaskError = 24 // Error File
	TaskErrorCameraError           TaskError = 25 // Camera Error. Please check if the camera is properly connected, or you can also disable this feature to continue printing
	TaskErrorNetworkError          TaskError = 26 // Network Connection Error. Please check if your network connection is stable, or you can also disable this feature to continue printing
	TaskErrorServerConnectFailed   TaskError = 27 // Server Connection Failed. Please contact our customer support, or you can also disable this feature to continue printing
	TaskErrorDisconnectApp         TaskError = 28 // This printer is not bound to an app. To perform time-lapse photography, please first enable the remote control feature, or you can also disable this feature to continue printing
	TaskErrorCheckAutoResinFeeder  TaskError = 29 // lease check the installation of the "automatic material extraction / feeding machine"
	TaskErrorContainerResinLow     TaskError = 30 // The resin in the container is running low. Add more resin to automatically close this notification, or click "Stop Auto Feeding" to continue printing
	TaskErrorBottleDisconnect      TaskError = 31 // Please ensure that the automatic material extraction/feeding machine is correctly installed and the data cable is connected
	TaskErrorFeedTimeout           TaskError = 32 // Automatic material extraction timeout, please check if the resin tube is blocked
	TaskErrorTankTempSensorOffline TaskError = 33 // Resin vat temperature sensor not connected
	TaskErrorTankTempSensorError   TaskError = 34 // Resin vat temperature sensor indicates an over-temperature condition
)

type ErrorCode uint8

const (
	ErrorCodeMD5Failed    ErrorCode = 1 // File Transfer MD5 Check Failed
	ErrorCodeFormatFailed ErrorCode = 2 // File format is incorrect
)

type NotificationType uint8

const (
	HistorySynchronizationSuccessful NotificationType = 1
)

type Command int

const (
	CommandStatusRefresh Command = 0
	CommandAttribute     Command = 1

	CommandStartPrint          Command = 128
	CommandPausePrint          Command = 129
	CommandStopPrint           Command = 130
	CommandResumePrint         Command = 131
	CommandStopFeedingMaterial Command = 132
	CommandSkipPreheating      Command = 133

	CommandChangePrinterName Command = 192

	CommandTerminateFileTransfer Command = 255

	CommandRetrieveFileList Command = 258
	CommandBatchDeleteFiles Command = 259

	CommandRetrieveHistoricalTasks Command = 320
	CommandRetrieveTaskDetails     Command = 321

	CommandEnableDisableVideoStream Command = 386
	CommandEnableDisableTimeLapse   Command = 387
)
