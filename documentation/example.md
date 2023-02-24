```
func #time: num {
    // whatever it is the system does to access time constants
}

func ~#glbinding: context [#driver: system.driver, x: num, y: num] {
    context? = #driver.init(:x, :y)

    if context? {
        return context
    }

    // other boring graphics library context setup
}
```