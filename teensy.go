package main

import (
    "fmt"
    "github.com/google/gousb"
    "github.com/marcinbor85/gohex"
    "runtime"
    "log"
    "os"
    "time"
    )

// TeensyFlash: Flashes Teensy boards.
// It opens the firmware file at the provided path, checks it's integrity, wait for the keyboard to be in Flash mode, flashes it and reboots the board.
func teensyFlash(firmwarePath string, s *state) {
    file, err := os.Open(firmwarePath)
    if err != nil {
        message := fmt.Sprintf("Error while opening firmware: %s", err)
        log.Fatal(message)
        return
    }
    defer file.Close()

    s.total = ergodoxCodeSize

    firmware := gohex.NewMemory()
    err = firmware.ParseIntelHex(file)
    if err != nil {
        message := fmt.Sprintf("Error while parsing firmware: %s", err)
        log.Fatal(message)
        return
    }

    ctx := gousb.NewContext()
    ctx.Debug(0)
    defer ctx.Close()
    var dev *gousb.Device

    // Loop until a keyboard is ready to flash
    for {
        devs, _ := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
            if desc.Vendor == gousb.ID(halfKayVendorID) && desc.Product == gousb.ID(halfKayProductID) {
                return true
            }
            return false
        })

        defer func() {
            for _, d := range devs {
                d.Close()
            }
        }()

        if len(devs) > 0 {
            dev = devs[0]
            break
        }
        time.Sleep(1 * time.Second)
    }

    if runtime.GOOS != "darwin" {
        dev.SetAutoDetach(true)
    }

    // Claim usb device
    cfg, err := dev.Config(1)
    defer cfg.Close()
    if err != nil {
        message := fmt.Sprintf("Error while claiming the usb interface: %s", err)
        log.Fatal(message)
        return
    }

    s.step = 1

    // Loop on the firmware data and program
    var addr uint32
    for addr = 0; addr < ergodoxCodeSize; addr += ergodoxBlockSize {
        // set a longer timeout when writing the first block
        if addr == 0 {
            dev.ControlTimeout = 5 * time.Second
        } else {
            dev.ControlTimeout = 500 * time.Millisecond
        }
        // Prepare and write a firmware block
        // https://www.pjrc.com/teensy/halfkay_protocol.html
        buf := make([]byte, ergodoxBlockSize+2)
        buf[0] = byte(addr & 255)
        buf[1] = byte((addr >> 8) & 255)
        block := firmware.ToBinary(addr, ergodoxBlockSize, 255)
        for index := range block {
            buf[index+2] = block[index]
        }

        bytes, err := dev.Control(0x21, 9, 0x0200, 0, buf)
        if err != nil {
            message := fmt.Sprintf("Error while sending firmware bytes: %s", err)
            log.Fatal(message)
            return
        }

        s.sent += bytes
    }

    buf := make([]byte, ergodoxBlockSize+2)
    buf[0] = byte(0xFF)
    buf[1] = byte(0xFF)
    buf[2] = byte(0xFF)
    _, err = dev.Control(0x21, 9, 0x0200, 0, buf)

    if err != nil {
        message := fmt.Sprintf("Error while rebooting device: %s", err)
        log.Fatal(message)
        return
    }
    s.step = 2
}
