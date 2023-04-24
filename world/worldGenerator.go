package world

//
//import (
//	"CyberCore/world/blocks"
//	"github.com/Tnze/go-mc/bot/world"
//	"github.com/Tnze/go-mc/level"
//	_ "github.com/Tnze/go-mc/level"
//	"github.com/Tnze/go-mc/net/packet"
//)
//
//type FlatWorldGenerator struct{}
//
//func (wg *FlatWorldGenerator) GenerateChunk(cx, cz int) (level.Chunk, error) {
//	chunk := level.Chunk{BlockEntity: make([]world.EventsListener, 65536)}
//
//	// Заполним чанк заданными блоками
//	for y := 0; y < 3; y++ {
//		var BlockID blocks.ID
//		switch y {
//		case 0:
//			BlockID = blocks.ID(blocks.Bedrock)
//		case 1:
//			BlockID = blocks.Dirt
//		case 2:
//			BlockID = blocks.GrassBlock
//		}
//
//		for z := 0; z < 16; z++ {
//			for x := 0; x < 16; x++ {
//				index := y*256 + z*16 + x
//				chunk.Blocks[index] = world.Block{ID: blockID}
//			}
//		}
//	}
//
//	// Отправим клиенту пакеты для генерации чанка
//	go func() {
//		defer func() {
//			if r := recover(); r != nil {
//				// При возникновении паники завершаем горутину без ошибки
//				// Таким образом, пакеты на отправку в очередь не попадут
//				// и приложение не упадет с ошибкой
//				return
//			}
//		}()
//
//		for i := 0; i < 16; i++ {
//			// Пакет с блоками
//			blockPacket := packet.NewBlockChunk(packet.BlockUpdateData{
//				ChunkX:      int32(cx),
//				ChunkZ:      int32(cz),
//				BlocksData:  chunk.Blocks[i*4096 : (i+1)*4096],
//				BlockEntity: chunk.BlockEntity[i*256 : (i+1)*256],
//			})
//			world.BroadcastPacket(blockPacket)
//
//			// Пакет со светом
//			lightPacket := packet.NewLightChunk(packet.LightUpdateData{
//				ChunkX: int32(cx),
//				ChunkZ: int32(cz),
//				Sky:    make([]uint8, 2048),
//				Block:  make([]uint8, 2048),
//			})
//			world.BroadcastPacket(lightPacket)
//		}
//
//		// Пакет с высотами блоков
//		heightPacket := packet.NewHeightMap(int32(cx), int32(cz), make([]int32, 256))
//		world.BroadcastPacket(heightPacket)
//	}()
//
//	return chunk, nil
//}
