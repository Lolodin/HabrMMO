import {MainScene} from "./MainScene.js";
let config = {
    type: Phaser.AUTO,
    width: 800,
    height: 600,
    disableContextMenu: true,
    background: 'black',
    physics: {
        default: 'arcade',
        arcadePhysics: {
            overlapBias: 1
        }
    },
    scene:[MainScene],
    pixelArt: true,
    roundPixels: true,
    antialias: true

}
let game = new Phaser.Game(config);
