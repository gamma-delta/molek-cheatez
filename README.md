# Molek-Cheatez

![A picture of the score box for Chloroform with a solution titled "Cheating: Score Hack" solved with 69/1/1](https://github.com/gamma-delta/molek-cheatez/blob/master/pictures/molek-cheatez.png) "A picture of the score box for Chloroform with a solution titled "Cheating: Score Hack" solved with 69/1/1")

I decided to try to reverse-engineer the Molek-Syntez solution file format, just to see if I could. I'd never reverse-engineered any proprietary formats like this before, so I figured it would be fun. Armed with a hexdumper and a C# decompiler, I went to work.

This Go library is used to read and write Molek-Syntez's `.solution` files. Using it, you can:
- Put precursors in puzzles you can't get them in
- Make your scores display anything you like.

Examples:
```go
//Get the bytes of the solution file
solutionBytes, _ := ioutil.ReadFile("amphetamine-1.solution")

solution, _ := molekcheatez.Unmarshal(solutionBytes)

//do nefarious things with the solution file...
hackedSolution := doNefariousThings(solution)

//and write it back to disk
hackedBytes := molekcheatez.Marshal(hackedSolution)
ioutil.WriteFile("amphetamine-2.solution", hackedBytes, 0644)

//In real life you probably want to handle the errors properly.
```

I also wrote a simple viewer and editor called [Submarine](https://github.com/gamma-delta/submarine).

[Reference.md](Reference.md) has the internal numbers used for the precursor IDs, puzzle IDs, and instruction IDs.

# Please don't do anything stupid with this
In fact, it's probably best if you turn off leaderboards. (You can do that in the settings.)

# Format Explanation

`.solution` files start with a header, then contain parts.

Bytes written in `code blocks` are hexadecimal.

Numbers are always stored as little-endian int32s except for instructions, which are uint8s.

### Header

The header always starts with the four bytes `27 27 00 00`. Next is an int32 for the puzzle ID. (Puzzle IDs are not in order as shown in-game.) Next is an int32 for the number of bytes in the solution name, followed by the name itself encoded in UTF-8. (UTF-8 is stored, but non-ASCII characters are displayed as `!`.)

Next is an int32 describing the state of the solution: 0x00 for an unsolved one, 0x03 for a solved one. If the puzzle is solved, there will be 6 bytes afterwards: 0x00, the cycle count of the solution, 0x01, the module count, 0x02, and the symbol count.

After that is the total number of parts in the solution. This includes emitters that aren't being used.

Here are some examples of headers:

```
Unsolved Amphetamine named "NEW SOLUTION 1"
27 27 00 00 | 13 00 00 00 | 0E 00 00 00 | 4E 45 57 20 53 4F 4C 55 54 49 4F 4E 20 31 | 00 00 00 00 | 07 00 00 00
Header        Puzzle ID     Name Length   Name                                        Solved?       Part Count


Solved Kojic Acid named "PRODUCTION"
27 27 00 00 | 06 00 00 00 | 0A 00 00 00 | 50 52 4F 44 55 43 54 49 4F 4E | 03 00 00 00 | 00 00 00 00 | E0 01 00 00 | 01 00 00 00 | 05 00 00 00 | 02 00 00 00 | 08 00 00 00 | 08 00 00 00
Header        Puzzle ID     Name Length  Name                             Cycles head   Cycle #       Modules head  Module #      Symbols head  Symbols #     Part Count

```

### Part

Parts always start with an int32 with their type (`01 00 00 00` for input, `03 00 00 00` for emitter), then 2 int32s for location (first is hexes right, second is hexes up-right, with (0, 0) being the center of the board), then an int32 for rotation. Rotation keeps increasing even if the part is rotated more than 360 degrees, and can be negative.

After is one byte showing the type, again. (I suspect that the first int32 is actually for something else internally and this byte is the type proper.)

Next is some mystery number. On emitters, it is always 0x18, and on inputs, usually one input is assigned to each number starting at 1 (but not always).

Finally are instructions. On emitters, it is 24 bytes with each byte being one instruction. On inputs, it is 28 bytes, with the first always being 0x18 and the rest 0x00. I don't know why.

```
Ammonia:
01 00 00 00 | 00 00 00 00 FF FF FF FF | 00 00 00 00 | 01  | 02 00 00 00 | 03 00 00 00 | 18 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 | 00 00 00 00 
Type          Position                  Rotation      Type  Prec. ID      Mystery       "Opcodes"                                                                 Mystery

Emitter:
03 00 00 00 | F9 FF FF FF 00 00 00 00 | FF FF FF FF | 00  | 00 00 00 00 | 18 00 00 00 | 0A 09 09 09 09 09 09 09 09 09 09 09 09 09 09 09 09 09 09 09 09 09 09 0A
Type          Position                  Rotation      Type  Arm ID        Always 0x18

```
