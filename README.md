A SIMPLE TEST COMPLETION/EDITING/AUTO-CORRECTION TOOL 
=============

This is part of the projects I am currently undertaking in learning Go.

The programs receives as arguments the name of a file containing a text that needs modifications (the input) and the name of the file that that the modified text is placed in (the output)

`go run . input.txt output.txt`

# Guidelines to follow

- The project is written in Go
- The code respects the [good practices.](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/good-practices/README.md)
- There are test files for [unit testing](https://go.dev/doc/tutorial/add-a-test).
- [Use Standard Go](https://golang.org/pkg/) packages

The modifications that the program does are :
 + Instances of **(hex)** replaces the word before it with the decimal version of the word. *(the word will be a hexadecimal number)* . (Ex: "42 (hex) books" > "66 books")
 + Instances of **(bin)** replaces the word before it with the decimal version of the word. *(the word will be a binary number)* . (Ex: "10 (bin) mice" > "2 mice")
 + Instances of **(up)** replaces the word before it with the Uppercase version of the word. (Ex: "The love of my life (up)" > "The love of my LIFE")
 + Instances of **(low)** replaces the word before it with the Lowercase version of the word. (Ex: "The last day on earth (up)" > "The last day on EARTH")
 + Instances of **(cap)** replaces the word before it with the capitalized version of the word. 
    * For **(up)**, **(low)**, **(cap)** if a number appears next to it like: **(low, number)** the program turns the specified number of words before the instance accordingly.
    (Ex: "My name is frank kim thommas .(cap, 3)" > "My name is Frank Kim Thommas." )
 + Instances of punctuation **(. , ! ? : ; )** are moved close to the previous word and a space put between it and the next word.
    * This also applies if there is a group of punctuations like: **...**.
 + Instances of a pair of apostrophe mark ***(' ')*** are placed to the right and left of the word in the middle of them, without any spaces.
    * Where there is more than one word between the mark, the mark is placed next to corresponding words. 
 + Instances of **a** or **A** are turned into **an** and **An** respectively if the next word begins with vowels *(a, e, i, o ,u, A, E, I, O, U )* or *h.* 

