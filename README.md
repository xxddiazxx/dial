# About
This is my first attempt at microcontrollers and I started this project using the Adafruit USB Rotary Media Dial project. This is an attempt at rewriting the code in tinygo.
Here is a link to the amazing project: https://learn.adafruit.com/usb-rotary-media-dial/overview

# Notes
If you are using vscode, add this to your `.vscode/setings.json` file in this directory to get rid of the gopls linting errors for `machine`:
```
{
    "go.toolsEnvVars": {
        "GOROOT": "/home/ddiaz/.cache/tinygo/goroot-3e9b0b724b1ac271ee187a3e7ef5b08eaf2f4c75d967e00556f93408596c695f",
        "GOFLAGS": "-tags=cortexm,baremetal,linux,arm,rp2040,rp,qtpy_rp2040,tinygo,purego,osusergo,math_big_pure_go,gc.conservative,scheduler.tasks,serial.usb",
        "GOOS": "linux",
        "GOARCH": "arm"
    }
}
```