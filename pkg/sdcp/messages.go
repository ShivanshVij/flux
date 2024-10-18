package sdcp

type DiscoverMessage struct {
	ID   string       `json:"Id"` // Machine brand identifier, 32-bit UUID
	Data DiscoverData `json:"Data"`
}

type TopicMessage struct {
	Topic string `json:"Topic"` // Topic, used to distinguish the type of reported message
}

type AttributesMessage struct {
	TopicMessage
	Attributes  Attributes `json:"Attributes"`
	MainboardID string     `json:"MainboardID"` // Motherboard ID (16-bit)
	TimeStamp   int        `json:"TimeStamp"`   // Timestamp
}

type StatusMessage struct {
	TopicMessage
	Status      Status `json:"Status"`
	MainboardID string `json:"MachineID"` // Motherboard ID (16-bit)
	TimeStamp   int    `json:"TimeStamp"` // Timestamp
}

type Request[T any] struct {
	TopicMessage
	Id   string         `json:"Id"`   // Machine brand identifier, 32-bit UUID
	Data RequestData[T] `json:"Data"` // Request Data
}

type Response[T any] struct {
	TopicMessage
	Id   string          `json:"Id"`   // Machine brand identifier, 32-bit UUID
	Data ResponseData[T] `json:"Data"` // Response Data
}

type Error struct {
	TopicMessage
	Id   string    `json:"Id"`   // Machine brand identifier, 32-bit UUID
	Data ErrorData `json:"Data"` // Error Data
}

type Notification struct {
	TopicMessage
	Id   string           `json:"Id"`   // Machine brand identifier, 32-bit UUID
	Data NotificationData `json:"Data"` // Notification Data
}
