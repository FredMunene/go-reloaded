A SIMPLE TEST COMPLETION/EDITING/AUTO-CORRECTION TOOL 
=============

This is part of the projects I am currently undertaking in learning Go.

The programs receives as arguments the name of a file containing a text that needs modifications (the input) and the name of the file that that the modified text is placed in (the output)

`go run . input.txt output.txt`

#Guidelines to follow

- The project is written in Go
- The code respects the [good practices.](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/good-practices/README.md)
- There are test files for [unit testing](https://go.dev/doc/tutorial/add-a-test).

The modifications that the program does are :
 + Instances of **(hex)** replaces the word before it with the decimal version of the word. *(the word will be a hexadecimal number)* . 
 + Instances of **(bin)** replaces the word before it with the decimal version of the word. *(the word will be a binary number)* . 
 + Instances of **(up)** replaces the word before it with the Uppercase version of the word. 
 + Instances of **(low)** replaces the word before it with the Lowercase version of the word. 
 + Instances of **(cap)** replaces the word before it with the capitalized version of the word. 
    * For **(up)**, **(low)**, **(cap)** if a number appears next to it like: **(low, number)** the program turns the specified number of words before the instance accordingly.
 + Instances of punctuation **(. , ! ? : ; )** are moved close to the previous word and a space put between it and the next word.
    * This applies if there is a group of punctuations like: **...** .

