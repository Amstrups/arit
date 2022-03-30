const { app, BrowserWindow } = require('electron')
const InputHandler = require("./handlers/InputHandler")
const mainWin = require("./components/mainWin")
const PackageHandler = require('./handlers/PackageHandler')
const ParseHandler = require("./handlers/ParseHandler")
const {TagExp} = require("./model/Exp")


var inputController = new InputHandler()

var history = new Array()

const createWindow = () => {
  const win = new BrowserWindow({
    width: 800,
    height: 600,
    frame: true,
    title: "AriT"
  })

  win.loadFile("index.html")
  win.webContents.on("before-input-event", (event, input) => {
        inputController.eval(input, (stack) => {
            ParseHandler.parse(stack, (parsed) => {
                history.push(parsed)
                console.log(parsed)
            }) 
        }, (key) => {
          console.log(`SHOULD_APPEND key ${key}`)
        }, () => {
          console.log("SHOULD_REMOVE")
        })
    });
}

app.whenReady().then(() => {
    createWindow()
})