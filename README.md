# tk-weight-calc

## Problem statement

Performing workouts that involve loading a 45 lb barbell with weights involves tedious mental math to calculate the number of plates to add onto each side given a total desired weight. It is generally much easier to document the total weight and use it to measure progress rather than the hogde-podge of plates. Example: to bench 180 lbs, the calculations are:

```
180 - 45 = 135 lbs remaining // 45 = bar weight
135 - (45 x 2) -> 135 - 90 = 45 lbs left // 45 lb plates are usually the largest, and one goes on each side of the bar
45 - ((10 x 2) x 2)-> 45 - 20 * 2 -> 45 - 40 = 5 lbs left // can't use 25s, gotta use 2 10 lb plates on each side of the bar
5 - (2.5 * 2) = 0 // finally we can go lift
```

## Solution

This project calculates the number of plates needed to add on each side of a standard 45lb barbell to reach a desired total weight. The intention is to demonstrate the below listed competencies given this relatively simple domain. 

```
Total weight: 180
Result: 
{
    "45":  1,
    "25":  0,
    "10":  2,
    "5":   0,
    "2pt5":1
}
```

## Non-functionals

This needs to be easily accessible while at the gym, and really only service about a dozen calls a week. To keep costs low, it'll be deployed on an AWS lambda and tied to a domain I own already. Golang was chosen as the language to keep the code footprint and maintenance as small and simple as possible. 

## Development workflow

### Steps

- Visit /server_gen/spec.yml 
- Modify as necessary
- `make generate` will generate the server code and models per the config in `server_gen/oapi-codegen-cfg.yml`
- Implementation of the server interface can be found in `server_impl`
- Meat'n potatoes of the project lives in the `calculator` module. 

### Philosophical Notes

I'm big on contracts being front and center during design discussions. It is the first indicator of success in understanding requirements, the simplest form of communication regarding them, and stands a guardian for users sturdier than any other when it comes to the introduction of change. 

I'm also a fan of keeping code as simple as possible, with things exactly where you expect them to be.

## Development Info

See the `Makefile` for local dev commands. 

`go version` -> `go1.25.3 darwin/arm64`

### Environment Variables 
| var name | type | required | description | default | examples |
| --- | --- | :---: | ---| --- | --- |
| `SLOG_LEVEL` | `string` | no | Log level, gets lowercased. Only `info` and `debug` supported for now. |`INFO` | `info`, `DEBUG` |

## Competencies Demonstrated

- Golang fundamentals (modules, lists, maps, unit tests)
- Openapi spec, generation of net/http server interface
- Minimal implementation of net/http server
- Class functions currying (withLogger, for example)
- Overall project organization
- Basics of AWS Lambda functions (design choice, implementation with Terraform) (WIP)
- Micro-front end (adheres to parent style, seamless integration) (WIP)
- Documentation style