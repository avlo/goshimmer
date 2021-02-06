package mqtt

import (
	"encoding/json"

	"github.com/iotaledger/goshimmer/packages/tangle"
	"github.com/iotaledger/hive.go/identity"
	"github.com/iotaledger/hive.go/objectstorage"
)

func publishMessage(ev *tangle.CachedMessageEvent) {
	defer ev.MessageMetadata.Release()

	for _, topic := range Topics() {
		if mqttBroker.HasSubscribers(topic) {
			switch topic {
			case "topicMessageMissing",
				"topicMessageUnsolidifiable",
				"topicMessageRemoved":
				ev.Message.Consume(sendMessageResponse(topic))
			default:
				ev.MessageMetadata.Consume(sendMessageMetadataResponse(topic))
			}
		}
	}
}

func sendMessageResponse(topic string) func(msg *tangle.Message) {
	return func(msg *tangle.Message) {
		messageResponse := message{
			MessageID:     msg.ID().String(),
			IssuerID:      identity.New(msg.IssuerPublicKey()).ID().String(),
			Timestamp:     msg.IssuingTime().UnixNano(),
			StrongParents: msg.StrongParents().ToStrings(),
			WeakParents:   msg.WeakParents().ToStrings(),
			Payload:       msg.Payload().Bytes(), // should this be here since no message, just ID?
			Nonce:         msg.Nonce(),
			Signature:     msg.Signature().String(),
		}

		msg.ForEachStrongParent(func(parent tangle.MessageID) {
			messageResponse.StrongParents = append(messageResponse.StrongParents, parent.String())
		})

		msg.ForEachWeakParent(func(parent tangle.MessageID) {
			messageResponse.WeakParents = append(messageResponse.WeakParents, parent.String())
		})

		// Serialize here instead of using publishOnTopic to avoid double JSON marshalling
		jsonPayload, err := json.Marshal(messageResponse)
		if err != nil {
			log.Warn(err.Error())
			return
		}
		mqttBroker.Send(topic, jsonPayload)
	}
}

func sendMessageMetadataResponse(topic string) func(objectStorage objectstorage.StorableObject) {
	return func(objectStorage objectstorage.StorableObject) {
		messageResponse := message{
			// needs filling in, i'm unaware of what should go here
		}

		// Serialize here instead of using publishOnTopic to avoid double JSON marshalling
		jsonPayload, err := json.Marshal(messageResponse)
		if err != nil {
			log.Warn(err.Error())
			return
		}
		mqttBroker.Send(topic, jsonPayload)
	}
}
