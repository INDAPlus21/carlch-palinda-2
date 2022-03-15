# Task 2 - Many Senders; Many Receivers

## What happens if you switch the order of the statement `wgp.Wait()` and `close(ch)` in the end of the `main` function?

* Hypothesis: Goroutines are using the channel, but as the program closes the channel before all the subroutines are finished, there is no channel to use and the program crashes.

* Result: As hypothesis states, the program crashes as there is no channel for the goroutines to use.


## What happens if you move the `close(ch)` from the `main` function and instead close the channel in the end of the function `Produce`?

* Hypothesis: When `Produce` finishes, the channel will close for all subroutines and if there is a subroutine still using the channel, the program will crash.

* Result: As hypothesis states, the program crashes as the other goroutines are trying to use the channel which already has been closed.


## What happens if you remove the statement `close(ch)` completetly?

* Hypothesis: Nothing, the channel will not close and will be collected by in-built GCC.

* Result: Cannot prove hypothesis, however, nothing happens but do not know if GCC collects it.


## What happens if you increase the number of consumer from 2 to 4?

* Hypothesis: Time is reduced by roughly 2 as each consumer loops roughly half of what it does with 4 consumers. As each loop, a consumer does the time is increased by `RandomSleep()`.

* Result: Time is reduced by roughly half.


## Can you be sure that all strings are printed before the program stops?

* Hypothesis: Not always, but almost always as `Consume` will finish faster than `Produce` becuase `strconv.Itoa()` is slow AF. But the program does not garuantee.

If you modify the code as per the instructions, the code will garuantee the strings are printed.
