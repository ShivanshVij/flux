package sdcp

type DiscoverData struct {
	MachineName     string `json:"Name"`            // Machine Name
	MachineModel    string `json:"MachineName"`     // Machine Model
	BrandName       string `json:"BrandName"`       // Brand Name
	MainboardIP     string `json:"MainboardIP"`     // Motherboard IP Address
	MainboardID     string `json:"MainboardID"`     // Motherboard ID (16-bit)
	ProtocolVersion string `json:"ProtocolVersion"` // Protocol Version
	FirmwareVersion string `json:"FirmwareVersion"` // Firmware Version
}

type DeviceStatus struct {
	TempSensorStatusOfUVLED TempSensorStatusOfUVLED `json:"TempSensorStatusOfUVLED"` // UVLED Temperature Sensor Status
	LCDStatus               LCDStatus               `json:"LCDStatus"`               // Exposure Screen Connection Status
	SgStatus                SgStatus                `json:"SgStatus"`                // Strain Gauge Status
	ZMotorStatus            ZMotorStatus            `json:"ZMotorStatus"`            // Z-Axis Motor Connection Status
	RotateMotorStatus       RotateMotorStatus       `json:"RotateMotorStatus"`       // Rotary Axis Motor Connection Status
	ReleaseFilmState        ReleaseFilmState        `json:"ReleaseFilmState"`        // Release Film Status
	XMotorStatus            XMotorStatus            `json:"XMotorStatus"`            // X-Axis Motor Connection Status
}

type Attributes struct {
	MachineName                  string              `json:"Name"`                         // Machine Name
	MachineModel                 string              `json:"MachineName"`                  // Machine Model
	BrandName                    string              `json:"BrandName"`                    // Brand Name
	ProtocolVersion              string              `json:"ProtocolVersion"`              // Protocol Version
	FirmwareVersion              string              `json:"FirmwareVersion"`              // Firmware Version
	Resolution                   string              `json:"Resolution"`                   // Resolution
	XYZsize                      string              `json:"XYZsize"`                      // Maximum printing dimensions in the XYZ directions of the machine (millimeters)
	MainboardIP                  string              `json:"MainboardIP"`                  // Motherboard IP Address
	MainboardID                  string              `json:"MainboardID"`                  // Motherboard ID (16-bit)
	NumberOfVideoStreamConnected int                 `json:"NumberOfVideoStreamConnected"` // Number of Connected Video Streams
	MaximumVideoStreamAllowed    int                 `json:"MaximumVideoStreamAllowed"`    // Maximum Number of Connections for Video Streams
	NetworkStatus                NetworkStatus       `json:"NetworkStatus"`                // Network Connection Status
	UsbDiskStatus                UsbDiskStatus       `json:"UsbDiskStatus"`                // USB Drive Connection Status
	Capabilities                 []Capabilities      `json:"Capabilities"`                 // Supported Sub-protocols on the Motherboard
	SupportFileType              []SupportedFileType `json:"SupportFileType"`              // Supported File Types
	DevicesStatus                DeviceStatus        `json:"DevicesStatus"`                // Device Self-Check Status
	ReleaseFilmMax               int                 `json:"ReleaseFilmMax"`               // Maximum number of uses (service life) for the release film
	TempOfUVLEDMax               float64             `json:"TempOfUVLEDMax"`               // Maximum operating temperature for UVLED (Celsius)
	CameraStatus                 CameraStatus        `json:"CameraStatus"`                 // Camera Connection Status
	RemainingMemory              int                 `json:"RemainingMemory"`              // Remaining File Storage Space Size (bits)
	TLPNoCapPos                  float64             `json:"TLPNoCapPos"`                  // Model height threshold for not performing time-lapse photography (millimeters)
	TLPStartCapPos               float64             `json:"TLPStartCapPos"`               // The print height at which time-lapse photography begins (millimeters)
	TLPInterLayers               int                 `json:"TLPInterLayers"`               // Time-lapse photography shooting interval layers
}

type PrintInfo struct {
	Status       PrintInfoStatus `json:"Status"`       // Printing Sub-status
	CurrentLayer int             `json:"CurrentLayer"` // Current Printing Layer
	TotalLayer   int             `json:"TotalLayer"`   // Total Number of Print Layers
	CurrentTicks int             `json:"CurrentTicks"` // Current Print Time (milliseconds)
	TotalTicks   int             `json:"TotalTicks"`   // Estimated Total Print Time (milliseconds)
	Filename     string          `json:"Filename"`     // Print File Name
	ErrorNumber  PrintInfoError  `json:"ErrorNumber"`  // Error Number
	TaskId       string          `json:"TaskId"`       // Current Task ID
}

type Status struct {
	CurrentStatus   []MachineStatus `json:"CurrentStatus"`   // Current Machine Status
	PreviousStatus  MachineStatus   `json:"PreviousStatus"`  // Previous Machine Status
	PrintScreen     float64         `json:"PrintScreen"`     // Total Exposure Screen Usage Time (seconds)
	ReleaseFilm     int             `json:"ReleaseFilm"`     // Total Release Film Usage Count
	TempOfUVLED     float64         `json:"TempOfUVLED"`     // Current UVLED Temperature (Celsius)
	TimeLapseStatus TimeLapseStatus `json:"TimeLapseStatus"` // Time-lapse Photography Switch Status
	TempOfBox       float64         `json:"TempOfBox"`       // Current Enclosure Temperature (Celsius)
	TempTargetBox   float64         `json:"TempTargetBox"`   // Target Enclosure Temperature (Celsius)
	PrintInfo       PrintInfo       `json:"PrintInfo"`
}

type FileList struct {
	Name        Path        `json:"name"`      // Current File or Folder Path
	UsedSize    int         `json:"usedSize"`  // Used Storage Space (bytes)
	TotalSize   int         `json:"totalSize"` // Total Storage Space (bytes)
	StorageType StorageType `json:"storageType"`
	Type        FileType    `json:"type"`
}

type TaskDetails struct {
	Thumbnail             string               `json:"Thumbnail"`             // Thumbnail Address
	TaskName              string               `json:"TaskName"`              // Task Name
	BeginTime             int                  `json:"BeginTime"`             // Start Time (Timestamp in Seconds)
	EndTime               int                  `json:"EndTime"`               // End Time (Timestamp in Seconds)
	TaskStatus            TaskStatus           `json:"TaskStatus"`            // Task Status
	SliceInformation      struct{}             `json:"SliceInformation"`      // Slice Information
	AlreadyPrintLayer     int                  `json:"AlreadyPrintLayer"`     // Printed Layer Count
	TaskId                string               `json:"TaskId"`                // Task ID
	MD5                   string               `json:"MD5"`                   // MD5 of the Sliced File
	CurrentLayerTalVolume float64              `json:"CurrentLayerTalVolume"` // Total Volume of Printed Layers (milliliters)
	TimeLapseVideoStatus  TimeLapseVideoStatus `json:"TimeLapseVideoStatus"`  // Time-lapse photography status
	TimeLapseVideoUrl     string               `json:"TimeLapseVideoUrl"`     // URL for the time-lapse photography video
	ErrorStatusReason     TaskError            `json:"ErrorStatusReason"`     // Status Code
}

type ErrorCodeData struct {
	ErrorCode ErrorCode `json:"ErrorCode"` // Error Code
}

type ErrorData struct {
	Data        ErrorCodeData `json:"Data"`        // Error Data
	MainboardID string        `json:"MainboardID"` // Motherboard ID (16-bit)
	TimeStamp   int           `json:"TimeStamp"`   // Timestamp
}

type NotificationTypeData struct {
	Message string           `json:"Message"` // Can be a string, can be JSON
	Type    NotificationType `json:"Type"`    // Notification Type
}

type NotificationData struct {
	Data        NotificationTypeData `json:"Data"`        // Notification Type Data
	MainboardID string               `json:"MainboardID"` // Motherboard ID (16-bit)
	TimeStamp   int                  `json:"TimeStamp"`   // Timestamp
}

type TerminateFileTransferRequest struct {
	Uuid     string `json:"Uuid"`     // UUID for File Sending
	FileName string `json:"FileName"` // File Name
}

type TerminateFileTransferResponse struct {
	Ack FileTransferAck `json:"Ack"` // Acknowledgement
}

type StatusRefreshRequest struct{}

type StatusRefreshResponse struct {
	Ack int `json:"Ack"` // Acknowledgement
}

type AttributesRequest struct{}

type AttributesResponse struct {
	Ack int `json:"Ack"` // Acknowledgement
}

type StartPrintingRequest struct {
	Filename   string `json:"Filename"`   // File Name or File Path
	StartLayer int    `json:"StartLayer"` // Start Printing Layer Number
}

type StartPrintingResponse struct {
	Ack ControlAck `json:"Ack"` // Acknowledgement
}

type PausePrintingRequest struct{}

type PausePrintingResponse struct {
	Ack int `json:"Ack"` // Acknowledgement
}

type StopPrintingRequest struct{}

type StopPrintingResponse struct {
	Ack int `json:"Ack"` // Acknowledgement
}

type ResumePrintingRequest struct{}

type ResumePrintingResponse struct {
	Ack int `json:"Ack"` // Acknowledgement
}

type StopFeedingMaterialRequest struct{}

type StopFeedingMaterialResponse struct {
	Ack int `json:"Ack"` // Acknowledgement
}

type SkipPreheatingRequest struct{}

type SkipPreheatingResponse struct {
	Ack int `json:"Ack"` // Acknowledgement
}

type ChangePrinterNameRequest struct {
	Name string `json:"Name"` // New Printer Name
}

type ChangePrinterNameResponse struct {
	Ack int `json:"Ack"` // Acknowledgement
}

type RetrieveFileListRequest struct {
	Url Path `json:"Url"` // URL for File List
}

type RetrieveFileListResponse struct {
	Ack      int        `json:"Ack"`
	FileList []FileList `json:"FileList"`
}

type BatchDeleteFilesRequest struct {
	FileList   []Path `json:"FileList"`   // List of File Names to Delete
	FolderList []Path `json:"FolderList"` // List of Folder Names to Delete
}

type BatchDeleteFilesResponse struct {
	Ack     int    `json:"Ack"`               // Acknowledgement
	ErrData []Path `json:"ErrData,omitempty"` // List of Files or Folders that Failed to Delete
}

type RetrieveHistoricalTasksRequest struct{}

type RetrieveHistoricalTasksResponse struct {
	Ack         int      `json:"Ack"`         // Acknowledgement
	HistoryData []string `json:"HistoryData"` // An ordered list of historical records, where the array elements are the taskid (UUID) of the historical records
}

type RetrieveTaskDetailsRequest struct {
	Id []string `json:"Id"` // Task ID List
}

type RetrieveTaskDetailsResponse struct {
	Ack               int           `json:"Ack"`               // Acknowledgement
	HistoryDetailList []TaskDetails `json:"HistoryDetailList"` // Task Details
}

type EnableDisableVideoStreamRequest struct {
	Enable EnableDisable `json:"Enable"` // Enable or Disable Video Stream
}

type EnableDisableVideoStreamResponse struct {
	Ack      StreamAck `json:"Ack"`      // Acknowledgement
	VideoUrl string    `json:"VideoUrl"` // When opening the video stream, return the RTSP protocol address
}

type EnableDisableTimeLapseRequest struct {
	Enable EnableDisable `json:"Enable"` // Enable or Disable Time-lapse Photography
}

type EnableDisableTimeLapseResponse struct {
	Ack int `json:"Ack"` // Acknowledgement
}

type RequestData[T any] struct {
	Cmd         Command `json:"Cmd"`         // Request Command
	Data        T       `json:"Data"`        // Request Data
	RequestID   string  `json:"RequestID"`   // Request ID
	MainboardID string  `json:"MainboardID"` // Motherboard ID
	TimeStamp   int     `json:"TimeStamp"`   // Timestamp
	From        From    `json:"From"`        // Identify the source of the command
}

type ResponseData[T any] struct {
	Cmd         Command `json:"Cmd"`         // Response Command
	Data        T       `json:"Data"`        // Response Data
	RequestID   string  `json:"RequestID"`   // Request ID
	MainboardID string  `json:"MainboardID"` // Motherboard ID
	TimeStamp   int     `json:"TimeStamp"`   // Timestamp
}
