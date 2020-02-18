export {MainScene}
class MainScene extends Phaser.Scene{
constructor() {
    super({key: 'MainScene'})

}
preload() {
    // Загружаем наши ресурсы в игру
    this.load.image("Ground", "Client/Content/sprGrass.png")
    this.load.image("Water", "Client/Content/sprWater1.png")
    this.load.image("Sand", "Client/Content/sprGrass.png")


}
create() {
    this.getGameMap()
}
update() {

}
// Получаем карту чанка
async getGameMap() {
    let res = await fetch("/map")
    let result = await res.json()
    console.log(result)
    this.drawChunk(result.map)

}
drawChunk(map) {
    for (let chunkKey in map) {
        this.add.image(map[chunkKey].x,map[chunkKey].y, map[chunkKey].key)
    }
}

}