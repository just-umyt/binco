# Binary Colors

Hi! My name is Umyt, and I’m excited to introduce my console game, **Binary Colors**, written in **Go**.

## Game Rules

The game is very simple:

When you start the game with the command `./binco -start`, you will see something like this:

```bash
                        🟢
             __________0|1__________
       ____0|1____             ____0|1____
    _0|1_       _0|1_       _0|1_       _0|1_
  0|1   0|1   0|1   0|1   0|1   0|1   0|1   0|1
|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
[⬜ 🟨 🟩 🟦 🟥 🟪 🟫 ⬛⬜ 🟨 🟩 🟦 🟥 🟪 🟫 ⬛]
  0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
```

Your goal is to fill all 32 slots with balls as quickly as possible. The color of the ball must match one of the colors shown at the bottom. For example, if the ball is blue:

```bash
                        🔵
            ___________0️⃣|1__________
       ___0️⃣|1____             ____0|1____
    _0|1️⃣_       _0|1_       _0|1_       _0|1_
  0|1   0|1️⃣   0|1   0|1   0|1   0|1   0|1   0|1
|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
|  |  |  |🔵|  |  |  |  |  |  |  |  |  |  |  |  |
[⬜ 🟨 🟩 🟦 🟥 🟪 🟫 ⬛⬜ 🟨 🟩 🟦 🟥 🟪 🟫 ⬛]
  0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
0011 //input
```

You need to input 0011 or 1011, since this is the path to the blue slot.

## Game Modes

The game has two modes:

- **Simple Mode**: The path to the slot is visible on the screen.
- **Pro Mode**: The path is hidden, and you must remember it.

Start Simple Mode:

`./binco -start`

```bash
                        🟢
             __________0|1__________
       ____0|1____             ____0|1____
    _0|1_       _0|1_       _0|1_       _0|1_
  0|1   0|1   0|1   0|1   0|1   0|1   0|1   0|1
|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
[⬜ 🟨 🟩 🟦 🟥 🟪 🟫 ⬛⬜ 🟨 🟩 🟦 🟥 🟪 🟫 ⬛]
  0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
```

Start Pro Mode:

`./binco -start pro`

```bash
         ____                 _____
        |  _ \   🟢          / ____|
        | |_) |  _   _ __   | |        ___
        |  _ <  | | | `_ \  | |       / _ \
        | |_) | | | | | | | | |____  | (_) |
        |____/  |_| |_| |_|  \_____|  \___/
|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
|  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
 ⬜ 🟨 🟩 🟦 🟥 🟪 🟫 ⬛⬜ 🟨 🟩 🟦 🟥 🟪 🟫 ⬛]
  0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
```

## Try Making Your Own Version!

I encourage you to try creating your own version of this game in your favorite programming language. It’s a great way to reinforce your knowledge, improve your skills, or just have some fun!
