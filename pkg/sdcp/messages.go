package sdcp

type DiscoveryMessage struct {
	ID   string        `json:"Id"` // Machine brand identifier, 32-bit UUID
	Data DiscoveryData `json:"Data"`
}

type AttributesMessage struct {
	Attributes  Attributes `json:"Attributes"`
	MainboardID string     `json:"MainboardID"` // Motherboard ID (16-bit)
	TimeStamp   int        `json:"TimeStamp"`   // Timestamp
	Topic       string     `json:"Topic"`       // Topic, used to distinguish the type of reported message
}

type StatusMessage struct {
	Status      Status `json:"Status"`
	MainboardID string `json:"MainboardID"` // Motherboard ID (16-bit)
	TimeStamp   int    `json:"TimeStamp"`   // Timestamp
	Topic       string `json:"Topic"`       // Topic, used to distinguish the type of reported message
}

type RequestData[T any] struct {
	Cmd         Command `json:"Cmd"`         // Request Command
	Data        T       `json:"Data"`        // Request Data
	RequestID   string  `json:"RequestID"`   // Request ID
	MainboardID string  `json:"MainboardID"` // Motherboard ID
	TimeStamp   int     `json:"TimeStamp"`   // Timestamp
	From        From    `json:"From"`        // Identify the source of the command
}

type Request[T any] struct {
	Id    string         `json:"Id"`    // Machine brand identifier, 32-bit UUID
	Data  RequestData[T] `json:"Data"`  // Request Data
	Topic string         `json:"Topic"` // Topic, used to distinguish the type of reported message
}

type ResponseData[T any] struct {
	Cmd         Command `json:"Cmd"`         // Response Command
	Data        T       `json:"Data"`        // Response Data
	RequestID   string  `json:"RequestID"`   // Request ID
	MainboardID string  `json:"MainboardID"` // Motherboard ID
	TimeStamp   int     `json:"TimeStamp"`   // Timestamp
}

type Response[T any] struct {
	Id    string          `json:"Id"`    // Machine brand identifier, 32-bit UUID
	Data  ResponseData[T] `json:"Data"`  // Response Data
	Topic string          `json:"Topic"` // Topic, used to distinguish the type of reported message
}

type Error struct {
	Id    string    `json:"Id"`    // Machine brand identifier, 32-bit UUID
	Data  ErrorData `json:"Data"`  // Error Data
	Topic string    `json:"Topic"` // Topic, used to distinguish the type of reported message
}

type Notification struct {
	Id    string           `json:"Id"`    // Machine brand identifier, 32-bit UUID
	Data  NotificationData `json:"Data"`  // Notification Data
	Topic string           `json:"Topic"` // Topic, used to distinguish the type of reported message
}
