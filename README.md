# Feel CLI happiness tracker #

Wanted a project to work on so I could get familiar with a new language and to create something I may find practical for myself.

Simple tool using a SQLite database to track your current mood and feelings. Feel can later provide you data on how you generally feel over time and facilitates reflection on issues you overcame.

<!--
Todo and thoughts:

Use sentiment analysis or models to identify tags for what the user types?
Try to find patterns in the user's comments?

Notes:

Issues compiling go packages using `go get` on my system. Always:

> cc1.exe: sorry, unimplemented: 64-bit mode not compiled in

Is an issue with the architecture, but I'm unsure if it's due to my Cygwin
install or what.

Downloaded Mingw-w64 from [here](https://mingw-w64.org/doku.php/download),
recommended by this [GitHub issue](https://github.com/golang/go/issues/3441).

Specified architecture x86_64 this time (vs i686; e.g., 32-bit).

Changed `C:\MinGW\bin` in env var to be `C:\Program Files\mingw-w64\..\mingw64\bin`

Refreshed shell, looks to have worked.
-->

## Configuration ##



## Examples ##

### now ###

The command for actual reporting.

```
$ feel now [--help|--amend|--delete]
```

Execution looks like:

```
$ feel now

How happy do you feel right now? Choose from 1 (Awful) to 10 (Elated):
5

Anything have you concerned? (Press `enter` to skip)
I feel like I'm going to fail my test

Do you feel grateful for anything? (Press `enter` to skip)
<skipped>

Did you learn anything new today? (Press `enter` to skip)
```

### log ###

```
$ feel log
```

Execution looks like:

```
$ feel log

     Date: Fri Jun 14 16:34:57 2019 -0700
    Score: 3
Concerned: Looked like an idiot to my friends
 Grateful: Memes
  Learned: How to chop onions

     Date: Thu Jun 13 12:31:02 2019 -0700
    Score: 5
Concerned: <no entry>
 Grateful: The Internet
  Learned: They call the Big Mac a Royale with Cheese in France
```

### stat ###

```
$ feel stat [--start|--stop|--ago]
```

The `stat` function provides stats from your self-reported happiness scores.

You can specify a range using the start and stop times. If `--start` is not present, it will use the first-ever entry. Likewise, if `--stop` is not present, it will use the most recent report. The `--ago` flag will provide values for the last number of days and it will ignore the `--start` and `--stop` flags, if also provided. If you don't provide a flag, the default is the last 7 days.

How statistically useful this is may depend on how consistent you are in reporting. Take it with a grain of salt.

Execution looks like:

```
$ feel stat

Stats for Jun 7 2019 to Jun 14 2019:

Score: 3.4

Concerns: 4 (57.1%)
Grateful: 7 (100%)
Learned: 5 (71.4%)
```

### reflect ###

In progress. Does not provide anything that isn't obtainable with `log`.

```
$ feel reflect [--help|--amount|--random]
```

Execution looks like:
