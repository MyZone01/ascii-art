#   🎨 ASCII ART

##  Description
ASCII-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. Time to write big.

<center>

| | | | | | | | | | | | | | | | |
|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|
|   |    !	|   "	|   #	|   $	|   %	|   &	|   '	|   (	|   )	|   *	|   +	|   ,	|   -	|   .	|   /
0	|   1	|   2	|   3	|   4	|   5	|   6	|   7	|   8	|   9	|   :	|   ;	|   <	|   =	|   >	|   ?
@	|   A	|   B	|   C	|   D	|   E	|   F	|   G	|   H	|   I	|   J	|   K	|   L	|   M	|   N	|   O
P	|   Q	|   R	|   S	|   T	|   U	|   V	|   W	|   X	|   Y	|   Z	|   [	|   \	|   ]	|   ^	|   _
`	|   a	|   b	|   c	|   d	|   e	|   f	|   g	|   h	|   i	|   j	|   k	|   l	|   m	|   n	|   o
p	|   q	|   r	|   s	|   t	|   u	|   v	|   w	|   x	|   y	|   z	|   {	|   \|	|   }	|   ~	|   

</center>

Some **banner** files with a specific graphical template representation using ASCII will be given. The files are formatted in a way that is not necessary to change them. Each character has a height of 8 lines. Characters are separated by a new line \n.
+   [shadow](fonts/shadow.txt)
+   [standard](fonts/standard.txt)
+   [thinkertoy](fonts/thinkertoy.txt)
+   [htag](fonts/htag.txt)
+   [zigzag](fonts/zigzag.txt)

##  Usage
Longest allowed command
```
go run . --align=justify --output=output --color=#980fd0 iuoae "Serigne Saliou\nMbaye" shadow
```

##  Steps
+   [x] Core
    +   [x] Parse banner files into a character array
    +   [x] Split the string to be printed into multiple parts
    +   [x] For each part iterate over the character and display it line by line
+   [x] Output
+   [ ] Reverse
    +   [x] " H"
    +   [x] "  "
    +   [x] "\n H"
    +   [x] "H \n H"
    +   [x] "\n"
    +   [x] "H\n"
+   [x] Font
    +   [x] Add Custom font
+   [x] Align
    +   [x] Center
    +   [x] Right
    +   [x] Left
    +   [x] Justify
+   [ ] Color
    +   [x] Name
    +   [x] Hex code
    +   [x] RGB code
    +   [ ] HSL code
+   [x] Combine flag
+   [x] Error Handling
+   [ ] Test file
    +   [ ] Unit test
    +   [ ] Font test
    +   [x] Basic test