console.log("audio.js loaded")

playButton = () => {}
stopButton = () => {}
pauseButton = () => {}


muteButton = (button, id) => {
    audio = document.getElementById(id)

    if (audio.muted) {
        audio.muted = false
        button.style.backgroundColor = "lightgray"
        return
    }
    
    audio.muted = true
    button.style.backgroundColor = "yellow"
}

soloButton = (button, id) => {
    // mute all tracks except this one
}