package mqtt

// Topic names
const (
	topicMessageSolid           = "message/solid"
	topicMessageAttached        = "message/attached"
	topicMissingMessageReceived = "message/missing_message_received"
	topicMessageMissing         = "message/missing"
	topicMessageUnsolidifiable  = "message/unsolidifiable"
	topicMessageRemoved         = "message/removed"
)

// Topics returns slice
func Topics() []string {
	return []string{
		topicMessageSolid,
		topicMessageAttached,
		topicMissingMessageReceived,
		topicMessageMissing,
		topicMessageUnsolidifiable,
		topicMessageRemoved,
	}
}
