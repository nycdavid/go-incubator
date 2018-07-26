### Pipelines
A pipeline is when a channel connects two goroutines together

* Every receive (`x := <-chan`) operation gets two values:
    * the value sent (`0` if it's a closed channel)
    * A boolean (usually declared as `ok` that's `true` when the channel is open
      and `false` when it's closed)
    ```go
    for {
      x, ok := <-fakechan
      if !ok {
        break // channel is closed, so break
      }
    }
    ```

---

* We can also use the builtin `range` operator to work on channels.
* `range` will continue to iterate as values are sent on a channel and then close
  the loop after the last one is sent.
    ```go
    for val := range fakechan {
      otherfakechan <-val
    }
    close(otherfakechan)
    ```
