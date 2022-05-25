# Getting Started

- Expects `golangci-lint` to be installed for linting purposes.
- `go module download` to install the dependencies.
    - Although the core features do not use any external, testing framework and command line flag helper and postgres is
      ued.

# Solutions:

## LCG

> Linear Congruential Generator

- Upside
    - Single state
    - Very efficient
    - Adding one more variable can increase the randomness.
        - Multiple services running in round-robin etc.
- Downsides
    - If more than 8 digits the requirement to find prime or square of prime increases one time calculation time to
      generate it.
    - Need to screen each generated number.
    - Single threaded as previous value is always a dependency.

## Brute Force

> Brute Force method initially made to test the sequence generation, when I had the idea of using postgres to generate
> the random numbers.

- Upsides
    - By loading up all the possible scenarios and removing the unfavourable hex codes, we get a dataset that is
      passively
      screened.
    - If it were a service, multiple services could leverage this system to generate random number in-line with the
      sequence. (Which is repeat only after exhausting all possible values)
    - Leverage Postgres's random selection
- Downsides
    - Postgres memory usage.
    - Though CPU usage is lower, space taken by the database is huge.
    - 