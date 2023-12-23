# Coding Challenges

This repository is a weekly coding challenge. These are challenges that I’ve used or am using as exercises to learn a new programming language or technology.

Each challenge will have you writing a full application or tool. Most of which will be based on real-world tools and utilities.

The challenges are picked to be small enough to be completed in your spare time over a week or two and yet large enough to be complete working projects. 

> Note: This repo is based on [Coding Challenges - John Crickett](https://codingchallenges.fyi/challenges)

# Table of Contents
1. [Write Your Own `wc` Tool](#write-our-own-wc-tool)
2. [Write Your Own JSON Parser](#write-our-own-json-parser)
3. [Write Your Own Compression Tool](#write-our-own-compression-tool)
4. [Write Your Own `cut` Tool](#fourth-examplehttpwwwfourthexamplecom)


## Write Your Own wc Tool <a name="write-our-own-wc-tool"></a>

This challenge is to build your version of the Unix command line tool `wc`!

The Unix command line tools are a great metaphor for good software engineering and they follow the Unix Philosophies of:

- Writing simple parts connected by clean interfaces - each tool does just one thing and provides a simple CLI that handles text input from either files or file streams.
- Design programs to be connected to other programs - each tool can be easily connected to other tools to create incredibly powerful compositions.

Following these philosophies has made the simple Unix command line tools some of the most widely used software engineering tools - allowing us to create very complex text data processing pipelines from simple command line tools. There’s even a Coursera course on [Linux and Bash for Data Engineering](https://gb.coursera.org/learn/linux-and-bash-for-data-engineering-duke).

You can read more about Unix Philosophy in the excellent book [The Art of Unix Programming](http://www.catb.org/~esr/writings/taoup/html/).

## The Challenge - Building wc

The functional requirements for `wc` are concisely described by its man page - give it a go in your local terminal now:

```shell
man wc
```

The TL/DR version is: `wc` – word, line, character, and byte count.

### Step Zero

Like all good software engineering, we’re zero-indexed! In this step, you’re going to set your environment up ready to begin developing and testing your solution.

I’ll leave you to set up your IDE/editor of choice and programming language of choice. After that here’s what I’d like you to do to be ready to test your solution.

Download this [text](https://www.dropbox.com/scl/fi/rbzvip5qru9hp9k6zxeyz/test.txt?rlkey=28vxnv3s7upewp1u95ov5zjxx&dl=0) and save it as `test.txt`.

### Step One

In this step, your goal is to write a simple version of `wc`, let’s call it `ccwc` (`cc` for Coding Challenges) that takes the command line option `-c` and outputs the number of bytes in a file.

If you’ve done it right your output should match this:

```shell
>ccwc -c test.txt
  342190 test.txt
```

If it doesn’t, check your code, fix any bugs, and try again. If it does, congratulations! On to…

### Step Two

In this step, your goal is to support the command line option `-l` that outputs the number of lines in a file.

If you’ve done it right your output should match this:

```shell
>ccwc -l test.txt
    7145 test.txt
```

If it doesn’t, check your code, fix any bugs, and try again. If it does, congratulations! On to…

### Step Three

In this step, your goal is to support the command line option `-w` that outputs the number of words in a file. If you’ve done it right your output should match this:

```shell
>ccwc -w test.txt
   58164 test.txt
```

If it doesn’t, check your code, fix any bugs, and try again. If it does, congratulations! On to…

### Step Four

In this step, your goal is to support the command line option `-m` that outputs the number of characters in a file. If the current locale does not support multibyte characters this will match the `-c` option.

You can learn more about programming for locales here

For this one your answer will depend on your locale, so if can, use `wc` itself and compare the output to your solution:

```shell
>wc -m test.txt
  339292 test.txt
>ccwc -m test.txt
  339292 test.txt
```
If it doesn’t, check your code, fix any bugs, and try again. If it does, congratulations! On to…

### Step Five

In this step your goal is to support the default option - i.e. no options are provided, which is equivalent to the `-c`, `-l` and `-w` options. If you’ve done it right your output should match this:

```shell
>ccwc test.txt
    7145   58164  342190 test.txt
```

If it doesn’t, check your code, fix any bugs, and try again. If it does, congratulations! On to…

### The Final Step

In this step, your goal is to support being able to read from standard input if no filename is specified. If you’ve done it right your output should match this:

```shell
>cat test.txt | ccwc -l
    7145
```

If it doesn’t, check your code, fix any bugs, and try again. If it does, congratulations! You’ve done it, pat yourself on the back, job well done!

## Write Your Own JSON Parser <a name="write-our-own-json-parser"></a>
## Write Your Own Compression Tool <a name="write-our-own-compression-tool"></a>
## Write Your Own cut Tool <a name="write-our-own-cut-tool"></a>
