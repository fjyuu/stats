# stats

`stats` command calculates statistics such as count, mean, standard deviation, min, max and sum.

# Usage

sample file:

```
% cat sample.txt
1
2
3
```

`stats` reads numbers from stdin and calculates statistics:

```
% stats < sample.txt
count	3
mean	2.000000
std	1.000000
min	1.000000
max	3.000000
sum	6.000000
```

# Installation

Download the binary from [Releases](https://github.com/fjyuu/stats/releases).
