# Number of 1 Bits

> NeetCode 150 · Bit Manipulation · Blind 75 — Kernighan's trick: n & (n−1) deletes the lowest set bit, so the loop runs once per set bit. Go, stdlib only, tested.

## Problem

Count the set bits (Hamming weight) of a 32-bit unsigned integer.

## Approach

Subtracting 1 flips the lowest set bit to 0 and everything below it to 1s; ANDing with the original therefore erases exactly that one bit. Loop until zero, counting erasures — Kernighan's method, O(set bits) rather than O(32), which is the interview point. The production answer in Go is `math/bits.OnesCount32`, which compiles to a single POPCNT instruction; the test suite uses it as a cross-checking oracle across ten thousand values, which is also a nice pattern — when the stdlib has the function you're reimplementing, it's your free reference implementation.

**Complexity:** O(set bits) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
