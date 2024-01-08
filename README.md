# HangManWeb

Hangman is a classic word guessing game where players have to guess a word or phrase by suggesting letters. Guess correctly to prevent the character from being hanged! A fun game to test your vocabulary and deduction skills."

***
## Table of contents

- [Installation](#installation)
- [Depployment](#depployment)
- [Using the Hangman game](#using-the-hangman-game)
- [Licence](#license)

***
## Installation

Before that, you'll need to download golang
https://go.dev/doc/install

```bash
$ git clone https://ytrack.learn.ynov.com/git/rsoleane/Hangman.git
$ go build -o hangman
```

***
## Depployment

To play the game, you'll first need to launch it on the web.
To do this, enter the following command:

```bash
$ ./hangman 
```
Then launch your favorite browser and enter the following url:
```
localhost:8080
```
And that's all there is to it! Now you can move on to the next step.

***
## Using the Hangman game

Now that you've launched the game, you can either log in, register or play without an account !

### Signup / Login

#### 1. Signup


#### 2. Signup
- Launch the game with ./hangman --classic
- You'll be presented with a word to guess in a classic game environment.
- Suggest letters to find the hidden word and save the character.


### Recovering a Savegame
If you've previously saved a game, you can recover it :
- Launch the game with ./hangman --startWith [file]

### Changing the ascii art font
You can choose the ascii art font :
- ./hangman --letterFile thinkertoy.txt
- ./hangman --letterFile shadow.txt
- ./hangman --letterFile standard.txt

***
## License

This game is protected by copyright and is under a proprietary license. Any use, distribution, or modification of this game without prior authorization from Rivier Soleane is strictly prohibited. For licensing or permission inquiries, please contact us at soleane.rivier@ynov.com.

All rights reserved © 2023