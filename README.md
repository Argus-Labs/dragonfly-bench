## Cardinal Tickrate Benchmark

This benchmark creates 15,000 player entities and increments its HP at every tick.

The time it takes to complete a tick is logged to the console with 100ms set as the target tick interval.

## Using different in-memory store

- Main branch -> Redis
- Dragonfly -> Dragonfly


**Running the benchmark**
```
docker compose up
```