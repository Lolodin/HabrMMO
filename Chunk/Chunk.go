package Chunk

import (
	"exampleMMO/PerlinNoise"
	"fmt"
	"strconv"
)

var TILE_SIZE = 16
var CHUNK_SIZE = 16 * 16
var PERLIN_SEED float32 = 160

type Chunk struct {
	ChunkID [2]int              `json:"chunkID"`
	Map     map[Coordinate]Tile `json:"map"`
}

/*
Тайтл игрового мира
*/
type Tile struct {
	Key string `json:"key"`
	X   int    `json:"x"`
	Y   int    `json:"y"`
}

/*
Универсальная структура для хранения координат
*/
type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (t Coordinate) MarshalText() ([]byte, error) {

	return []byte("[" + strconv.Itoa(t.X) + "," + strconv.Itoa(t.Y) + "]"), nil
}

/*
Создает карту чанка из тайтлов, генерирует карту на основе координаты чанка
Например [1,1]
*/
func NewChunk(idChunk Coordinate) Chunk {
	fmt.Println("New Chank", idChunk)
	chunk := Chunk{ChunkID: [2]int{idChunk.X, idChunk.Y}}
	var chunkXMax, chunkYMax int
	var chunkMap map[Coordinate]Tile
	chunkMap = make(map[Coordinate]Tile)
	chunkXMax = idChunk.X * CHUNK_SIZE
	chunkYMax = idChunk.Y * CHUNK_SIZE

	switch {
	case chunkXMax < 0 && chunkYMax < 0:
		{
			for x := chunkXMax + CHUNK_SIZE; x > chunkXMax; x -= TILE_SIZE {
				for y := chunkYMax + CHUNK_SIZE; y > chunkYMax; y -= TILE_SIZE {

					posX := float32(x - (TILE_SIZE / 2))
					posY := float32(y + (TILE_SIZE / 2))
					tile := Tile{}
					tile.X = int(posX)
					tile.Y = int(posY)
					perlinValue := PerlinNoise.Noise(posX/PERLIN_SEED, posY/PERLIN_SEED)
					switch {
					case perlinValue < -0.01:
						tile.Key = "Water"
					case perlinValue >= -0.01 && perlinValue < 0:
						tile.Key = "Sand"
					case perlinValue >= 0 && perlinValue <= 0.5:
						tile.Key = "Ground"
					case perlinValue > 0.5:
						tile.Key = "Mount"
					}
					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile

				}
			}
		}
	case chunkXMax < 0:
		{
			for x := chunkXMax + CHUNK_SIZE; x > chunkXMax; x -= TILE_SIZE {
				for y := chunkYMax - CHUNK_SIZE; y < chunkYMax; y += TILE_SIZE {
					posX := float32(x - (TILE_SIZE / 2))
					posY := float32(y + (TILE_SIZE / 2))
					tile := Tile{}
					tile.X = int(posX)
					tile.Y = int(posY)
					perlinValue := PerlinNoise.Noise(posX/PERLIN_SEED, posY/PERLIN_SEED)
					switch {
					case perlinValue < -0.12:
						tile.Key = "Water"
					case perlinValue >= -0.12 && perlinValue <= 0.5:
						tile.Key = "Ground"
					case perlinValue > 0.5:
						tile.Key = "Mount"
					}
					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile
				}
			}
		}
	case chunkYMax < 0:
		{
			for x := chunkXMax - CHUNK_SIZE; x < chunkXMax; x += TILE_SIZE {
				for y := chunkYMax + CHUNK_SIZE; y > chunkYMax; y -= TILE_SIZE {
					posX := float32(x + (TILE_SIZE / 2))
					posY := float32(y - (TILE_SIZE / 2))
					tile := Tile{}
					tile.X = int(posX)
					tile.Y = int(posY)
					perlinValue := PerlinNoise.Noise(posX/PERLIN_SEED, posY/PERLIN_SEED)
					switch {
					case perlinValue < -0.12:
						tile.Key = "Water"
					case perlinValue >= -0.12 && perlinValue <= 0.5:
						tile.Key = "Ground"
					case perlinValue > 0.5:
						tile.Key = "Mount"
					}
					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile

				}
			}
		}
	default:
		{
			for x := chunkXMax - CHUNK_SIZE; x < chunkXMax; x += TILE_SIZE {
				for y := chunkYMax - CHUNK_SIZE; y < chunkYMax; y += TILE_SIZE {
					posX := float32(x + (TILE_SIZE / 2))
					posY := float32(y + (TILE_SIZE / 2))
					tile := Tile{}
					tile.X = int(posX)
					tile.Y = int(posY)
					perlinValue := PerlinNoise.Noise(posX/PERLIN_SEED, posY/PERLIN_SEED)
					switch {
					case perlinValue < -0.12:
						tile.Key = "Water"
					case perlinValue >= -0.12 && perlinValue <= 0.5:
						tile.Key = "Ground"
					case perlinValue > 0.5:
						tile.Key = "Mount"
					}
					chunkMap[Coordinate{X: tile.X, Y: tile.Y}] = tile

				}
			}
		}

	}
	chunk.Map = chunkMap
	return chunk
}
