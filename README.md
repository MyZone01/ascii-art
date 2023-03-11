#   ASCII ART
##  Description
ASCII-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. Time to write big.

| | | | | | | | | | | | | | | | |
|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|-|
|   |    !	|   "	|   #	|   $	|   %	|   &	|   '	|   (	|   )	|   *	|   +	|   ,	|   -	|   .	|   /
0	|   1	|   2	|   3	|   4	|   5	|   6	|   7	|   8	|   9	|   :	|   ;	|   <	|   =	|   >	|   ?
@	|   A	|   B	|   C	|   D	|   E	|   F	|   G	|   H	|   I	|   J	|   K	|   L	|   M	|   N	|   O
P	|   Q	|   R	|   S	|   T	|   U	|   V	|   W	|   X	|   Y	|   Z	|   [	|   \	|   ]	|   ^	|   _
`	|   a	|   b	|   c	|   d	|   e	|   f	|   g	|   h	|   i	|   j	|   k	|   l	|   m	|   n	|   o
p	|   q	|   r	|   s	|   t	|   u	|   v	|   w	|   x	|   y	|   z	|   {	|   \|	|   }	|   ~	|   

Some **banner** files with a specific graphical template representation using ASCII will be given. The files are formatted in a way that is not necessary to change them. Each character has a height of 8 lines. Characters are separated by a new line \n.
+   [shadow](templates/shadow.txt)
+   [standard](templates/standard.txt)
+   [thinkertoy](templates/thinkertoy.txt)

##  Steps
+   [ ] Parse banner files into a character array
+   [ ] Split the string to be printed into multiple parts
+   [ ] For each part iterate over the character and display it line by line