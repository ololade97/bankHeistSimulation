package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())

  isHeistOn := true
  eludedGuards := rand.Intn(100)

  if eludedGuards >= 50 {
    fmt.Println("You just made it past the guards. Move on.")
  } else {
    isHeistOn = false
    fmt.Println("You just failed to overpower the guards.")
  }
 openedVault := rand.Intn(100)

 if isHeistOn && openedVault >= 70 {
  fmt.Println("Vault opened. Grab and go, go!")
 } else if isHeistOn == false {
  fmt.Println("Vault can't be opened.")
 }

leftSafely := rand.Intn(5)

if isHeistOn {
  switch leftSafely {
    case 1:
    isHeistOn = false
    fmt.Println("The heist failed woefully.")
    case 2: 
    isHeistOn = false
    fmt.Println("Looks like you tripped an alarm")
    case 3:
    isHeistOn = false
    fmt.Println("Holyshit, failed asses.")
    case 4:
    isHeistOn = false
    fmt.Println("The vault refused to open.")
    case 5:
    isHeistOn = false
    fmt.Println("You failed!")
    default:
    fmt.Println("Start the damn car. We've cashed out.")
  }

  if isHeistOn {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Println(amtStolen)
  }
}
  fmt.Println("The heist is now", isHeistOn)

}
