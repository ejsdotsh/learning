package jedlik

import "fmt"

// Drive updates the number of meters driven based on the car's speed,
// and reduces the battery according to the battery drainage.
func (c *Car) Drive() {
	    if c.battery - c.batteryDrain > 0 {
            c.distance += c.speed
            c.battery -= c.batteryDrain
        }
}

// DisplayDistance returns the distance as a string
func (c *Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery returns the battery fill percentage as a string
func (c *Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish takes an integer trackDistance, and returns a bool if the car
// can finish the race
func (c *Car) CanFinish(trackDistance int) bool {
    d := trackDistance / c.speed
    if trackDistance % c.speed != 0 {
        d += 1
    }
    if d * c.batteryDrain <= c.battery {
        return true
    }
    return false
}