package main


// Part 3: Methods and Receivers

type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64 { return r.Width * r.Height }

func Area(r Rectangle) float64 { return r.Width * r.Height }


////////////////////////////////////////////////////////////////


type Celsius float64
func (c Celsius) ToFahrenheit() float64 { return float64(c)*9/5 + 32 }