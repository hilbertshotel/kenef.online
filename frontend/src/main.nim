import jsconsole
import dom

console.log("nim")

proc getId*(id: string): Node =
    document.getElementById(id)

proc getAudio*(id: string): EmbedElement =
    return getId(id).EmbedElement

proc playSound*(id: string) =
    getAudio(id).play()

proc play() {.exportc.} = 
    playSound("tri_e_mnogo.mp3")
