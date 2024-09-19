package character

import (
	"github.com/Chronicle20/atlas-kafka/producer"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/segmentio/kafka-go"
)

func enableActionsProvider(worldId byte, channelId byte, characterId uint32) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(characterId))
	value := &statusEvent[statusEventStatChangedBody]{
		CharacterId: characterId,
		Type:        EventCharacterStatusTypeStatChanged,
		WorldId:     worldId,
		Body: statusEventStatChangedBody{
			ChannelId:       channelId,
			ExclRequestSent: true,
		},
	}
	return producer.SingleMessageProvider(key, value)
}

func ChangeMapProvider(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) model.Provider[[]kafka.Message] {
	key := producer.CreateKey(int(characterId))
	value := &commandEvent[changeMapBody]{
		WorldId:     worldId,
		CharacterId: characterId,
		Type:        CommandCharacterChangeMap,
		Body: changeMapBody{
			ChannelId: channelId,
			MapId:     mapId,
			PortalId:  portalId,
		},
	}
	return producer.SingleMessageProvider(key, value)
}
