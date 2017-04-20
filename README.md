# dots

Turn lines into dots.

Sometimes a command produces a lot of output that you don't care about,
but you still want some vague indication of its progress.

### Install

Grab a binary from the [releases](https://github.com/minrk/dots/releases) page, or build it yourself:

    make && make install

### Use

    $ program-that-makes-a-lot-of-lines | dots
    ..............

Pass an arg to `dots` and it will be used instead of `.`:

    build | dots 💩
    💩💩💩💩💩💩💩💩💩

---

I wrote this in Go because I've never written anything in Go before, and interpreter startup time has been bugging me recently. It even seems to work!
